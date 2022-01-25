package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddressDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"address_id,omitempty"`
	UserID      string             `bson:"user_id" json:"user_id,omitempty"`
	Address     string             `bson:"address" json:"address,omitempty"`
	Country     string             `bson:"country" json:"country,omitempty"`
	State       string             `bson:"state,omitempty" json:"state,omitempty"`
	City        string             `bson:"city,omitempty" json:"city,omitempty"`
	Pin         string             `bson:"pin" json:"pin,omitempty"`
	Primary     bool               `bson:"primary" json:"primary,omitempty"`
	Type        string             `bson:"type" json:"type,omitempty"`
	Note        string             `bson:"note" json:"note,omitempty"`
	GeoLocation bson.M             `bson:"geo_location" json:"geo_location,omitempty"`
	CreatedTime time.Time          `bson:"created_time" json:"created_time,omitempty"`
	UpdatedTime time.Time          `bson:"updated_time" json:"updated_time,omitempty"`
}
type UserDB struct {
	ID                primitive.ObjectID `bson:"_id" json:"user_id,omitempty"`
	Email             string             `bson:"email,omitempty" json:"email,omitempty"`
	Status            string             `bson:"status,omitempty" json:"status,omitempty"`
	Name              string             `bson:"name" json:"name,omitempty"`
	Designation       string             `bson:"designation" json:"designation,omitempty"`
	PhoneNo           string             `bson:"phone_number" json:"phone_number,omitempty"`
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
	Availability      string             `bson:"availability" json:"availability,omitempty"`
	Deleted           bool               `bson:"deleted" json:"deleted,omitempty"`
	NationalID        string             `bson:"national_id" json:"national_id,omitempty"`
	PictureID         string             `bson:"picture_id" json:"picture_id,omitempty"`
	VehicleNumber     string             `bson:"vehicle_number" json:"vehicle_number,omitempty"`
	VehicleType       string             `bson:"vehicle_type" json:"vehicle_type,omitempty"`
	Wallet            string             `bson:"wallet" json:"wallet,omitempty"`
}
