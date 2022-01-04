package notification

import (
	entity "Drop/Droporder/entities"

	"go.mongodb.org/mongo-driver/bson"
)

type NotificationRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.MessageData, error)
	Find(filter, projection bson.M, limit, skip int) ([]entity.MessageData, error)
}
