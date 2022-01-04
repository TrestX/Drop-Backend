package admin

import (
	entity "Drop/DropAdminSupport/entities"
	"context"
	"errors"

	"github.com/aekam27/trestCommon"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type repo struct {
	CollectionName string
}

func NewAdminRepository(collectionName string) AdminRepository {
	return &repo{
		CollectionName: collectionName,
	}
}
func (r *repo) InsertOne(document interface{}) (string, error) {
	_, err := trestCommon.InsertOne(document, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"insert acct",
			err,
			logrus.Fields{
				"document":        document,
				"collection name": r.CollectionName,
			})
		return "", err
	}
	return "added successfully", nil
}

func (r *repo) UpdateOne(filter, update bson.M) (string, error) {
	result, err := trestCommon.UpdateOne(filter, update, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"update acct",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})

		return "", err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		err = errors.New("acct not found(404)")
		trestCommon.ECLog3(
			"update acct",
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

func (r *repo) FindOne(filter, projection bson.M) (entity.AdminSupportDB, error) {
	var adminsupp entity.AdminSupportDB
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&adminsupp)
	if err != nil {
		trestCommon.ECLog3(
			"Find acct",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return adminsupp, err
	}
	return adminsupp, err
}

func (r *repo) Find(filter, projection bson.M, limit, skip int) ([]entity.AdminSupportDB, error) {
	var adminsupps []entity.AdminSupportDB
	cursor, err := trestCommon.FindWithLimitAndOffSet(filter, projection, limit, skip, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find acct",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var adminsupp entity.AdminSupportDB
		if err = cursor.Decode(&adminsupp); err != nil {
			trestCommon.ECLog3(
				"Find accts",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return adminsupps, nil
		}
		adminsupps = append(adminsupps, adminsupp)
	}
	return adminsupps, nil
}

func (r *repo) DeleteOne(filter bson.M) error {
	deleteResult, err := trestCommon.DeleteOne(filter, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"delete acct",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	if deleteResult.DeletedCount == 0 {
		err = errors.New("acct not found(404)")
		trestCommon.ECLog3(
			"delete acct",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	return nil
}
