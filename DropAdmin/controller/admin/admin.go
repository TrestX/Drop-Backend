package admin

import (
	"errors"
	"strings"
	"time"

	"github.com/aekam27/trestCommon"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"Drop/DropAdmin/api"
	entity "Drop/DropAdmin/entities"
	"Drop/DropAdmin/repository/admin"
)

var (
	repo = admin.NewAdminRepository("banner")
)

type adminService struct{}

func NewAdminService(repository admin.AdminRepository) AdminService {
	repo = repository
	return &adminService{}
}
func (add *adminService) AddBanner(banners *Banners) (string, error) {
	var bannerEntity entity.BannerDB
	bannerEntity.Status = banners.Status
	bannerEntity.CreatedTime = time.Now()
	bannerEntity.ID = primitive.NewObjectID()
	bannerEntity.Name = banners.Name
	bannerEntity.Description = banners.Description
	bannerEntity.Text = banners.Text
	bannerEntity.ButtonText = banners.ButtonText
	bannerEntity.Dimensions = banners.Dimensions
	bannerEntity.Category = banners.Category
	bannerEntity.ShopName = banners.ShopName
	url := ""
	if banners.PresignedUrl != "" {
		url = createPreSignedDownloadUrl(banners.PresignedUrl)
	}
	bannerEntity.PresignedDownloadUrl = url
	return repo.InsertOne(bannerEntity)
}
func createPreSignedDownloadUrl(url string) string {
	s := strings.Split(url, "?")
	if len(s) > 0 {
		o := strings.Split(s[0], "/")
		if len(o) > 3 {
			fileName := o[4]
			path := o[3]
			downUrl, _ := trestCommon.PreSignedDownloadUrl(fileName, path)
			return downUrl
		}
	}
	return ""
}
func (*adminService) UpdateBannerStatus(banners *Banners, bannerid string) (string, error) {
	if bannerid == "" {
		err := errors.New("banner id missing")
		trestCommon.ECLog2(
			"update banner location",
			err,
		)
		return "", err
	}
	id, _ := primitive.ObjectIDFromHex(bannerid)
	_, err := checkByBannerID(id)
	if err != nil {
		return "", errors.New("invalid banner Id")
	}
	setParameters := bson.M{}
	if banners.Status != "" {
		setParameters["status"] = banners.Status
	}
	if banners.Name != "" {
		setParameters["name"] = banners.Name
	}
	if banners.Description != "" {
		setParameters["description"] = banners.Description
	}
	if banners.Text != "" {
		setParameters["text"] = banners.Text
	}
	if banners.ButtonText != "" {
		setParameters["button_text"] = banners.ButtonText
	}
	if banners.ShopName != "" {
		setParameters["shop_name"] = banners.ShopName
	}
	if banners.PresignedUrl != "" {
		setParameters["presignedurl"] = createPreSignedDownloadUrl(banners.PresignedUrl)
	}
	setParameters["updated_time"] = time.Now()
	filter := bson.M{"_id": id}
	set := bson.M{
		"$set": setParameters,
	}
	result, err := repo.UpdateOne(filter, set)
	if err != nil {
		trestCommon.ECLog3(
			"update banner location",
			err,
			logrus.Fields{
				"banner_id": bannerid,
			})
		return "", err
	}

	return result, nil
}

func checkByBannerID(id primitive.ObjectID) (entity.BannerDB, error) {
	banner, err := repo.FindOne(bson.M{"_id": id}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get banner Details section",
			err,
		)
		return banner, err
	}
	return banner, nil
}
func getUserDetails(token string) error {
	user, err := api.GetUserDetails(token)
	if err != nil {
		return err
	}
	if strings.ToLower(user.AccountType) != "admin" {
		return errors.New("user doesnot have admin privilages")
	}
	return nil
}
func (*adminService) GetAllBanners(token string, limit, skip int) ([]entity.BannerDB, error) {
	err := getUserDetails(token)
	if err != nil {
		return []entity.BannerDB{}, err
	}
	banners, err := repo.Find(bson.M{}, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"Get banners section",
			err,
		)
		return banners, err
	}
	for i := 0; i < len(banners); i++ {
		newdownloadurl := createPreSignedDownloadUrl(banners[i].PresignedDownloadUrl)
		banners[i].PresignedDownloadUrl = newdownloadurl
	}
	return banners, nil
}

func (*adminService) GetActivebanners(limit, skip int, bannerType string) ([]entity.BannerDB, error) {
	filter := bson.M{"status": "Active"}
	if bannerType != "" {
		filter["category"] = bannerType
	}
	banners, err := repo.Find(filter, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"Get banners Details section",
			err,
		)
		return banners, err
	}
	for i := 0; i < len(banners); i++ {
		newdownloadurl := createPreSignedDownloadUrl(banners[i].PresignedDownloadUrl)
		banners[i].PresignedDownloadUrl = newdownloadurl
	}
	return banners, nil
}
