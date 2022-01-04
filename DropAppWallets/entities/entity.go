package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AppWallet struct {
	ID                primitive.ObjectID `bson:"_id" json:"settelment_id,omitempty"`
	OrderId           string             `bson:"order_id" json:"order_id,omitempty"`
	ShopId            string             `bson:"shop_id" json:"shop_id,omitempty"`
	OrderAmount       int64              `bson:"order_amount" json:"order_amount,omitempty"`
	SellerAmount      int64              `bson:"seller_amount" json:"seller_amount,omitempty"`
	DropAmount        int64              `bson:"drop_amount" json:"drop_amount,omitempty"`
	SellerID          string             `bson:"seller_id" json:"seller_id,omitempty"`
	DeliveryID        string             `bson:"delivery_id" json:"delivery_id,omitempty"`
	DeliveryCharge    int64              `bson:"delivery_charge" json:"delivery_charge,omitempty"`
	DeliveryPersonCut int64              `bson:"delivery_person_cut" json:"delivery_person_cut,omitempty"`
	Status            string             `bson:"status" json:"status,omitempty"`
	SettledTime       time.Time          `bson:"settled_time" json:"settled_time,omitempty"`
	UpdatedTime       time.Time          `bson:"updated_time" json:"updated_time,omitempty"`
	AddedTime         time.Time          `bson:"added_time" json:"added_time,omitempty"`
	Type              string             `bson:"type" json:"type,omitempty"`
	TipAmount         int64              `bson:"tip_amount" json:"tip_amount,omitempty"`
}

type ShopDB struct {
	ID              primitive.ObjectID `bson:"_id" json:"shop_id,omitempty"`
	SellerID        string             `bson:"seller_id" json:"seller_id,omitempty"`
	Address         string             `bson:"address" json:"address,omitempty"`
	Country         string             `bson:"country" json:"country,omitempty"`
	State           string             `bson:"state,omitempty" json:"state,omitempty"`
	City            string             `bson:"city,omitempty" json:"city,omitempty"`
	Pin             string             `bson:"pin" json:"pin,omitempty"`
	Primary         bool               `bson:"primary" json:"primary,omitempty"`
	Type            string             `bson:"type" json:"type,omitempty"`
	Timing          string             `bson:"timing" json:"timing,omitempty"`
	ShopName        string             `bson:"shop_name" json:"shop_name,omitempty"`
	ShopLogo        string             `bson:"shop_logo" json:"shop_logo,omitempty"`
	ShopBanner      string             `bson:"shop_banner" json:"shop_banner,omitempty"`
	ShopPhotos      []string           `bson:"shop_photos" json:"shop_photos,omitempty"`
	ShopStatus      string             `bson:"shop_status" json:"shop_status,omitempty"`
	Featured        bool               `bson:"featured" json:"featured,omitempty"`
	ShopDescription string             `bson:"shop_description" json:"shop_description,omitempty"`
	GeoLocation     bson.M             `bson:"geo_location" json:"geo_location,omitempty"`
	CreatedTime     time.Time          `bson:"created_time" json:"created_time,omitempty"`
	UpdatedTime     time.Time          `bson:"updated_time" json:"updated_time,omitempty"`
	Deal            string             `bson:"deal" json:"deal,omitempty"`
	DeliveryType    string             `bson:"delivery" json:"delivery,omitempty"`
	Cuisine         string             `bson:"cuisine" json:"cuisine,omitempty"`
}
type OrderDB struct {
	ID        primitive.ObjectID `bson:"_id" json:"order_id,omitempty"`
	PaymentID string             `bson:"payment_id" json:"payment_id"`
	ShopID    string             `bson:"shop_id" json:"shop_id"`
}
type PaymentEntityDB struct {
	ID                 primitive.ObjectID `bson:"_id" json:"payment_id,omitempty"`
	Amount             int64              `bson:"amount" json:"amount,omitempty"`
	Currency           string             `bson:"currency" json:"currency,omitempty"`
	CouponCode         string             `bson:"coupon_code" json:"coupon_code,omitempty"`
	PaymentMethodTypes string             `json:"payment_method,omitempty"`
}

type SettingPaymentHistoryDB struct {
	ID      primitive.ObjectID `bson:"_id" json:"payment_id,omitempty"`
	Amount  int64              `bson:"amount" json:"amount,omitempty"`
	Name    string             `bson:"name" json:"name,omitempty"`
	Email   string             `bson:"email" json:"email,omitempty"`
	PhoneNo string             `bson:"phone_no" json:"phone_no,omitempty"`
	DoneBy  string             `bson:"done_by" json:"done_by,omitempty"`
	DoneAt  time.Time          `bson:"done_at" json:"done_at,omitempty"`
	Type    string             `bson:"type" json:"type,omitempty"`
}

type UserDB struct {
	ID           primitive.ObjectID `bson:"_id" json:"user_id,omitempty"`
	Email        string             `bson:"email,omitempty" json:"email,omitempty"`
	Name         string             `bson:"name" json:"name,omitempty"`
	PhoneNo      string             `bson:"phone_number" json:"phone_number,omitempty"`
	ProfilePhoto string             `bson:"profile_photo" json:"profile_photo,omitempty"`
	Type         []ShopType         `bson:"type" json:"type,omitempty"`
}

type ShopType struct {
	Name string `bson:"name" json:"name,omitempty"`
}

type SettingDB struct {
	ID                       primitive.ObjectID `bson:"_id" json:"setting_id"`
	Drop                     string             `bson:"drop" json:"drop"`
	CutType                  string             `bson:"cut_type" json:"cut_type"`
	DeliveryCharge           []interface{}      `bson:"delivery_charge" json:"delivery_charge"`
	DeliveryPersonPercentage string             `bson:"delivery_person_percentage" json:"delivery_person_percentage"`
	UpdatedBy                string             `bson:"updated_by" json:"updated_by"`
	Current                  bool               `bson:"current" json:"current"`
	CreatedTime              time.Time          `bson:"created_time" json:"created_time"`
	UpdatedTime              time.Time          `bson:"updated_time" json:"updated_time"`
}
