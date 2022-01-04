package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type PaymentDB struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserId     string             `bson:"user_id" json:"user,omitempty"`
	CustomerId string             `bson:"customer_id" json:"customer_id,omitempty"`
}
