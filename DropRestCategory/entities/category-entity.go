package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryDB struct {
	ID                   primitive.ObjectID `bson:"_id" json:"category_id,omitempty"`
	PresignedDownloadUrl string             `bson:"presignedurl" json:"presignedurl,omitempty"`
	Status               string             `bson:"status" json:"status,omitempty"`
	ShopType             string             `bson:"shop_type" json:"shop_type,omitempty"`
	CategoryName         string             `bson:"category_name" json:"category_name,omitempty"`
	CreatedTime          time.Time          `bson:"created_time" json:"created_time,omitempty"`
	UpdatedTime          time.Time          `bson:"updated_time" json:"updated_time,omitempty"`
}

type TagDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"category_id,omitempty"`
	Status      string             `bson:"status" json:"status,omitempty"`
	ShopType    string             `bson:"shop_type" json:"shop_type,omitempty"`
	TagName     string             `bson:"tag_name" json:"tag_name,omitempty"`
	CreatedTime time.Time          `bson:"created_time" json:"created_time,omitempty"`
	UpdatedTime time.Time          `bson:"updated_time" json:"updated_time,omitempty"`
}

type UserDB struct {
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
	Availability      string             `bson:"availability" json:"availability,omitempty"`
	Deleted           bool               `bson:"deleted" json:"deleted,omitempty"`
	NationalID        string             `bson:"national_id" json:"national_id,omitempty"`
	PictureID         string             `bson:"picture_id" json:"picture_id,omitempty"`
	VehicleNumber     string             `bson:"vehicle_number" json:"vehicle_number,omitempty"`
	VehicleType       string             `bson:"vehicle_type" json:"vehicle_type,omitempty"`
}
