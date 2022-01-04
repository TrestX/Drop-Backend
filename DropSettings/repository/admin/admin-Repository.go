package admin

import (
	entity "Drop/DropSettings/entities"

	"go.mongodb.org/mongo-driver/bson"
)

type AdminSettingRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.SettingDB, error)
	Find(filter, projection bson.M, limit, skip int) ([]entity.SettingDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
