package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FavouriteDB struct {
	ID        primitive.ObjectID `bson:"_id" json:"favourite_id,omitempty"`
	UserID    string             `bson:"user_id" json:"user_id,omitempty"`
	ItemID    string             `bson:"item_id" json:"item_id,omitempty"`
	AddedTime time.Time          `bson:"added_time" json:"added_time,omitempty"`
}
type ItemDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"item,omitempty"`
	UserID      string             `bson:"user_id" json:"user_id,omitempty"`
	SellerID    string             `bson:"seller_id" json:"seller_id,omitempty"`
	Category    string             `bson:"category" json:"category,omitempty"`
	SubCategory string             `bson:"sub_category" json:"sub_category,omitempty"`
	Title       string             `bson:"title" json:"title,omitempty"`
	SubTitle    string             `bson:"sub_title" json:"sub_title,omitempty"`
	Description string             `bson:"description" json:"description,omitempty"`
	Approved    bool               `bson:"approved" json:"approved,omitempty"`
	Rejected    bool               `bson:"rejected" json:"rejected,omitempty"`
	Images      []string           `bson:"images" json:"images,omitempty"`
	CreatedTime time.Time          `bson:"created_time" json:"created_time,omitempty"`
	UpdatedTime time.Time          `bson:"updated_time" json:"updated_time,omitempty"`
}
