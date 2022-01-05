package cart

import (
	"go.mongodb.org/mongo-driver/bson"
)

type CartRepository interface {
	UpdateOne(filter, update bson.M) (string, error)
}
