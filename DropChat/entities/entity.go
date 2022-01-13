package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatDB struct {
	ID               primitive.ObjectID `bson:"_id" json:"hat_id"`
	SenderID         string             `bson:"sender_id" json:"sender_id"`
	SenderName       string             `bson:"sender_name" json:"sender_name"`
	ReceiverName     string             `bson:"receiver_name" json:"receiver_name"`
	SenderImage      string             `bson:"sender_image" json:"sender_image"`
	ReceiverImage    string             `bson:"receiver_image" json:"receiver_image"`
	ReceiverID       string             `bson:"receiver_id" json:"receiver_id"`
	ReceiverJoinTime time.Time          `bson:"receiver_join_time" json:"receiver_join_time"`
	StopTime         time.Time          `bson:"stop_time" json:"stop_time"`
	Chat             []Chat             `bson:"chat" json:"chat"`
	Status           string             `bson:"status" json:"status"`
	UpdatedTime      time.Time          `bson:"updated_time" json:"updated_time"`
	AddedTime        time.Time          `bson:"added_time" json:"added_time"`
}

type Chat struct {
	Sender  string    `bson:"key" json:"key"`
	Message string    `bson:"message" json:"message"`
	Time    time.Time `bson:"time" json:"time"`
	Name    string    `bson:"name" json:"name"`
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
