package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingReviewDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID      string             `bson:"user_id" json:"user_id,omitempty"`
	EntityID    string             `bson:"entity_id" json:"entity_id,omitempty"`
	Rating      float64            `bson:"rating" json:"rating,omitempty"`
	Review      string             `bson:"review" json:"review,omitempty"`
	Deleted     bool               `bson:"deleted" json:"deleted,omitempty"`
	UpdatedTime time.Time          `bson:"updated_time" json:"updated_time,omitempty"`
	AddedTime   time.Time          `bson:"added_time" json:"added_time,omitempty"`
}
