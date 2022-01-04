package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TrackingDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"tracking_id,omitempty"`
	DeliveryID  string             `bson:"delivery_id" json:"delivery_id,omitempty"`
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
}
