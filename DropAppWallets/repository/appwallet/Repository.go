package appwallet

import (
	entity "Drop/DropAppWallets/entities"

	"go.mongodb.org/mongo-driver/bson"
)

type AppWalletRepository interface {
	InsertOne(document interface{}) (string, error)
	InsertPHistory(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.AppWallet, error)
	Find(filter, projection bson.M, limit, skip int) ([]entity.AppWallet, error)
	FindPHistory(filter, projection bson.M, limit, skip int) ([]entity.SettingPaymentHistoryDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	FindWithIDs(filter, projection bson.M) ([]entity.AppWallet, error)
	DeleteOne(filter bson.M) error
	DeleteMany(filter bson.M) error
}
