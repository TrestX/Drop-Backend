package tracking

import (
	entity "Drop/DropDeliveryTracking/entities"

	"go.mongodb.org/mongo-driver/bson"
)

type TrackingRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.TrackingDB, error)
	Find(filter, projection bson.M, limit, skip int) ([]entity.TrackingDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
