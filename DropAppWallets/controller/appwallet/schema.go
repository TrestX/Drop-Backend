package appwallet

import (
	entity "Drop/DropAppWallets/entities"
)

type AppWalletService interface {
	AddTransaction(transactionDetails AppWalletSchema, token string) (string, error)
	UpdateTrans(transactionId, status string) (string, error)
	GetTransaction(transactionId, status, entity, entityid, orderid string) (entity.AppWallet, error)
	GetTransactions(transactionId, status, entity, entityid, orderid string, limit, skip int) ([]entity.AppWallet, error)
	GetDeliveryPersonBalance(transactionId, status, entityid, orderid, token, fromD, endD, typeD string, limit, skip int) ([]entity.AppWallet, int64, int64, int64, int64, int64, int64, float64, string, error)
	GetSellerPersonBalance(transactionId, status, entityid, orderid, fromD, endD string, limit, skip int) ([]entity.AppWallet, int64, int64, int64, error)
	GetTotalTransactions() (int64, error)
	GetAppEarning() (int64, error)
	GetSellerPersonS(token, fromD, endD string, limit, skip int) ([]TransactionCustomOutput, error)
	GetDeliveryPersonS(token, fromD, endD string, limit, skip int) ([]TransactionCustomOutputD, error)
	UpdateSellerTransPer(sellerId, per, c string) (int, error)
	GetSellerPersonShops(token, sid, fromD, endD string, limit, skip int) ([]TransactionCustomShopOutput, error)

	UpdateSellerPaymentHistory(sid, name, email, phone, doneby, token string, amount int64) (string, error)
	UpdateShopPaymentHistory(shid, name, email, phone, doneby, token string, amount int64) (string, error)
	UpdateDeliveryPaymentHistory(did, name, email, phone, doneby, dtype, token string, amount int64) (string, error)
	GetPAymentsHistory(limit, skip int) ([]entity.SettingPaymentHistoryDB, error)
}

type AppWalletSchema struct {
	OrderId           string `bson:"order_id" json:"order_id,omitempty"`
	PaymentId         string `bson:"payment_id" json:"payment_id,omitempty"`
	SellerID          string `bson:"seller_id" json:"seller_id,omitempty"`
	ShopId            string `bson:"shop_id" json:"shop_id,omitempty"`
	DeliveryID        string `bson:"delivery_id" json:"delivery_id,omitempty"`
	Status            string `bson:"status" json:"status,omitempty"`
	SellerAmount      int64  `bson:"seller_amount" json:"seller_amount,omitempty"`
	DropAmount        int64  `bson:"drop_amount" json:"drop_amount,omitempty"`
	DeliveryCharge    int64  `bson:"delivery_charge" json:"delivery_charge,omitempty"`
	DeliveryPersonCut int64  `bson:"delivery_person_cut" json:"delivery_person_cut,omitempty"`
	TipAmount         int64  `bson:"tip_amount" json:"tip_amount,omitempty"`
}
type SettingPaymentHistory struct {
	ID      string `bson:"id" json:"id"`
	Amount  int64  `bson:"amount" json:"amount,omitempty"`
	Name    string `bson:"name" json:"name,omitempty"`
	Email   string `bson:"email" json:"email,omitempty"`
	PhoneNo string `bson:"phone_no" json:"phone_no,omitempty"`
	DoneBy  string `bson:"done_by" json:"done_by,omitempty"`
	Type    string `bson:"type" json:"type,omitempty"`
}
type TransactionCustomOutput struct {
	EntityID        string            `bson:"entity_id" json:"entity_id"`
	Status          string            `bson:"status" json:"status,omitempty"`
	Email           string            `bson:"email" json:"email,omitempty"`
	ShopName        string            `bson:"shop_name" json:"shop_name"`
	ProfilePhoto    string            `bson:"profile_photo" json:"profile_photo,omitempty"`
	Name            string            `bson:"name" json:"name,omitempty"`
	ShopType        string            `bson:"shop_type" json:"shop_type,omitempty"`
	PhoneNumber     string            `bson:"phone_number" json:"phone_number,omitempty"`
	Type            []entity.ShopType `bson:"type" json:"type,omitempty"`
	DropAmount      int64             `bson:"drop_amount" json:"drop_amount,omitempty"`
	UnSettledAmount int64             `bson:"unsettled_amount" json:"unsettled_amount,omitempty"`
	SettledAmount   int64             `bson:"settled_amount" json:"settled_amount,omitempty"`
	TotalAmount     int64             `bson:"total_amount" json:"total_amount,omitempty"`
}

type TransactionCustomOutputD struct {
	EntityID     string            `bson:"entity_id" json:"entity_id"`
	Status       string            `bson:"status" json:"status,omitempty"`
	Email        string            `bson:"email" json:"email,omitempty"`
	ShopName     string            `bson:"shop_name" json:"shop_name"`
	ProfilePhoto string            `bson:"profile_photo" json:"profile_photo,omitempty"`
	Name         string            `bson:"name" json:"name,omitempty"`
	ShopType     string            `bson:"shop_type" json:"shop_type,omitempty"`
	PhoneNumber  string            `bson:"phone_number" json:"phone_number,omitempty"`
	Type         []entity.ShopType `bson:"type" json:"type,omitempty"`

	CashSettled        int64 `bson:"cash_settled" json:"cash_settled,omitempty"`
	CashUnSettled      int64 `bson:"cash_unsettled" json:"cash_unsettled,omitempty"`
	UnSettledAmount    int64 `bson:"unsettled_amount" json:"unsettled_amount,omitempty"`
	SettledAmount      int64 `bson:"settled_amount" json:"settled_amount,omitempty"`
	TotalAmount        int64 `bson:"total_amount" json:"total_amount,omitempty"`
	UnSettledTipAmount int64 `bson:"unsettled_tip_amount" json:"unsettled_tip_amount,omitempty"`
	SettleTipdAmount   int64 `bson:"settled_tip_amount" json:"settled_tip_amount,omitempty"`
	TotalTipAmount     int64 `bson:"total_tip_amount" json:"total_tip_amount,omitempty"`
}

type TransactionCustomShopOutput struct {
	ShopName        string `bson:"shop_name" json:"shop_name"`
	ProfilePhoto    string `bson:"profile_photo" json:"profile_photo,omitempty"`
	ShopType        string `bson:"shop_type" json:"shop_type,omitempty"`
	UnSettledAmount int64  `bson:"unsettled_amount" json:"unsettled_amount,omitempty"`
	SettledAmount   int64  `bson:"settled_amount" json:"settled_amount,omitempty"`
	TotalAmount     int64  `bson:"total_amount" json:"total_amount,omitempty"`
}
