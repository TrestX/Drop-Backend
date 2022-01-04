package coupon

import (
	entity "Drop/DropCoupons/entities"
)

type CouponService interface {
	AddCoupon(coupon Coupon, token string) (string, error)
	UpdateCoupon(coupon Coupon, couponId string) (string, error)
	GetCoupon(code, validAmt, maxDis, usagePD, maxUsage, disPer, status, cId string) (entity.CouponDB, error)
	GetCouponWithIDs(cartId []string) ([]entity.CouponDB, error)
	GetCoupons(code, validAmt, maxDis, usagePD, maxUsage, disPer, status, cId string, limit, skip int) ([]entity.CouponDB, error)
}

type Coupon struct {
	CouponCode         string `bson:"coupon_code" json:"coupon_code,omitempty"`
	Description        string `bson:"description" json:"description,omitempty"`
	ValidAmount        string `bson:"valid_amount" json:"valid_amount,omitempty"`
	MaxDiscount        string `bson:"max_discount" json:"max_discount,omitempty"`
	UsagePerDay        string `bson:"usage_per_day" json:"usage_per_day,omitempty"`
	MaximumUsage       string `bson:"maximum_usage" json:"maximum_usage,omitempty"`
	DiscountPercentage string `bson:"discount_percentage" json:"discount_percentage,omitempty"`
	Status             string `bson:"status" json:"status,omitempty"`
}
