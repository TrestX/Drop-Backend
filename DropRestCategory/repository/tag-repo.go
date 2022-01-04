package tag

import (
	entity "Drop/DropItemCategories/entities"

	"go.mongodb.org/mongo-driver/bson"
)

type TagRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.TagDB, error)
	Find(filter, projection bson.M, limit, skip int) ([]entity.TagDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
