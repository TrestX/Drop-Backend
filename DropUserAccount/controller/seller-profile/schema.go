package seller_registration

import (
	entity "Drop/DropUserAccount/entities"
	"time"
)

type SellerAccountService interface {
	RegisterSellerPerson(cred Credentials) (string, error)
}
type Credentials struct {
	Email                      string         `bson:"email" json:"email,omitempty"`
	Password                   string         `bson:"password" json:"password,omitempty"`
	CreatedDate                time.Time      `bson:"created_date" json:"created_date,omitempty"`
	Name                       string         `bson:"name" json:"name,omitempty"`
	Status                     string         `bson:"status" json:"status,omitempty"`
	Gender                     string         `bson:"gender" json:"gender,omitempty"`
	DOB                        string         `bson:"dob" json:"dob,omitempty"`
	Type                       []ShopType     `bson:"type" json:"type,omitempty"`
	PhoneNumber                string         `bson:"phone_number" json:"phone_number,omitempty"`
	ProfilePhoto               string         `bson:"profile_photo" json:"profile_photo,omitempty"`
	Address                    entity.Address `bson:"address" json:"address,omitempty"`
	NationalID                 string         `bson:"national_id" json:"national_id,omitempty"`
	PictureID                  string         `bson:"picture_id" json:"picture_id,omitempty"`
	CertificateOfIncorporation string         `bson:"certificate_of_incorporation" json:"certificate_of_incorporation,omitempty"`
	Cuisine                    []string       `bson:"cuisine" json:"cuisine,omitempty"`
	DeliveryType               string         `bson:"delivery_type" json:"delivery_type,omitempty"`
	AccountType                string         `bson:"account_type" json:"account_type,omitempty"`
	VerificationCode           string         `bson:"verification_code" json:"verification_code,omitempty"`
	EmailSentTime              time.Time      `bson:"email_sent_time,omitempty" json:"email_sent_time,omitempty"`
	VerifiedTime               time.Time      `bson:"verified_time,omitempty" json:"verified_time,omitempty"`
	TermsChecked               bool           `bson:"terms_checked" json:"terms_checked,omitempty"`
	PasswordResetCode          string         `bson:"password_reset_code" json:"password_reset_code,omitempty"`
	PasswordResetTime          time.Time      `bson:"password_reset_time,omitempty" json:"password_reset_time,omitempty"`
	LoggedInUsing              string         `bson:"logged_in_using" json:"logged_in_using,omitempty"`
	Deleted                    bool           `bson:"deleted" json:"deleted,omitempty"`
	BankName                   string         `bson:"bankName" json:"bankName,omitempty"`
	AccountNumber              string         `bson:"account_number" json:"account_number,omitempty"`
	IFSC                       string         `bson:"ifsc" json:"ifsc,omitempty"`
	Tags                       string         `bson:"tags" json:"tags"`
	MinOrderAmount             int64          `bson:"minorderamount" json:"minorderamount"`
}

type ShopType struct {
	Name string `bson:"name" json:"name,omitempty"`
}

type Address struct {
	UserID          string   `bson:"user_id" json:"user_id,omitempty"`
	Address         string   `bson:"address" json:"address,omitempty"`
	Country         string   `bson:"country" json:"country,omitempty"`
	State           string   `bson:"state,omitempty" json:"state,omitempty"`
	City            string   `bson:"city,omitempty" json:"city,omitempty"`
	Pin             string   `bson:"pin" json:"pin,omitempty"`
	Primary         bool     `bson:"primary" json:"primary,omitempty"`
	Longitude       float64  `bson:"longitude" json:"longitude,omitempty"`
	Latitude        float64  `bson:"latitude" json:"latitude,omitempty"`
	Timing          string   `bson:"timing" json:"timing,omitempty"`
	ShopName        string   `bson:"shop_name" json:"shop_name,omitempty"`
	ShopPhotos      []string `bson:"shop_photos" json:"shop_photos,omitempty"`
	ShopStatus      string   `bson:"shop_status" json:"shop_status,omitempty"`
	ShopDescription string   `bson:"shop_description" json:"shop_description,omitempty"`
	ShopMenuImages  []string `bson:"shop_menu_images" json:"shop_menu_images,omitempty"`
}
