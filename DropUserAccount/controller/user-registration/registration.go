package user_registration

import (
	entity "Drop/DropUserAccount/entities"
	"Drop/DropUserAccount/repository/user"
	"errors"
	"time"

	"firebase.google.com/go/auth"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"

	"github.com/aekam27/trestCommon"
)

var (
	repo = user.NewProfileRepository("users")
)

type accountService struct{}

func NewSignUpService(repository user.UserRepository) AccountService {
	repo = repository
	return &accountService{}
}

func (*accountService) GSignUp(token *auth.Token) (string, error) {
	_, err := repo.FindOne(bson.M{"uid": token.UID}, bson.M{})
	if err != nil {
		trestCommon.ECLog2("sign up user not found", err)
		if err.Error() == "mongo: no documents in result" {
			res, _ := repo.InsertOne(token)
			tokenString, err := trestCommon.CreateToken(res, "", "", "created")
			if err != nil {
				trestCommon.ECLog3("sign up not successful", err, logrus.Fields{"email": ""})
			}
			return tokenString, err
		} else {
			return "", err

		}
	}
	return "", errors.New("email already registed")
}
func (*accountService) GLogin(token *auth.Token) (string, error) {
	uData, err := repo.FindOne(bson.M{"uid": token.UID}, bson.M{})
	if err != nil {
		trestCommon.ECLog2("login failed password hash doesn't match", err)
		return "", err
	}
	tokenString, err := trestCommon.CreateToken(uData.ID.Hex(), "", "", uData.Status)
	if err != nil {
		trestCommon.ECLog3("login failed unable to create token", err, logrus.Fields{})
		return "", err
	}
	repo.UpdateOne(bson.M{"_id": uData.ID}, bson.M{"$set": bson.M{"login_time": time.Now()}})
	return tokenString, nil
}

