package ratingReview

import (
	entity "Drop/DropRatingReview/entities"

	"go.mongodb.org/mongo-driver/bson"
)

type RatingReviewRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.RatingReviewDB, error)
	Find(filter, projection bson.M, limit, skip int) ([]entity.RatingReviewDB, error)
	UpdateOne(filter, update bson.M) (string, error)
}
