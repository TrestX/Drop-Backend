package address

import (
	entity "Drop/DropAddress/entities"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.AddressDB, error)
	Find(filter, projection bson.M, limit, skip int) ([]entity.AddressDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
