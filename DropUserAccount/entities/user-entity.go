package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDB struct {
	ID                         primitive.ObjectID `bson:"_id" json:"user_id,omitempty"`
	Email                      string             `bson:"email,omitempty" json:"email,omitempty"`
	Status                     string             `bson:"status,omitempty" json:"status,omitempty"`
	Name                       string             `bson:"name" json:"name,omitempty"`
	Designation                string             `bson:"designation" json:"designation,omitempty"`
	PhoneNo                    string             `bson:"phone_number" json:"phone_number,omitempty"`
	AccountType                string             `bson:"account_type" json:"account_type,omitempty"`
	AlternateEmail             string             `bson:"alternate_email,omitempty" json:"alternate_email,omitempty"`
	TermsChecked               bool               `bson:"terms_checked" json:"terms_checked,omitempty"`
	Password                   string             `bson:"password" json:"password,omitempty"`
	EmailSentTime              time.Time          `bson:"email_sent_time,omitempty" json:"email_sent_time,omitempty"`
	VerificationCode           string             `bson:"verification_code" json:"verification_code,omitempty"`
	PasswordResetCode          string             `bson:"password_reset_code" json:"password_reset_code,omitempty"`
	PasswordResetTime          time.Time          `bson:"password_reset_time,omitempty" json:"password_reset_time,omitempty"`
	LoggedInUsing              string             `bson:"logged_in_using" json:"logged_in_using,omitempty"`
	Theme                      string             `bson:"theme" json:"theme,omitempty"`
	Language                   string             `bson:"language" json:"language,omitempty"`
	ProfilePhoto               string             `bson:"profile_photo" json:"profile_photo,omitempty"`
	Type                       []ShopType         `bson:"type" json:"type,omitempty"`
	Gender                     string             `bson:"gender" json:"gender,omitempty"`
	DOB                        string             `bson:"dob" json:"dob,omitempty"`
	CertificateOfIncorporation string             `bson:"certificate_of_incorporation" json:"certificate_of_incorporation,omitempty"`
	Cuisine                    []string           `bson:"cuisine" json:"cuisine,omitempty"`
	DeliveryType               string             `bson:"delivery_type" json:"delivery_type,omitempty"`
	ApprovalStatus             string             `bson:"approval_status" json:"approval,omitempty"`
	Availability               string             `bson:"availability" json:"availability,omitempty"`
	Deleted                    bool               `bson:"deleted" json:"deleted,omitempty"`
	NationalID                 string             `bson:"national_id" json:"national_id,omitempty"`
	PictureID                  string             `bson:"picture_id" json:"picture_id,omitempty"`
	VehiclePhoto               string             `bson:"vehicle_photo" json:"vehicle_photo,omitempty"`
	VehicleRegistration        string             `bson:"vehicle_registration_document" json:"vehicle_registration_document,omitempty"`
	VehicleNumber              string             `bson:"vehicle_number" json:"vehicle_number,omitempty"`
	VehicleType                string             `bson:"vehicle_type" json:"vehicle_type,omitempty"`
	Wallet                     string             `bson:"wallet" json:"wallet,omitempty"`
	BankName                   string             `bson:"bankName" json:"bankName,omitempty"`
	AccountNumber              string             `bson:"account_number" json:"account_number,omitempty"`
	IFSC                       string             `bson:"ifsc" json:"ifsc,omitempty"`
	Tags                       string             `bson:"tags" json:"tags"`
	MinOrderAmount             int64              `bson:"minorderamount" json:"minorderamount"`
}
type ShopType struct {
	Name string `bson:"name" json:"name,omitempty"`
}
type Address struct {
	ID          primitive.ObjectID `bson:"_id" json:"address_id,omitempty"`
	UserID      string             `bson:"user_id" json:"user_id,omitempty"`
	Address     string             `bson:"address" json:"address,omitempty"`
	Country     string             `bson:"country" json:"country,omitempty"`
	State       string             `bson:"state,omitempty" json:"state,omitempty"`
	City        string             `bson:"city,omitempty" json:"city,omitempty"`
	Pin         string             `bson:"pin" json:"pin,omitempty"`
	Primary     bool               `bson:"primary" json:"primary,omitempty"`
	Type        string             `bson:"type" json:"type,omitempty"`
	GeoLocation bson.M             `bson:"geo_location" json:"geo_location,omitempty"`
	CreatedTime time.Time          `bson:"created_time" json:"created_time,omitempty"`
	UpdatedTime time.Time          `bson:"updated_time" json:"updated_time,omitempty"`
}

