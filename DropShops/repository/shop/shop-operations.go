package shop

import (
	"context"
	"errors"
	"strings"

	"github.com/aekam27/trestCommon"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"

	entity "Drop/DropShop/entities"
)

type repo struct {
	CollectionName string
}

func NewShopRepository(collectionName string) ShopRepository {
	return &repo{
		CollectionName: collectionName,
	}
}
func (r *repo) FindOneSetting(filter, projection bson.M) (entity.SettingDB, error) {
	var setting entity.SettingDB
	err := trestCommon.FindOne(filter, projection, "setting").Decode(&setting)
	if err != nil {
		trestCommon.ECLog3(
			"Find setting",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return setting, err
	}
	return setting, err
}

func (r *repo) InsertOne(document interface{}) (string, error) {
	_, err := trestCommon.InsertOne(document, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"insert shop",
			err,
			logrus.Fields{
				"document":        document,
				"collection name": r.CollectionName,
			})
		return "", err
	}
	return "Shop added successfully", nil
}

func (r *repo) UpdateOne(filter, update bson.M) (string, error) {
	result, err := trestCommon.UpdateOne(filter, update, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"update shop",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})

		return "", err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		err = errors.New("shop not found(404)")
		trestCommon.ECLog3(
			"update shop",
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

func (r *repo) FindOne(filter, projection bson.M) (entity.ShopDB, error) {
	var shop entity.ShopDB
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&shop)
	if err != nil {
		trestCommon.ECLog3(
			"Find shop",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return shop, err
	}
	return shop, err
}

func (r *repo) Find(filter, projection bson.M, limit, skip int) ([]entity.ShopDB, error) {
	var shops []entity.ShopDB
	cursor, err := trestCommon.FindWithLimitAndOffSet(filter, projection, limit, skip, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find shop",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var shop entity.ShopDB
		if err = cursor.Decode(&shop); err != nil {
			trestCommon.ECLog3(
				"Find shops",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return shops, nil
		}
		shops = append(shops, shop)
	}
	return shops, nil
}

func (r *repo) FindSort(filter, projection, filter1 bson.M, limit, skip int) ([]entity.ShopDB, error) {
	var shops []entity.ShopDB
	cursor, err := trestCommon.FindSort(filter, projection, filter1, limit, skip, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find shop",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var shop entity.ShopDB
		if err = cursor.Decode(&shop); err != nil {
			trestCommon.ECLog3(
				"Find shops",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return shops, nil
		}
		shops = append(shops, shop)
	}
	return shops, nil
}

func (r *repo) DeleteOne(filter bson.M) error {
	deleteResult, err := trestCommon.DeleteOne(filter, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"delete shops",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	if deleteResult.DeletedCount == 0 {
		err = errors.New("shop not found(404)")
		trestCommon.ECLog3(
			"delete shop",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	return nil
}

func (r *repo) FindWithIDs(filter, projection bson.M) ([]entity.ShopDB, error) {
	var users []entity.ShopDB
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
		var user entity.ShopDB
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
func (r *repo) FindUsingAggregaye(aggreateStruct bson.A) ([]entity.OutPutCategorySchema, error) {
	var shops []entity.OutPutCategorySchema
	cursor, err := trestCommon.Aggregate(aggreateStruct, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find shop",
			err,
			logrus.Fields{
				"filter":          aggreateStruct,
				"collection name": r.CollectionName,
			})
		return shops, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var shop entity.ShopAggreate
		var shopStruc entity.OutPutCategorySchema
		if err = cursor.Decode(&shop); err != nil {
			trestCommon.ECLog3(
				"Find shops",
				err,
				logrus.Fields{
					"filter":          aggreateStruct,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return shops, nil
		}
		categoryS := ""
		for _, item := range shop.Item {
			if categoryS != "" {
				categoryS += ","
			}
			for i := range item.Images {
				newdownloadurl := createPreSignedDownloadUrl(item.Images[i])
				item.Images[i] = newdownloadurl
			}
			if strings.Contains(categoryS, item.Category) {
				categoryS += item.Category
			}
		}
		newdownloadurl := createPreSignedDownloadUrl(shop.ID[0].ShopBanner)
		shop.ID[0].ShopBanner = newdownloadurl
		newdownloadurl = createPreSignedDownloadUrl(shop.ID[0].ShopLogo)
		shop.ID[0].ShopLogo = newdownloadurl
		for i := range shop.ID[0].ShopPhotos {
			newdownloadurl = createPreSignedDownloadUrl(shop.ID[0].ShopPhotos[i])
			shop.ID[0].ShopPhotos[i] = newdownloadurl
		}
		shopStruc = entity.OutPutCategorySchema{shop.ID[0], "", shop.Item}

		shops = append(shops, shopStruc)
	}
	return shops, nil

}
func createPreSignedDownloadUrl(url string) string {
	s := strings.Split(url, "?")
	if len(s) > 0 {
		o := strings.Split(s[0], "/")
		if len(o) > 3 {
			fileName := o[4]
			path := o[3]
			downUrl, _ := trestCommon.PreSignedDownloadUrl(fileName, path)
			return downUrl
		}
	}
	return ""
}
