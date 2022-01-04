package payment

import (
	entity "Drop/DropPayments/entities"

	"go.mongodb.org/mongo-driver/bson"
)

type PaymentRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.PaymentEntityDB, error)
	Find(filter, projection bson.M, limit, skip int) ([]entity.PaymentEntityDB, error)
	FindWithIDs(filter, projection bson.M) ([]entity.PaymentEntityDB, error)
	UpdateOne(filter, update bson.M) (string, error)
}
