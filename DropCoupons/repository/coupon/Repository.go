package coupon

import (
	entity "Drop/DropCoupons/entities"

	"go.mongodb.org/mongo-driver/bson"
)

type CouponRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.CouponDB, error)
	Find(filter, projection bson.M, limit, skip int) ([]entity.CouponDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	FindWithIDs(filter, projection bson.M) ([]entity.CouponDB, error)
	DeleteOne(filter bson.M) error
	DeleteMany(filter bson.M) error
}