type Seller struct {
	ID                primitive.ObjectID `bson:"_id" json:"user_id,omitempty"`
	Email             string             `bson:"email,omitempty" json:"email,omitempty"`
	Status            string             `bson:"status,omitempty" json:"status,omitempty"`
	Name              string             `bson:"name" json:"name,omitempty"`
	Designation       string             `bson:"designation" json:"designation,omitempty"`
	PhoneNo           string             `bson:"phone_no" json:"phone_no,omitempty"`
	AccountType       string             `bson:"account_type" json:"account_type,omitempty"`
	AlternateEmail    string             `bson:"alternate_email,omitempty" json:"alternate_email,omitempty"`
	TermsChecked      bool               `bson:"terms_checked" json:"terms_checked,omitempty"`
	Password          string             `bson:"password" json:"password,omitempty"`
	EmailSentTime     time.Time          `bson:"email_sent_time,omitempty" json:"email_sent_time,omitempty"`
	VerificationCode  string             `bson:"verification_code" json:"verification_code,omitempty"`
	PasswordResetCode string             `bson:"password_reset_code" json:"password_reset_code,omitempty"`
	PasswordResetTime time.Time          `bson:"password_reset_time,omitempty" json:"password_reset_time,omitempty"`
	LoggedInUsing     string             `bson:"logged_in_using" json:"logged_in_using,omitempty"`
	Theme             string             `bson:"theme" json:"theme,omitempty"`
	Language          string             `bson:"language" json:"language,omitempty"`
	ProfilePhoto      string             `bson:"profile_photo" json:"profile_photo,omitempty"`
	Gender            string             `bson:"gender" json:"gender,omitempty"`
	DOB               string             `bson:"dob" json:"dob,omitempty"`
	ApprovalStatus    string             `bson:"approval_status" json:"approval,omitempty"`
	Deleted           bool               `bson:"deleted" json:"deleted,omitempty"`
	NationalID        string             `bson:"national_id" json:"national_id,omitempty"`
	PictureID         string             `bson:"picture_id" json:"picture_id,omitempty"`
}

type OrderOutput struct {
	ID               string          `json:"order_Id"`
	OrderStatus      string          `json:"orderStatus"`
	OrderedAt        time.Time       `json:"orderedAt"`
	ItemsOrdered     Item            `json:"itemsOrdered"`
	PaymentAmount    int64           `json:"paymentAmount"`
	PaymentCurrency  string          `json:"paymentCurrency"`
	OrderDescription string          `json:"orderDescription"`
	DeliveryDetails  ShippingDetails `json:"deliveryDetails"`
}

type Item struct {
	ItemID          string `json:"item_id,omitempty"`
	ItemName        string `json:"item_name,omitempty"`
	ItemDescription string `json:"item_description,omitempty"`
	Status          string `json:"status,omitempty"`
	Quantity        int64  `json:"quantity,omitempty"`
}

type ShippingDetails struct {
	Address AddressDB `bson:"address" json:"address,omitempty"`
	Name    string    `bson:"name" json:"name,omitempty"`
	Phone   string    `bson:"phone" json:"phone,omitempty"`
}

type AddressDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"address_id,omitempty"`
	UserID      string             `bson:"user_id" json:"user_id,omitempty"`
	Address     string             `bson:"address" json:"address,omitempty"`
	Country     string             `bson:"country" json:"country,omitempty"`
	State       string             `bson:"state,omitempty" json:"state,omitempty"`
	City        string             `bson:"city,omitempty" json:"city,omitempty"`
	Pin         string             `bson:"pin" json:"pin,omitempty"`
	Primary     bool               `bson:"primary" json:"primary,omitempty"`
	GeoLocation bson.M             `bson:"geo_location" json:"geo_location,omitempty"`
	CreatedTime time.Time          `bson:"created_time" json:"created_time,omitempty"`
	UpdatedTime time.Time          `bson:"updated_time" json:"updated_time,omitempty"`
}
type OrderInteface struct {
	OrderList   []OrderDB   `json:"order_list,omitempty"`
	PaymentList []PaymentDB `json:"payment_list,omitempty"`
	CartList    []CartDB    `json:"carts_list,omitempty"`
}
type OrderDB struct {
	ID              primitive.ObjectID `bson:"_id" json:"order_id,omitempty"`
	DeliveryDetails DeliveryDB         `bson:"delivery_details" json:"delivery_details,omitempty"`
	PaymentID       string             `bson:"payment_id" json:"payment_id"`
	UserID          string             `bson:"user_id" json:"user_id,omitempty"`
	ShopID          string             `bson:"shop_id" json:"shop_id,omitempty"`
	CartID          string             `bson:"cart_id" json:"cart_id,omitempty"`
	Status          string             `bson:"status" json:"status,omitempty"`
	OrderedTime     time.Time          `bson:"ordered_time" json:"ordered_time,omitempty"`
	DeliveredTime   time.Time          `bson:"delivered_time" json:"delivered_time,omitempty"`
	OrderPickUpTime time.Time          `bson:"order_pickup_time" json:"order_pickup_time,omitempty"`
	UpdatedTime     time.Time          `bson:"updated_time" json:"updated_time,omitempty"`
	AddedTime       time.Time          `bson:"added_time" json:"added_time,omitempty"`
}
type DeliveryDB struct {
	TrackingID  string    `bson:"tracking_id" json:"tracking_id,omitempty"`
	DeliveryID  string    `bson:"delivery_id" json:"delivery_id,omitempty"`
	UserAddress AddressDB `bson:"user_address" json:"user_address"`
	ShopAddress AddressDB `bson:"shop_address" json:"shop_address"`
	UpdatedTime time.Time `bson:"updated_time" json:"updated_time,omitempty"`
	AddedTime   time.Time `bson:"added_time" json:"added_time,omitempty"`
}

type DeliveryPersonDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"delivery_person_id,omitempty"`
	Name        string             `bson:"name" json:"name,omitempty"`
	Age         string             `bson:"age" json:"age,omitempty"`
	Ratings     string             `bson:"ratings" json:"ratings,omitempty"`
	UpdatedTime time.Time          `bson:"updated_time" json:"updated_time,omitempty"`
	AddedTime   time.Time          `bson:"added_time" json:"added_time,omitempty"`
}

type CartDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"cart_id,omitempty"`
	UserID      string             `bson:"user_id" json:"user_id,omitempty"`
	ShopID      string             `bon:"shop_id" json:"shop_id,omitempty"`
	Items       []Item             `bson:"items" json:"items,omitempty"`
	UpdatedTime time.Time          `bson:"updated_time" json:"updated_time,omitempty"`
	AddedTime   time.Time          `bson:"added_time" json:"added_time,omitempty"`
	Status      string             `bson:"status" json:"status,omitempty"`
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

type PaymentDB struct {
	ID                 primitive.ObjectID `bson:"_id" json:"payment_id,omitempty"`
	UserID             string             `bson:"user_id" json:"user_id,omitempty"`
	Amount             int64              `bson:"amount" json:"amount,omitempty"`
	Currency           string             `bson:"currency" json:"currency,omitempty"`
	PaymentMethodTypes []string           `bson:"payment_method_types" json:"payment_method,omitempty"`
	Description        string             `bson:"description" json:"description,omitempty"`
	ReceiptEmail       string             `bson:"receipt_email" json:"receipt_email,omitempty"`
	Shipping           ShippingDetails    `bson:"shipping_details" json:"shipping_details,omitempty"`
	Status             string             `bson:"status" json:"status,omitempty"`
	UpdatedTime        time.Time          `bson:"updated_time" json:"updated_time,omitempty"`
	AddedTime          time.Time          `bson:"added_time" json:"added_time,omitempty"`
}
