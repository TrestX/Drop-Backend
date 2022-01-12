package ratingReview

import (
	entity "Drop/DropRatingReview/entities"
	"context"
	"errors"

	"github.com/aekam27/trestCommon"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type repo struct {
	CollectionName string
}

func NewRatingReviewRepository(collectionName string) RatingReviewRepository {
	return &repo{
		CollectionName: collectionName,
	}
}

func (r *repo) InsertOne(document interface{}) (string, error) {
	_, err := trestCommon.InsertOne(document, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"insert ratng and review details",
			err,
			logrus.Fields{
				"document":        document,
				"collection name": r.CollectionName,
			})
		return "", err
	}
	return "Added Successfully", nil
}

func (r *repo) UpdateOne(filter, update bson.M) (string, error) {
	result, err := trestCommon.UpdateOne(filter, update, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"update ratings and review",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})

		return "", err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		err = errors.New("record not found(404)")
		trestCommon.ECLog3(
			"update ratings and review",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})
		return "", err
	}
	return "updated successfully", nil
}

func (r *repo) FindOne(filter, projection bson.M) (entity.RatingReviewDB, error) {
	var payment entity.RatingReviewDB
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&payment)
	if err != nil {
		trestCommon.ECLog3(
			"Find a rating",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return payment, err
	}
	return payment, err
}

func (r *repo) Find(filter, projection bson.M, limit, skip int) ([]entity.RatingReviewDB, error) {
	var ratingsReviews []entity.RatingReviewDB
	cursor, err := trestCommon.FindWithLimitAndOffSet(filter, projection, limit, skip, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find ratings and Reviews details",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var ratingReview entity.RatingReviewDB
		if err = cursor.Decode(&ratingReview); err != nil {
			trestCommon.ECLog3(
				"Find ratings and Reviews details",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return nil, err
		}
		ratingsReviews = append(ratingsReviews, ratingReview)
	}
	return ratingsReviews, nil
}

func (r *repo) FindWithIDs(filter, projection bson.M) ([]entity.RatingReviewDB, error) {
	var users []entity.RatingReviewDB
	cursor, err := trestCommon.Find(filter, projection, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find users",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var user entity.RatingReviewDB
		if err = cursor.Decode(&user); err != nil {
			trestCommon.ECLog3(
				"Find users",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return users, nil
		}
		users = append(users, user)
	}
	return users, nil
}
