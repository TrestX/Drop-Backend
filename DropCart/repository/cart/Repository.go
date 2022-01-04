package cart

import (
	entity "Drop/DropCart/entities"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.CartDB, error)
	Find(filter, projection bson.M, limit, skip int) ([]entity.CartDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	FindWithIDs(filter, projection bson.M) ([]entity.CartDB, error)
	DeleteOne(filter bson.M) error
	DeleteMany(filter bson.M) error
}
