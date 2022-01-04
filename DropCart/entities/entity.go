package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"cart_id,omitempty"`
	UserID      string             `bson:"user_id" json:"user_id,omitempty"`
	ShopID      string             `bon:"shop_id" json:"shop_id,omitempty"`
	Items       []Item             `bson:"items" json:"items,omitempty"`
	AddOns      []AddOns           `bson:"addons" json:"addons"`
	UpdatedTime time.Time          `bson:"updated_time" json:"updated_time,omitempty"`
	AddedTime   time.Time          `bson:"added_time" json:"added_time,omitempty"`
	Status      string             `bson:"status" json:"status,omitempty"`
}
type Item struct {
	ItemID          string            `bson:"item_id" json:"item_id,omitempty"`
	ItemName        string            `json:"item_name,omitempty"`
	ItemDescription string            `json:"item_description,omitempty"`
	Status          string            `bson:"status" json:"status,omitempty"`
	Quantity        int64             `bson:"quantity" json:"quantity,omitempty"`
	Size            Optname           `bson:"size" json:"size"`
	ChoiceSelected  []ChoicesSelected `bson:"choiceSelected" json:"choiceSelected"`
}
type AddOns struct {
	ItemID string `bson:"item_id" json:"item_id"`
	Name   string `bson:"name" json:"name"`
	Price  int64  `bson:"price" json:"price"`
	Note   string `bson:"note" json:"note"`
}
type Choices struct {
	Name    string    `bson:"name" json:"name"`
	Options []Optname `bson:"options" json:"options"`
}

type ChoicesSelected struct {
	Name    string  `bson:"name" json:"name"`
	Options Optname `bson:"options" json:"options"`
}

type ItemDB struct {
	ID             primitive.ObjectID `bson:"_id" json:"item"`
	SellerID       string             `bson:"seller_id" json:"seller_id"`
	ShopID         string             `bson:"shop_id" json:"shop_id"`
	Category       string             `bson:"category" json:"category"`
	Name           string             `bson:"name" json:"name"`
	Description    string             `bson:"description" json:"description"`
	Approved       bool               `bson:"approved" json:"approved"`
	Rejected       bool               `bson:"rejected" json:"rejected"`
	ShopType       string             `bson:"shop_type" json:"shop_type"`
	Images         []string           `bson:"images" json:"images"`
	AddOns         []AddOns           `bson:"add_ons" json:"add_ons"`
	Quantity       int64              `bson:"quantity" json:"quantity"`
	Featured       bool               `bson:"featured" json:"featured"`
	FeaturedApp    bool               `bson:"featured_app" json:"featured_app"`
	Price          int64              `bson:"price" json:"price"`
	Type           string             `bson:"type" json:"type"`
	CreatedTime    time.Time          `bson:"created_time" json:"created_time"`
	UpdatedTime    time.Time          `bson:"updated_time" json:"updated_time"`
	Deal           string             `bson:"deal" json:"deal"`
	Sizes          []Optname          `bson:"sizes" json:"sizes"`
	Matrix         string             `bson:"matrix" json:"matrix"`
	Choices        []Choices          `bson:"choices" json:"choices"`
	Size           Optname            `bson:"size" json:"size"`
	ChoiceSelected []ChoicesSelected  `bson:"choiceSelected" json:"choiceSelected"`
}
type Optname struct {
	Name  string  `bson:"name" json:"name"`
	Price float64 `bson:"price" json:"price"`
}
type ItemAdOn struct {
	Name  string `bson:"name" json:"name,omitempty"`
	Price int64  `bson:"price" json:"price,omitempty"`
	Note  string `bson:"note" json:"note,omitempty"`
}
