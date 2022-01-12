package ratingreview

import (
	"Drop/DropRatingReview/api"
	entity "Drop/DropRatingReview/entities"
	ratingReview "Drop/DropRatingReview/repository/ratingReview"
	"errors"
	"time"

	"github.com/aekam27/trestCommon"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = ratingReview.NewRatingReviewRepository("ratingreview")
)

type ratingReviewService struct{}

func NewRatingReviewService(repository ratingReview.RatingReviewRepository) RatingReviewService {
	repo = repository
	return &ratingReviewService{}
}

func (*ratingReviewService) AddRatingReview(userId, entityId, token string, ratingReview RatingReviewSchema) (string, error) {
	if userId == "" {
		return "", errors.New("UserId missing")
	}
	if entityId == "" {
		return "", errors.New("EntityId missing")
	}
	if ratingReview.For == "Shop" {
		var shop Shop
		shop.Rating = ratingReview.Rating
		oldReview, _ := repo.Find(bson.M{"entity_id": entityId}, bson.M{}, 10000, 0)
		shop.NumbofRating = int64(len(oldReview))
		_, _ = api.UpdateShopRating(entityId, shop, token)
	}
	ratingReviewEntity := createRatingReviewEntity(ratingReview, userId, entityId)
	return repo.InsertOne(ratingReviewEntity)
}

func (*ratingReviewService) UpdateRatingReview(Id string, ratingReview RatingReviewSchema) (string, error) {
	if Id == "" {
		return "", errors.New("Rating Review Id missing")
	}
	_, err := checkRatingsReview(Id)
	if err != nil {
		return "", errors.New("ratingsreview doesnot exist")
	}
	id, _ := primitive.ObjectIDFromHex(Id)
	set := bson.M{}
	if ratingReview.Rating != 0 {
		set["rating"] = ratingReview.Rating
	}
	if ratingReview.Review != "" {
		set["review"] = ratingReview.Review
	}
	set["updated_time"] = time.Now()
	return repo.UpdateOne(bson.M{"_id": id}, bson.M{"$set": set})
}

func createRatingReviewEntity(ratingReview RatingReviewSchema, userId, entityId string) entity.RatingReviewDB {
	var ratinReviewEntity entity.RatingReviewDB
	ratinReviewEntity.ID = primitive.NewObjectID()
	ratinReviewEntity.UserID = userId
	ratinReviewEntity.EntityID = entityId
	ratinReviewEntity.Rating = ratingReview.Rating
	ratinReviewEntity.Review = ratingReview.Review
	ratinReviewEntity.Deleted = false
	ratinReviewEntity.AddedTime = time.Now()
	return ratinReviewEntity
}

func checkRatingsReview(Id string) (entity.RatingReviewDB, error) {
	id, _ := primitive.ObjectIDFromHex(Id)
	filter := bson.M{"_id": id}
	ratingsReviews, err := repo.FindOne(filter, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get Ratings And Review",
			err,
		)
		return ratingsReviews, err
	}
	return ratingsReviews, nil
}

func (*ratingReviewService) GetRatingsReview(userId, entityId string, limit, skip int) ([]entity.RatingReviewDB, error) {
	uId := ""
	eId := ""
	filter := bson.M{}
	if userId != "" {
		uId = userId
		filter["user_id"] = uId
	}
	if entityId != "" {
		eId = entityId
		filter["entity_id"] = eId
	}

	ratingsReviews, err := repo.Find(filter, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"Get Ratings And Review",
			err,
		)
		return ratingsReviews, err
	}
	return ratingsReviews, nil
}

func (*ratingReviewService) GetReviewRatingsWithIDs(userIds []string) ([]entity.RatingReviewDB, error) {
	subFilter := bson.A{}
	for _, item := range userIds {
		subFilter = append(subFilter, bson.M{"entity_id": item})
	}
	filter := bson.M{"$or": subFilter}
	users, err := repo.FindWithIDs(filter, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get Carts section",
			err,
		)
		return []entity.RatingReviewDB{}, err
	}
	return users, nil
}
