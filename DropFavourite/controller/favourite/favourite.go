package favourite

import (
	entity "Drop/DropFavourite/entities"

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
	if err != nil {
		trestCommon.ECLog2(
			"GetFavourite section",
			err,
		)
		return favourite, err
	}
	return favourite, nil
}
