package cart

import (
	"errors"

	"github.com/aekam27/trestCommon"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type repo struct {
	CollectionName string
}

func NewCartRepository(collectionName string) CartRepository {
	return &repo{
		CollectionName: collectionName,
	}
}

func (r *repo) UpdateOne(filter, update bson.M) (string, error) {
	result, err := trestCommon.UpdateOne(filter, update, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"update cart",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})

		return "", err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		err = errors.New("cart not found(404)")
		trestCommon.ECLog3(
			"update cart",
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
