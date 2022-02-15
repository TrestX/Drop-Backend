package seller_registration

import (
	"errors"
	"strings"
	"time"

	"github.com/aekam27/trestCommon"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"

	api "Drop/DropUserAccount/api"
	controller "Drop/DropUserAccount/controller/user-registration"
	entity "Drop/DropUserAccount/entities"
	"Drop/DropUserAccount/repository/user"
)

var (
	repo = user.NewProfileRepository("users")
)
var (
	accountService = controller.NewSignUpService(user.NewProfileRepository("users"))
)

type sellerAccountService struct{}

func NewSellerRegisterationService(repository user.UserRepository) SellerAccountService {
	repo = repository
	return &sellerAccountService{}
}
func (*sellerAccountService) RegisterSellerPerson(cred Credentials) (string, error) {
	if cred.Password == "" {
		err := errors.New("password missing")
		trestCommon.ECLog2("sign up failed no password", err)
		return "", err
	}
	_, err := checkUser(cred)
	if err != nil {
		trestCommon.ECLog2("sign up user not found", err)
		if err.Error() == "mongo: no documents in result" {
			var serv *sellerAccountService
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

func (*sellerAccountService) hashAndInsertData(cred Credentials) (string, error) {
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
	pictureId := ""
	if cred.PictureID != "" {
		pictureId = createPreSignedDownloadUrl(cred.PictureID)
	}
	cred.PictureID = pictureId
	nationalId := ""
	if cred.NationalID != "" {
		nationalId = createPreSignedDownloadUrl(cred.NationalID)
	}
	cred.NationalID = nationalId
	certificate := ""
	if cred.CertificateOfIncorporation != "" {
		certificate = createPreSignedDownloadUrl(cred.CertificateOfIncorporation)
	}
	cred.CertificateOfIncorporation = certificate
	cred.Password = string(hash)
	cred.CreatedDate = time.Now()
	cred.AccountType = "Seller"
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
