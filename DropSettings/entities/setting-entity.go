package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
