package item

import (
	entity "Drop/DropItems/entities"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.ItemDB, error)
	Find(filter, projection bson.M, limit, skip int) ([]entity.ItemDB, error)
	FindWithIDs(filter, projection bson.M) ([]entity.ItemDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
