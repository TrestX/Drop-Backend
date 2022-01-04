package admin

import (
	entity "Drop/DropAdmin/entities"
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
			"insert banner",
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
			"update banner",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})

		return "", err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		err = errors.New("banner not found(404)")
		trestCommon.ECLog3(
			"update banner",
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

func (r *repo) FindOne(filter, projection bson.M) (entity.BannerDB, error) {
	var banner entity.BannerDB
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&banner)
	if err != nil {
		trestCommon.ECLog3(
			"Find banner",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return banner, err
	}
	return banner, err
}

func (r *repo) Find(filter, projection bson.M, limit, skip int) ([]entity.BannerDB, error) {
	var banners []entity.BannerDB
	cursor, err := trestCommon.FindWithLimitAndOffSet(filter, projection, limit, skip, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find banner",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var banner entity.BannerDB
		if err = cursor.Decode(&banner); err != nil {
			trestCommon.ECLog3(
				"Find banner",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return banners, nil
		}
		banners = append(banners, banner)
	}
	return banners, nil
}

func (r *repo) DeleteOne(filter bson.M) error {
	deleteResult, err := trestCommon.DeleteOne(filter, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"delete banner",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	if deleteResult.DeletedCount == 0 {
		err = errors.New("banner not found(404)")
		trestCommon.ECLog3(
			"delete banner",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	return nil
}
