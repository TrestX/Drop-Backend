package order

import (
	entity "Drop/Droporder/entities"

	"go.mongodb.org/mongo-driver/bson"
)

type OrderRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.OrderDB, error)
	Find(filter, projection bson.M, limit, skip int) ([]entity.OrderDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	FindWithIDs(filter, projection bson.M) ([]entity.OrderDB, error)
}
