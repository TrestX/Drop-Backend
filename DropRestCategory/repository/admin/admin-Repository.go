package admin

import (
	entity "Drop/DropItemCategories/entities"

	"go.mongodb.org/mongo-driver/bson"
)

type AdminRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.CategoryDB, error)
	Find(filter, projection bson.M, limit, skip int) ([]entity.CategoryDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
