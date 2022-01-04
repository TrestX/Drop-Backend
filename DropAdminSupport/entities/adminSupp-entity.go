package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AdminSupportDB struct {
	ID        primitive.ObjectID `bson:"_id" json:"user_id,omitempty"`
	Email     string             `bson:"email,omitempty" json:"email,omitempty"`
	Status    string             `bson:"status,omitempty" json:"status,omitempty"`
	Name      string             `bson:"name" json:"name,omitempty"`
	Role      string             `bson:"role" json:"role,omitempty"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Type      string             `bson:"type" json:"type,omitempty"`
	CreatedOn time.Time          `bson:"created_on" json:"created_on,omitempty"`
}
