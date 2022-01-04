package cart

import (
	entity "Drop/DropCart/entities"
	"context"
	"errors"

	"github.com/aekam27/trestCommon"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type repo struct {
	CollectionName string
}

//NewFirestoreRepository creates a new repo
func NewCartRepository(collectionName string) UserRepository {
	return &repo{
		CollectionName: collectionName,
	}
}

//used by signup
func (r *repo) InsertOne(document interface{}) (string, error) {
	_, err := trestCommon.InsertOne(document, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"insert item",
			err,
			logrus.Fields{
				"document":        document,
				"collection name": r.CollectionName,
			})
		return "", err
	}
	return "Added Successfully", nil
}

//used by update cart ,login and email verifcation
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

//used by get cart ,login and email verification
func (r *repo) FindOne(filter, projection bson.M) (entity.CartDB, error) {
	var cart entity.CartDB
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&cart)
	if err != nil {
		trestCommon.ECLog3(
			"Find cart",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return cart, err
	}
	return cart, err
}

//not used may use in future for gettin list of cart
func (r *repo) Find(filter, projection bson.M, limit, skip int) ([]entity.CartDB, error) {
	var carts []entity.CartDB
	cursor, err := trestCommon.FindWithLimitAndOffSet(filter, projection, limit, skip, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find cart",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var cart entity.CartDB
		if err = cursor.Decode(&cart); err != nil {
			trestCommon.ECLog3(
				"Find cart",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return nil, err
		}
		carts = append(carts, cart)
	}
	return carts, nil
}

//not using
func (r *repo) DeleteOne(filter bson.M) error {
	deleteResult, err := trestCommon.DeleteOne(filter, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"delete cart",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	if deleteResult.DeletedCount == 0 {
		err = errors.New("cart not found(404)")
		trestCommon.ECLog3(
			"delete cart",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	return nil
}
func (r *repo) DeleteMany(filter bson.M) error {
	deleteResult, err := trestCommon.DeleteOne(filter, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"delete cart",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	if deleteResult.DeletedCount == 0 {
		err = errors.New("cart not found(404)")
		trestCommon.ECLog3(
			"delete cart",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	return nil
}

func (r *repo) FindWithIDs(filter, projection bson.M) ([]entity.CartDB, error) {
	var carts []entity.CartDB
	cursor, err := trestCommon.Find(filter, projection, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find item",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var cart entity.CartDB
		if err = cursor.Decode(&cart); err != nil {
			trestCommon.ECLog3(
				"Find item",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return carts, nil
		}
		carts = append(carts, cart)
	}
	return carts, nil
}
