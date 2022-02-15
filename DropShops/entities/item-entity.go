package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Optname struct {
	Name  string  `bson:"name" json:"name"`
	Price float64 `bson:"price" json:"price"`
}
type Choices struct {
	Name    string    `bson:"name" json:"name"`
	Options []Optname `bson:"options" json:"options"`
}

type ItemDB struct {
	ID            primitive.ObjectID `bson:"_id" json:"item"`
	SellerID      string             `bson:"seller_id" json:"seller_id"`
	ShopID        string             `bson:"shop_id" json:"shop_id"`
	Category      string             `bson:"category" json:"category"`
	Name          string             `bson:"name" json:"name"`
	Description   string             `bson:"description" json:"description"`
	Approved      bool               `bson:"approved" json:"approved"`
	Rejected      bool               `bson:"rejected" json:"rejected"`
	ShopType      string             `bson:"shop_type" json:"shop_type"`
	Images        []string           `bson:"images" json:"images"`
	AddOns        []ItemAdOn         `bson:"add_ons" json:"add_ons"`
	Quantity      int64              `bson:"quantity" json:"quantity"`
	Featured      bool               `bson:"featured" json:"featured"`
	FeaturedApp   bool               `bson:"featured_app" json:"featured_app"`
	Price         int64              `bson:"price" json:"price"`
	Type          string             `bson:"type" json:"type"`
	CreatedTime   time.Time          `bson:"created_time" json:"created_time"`
	UpdatedTime   time.Time          `bson:"updated_time" json:"updated_time"`
	Deal          string             `bson:"deal" json:"deal"`
	Sizes         []Optname          `bson:"sizes" json:"sizes"`
	Matrix        string             `bson:"matrix" json:"matrix"`
	Choices       []Choices          `bson:"choices" json:"choices"`
	DiscountPrice float64            `bson:"discount_price" json:"discount_price,omitempty"`
}

type ItemAdOn struct {
	ID    primitive.ObjectID `bson:"_id" json:"item"`
	Name  string             `bson:"name" json:"name"`
	Price int64              `bson:"price" json:"price"`
	Note  string             `bson:"note" json:"note"`
}

type RatingReviewDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	UserID      string             `bson:"user_id" json:"user_id"`
	EntityID    string             `bson:"entity_id" json:"entity_id"`
	Rating      float64            `bson:"rating" json:"rating"`
	Review      string             `bson:"review" json:"review"`
	Deleted     bool               `bson:"deleted" json:"deleted"`
	UpdatedTime time.Time          `bson:"updated_time" json:"updated_time"`
	AddedTime   time.Time          `bson:"added_time" json:"added_time"`
}
