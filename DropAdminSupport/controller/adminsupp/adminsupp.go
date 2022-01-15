package admin

import (
	entity "Drop/DropAdminSupport/entities"
	"Drop/DropAdminSupport/repository/admin"
	"errors"
	"time"

	"github.com/aekam27/trestCommon"
	"github.com/dgrijalva/jwt-go"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var (
	repo = admin.NewAdminRepository("admin")
)

type adminService struct{}

func NewAdminService(repository admin.AdminRepository) AdminService {
	repo = repository
	return &adminService{}
}
func (add *adminService) AddAcct(acct *Acct) (string, error) {
	var acctEntity entity.AdminSupportDB
	acctEntity.Role = acct.Role
	acctEntity.Status = "Active"
	acctEntity.CreatedOn = time.Now()
	acctEntity.ID = primitive.NewObjectID()
	acctEntity.Name = acct.Name
	acctEntity.Email = acct.Email
	acctEntity.Type = acct.Type
	salt := viper.GetString("salt")
	hash, err := bcrypt.GenerateFromPassword([]byte(acct.Password+salt), 5)
	if err != nil {
		trestCommon.ECLog3("hash password", err, logrus.Fields{"email": acct.Email})
		return "", err
	}
	acctEntity.Password = string(hash)
	return repo.InsertOne(acctEntity)
}

func (*adminService) UpdateAcctStatus(acct *Acct, iid string) (string, error) {
	if iid == "" {
		err := errors.New("id missing")
		trestCommon.ECLog2(
			"update",
			err,
		)
		return "", err
	}
	id, _ := primitive.ObjectIDFromHex(iid)
	_, err := checkByCategoryID(id)
	if err != nil {
		return "", errors.New("invalid Id")
	}
	setParameters := bson.M{}
	if acct.Status != "" {
		setParameters["status"] = acct.Status
	}
	if acct.Name != "" {
		setParameters["name"] = acct.Name
	}
	if acct.Role != "" {
		setParameters["role"] = acct.Role
	}
	setParameters["updated_time"] = time.Now()
	filter := bson.M{"_id": id}
	set := bson.M{
		"$set": setParameters,
	}
	result, err := repo.UpdateOne(filter, set)
	if err != nil {
		trestCommon.ECLog3(
			"update ",
			err,
			logrus.Fields{
				"id": iid,
			})
		return "", err
	}

	return result, nil
}

func checkByCategoryID(id primitive.ObjectID) (entity.AdminSupportDB, error) {
	acct, err := repo.FindOne(bson.M{"_id": id}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get Acct Details section",
			err,
		)
		return acct, err
	}
	return acct, nil
}

func (*adminService) GetAllAcct(limit, skip int) ([]entity.AdminSupportDB, error) {
	return repo.Find(bson.M{"type": "Drop"}, bson.M{}, limit, skip)
}

func (*adminService) GetAllShopAcct(limit, skip int) ([]entity.AdminSupportDB, error) {
	return repo.Find(bson.M{"type": "Seller"}, bson.M{}, limit, skip)
}

func (*adminService) Login(cred Credentials) (string, error) {
	if cred.Password == "" {
		err := errors.New("password missing")
		trestCommon.ECLog2("login failed no password", err)
		return "", err
	}
	salt := viper.GetString("salt")
	userData, err := checkUser(cred)
	if err != nil {
		trestCommon.ECLog2("login failed user not found", err)
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(cred.Password+salt))
	if err != nil {
		trestCommon.ECLog2("login failed password hash doesn't match", err)
		return "", err
	}
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["userid"] = userData.ID.Hex()
	atClaims["email"] = cred.Email
	atClaims["status"] = userData.Status
	atClaims["role"] = userData.Role
	atClaims["name"] = userData.Name
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tokenString, err := at.SignedString([]byte(viper.GetString("tokensecret")))
	if err != nil {
		trestCommon.ECLog3("login failed unable to create token", err, logrus.Fields{"email": cred.Email, "name": userData.Name, "status": userData.Status})
		return "", err
	}
	repo.UpdateOne(bson.M{"_id": userData.ID}, bson.M{"$set": bson.M{"login_time": time.Now()}})

	return tokenString, nil
}

func checkUser(cred Credentials) (entity.AdminSupportDB, error) {
	var userData entity.AdminSupportDB
	if cred.Email == "" {
		err := errors.New("email missing")
		trestCommon.ECLog2("check user failed no email", err)
		return userData, err
	}
	if !trestCommon.ValidateEmail(cred.Email) {
		err := errors.New("invalid email")
		trestCommon.ECLog2("check user failed invalid email", err)
		return userData, err
	}
	return repo.FindOne(bson.M{"email": cred.Email}, bson.M{})
}
