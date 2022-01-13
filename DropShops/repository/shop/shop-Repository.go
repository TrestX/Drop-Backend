package shop

import (
	entity "Drop/DropShop/entities"

	"go.mongodb.org/mongo-driver/bson"
)

type ShopRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.ShopDB, error)
	Find(filter, projection bson.M, limit, skip int) ([]entity.ShopDB, error)
	FindSort(filter, projection, filter1 bson.M, limit, skip int) ([]entity.ShopDB, error)
	FindWithIDs(filter, projection bson.M) ([]entity.ShopDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
