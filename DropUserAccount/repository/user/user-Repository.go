package user

import (
	entity "Drop/DropUserAccount/entities"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.UserDB, error)
	Find(filter, projection bson.M) ([]entity.UserDB, error)
	FindWithIDs(filter, projection bson.M) ([]entity.UserDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
