package address

import (
	entity "Drop/DropAddress/entities"
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
func NewAddressRepository(collectionName string) UserRepository {
	return &repo{
		CollectionName: collectionName,
	}
}

//used by signup
func (r *repo) InsertOne(document interface{}) (string, error) {
	_, err := trestCommon.InsertOne(document, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"insert address",
			err,
			logrus.Fields{
				"document":        document,
				"collection name": r.CollectionName,
			})
		return "", err
	}
	return "Address added successfully", nil
}

//used by update address ,login and email verifcation
func (r *repo) UpdateOne(filter, update bson.M) (string, error) {
	result, err := trestCommon.UpdateOne(filter, update, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"update address",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})

		return "", err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		err = errors.New("address not found(404)")
		trestCommon.ECLog3(
			"update address",
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

//used by get address ,login and email verification
func (r *repo) FindOne(filter, projection bson.M) (entity.AddressDB, error) {
	var address entity.AddressDB
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&address)
	if err != nil {
		trestCommon.ECLog3(
			"Find address",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return address, err
	}
	return address, err
}

//not used may use in future for gettin list of address
func (r *repo) Find(filter, projection bson.M, limit, skip int) ([]entity.AddressDB, error) {
	var addresses []entity.AddressDB
	cursor, err := trestCommon.FindWithLimitAndOffSet(filter, projection, limit, skip, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find address",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var address entity.AddressDB
		if err = cursor.Decode(&address); err != nil {
			trestCommon.ECLog3(
				"Find address",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return addresses, nil
		}
		addresses = append(addresses, address)
	}
	return addresses, nil
}

//not using
func (r *repo) DeleteOne(filter bson.M) error {
	deleteResult, err := trestCommon.DeleteOne(filter, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"delete address",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	if deleteResult.DeletedCount == 0 {
		err = errors.New("address not found(404)")
		trestCommon.ECLog3(
			"delete address",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	return nil
}
