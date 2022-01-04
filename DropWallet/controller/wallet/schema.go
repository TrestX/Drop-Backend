package wallet

import (
	entity "Drop/DropWallet/entities"
)

type WalletService interface {
	AddWalletTransaction(wallet Wallet, token string) (string, error)
	GetWallet(userId, status string) (entity.WalletDB, error)
	GetWalletWithIDs(cartId []string) ([]entity.WalletDB, error)
	GetWalletByUserId(userId string) (entity.WalletDB, error)
}

type Wallet struct {
	UserID      string `bson:"user_id" json:"user_id,omitempty"`
	Amount      string `bson:"amount" json:"amount,omitempty"`
	Description string `bson:"description" json:"description,omitempty"`
	Type        string `bson:"type" json:"type,omitempty"`
	Status      string `bson:"status" json:"status,omitempty"`
}
