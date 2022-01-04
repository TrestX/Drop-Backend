package payment

import (
	entity "Nailzee/NailzeePayments/entities"

	"go.mongodb.org/mongo-driver/bson"
)

type PaymentRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.PaymentDB, error)
	Find(filter, projection bson.M, limit, skip int) ([]entity.PaymentDB, error)
	UpdateOne(filter, update bson.M) (string, error)
}
