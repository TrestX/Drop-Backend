package favourite

import (
	entity "Drop/DropFavourite/entities"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.FavouriteDB, error)
	Find(filter, projection bson.M, limit, skip int) ([]string, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
