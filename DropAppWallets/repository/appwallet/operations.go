package appwallet

import (
	entity "Drop/DropAppWallets/entities"
	"context"
	"errors"

	"github.com/aekam27/trestCommon"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type repo struct {
	CollectionName string
}

func NewAppWalletRepository(collectionName string) AppWalletRepository {
	return &repo{
		CollectionName: collectionName,
	}
}
func (r *repo) InsertOne(document interface{}) (string, error) {
	_, err := trestCommon.InsertOne(document, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"insert transaction",
			err,
			logrus.Fields{
				"document":        document,
				"collection name": r.CollectionName,
			})
		return "", err
	}
	return "Added Successfully", nil
}
func (r *repo) InsertPHistory(document interface{}) (string, error) {
	_, err := trestCommon.InsertOne(document, "paymentsHistory")
	if err != nil {
		trestCommon.ECLog3(
			"insert transaction",
			err,
			logrus.Fields{
				"document":        document,
				"collection name": "paymentsHistory",
			})
		return "", err
	}
	return "Added Successfully", nil
}
func (r *repo) FindPHistory(filter, projection bson.M, limit, skip int) ([]entity.SettingPaymentHistoryDB, error) {
	var transactions []entity.SettingPaymentHistoryDB
	cursor, err := trestCommon.FindWithLimitAndOffSet(filter, projection, limit, skip, "paymentsHistory")
	if err != nil {
		trestCommon.ECLog3(
			"Find paymentsHistory",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": "paymentsHistory",
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var transaction entity.SettingPaymentHistoryDB
		if err = cursor.Decode(&transaction); err != nil {
			trestCommon.ECLog3(
				"Find transaction",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": "paymentsHistory",
					"error at":        cursor.RemainingBatchLength(),
				})
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (r *repo) UpdateOne(filter, update bson.M) (string, error) {
	result, err := trestCommon.UpdateOne(filter, update, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"update transaction",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})

		return "", err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		err = errors.New("transaction not found(404)")
		trestCommon.ECLog3(
			"update transaction",
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

func (r *repo) FindOne(filter, projection bson.M) (entity.AppWallet, error) {
	var transaction entity.AppWallet
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&transaction)
	if err != nil {
		trestCommon.ECLog3(
			"Find transaction",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return transaction, err
	}
	return transaction, err
}
func (r *repo) Find(filter, projection bson.M, limit, skip int) ([]entity.AppWallet, error) {
	var transactions []entity.AppWallet
	cursor, err := trestCommon.FindWithLimitAndOffSet(filter, projection, limit, skip, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find transaction",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var transaction entity.AppWallet
		if err = cursor.Decode(&transaction); err != nil {
			trestCommon.ECLog3(
				"Find transaction",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

//not using
func (r *repo) DeleteOne(filter bson.M) error {
	deleteResult, err := trestCommon.DeleteOne(filter, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"delete transaction",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	if deleteResult.DeletedCount == 0 {
		err = errors.New("transaction not found(404)")
		trestCommon.ECLog3(
			"delete transaction",
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
			"delete transaction",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	if deleteResult.DeletedCount == 0 {
		err = errors.New("transaction not found(404)")
		trestCommon.ECLog3(
			"delete transaction",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	return nil
}

func (r *repo) FindWithIDs(filter, projection bson.M) ([]entity.AppWallet, error) {
	var transactions []entity.AppWallet
	cursor, err := trestCommon.Find(filter, projection, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find transaction",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var transaction entity.AppWallet
		if err = cursor.Decode(&transaction); err != nil {
			trestCommon.ECLog3(
				"Find transaction",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return transactions, nil
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
