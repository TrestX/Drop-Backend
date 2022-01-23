package payment

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	entity "Drop/DropPayments/entities"

)

type PaymentService interface {
	CreatePaymentIntent(userId, token string, payments PaymentIntentSchema) (string, string, error)
	UpdatePaymentStatus(userId, paymentId, status string) (string, error)
	UpdatePaymentStatusSuccess(paymentId string) (string, error)
	GetPaymentsDetails(userId, status string, limit, skip int) ([]entity.PaymentEntityDB, error)
	GetPaymentDetails(userId, paymentId string) (entity.PaymentEntityDB, error)
	GetPaymentWithIDs(paymentIds []string) ([]entity.PaymentEntityDB, error)
	GetAdminPaymentDetails(status, user, token, seller, shop string, limit, skip int) ([]PaymentOP, error)
}

type PaymentIntentSchema struct {
	Amount             string `json:"amount"`
	Currency           string `json:"currency"`
	Description        string `json:"description"`
	PaymentMethodTypes string `json:"payment_method_types"`
	CustomerEmail      string `json:"customer_email"`
	RedirectUrl        string `json:"redirect_url"`
	SellerID           string `bson:"seller_id" json:"seller_id"`
	ShopID             string `bson:"shop_id" json:"shop_id"`
	CouponCode         string `bson:"coupon_code" json:"coupon_code"`
	UserAddressId      string `json:"shipping"`
	Type               string `bson:"type" json:"type"`
}

type ShippingDetails struct {
	Address        string `json:"address"`
	Name           string `json:"name"`
	Phone          string `json:"phone"`
	TrackingNumber string `json:"trackingNumber"`
}

type PaymentOP struct {
	ID                 primitive.ObjectID     `bson:"_id" json:"payment_id"`
	UserID             string                 `bson:"user_id" json:"user_id"`
	SellerID           string                 `bson:"seller_id" json:"seller_id"`
	ShopID             string                 `bson:"shop_id" json:"shop_id"`
	Amount             int64                  `bson:"amount" json:"amount"`
	Currency           string                 `bson:"currency" json:"currency"`
	PaymentMethodTypes string                 `bson:"payment_method_types" json:"payment_method"`
	Description        string                 `bson:"description" json:"description"`
	ReceiptEmail       string                 `bson:"receipt_email" json:"receipt_email"`
	Shipping           entity.ShippingDetails `bson:"shipping_details" json:"shipping_details"`
	Status             string                 `bson:"status" json:"status"`
	CouponCode         string                 `bson:"coupon_code" json:"coupon_code"`
	UpdatedTime        time.Time              `bson:"updated_time" json:"updated_time"`
	AddedTime          time.Time              `bson:"added_time" json:"added_time"`
	SellerDetails      ShopDetails            `bson:"seller_details" json:"seller_details"`
}
type ShopDetails struct {
	ShopName      string `bson:"shop_name" json:"shop_name"`
	Deal          string `bson:"deal" json:"deal"`
	SellerName    string `bson:"seller_name" json:"seller_name"`
	AccountNumber string `bson:"account_number" json:"account_number"`
	IFSC          string `bson:"ifsc" json:"ifsc"`
	Address       string `bson:"address" json:"address"`
}
