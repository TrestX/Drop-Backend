package payment

import (
	entity "Nailzee/NailzeePayments/entities"
	"context"
	"errors"

	"github.com/aekam27/trestCommon"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type repo struct {
	CollectionName string
}

func NewPaymentRepository(collectionName string) PaymentRepository {
	return &repo{
		CollectionName: collectionName,
	}
}

func (r *repo) InsertOne(document interface{}) (string, error) {
	_, err := trestCommon.InsertOne(document, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"insert payment details",
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
			"update payment status",
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
			"update payments status",
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

func (r *repo) FindOne(filter, projection bson.M) (entity.PaymentDB, error) {
	var payment entity.PaymentDB
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&payment)
	if err != nil {
		trestCommon.ECLog3(
			"Find payment details",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return payment, err
	}
	return payment, err
}

func (r *repo) Find(filter, projection bson.M, limit, skip int) ([]entity.PaymentDB, error) {
	var payments []entity.PaymentDB
	cursor, err := trestCommon.FindWithLimitAndOffSet(filter, projection, limit, skip, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find payment details",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var payment entity.PaymentDB
		if err = cursor.Decode(&payment); err != nil {
			trestCommon.ECLog3(
				"Find payments details",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return nil, err
		}
		payments = append(payments, payment)
	}
	return payments, nil
}
