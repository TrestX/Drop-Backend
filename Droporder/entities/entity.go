package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShopDB struct {
	ID              primitive.ObjectID `bson:"_id" json:"shop_id"`
	SellerID        string             `bson:"seller_id" json:"seller_id"`
	Address         string             `bson:"address" json:"address"`
	Country         string             `bson:"country" json:"country"`
	State           string             `bson:"state" json:"state"`
	City            string             `bson:"city" json:"city"`
	Pin             string             `bson:"pin" json:"pin"`
	Primary         bool               `bson:"primary" json:"primary"`
	Type            string             `bson:"type" json:"type"`
	Timing          string             `bson:"timing" json:"timing"`
	ShopName        string             `bson:"shop_name" json:"shop_name"`
	ShopLogo        string             `bson:"shop_logo" json:"shop_logo"`
	ShopBanner      string             `bson:"shop_banner" json:"shop_banner"`
	ShopPhotos      []string           `bson:"shop_photos" json:"shop_photos"`
	ShopStatus      string             `bson:"shop_status" json:"shop_status"`
	Featured        bool               `bson:"featured" json:"featured"`
	ShopDescription string             `bson:"shop_description" json:"shop_description"`
	GeoLocation     bson.M             `bson:"geo_location" json:"geo_location"`
	CreatedTime     time.Time          `bson:"created_time" json:"created_time"`
	UpdatedTime     time.Time          `bson:"updated_time" json:"updated_time"`
}
type OrderDB struct {
	ID                primitive.ObjectID `bson:"_id" json:"order_id"`
	DeliveryDetails   DeliveryDB         `bson:"delivery_details" json:"delivery_details"`
	PaymentID         string             `bson:"payment_id" json:"payment_id"`
	UserID            string             `bson:"user_id" json:"user_id"`
	SellerID          string             `bson:"seller_id" json:"seller_id"`
	ShopID            string             `bson:"shop_id" json:"shop_id"`
	CartID            string             `bson:"cart_id" json:"cart_id"`
	Status            string             `bson:"status" json:"status"`
	OrderPlacedTime   time.Time          `bson:"order_placed_time" json:"order_placed_time"`
	OrderAcceptedTime time.Time          `bson:"order_accepted_time" json:"order_accepted_time"`
	DeliveredTime     time.Time          `bson:"delivered_time" json:"delivered_time"`
	OrderPickUpTime   time.Time          `bson:"order_pickup_time" json:"order_pickup_time"`
	UpdatedTime       time.Time          `bson:"updated_time" json:"updated_time"`
	AddedTime         time.Time          `bson:"added_time" json:"added_time"`
	TipAmount         int64              `bson:"tip_amount" json:"tip_amount"`
	DeliveryCode      int64              `bson:"delivery_code" json:"delivery_code"`
}

type UserDB struct {
	ID                primitive.ObjectID `bson:"_id" json:"user_id"`
	Email             string             `bson:"email" json:"email"`
	Status            string             `bson:"status" json:"status"`
	Name              string             `bson:"name" json:"name"`
	Designation       string             `bson:"designation" json:"designation"`
	PhoneNo           string             `bson:"phone_number" json:"phone_number"`
	AccountType       string             `bson:"account_type" json:"account_type"`
	AlternateEmail    string             `bson:"alternate_email" json:"alternate_email"`
	TermsChecked      bool               `bson:"terms_checked" json:"terms_checked"`
	Password          string             `bson:"password" json:"password"`
	EmailSentTime     time.Time          `bson:"email_sent_time" json:"email_sent_time"`
	VerificationCode  string             `bson:"verification_code" json:"verification_code"`
	PasswordResetCode string             `bson:"password_reset_code" json:"password_reset_code"`
	PasswordResetTime time.Time          `bson:"password_reset_time" json:"password_reset_time"`
	LoggedInUsing     string             `bson:"logged_in_using" json:"logged_in_using"`
	Theme             string             `bson:"theme" json:"theme"`
	Language          string             `bson:"language" json:"language"`
	ProfilePhoto      string             `bson:"profile_photo" json:"profile_photo"`
	Gender            string             `bson:"gender" json:"gender"`
	DOB               string             `bson:"dob" json:"dob"`
	ApprovalStatus    string             `bson:"approval_status" json:"approval"`
	Availability      string             `bson:"availability" json:"availability"`
	Deleted           bool               `bson:"deleted" json:"deleted"`
	NationalID        string             `bson:"national_id" json:"national_id"`
	PictureID         string             `bson:"picture_id" json:"picture_id"`
	VehicleNumber     string             `bson:"vehicle_number" json:"vehicle_number"`
	VehicleType       string             `bson:"vehicle_type" json:"vehicle_type"`
	Wallet            string             `bson:"wallet" json:"wallet"`
}

