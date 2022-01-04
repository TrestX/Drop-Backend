package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderDB struct {
	ID              primitive.ObjectID `bson:"_id" json:"order_id"`
	DeliveryDetails DeliveryDB         `bson:"delivery_details" json:"delivery_details"`
	PaymentID       string             `bson:"payment_id" json:"payment_id"`
	PaymentsDetails PaymentDB          `bson:"payments_details" json:"payments_details"`
	UserID          string             `bson:"user_id" json:"user_id"`
	ShopID          string             `bson:"shop_id" json:"shop_id"`
	CartID          string             `bson:"cart_id" json:"cart_id"`
	CartDetails     CartDB             `bson:"cart_details" json:"cart_details"`
	Status          string             `bson:"status" json:"status"`
	UpdatedTime     time.Time          `bson:"updated_time" json:"updated_time"`
	AddedTime       time.Time          `bson:"added_time" json:"added_time"`
}

type DeliveryDB struct {
	DeliveryPersonID      string           `bson:"delivery_person_id" json:"delivery_person_id"`
	DeliveryPersonDetails DeliveryPersonDB `bson:"delivery_person_details" json:"delivery_person_details"`
	UserAddress           AddressDB        `bson:"user_address" json:"user_address"`
	ShopAddress           AddressDB        `bson:"shop_address" json:"shop_address"`
	UpdatedTime           time.Time        `bson:"updated_time" json:"updated_time"`
	AddedTime             time.Time        `bson:"added_time" json:"added_time"`
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
	GeoLocation bson.M             `bson:"geo_location" json:"geo_location"`
	CreatedTime time.Time          `bson:"created_time" json:"created_time"`
	UpdatedTime time.Time          `bson:"updated_time" json:"updated_time"`
}

type UserDB struct {
	ID                         primitive.ObjectID `bson:"_id" json:"user_id"`
	Email                      string             `bson:"email" json:"email"`
	Status                     string             `bson:"status" json:"status"`
	Name                       string             `bson:"name" json:"name"`
	Designation                string             `bson:"designation" json:"designation"`
	PhoneNo                    string             `bson:"phone_number" json:"phone_number"`
	AccountType                string             `bson:"account_type" json:"account_type"`
	AlternateEmail             string             `bson:"alternate_email" json:"alternate_email"`
	TermsChecked               bool               `bson:"terms_checked" json:"terms_checked"`
	Password                   string             `bson:"password" json:"password"`
	EmailSentTime              time.Time          `bson:"email_sent_time" json:"email_sent_time"`
	VerificationCode           string             `bson:"verification_code" json:"verification_code"`
	PasswordResetCode          string             `bson:"password_reset_code" json:"password_reset_code"`
	PasswordResetTime          time.Time          `bson:"password_reset_time" json:"password_reset_time"`
	LoggedInUsing              string             `bson:"logged_in_using" json:"logged_in_using"`
	Theme                      string             `bson:"theme" json:"theme"`
	Language                   string             `bson:"language" json:"language"`
	ProfilePhoto               string             `bson:"profile_photo" json:"profile_photo"`
	Type                       []ShopType         `bson:"type" json:"type"`
	Gender                     string             `bson:"gender" json:"gender"`
	DOB                        string             `bson:"dob" json:"dob"`
	CertificateOfIncorporation string             `bson:"certificate_of_incorporation" json:"certificate_of_incorporation"`
	Cuisine                    []string           `bson:"cuisine" json:"cuisine"`
	DeliveryType               string             `bson:"delivery_type" json:"delivery_type"`
	ApprovalStatus             string             `bson:"approval_status" json:"approval"`
	Availability               string             `bson:"availability" json:"availability"`
	Deleted                    bool               `bson:"deleted" json:"deleted"`
	NationalID                 string             `bson:"national_id" json:"national_id"`
	PictureID                  string             `bson:"picture_id" json:"picture_id"`
	VehiclePhoto               string             `bson:"vehicle_photo" json:"vehicle_photo"`
	VehicleRegistration        string             `bson:"vehicle_registration_document" json:"vehicle_registration_document"`
	VehicleNumber              string             `bson:"vehicle_number" json:"vehicle_number"`
	VehicleType                string             `bson:"vehicle_type" json:"vehicle_type"`
	Wallet                     string             `bson:"wallet" json:"wallet"`
	BankName                   string             `bson:"bankName" json:"bankName"`
	AccountNumber              string             `bson:"account_number" json:"account_number"`
	IFSC                       string             `bson:"ifsc" json:"ifsc"`
}
type ShopType struct {
	Name string `bson:"name" json:"name"`
}

