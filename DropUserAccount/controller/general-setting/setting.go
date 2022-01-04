package general

import (
	entity "Drop/DropUserAccount/entities"
	"Drop/DropUserAccount/repository/user"
	"errors"
	"time"

	"github.com/aekam27/trestCommon"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = user.NewProfileRepository("users")
)

type settingService struct{}

func NewSettingService(repository user.UserRepository) SettingService {
	repo = repository
	return &settingService{}
}

func (*settingService) SetSetting(setting *Setting, userid string) (string, error) {
	if userid == "" {
		err := errors.New("user id missing")
		trestCommon.ECLog2(
			"update setting section",
			err,
		)
		return "", err
	}
	id, _ := primitive.ObjectIDFromHex(userid)
	setParameters := bson.M{}
	_, err := checkUser(userid)
	if err != nil {
		trestCommon.ECLog2(
			"update setting section",
			err,
		)
		return "", err
	}

	if setting.Language != "" {
		setParameters["language"] = setting.Language
	}
	if setting.Theme != "" {
		setParameters["theme"] = setting.Theme
	}
	setParameters["update_time"] = time.Now()
	filter := bson.M{"_id": id}
	set := bson.M{
		"$set": setParameters,
	}
	result, err := repo.UpdateOne(filter, set)
	if err != nil {
		trestCommon.ECLog3(
			"update setting section",
			err,
			logrus.Fields{
				"user_id": userid,
				"setting": setting,
			})
		return "", err
	}
	return result, nil
}

func (*settingService) GetSetting(userID string) (entity.UserDB, error) {
	if userID == "" {
		err := errors.New("user id missing")
		trestCommon.ECLog2(
			"GetProfile section",
			err,
		)
		return entity.UserDB{}, err
	}
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		trestCommon.ECLog3(
			"GetProfile section",
			err,
			logrus.Fields{
				"user_id": userID,
			},
		)
		return entity.UserDB{}, err
	}
	setting, err := repo.FindOne(bson.M{"_id": id}, bson.M{"language": 1, "theme": 1})
	if err != nil {
		trestCommon.ECLog2(
			"GetProfile section",
			err,
		)
		return setting, err
	}
	setting.Password = ""
	return setting, nil
}

func checkUser(userID string) (entity.UserDB, error) {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		trestCommon.ECLog3(
			"CheckUser section",
			err,
			logrus.Fields{
				"user_id": userID,
			},
		)
		return entity.UserDB{}, err
	}
	return repo.FindOne(bson.M{"_id": id}, bson.M{})
}
