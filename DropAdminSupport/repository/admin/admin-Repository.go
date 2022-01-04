package admin

import (
	entity "Drop/DropAdminSupport/entities"

	"go.mongodb.org/mongo-driver/bson"
)

type AdminRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.AdminSupportDB, error)
	Find(filter, projection bson.M, limit, skip int) ([]entity.AdminSupportDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