type PaymentEntityDB struct {
	ID                 primitive.ObjectID `bson:"_id" json:"payment_id"`
	UserID             string             `bson:"user_id" json:"user_id"`
	SellerID           string             `bson:"seller_id" json:"seller_id"`
	ShopID             string             `bson:"shop_id" json:"shop_id"`
	Amount             int64              `bson:"amount" json:"amount"`
	Currency           string             `bson:"currency" json:"currency"`
	PaymentMethodTypes string             `bson:"payment_method_types" json:"payment_method"`
	Description        string             `bson:"description" json:"description"`
	CustomerEmail      string             `bson:"customer_email" json:"customer_email"`
	RedirectUrl        string             `bson:"redirect_url" json:"redirect_url"`
	Shipping           ShippingDetails    `bson:"shipping_details" json:"shipping_details"`
	Status             string             `bson:"status" json:"status"`
	CouponCode         string             `bson:"coupon_code" json:"coupon_code"`
	UpdatedTime        time.Time          `bson:"updated_time" json:"updated_time"`
	AddedTime          time.Time          `bson:"added_time" json:"added_time"`
}

type PaymentIntent struct {
	Tx_Ref         string        `bson:"tx_ref" json:"tx_ref"`
	Amount         string        `bson:"amount" json:"amount"`
	Currency       string        `bson:"currency" json:"currency"`
	Redirect_Url   string        `bson:"redirect_url" json:"redirect_url"`
	PaymentMethod  string        `bson:"payment_options" json:"payment_options"`
	Customer       Customerz     `bson:"customer" json:"customer"`
	Customizations Customization `bson:"customizations" json:"customizations"`
}
type Customerz struct {
	Email       string `bson:"email" json:"email"`
	PhoneNumber string `bson:"phonenumber" json:"phonenumber"`
	Name        string `bson:"name" json:"name"`
}
type Customization struct {
	Title       string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	Logo        string `bson:"logo" json:"logo"`
}

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
	Deal            string             `bson:"deal" json:"deal"`
	DeliveryType    string             `bson:"delivery" json:"delivery"`
	Cuisine         string             `bson:"cuisine" json:"cuisine"`
}
type ShippingDetails struct {
	Address      AddressDB `bson:"address" json:"address"`
	Name         string    `bson:"name" json:"name"`
	Phone        string    `bson:"phone" json:"phone"`
	Email        string    `bson:"email" json:"email"`
	ProfilePhoto string    `bson:"profilePhoto" json:"profilePhoto"`
}

type DeliveryPersonDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"delivery_person_id"`
	Name        string             `bson:"name" json:"name"`
	Age         string             `bson:"age" json:"age"`
	Ratings     string             `bson:"ratings" json:"ratings"`
	UpdatedTime time.Time          `bson:"updated_time" json:"updated_time"`
	AddedTime   time.Time          `bson:"added_time" json:"added_time"`
}

type PaymentDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"payment_id"`
	UserID      string             `bson:"user_id" json:"user_id"`
	ShopID      string             `bson:"shop_id" json:"shop_id"`
	OrderID     string             `bson:"order_id" json:"order_id"`
	UpdatedTime time.Time          `bson:"updated_time" json:"updated_time"`
	AddedTime   time.Time          `bson:"added_time" json:"added_time"`
	Status      string             `bson:"status" json:"status"`
}

type CartDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"cart_id"`
	UserID      string             `bson:"user_id" json:"user_id"`
	ShopID      string             `bon:"shop_id" json:"shop_id"`
	Items       []Item             `bson:"items" json:"items"`
	UpdatedTime time.Time          `bson:"updated_time" json:"updated_time"`
	AddedTime   time.Time          `bson:"added_time" json:"added_time"`
	Status      string             `bson:"status" json:"status"`
}

type Item struct {
	ItemID   string `bson:"item_id" json:"item_id"`
	Status   string `bson:"status" json:"status"`
	Quantity int64  `bson:"quantity" json:"quantity"`
}

type ItemDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"item"`
	UserID      string             `bson:"user_id" json:"user_id"`
	SellerID    string             `bson:"seller_id" json:"seller_id"`
	Category    string             `bson:"category" json:"category"`
	SubCategory string             `bson:"sub_category" json:"sub_category"`
	Title       string             `bson:"title" json:"title"`
	SubTitle    string             `bson:"sub_title" json:"sub_title"`
	Description string             `bson:"description" json:"description"`
	Approved    bool               `bson:"approved" json:"approved"`
	Rejected    bool               `bson:"rejected" json:"rejected"`
	Images      []string           `bson:"images" json:"images"`
	CreatedTime time.Time          `bson:"created_time" json:"created_time"`
	UpdatedTime time.Time          `bson:"updated_time" json:"updated_time"`
}
