package chat

import (
	entity "Drop/DropChat/entities"

	"go.mongodb.org/mongo-driver/bson"
)

type ChatRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.ChatDB, error)
	Find(filter, projection bson.M, limit, skip int) ([]entity.ChatDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
	DeleteMany(filter bson.M) error
}
