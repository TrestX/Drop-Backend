package notification

import (
	entity "Drop/Droporder/entities"
	"context"
	"errors"

	"github.com/aekam27/trestCommon"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type repo struct {
	CollectionName string
}

func NewNotificationRepository(collectionName string) NotificationRepository {
	return &repo{
		CollectionName: collectionName,
	}
}

func (r *repo) InsertOne(document interface{}) (string, error) {
	_, err := trestCommon.InsertOne(document, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"insert order",
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
			"update order",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})

		return "", err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		err = errors.New("order not found(404)")
		trestCommon.ECLog3(
			"update order",
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

func (r *repo) FindOne(filter, projection bson.M) (entity.MessageData, error) {
	var order entity.MessageData
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&order)
	if err != nil {
		trestCommon.ECLog3(
			"Find order",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return order, err
	}
	return order, err
}

func (r *repo) Find(filter, projection bson.M, limit, skip int) ([]entity.MessageData, error) {
	var orders []entity.MessageData
	cursor, err := trestCommon.FindSort(filter, projection, bson.M{"_id": -1}, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find orders",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var order entity.MessageData
		if err = cursor.Decode(&order); err != nil {
			trestCommon.ECLog3(
				"Find order",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
