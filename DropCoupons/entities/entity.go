package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CouponDB struct {
	ID                 primitive.ObjectID `bson:"_id" json:"coupon_id,omitempty"`
	CouponCode         string             `bson:"coupon_code" json:"coupon_code,omitempty"`
	Description        string             `bson:"description" json:"description,omitempty"`
	ValidAmount        string             `bson:"valid_amount" json:"valid_amount,omitempty"`
	MaxDiscount        string             `bson:"max_discount" json:"max_discount,omitempty"`
	UsagePerDay        string             `bson:"usage_per_day" json:"usage_per_day,omitempty"`
	MaximumUsage       string             `bson:"maximum_usage" json:"maximum_usage,omitempty"`
	DiscountPercentage string             `bson:"discount_percentage" json:"discount_percentage,omitempty"`
	UpdatedTime        time.Time          `bson:"updated_time" json:"updated_time,omitempty"`
	CreatedTime        time.Time          `bson:"created_time" json:"created_time,omitempty"`
	Status             string             `bson:"status" json:"status,omitempty"`
}
