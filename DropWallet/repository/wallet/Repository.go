package wallet

import (
	entity "Drop/DropWallet/entities"

	"go.mongodb.org/mongo-driver/bson"
)

type WalletRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.WalletDB, error)
	Find(filter, projection bson.M, limit, skip int) ([]entity.WalletDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	FindWithIDs(filter, projection bson.M) ([]entity.WalletDB, error)
	DeleteOne(filter bson.M) error
	DeleteMany(filter bson.M) error
}
