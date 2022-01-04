package admin

import (
	entity "Drop/DropSettings/entities"
	"Drop/DropSettings/repository/admin"

	"github.com/aekam27/trestCommon"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = admin.NewAdminSettingRepository("setting")
)

type adminService struct{}

func NewAdminSettingService(repository admin.AdminSettingRepository) AdminSettingService {
	repo = repository
	return &adminService{}
}
func (add *adminService) AddSettings(setting Settings, token string) (string, error) {
	res, _ := repo.Find(bson.M{}, bson.M{}, 100, 0)
	if len(res) > 0 {
		filter := bson.M{}
		settingss, _ := repo.FindOne(bson.M{"current": true}, bson.M{})
		id, _ := primitive.ObjectIDFromHex(settingss.ID.Hex())
		filter["_id"] = id
		setParameters := bson.M{}
		setParameters["updated_time"] = time.Now()
		setParameters["current"] = false
		set := bson.M{
			"$set": setParameters,
		}
		_, err := repo.UpdateOne(filter, set)
		if err != nil {
			trestCommon.ECLog2(
				"update settings section",
				err)
			return "", err
		}
	}
	var settingEntity entity.SettingDB
	settingEntity.CreatedTime = time.Now()
	settingEntity.ID = primitive.NewObjectID()
	settingEntity.UpdatedBy = setting.UpdatedBy
	settingEntity.UpdatedTime = time.Now()
	settingEntity.Drop = setting.Drop
	settingEntity.DeliveryPersonPercentage = setting.DeliveryPersonPercentage
	settingEntity.CutType = setting.CutType
	settingEntity.DeliveryCharge = setting.DeliveryCharge
	settingEntity.Current = true
	re, err := repo.InsertOne(settingEntity)
	return re, err
}

func (*adminService) GetSettingsHistory(token string, limit, skip int) ([]entity.SettingDB, error) {
	return repo.Find(bson.M{}, bson.M{}, limit, skip)
}

func (*adminService) GetCurrentSettings(token string, limit, skip int) ([]entity.SettingDB, error) {
	return repo.Find(bson.M{"current": true}, bson.M{}, limit, skip)
}
