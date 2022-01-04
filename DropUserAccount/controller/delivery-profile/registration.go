package delivery_registration

import (
	api "Drop/DropUserAccount/api"
	controller "Drop/DropUserAccount/controller/user-registration"
	entity "Drop/DropUserAccount/entities"

	"Drop/DropUserAccount/repository/user"
	"errors"
	"strings"
	"time"

	"github.com/aekam27/trestCommon"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var (
	repo = user.NewProfileRepository("users")
)
var (
	accountService = controller.NewSignUpService(user.NewProfileRepository("users"))
)

type deliveryAccountService struct{}

func NewDeliveryRegisterationService(repository user.UserRepository) DeliveryAccountService {
	repo = repository
	return &deliveryAccountService{}
}
func (*deliveryAccountService) RegisterDeliveryPerson(cred Credentials) (string, error) {
	if cred.Password == "" {
		err := errors.New("password missing")
		trestCommon.ECLog2("sign up failed no password", err)
		return "", err
	}
	_, err := checkUser(cred)
	if err != nil {
		trestCommon.ECLog2("sign up user not found", err)
		if err.Error() == "mongo: no documents in result" {
			var serv *deliveryAccountService
			tokenString, err := serv.hashAndInsertData(cred)
			if err != nil {
				trestCommon.ECLog3("sign up not successful", err, logrus.Fields{"email": cred.Email})
			}
			return tokenString, err
		} else {
			return "", err
		}
	}
	return "", errors.New("email already registed")
}

func checkUser(cred Credentials) (entity.UserDB, error) {
	var userData entity.UserDB
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

func (*deliveryAccountService) hashAndInsertData(cred Credentials) (string, error) {
	salt := viper.GetString("salt")
	hash, err := bcrypt.GenerateFromPassword([]byte(cred.Password+salt), 5)
	if err != nil {
		trestCommon.ECLog3("hash password", err, logrus.Fields{"email": cred.Email})
		return "", err
	}
	profilePicUrl := ""
	if cred.ProfilePhoto != "" {
		profilePicUrl = createPreSignedDownloadUrl(cred.ProfilePhoto)
	}
	cred.ProfilePhoto = profilePicUrl
	vehiclePhoto := ""
	if cred.VehiclePhoto != "" {
		vehiclePhoto = createPreSignedDownloadUrl(cred.VehiclePhoto)
	}
	vehicleRegistration := ""
	if cred.VehicleRegistration != "" {
		vehicleRegistration = createPreSignedDownloadUrl(cred.VehicleRegistration)
	}
	nationalId := ""
	if cred.NationalID != "" {
		nationalId = createPreSignedDownloadUrl(cred.NationalID)
	}
	cred.VehiclePhoto = vehiclePhoto
	cred.VehicleRegistration = vehicleRegistration
	cred.NationalID = nationalId
	cred.Password = string(hash)
	cred.CreatedDate = time.Now()
	cred.AccountType = "Delivery"
	cred.Status = "created"
	cred.Deleted = false
	var address entity.Address
	address = cred.Address
	cred.Address = entity.Address{}
	userid, err := repo.InsertOne(cred)
	if err != nil {
		trestCommon.ECLog3("hashAndInsertData Insert failed", err, logrus.Fields{"email": cred.Email})
		return "", nil
	}
	token, _ := trestCommon.CreateToken(userid, cred.Email, "", cred.Status)
	_, err = api.AddUserAddress(address, token)
	if err != nil {
		trestCommon.ECLog3("hashAndInsertData Insert failed", err, logrus.Fields{"email": cred.Email})
	}
	_, err = accountService.SendVerificationEmail(cred.Email)
	if err != nil {
		trestCommon.ECLog3("hashAndInsertData Insert failed", err, logrus.Fields{"email": cred.Email})
	}
	return token, nil
}
func createPreSignedDownloadUrl(url string) string {
	s := strings.Split(url, "?")
	o := strings.Split(s[0], "/")
	fileName := o[4]
	path := o[3]
	downUrl, _ := trestCommon.PreSignedDownloadUrl(fileName, path)
	return downUrl
}