type DeliveryDB struct {
	TrackingID  string    `bson:"tracking_id" json:"tracking_id"`
	DeliveryID  string    `bson:"delivery_id" json:"delivery_id"`
	UserAddress AddressDB `bson:"user_address" json:"user_address"`
	ShopAddress AddressDB `bson:"shop_address" json:"shop_address"`
	UpdatedTime time.Time `bson:"updated_time" json:"updated_time"`
	AddedTime   time.Time `bson:"added_time" json:"added_time"`
}

type AddressDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"address_id"`
	UserID      string             `bson:"user_id" json:"user_id"`
	Address     string             `bson:"address" json:"address"`
	Country     string             `bson:"country" json:"country"`
	State       string             `bson:"state" json:"state"`
	City        string             `bson:"city" json:"city"`
	Pin         string             `bson:"pin" json:"pin"`
	Primary     bool               `bson:"primary" json:"primary"`
	Type        string             `bson:"type" json:"type"`
	GeoLocation bson.M             `bson:"geo_location" json:"geo_location"`
	CreatedTime time.Time          `bson:"created_time" json:"created_time"`
	UpdatedTime time.Time          `bson:"updated_time" json:"updated_time"`
}

type DeliveryPersonDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"delivery_person_id"`
	Name        string             `bson:"name" json:"name"`
	Age         string             `bson:"age" json:"age"`
	Ratings     string             `bson:"ratings" json:"ratings"`
	UpdatedTime time.Time          `bson:"updated_time" json:"updated_time"`
	AddedTime   time.Time          `bson:"added_time" json:"added_time"`
}

type CartDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"cart_id"`
	UserID      string             `bson:"user_id" json:"user_id"`
	ShopID      string             `bon:"shop_id" json:"shop_id"`
	Items       []ItemDB           `bson:"items" json:"items"`
	UpdatedTime time.Time          `bson:"updated_time" json:"updated_time"`
	AddedTime   time.Time          `bson:"added_time" json:"added_time"`
	Status      string             `bson:"status" json:"status"`
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
	Images         []string           `bson:"images" json:"images"`
	AddOns         []ItemAdOn         `bson:"add_ons" json:"add_ons"`
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
type ItemAdOn struct {
	ID    primitive.ObjectID `bson:"_id" json:"item"`
	Name  string             `bson:"name" json:"name"`
	Price int64              `bson:"price" json:"price"`
	Note  string             `bson:"note" json:"note"`
}

type PaymentDB struct {
	ID                 primitive.ObjectID `bson:"_id" json:"payment_id"`
	UserID             string             `bson:"user_id" json:"user_id"`
	SellerID           string             `bson:"seller_id" json:"seller_id"`
	ShopID             string             `bson:"shop_id" json:"shop_id"`
	Amount             int64              `bson:"amount" json:"amount"`
	Currency           string             `bson:"currency" json:"currency"`
	PaymentMethodTypes string             `bson:"payment_method_types" json:"payment_method"`
	Description        string             `bson:"description" json:"description"`
	ReceiptEmail       string             `bson:"receipt_email" json:"receipt_email"`
	Shipping           ShippingDetails    `bson:"shipping_details" json:"shipping_details"`
	Status             string             `bson:"status" json:"status"`
	CouponCode         string             `bson:"coupon_code" json:"coupon_code"`
	UpdatedTime        time.Time          `bson:"updated_time" json:"updated_time"`
	AddedTime          time.Time          `bson:"added_time" json:"added_time"`
}

type ShippingDetails struct {
	Address      AddressDB `bson:"address" json:"address"`
	Name         string    `bson:"name" json:"name"`
	Phone        string    `bson:"phone" json:"phone"`
	Email        string    `bson:"email" json:"email"`
	ProfilePhoto string    `bson:"profilePhoto" json:"profilePhoto"`
}

type MessageData struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Title      string             `bson:"title" json:"title"`
	Body       string             `bson:"body" json:"body"`
	Topic      string             `bson:"topic" json:"topic"`
	UserId     string             `bson:"userId" json:"userId"`
	CategoryId string             `bson:"categoryId" json:"categoryId"`
	SentTime   time.Time          `bson:"sentTime" json:"sentTime"`
	Status     string             `bson:"status" json:"status"`
	NStatus    string             `bson:"nStatus" json:"nStatus"`
}

type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}