func (*accountService) SignUp(cred Credentials) (string, error) {
	if cred.Password == "" {
		err := errors.New("password missing")
		trestCommon.ECLog2("sign up failed no password", err)
		return "", err
	}
	_, err := checkUser(cred)

	if err != nil {
		trestCommon.ECLog2("sign up user not found", err)
		if err.Error() == "mongo: no documents in result" {
			var serv *accountService
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

func (*accountService) SendVerificationEmail(email string) (string, error) {
	emailSentTime := time.Now()
	verificationCode := trestCommon.GetRandomString(16)
	sendCode, err := trestCommon.Encrypt(email + ":" + verificationCode)
	if err != nil {
		trestCommon.ECLog2("send verification email encryption failed", err)
		return "", err
	}
	_, err = trestCommon.SendVerificationCode(email, sendCode)
	if err != nil {
		trestCommon.ECLog2("send verification email failed", err)
		return "", err
	}
	_, err = repo.UpdateOne(bson.M{"email": email}, bson.M{"$set": bson.M{"email_sent_time": emailSentTime, "verification_code": verificationCode}})
	if err != nil {
		trestCommon.ECLog2("send verification email update failed", err)
		return "", err
	}
	return "email sent successfully", nil
}

func (*accountService) Login(cred Credentials) (string, string, error) {
	if cred.Password == "" {
		err := errors.New("password missing")
		trestCommon.ECLog2("login failed no password", err)
		return "", "", err
	}
	salt := viper.GetString("salt")
	userData, err := checkUser(cred)
	if err != nil {
		trestCommon.ECLog2("login failed user not found", err)
		return "", "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(cred.Password+salt))
	if err != nil {
		trestCommon.ECLog2("login failed password hash doesn't match", err)
		return "", "", err
	}
	tokenString, err := trestCommon.CreateToken(userData.ID.Hex(), cred.Email, "", userData.Status)
	if err != nil {
		trestCommon.ECLog3("login failed unable to create token", err, logrus.Fields{"email": cred.Email, "name": userData.Name, "status": userData.Status})
		return "", "", err
	}
	repo.UpdateOne(bson.M{"_id": userData.ID}, bson.M{"$set": bson.M{"login_time": time.Now()}})

	return tokenString, userData.ID.Hex(), nil
}

func (*accountService) VerifyEmail(cred Credentials) (string, error) {

	userData, err := checkUser(cred)
	if err != nil {
		trestCommon.ECLog3("verify user not found", err, logrus.Fields{"email": cred.Email})
		return "", err
	}

	if cred.VerificationCode != userData.VerificationCode {
		err = errors.New("unauthorized user")
		trestCommon.ECLog3("verify user verification code didn't match", err, logrus.Fields{"email": cred.Email, "db verify code": userData.VerificationCode, "code provided by user": cred.VerificationCode})
		return "", err
	}
	if userData.Status == "verified" {
		err = errors.New("user already verified")
		trestCommon.ECLog3("verify user verification user already verified", err, logrus.Fields{"email": cred.Email})
		return "", err
	}
	_, err = repo.UpdateOne(bson.M{"_id": userData.ID}, bson.M{"$set": bson.M{"verified_time": time.Now(), "status": "verified"}})
	if err != nil {
		trestCommon.ECLog3("verify user unable to update status", err, logrus.Fields{"email": cred.Email})
		return "", err
	}
	return "verified", nil
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

func (*accountService) hashAndInsertData(cred Credentials) (string, error) {
	salt := viper.GetString("salt")

	hash, err := bcrypt.GenerateFromPassword([]byte(cred.Password+salt), 5)
	if err != nil {
		trestCommon.ECLog3("hash password", err, logrus.Fields{"email": cred.Email})
		return "", err
	}
	cred.Password = string(hash)
	cred.CreatedDate = time.Now()
	cred.Status = "created"
	cred.Deleted = false
	userid, err := repo.InsertOne(cred)
	if err != nil {
		trestCommon.ECLog3("hashAndInsertData Insert failed", err, logrus.Fields{"email": cred.Email})
		return "", nil
	}
	var serv accountService
	_, err = serv.SendVerificationEmail(cred.Email)
	if err != nil {
		trestCommon.ECLog3("hashAndInsertData Insert failed", err, logrus.Fields{"email": cred.Email})
	}
	return trestCommon.CreateToken(userid, cred.Email, "", cred.Status)
}
func (*accountService) SendResetLink(email string) (string, error) {
	var cred Credentials
	cred.Email = email
	_, err := checkUser(cred)
	if err != nil {
		trestCommon.ECLog2("user not found", err)
		return "", err

	}
	emailSentTime := time.Now()
	verificationCode := trestCommon.GetRandomString(16)
	resetCode, err := trestCommon.Encrypt(email + ":" + verificationCode)
	if err != nil {
		trestCommon.ECLog2("send reset link encryption failed", err)
		return "", err
	}
	_, err = trestCommon.SendResetPasswordLink(email, resetCode)
	if err != nil {
		trestCommon.ECLog2("send reset password link failed", err)
		return "", errors.New("send reset password link failed")
	}
	_, err = repo.UpdateOne(bson.M{"email": email}, bson.M{"$set": bson.M{"email_sent_time": emailSentTime, "password_reset_code": verificationCode}})
	if err != nil {
		trestCommon.ECLog2("send reset link update failed", err)
		return "", err
	}
	return "Reset link sent successfully", nil
}

func (*accountService) VerifyResetLink(cred Credentials) (string, string, error) {

	userData, err := checkUser(cred)
	if err != nil {
		trestCommon.ECLog3("verify user not found", err, logrus.Fields{"email": cred.Email})
		return "", "", err
	}

	if cred.PasswordResetCode != userData.PasswordResetCode {
		err = errors.New("unauthorized user")
		trestCommon.ECLog3("verify user password reset code didn't match", err, logrus.Fields{"email": cred.Email, "db verify code": userData.PasswordResetCode, "code provided by user": cred.PasswordResetCode})
		return "", "", err
	}
	_, err = repo.UpdateOne(bson.M{"_id": userData.ID}, bson.M{"$set": bson.M{"password_reset_time": time.Now()}})
	if err != nil {
		trestCommon.ECLog3("verify user unable to update status", err, logrus.Fields{"email": cred.Email})
		return "", "", err
	}

	return "verified", userData.Email, nil
}

func (*accountService) UpdatePassword(cred Credentials) (string, error) {
	if cred.Password == "" {
		err := errors.New("password missing")
		trestCommon.ECLog2("update password failed no password", err)
		return "", err
	}
	if cred.CurrentPassword == "" {
		err := errors.New("current password missing")
		trestCommon.ECLog2("update password failed no password", err)
		return "", err
	}
	salt := viper.GetString("salt")
	userData, err := checkUser(cred)
	if err != nil {
		trestCommon.ECLog2("update password failed user not found", err)
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(cred.CurrentPassword+salt))
	if err != nil {
		trestCommon.ECLog3("hash password", err, logrus.Fields{"email": cred.Email})
		return "", errors.New("password doesnot match")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(cred.Password+salt), 5)
	if err != nil {
		trestCommon.ECLog3("hash password", err, logrus.Fields{"email": cred.Email})
		return "", err
	}
	cred.Password = string(hash)
	cred.CreatedDate = time.Now()
	_, err = repo.UpdateOne(bson.M{"email": cred.Email}, bson.M{"$set": bson.M{"password": cred.Password, "update_time": time.Now()}})
	if err != nil {
		trestCommon.ECLog2("password update failed", err)
		return "", err
	}

	return "password updated successfully", nil
}
