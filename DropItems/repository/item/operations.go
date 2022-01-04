package item

import (
	entity "Drop/DropItems/entities"

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
func NewItemRepository(collectionName string) UserRepository {
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

//used by update item ,login and email verifcation
func (r *repo) UpdateOne(filter, update bson.M) (string, error) {
	result, err := trestCommon.UpdateOne(filter, update, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"update item",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})

		return "", err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		err = errors.New("item not found(404)")
		trestCommon.ECLog3(
			"update item",
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

//used by get item ,login and email verification
func (r *repo) FindOne(filter, projection bson.M) (entity.ItemDB, error) {
	var item entity.ItemDB
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&item)
	if err != nil {
		trestCommon.ECLog3(
			"Find item",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return item, err
	}
	return item, err
}

//not used may use in future for gettin list of item
func (r *repo) Find(filter, projection bson.M, limit, skip int) ([]entity.ItemDB, error) {
	var itemes []entity.ItemDB
	cursor, err := trestCommon.FindWithLimitAndOffSet(filter, projection, limit, skip, r.CollectionName)
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
		var item entity.ItemDB
		if err = cursor.Decode(&item); err != nil {
			trestCommon.ECLog3(
				"Find item",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return itemes, nil
		}
		itemes = append(itemes, item)
	}
	return itemes, nil
}

//not used may use in future for gettin list of item
func (r *repo) FindWithIDs(filter, projection bson.M) ([]entity.ItemDB, error) {
	var itemes []entity.ItemDB
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
		var item entity.ItemDB
		if err = cursor.Decode(&item); err != nil {
			trestCommon.ECLog3(
				"Find item",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return itemes, nil
		}
		itemes = append(itemes, item)
	}
	return itemes, nil
}

//not using
func (r *repo) DeleteOne(filter bson.M) error {
	deleteResult, err := trestCommon.DeleteOne(filter, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"delete item",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	if deleteResult.DeletedCount == 0 {
		err = errors.New("item not found(404)")
		trestCommon.ECLog3(
			"delete item",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	return nil
}
