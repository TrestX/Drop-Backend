package favourite

import (
	"Drop/DropFavourite/api"
	entity "Drop/DropFavourite/entities"
	"strings"

	"Drop/DropFavourite/repository/favourite"
	"errors"
	"time"

	"github.com/aekam27/trestCommon"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = favourite.NewFavouriteRepository("favourite")
)

type favouriteService struct{}

func NewFavouriteService(repository favourite.UserRepository) FavouriteService {
	repo = repository
	return &favouriteService{}
}

func (*favouriteService) AddFavourite(itemID, userId string) (string, error) {
	var favouriteEntity entity.FavouriteDB
	if userId == "" || itemID == "" {
		return "", errors.New("Something went wrong")
	}
	favouriteEntity.ID = primitive.NewObjectID()
	favouriteEntity.UserID = userId
	favouriteEntity.ItemID = itemID
	favouriteEntity.AddedTime = time.Now()
	return repo.InsertOne(favouriteEntity)
}

func (*favouriteService) DeleteFavourite(itemID, userId string) error {
	if userId == "" || itemID == "" {
		return errors.New("Something went wrong")
	}
	return repo.DeleteOne(bson.M{"user_id": userId, "item_id": itemID})
}
func (*favouriteService) GetFavourite(userId string, limit, skip int) ([]entity.ItemDB, error) {
	favourite, err := repo.Find(bson.M{"user_id": userId}, bson.M{}, limit, skip)
	items, err := api.GetItems(favourite)
	for i := 0; i < len(items); i++ {
		url := createPreSignedDownloadUrl(items[i].Images[0])
		items[i].Images[0] = url
	}
	if err != nil {
		trestCommon.ECLog2(
			"GetFavourite section",
			err,
		)
		return items, err
	}
	return items, nil
}

func (*favouriteService) GetFavouriteShop(userId string, limit, skip int) ([]entity.ShopDB, error) {
	favourite, err := repo.Find(bson.M{"user_id": userId}, bson.M{}, limit, skip)
	shop, err := api.GetShopDetailsByIDs(favourite)
	for i := 0; i < len(shop); i++ {
		url := createPreSignedDownloadUrl(shop[i].ShopLogo)
		url1 := createPreSignedDownloadUrl(shop[i].ShopBanner)
		shop[i].ShopLogo = url
		shop[i].ShopBanner = url1
	}
	if err != nil {
		trestCommon.ECLog2(
			"GetFavourite section",
			err,
		)
		return shop, err
	}
	return shop, nil
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
