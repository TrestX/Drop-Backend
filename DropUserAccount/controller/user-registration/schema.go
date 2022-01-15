package user_registration

import (
	"time"

	"firebase.google.com/go/auth"
)

type AccountService interface {
	SignUp(cred Credentials) (string, error)
	GSignUp(token *auth.Token) (string, error)
	GLogin(token *auth.Token) (string, error)
	Login(cred Credentials) (string, string, error)
	VerifyEmail(cred Credentials) (string, error)
	SendVerificationEmail(email string) (string, error)
	SendResetLink(email string) (string, error)
	VerifyResetLink(cred Credentials) (string, string, error)
	UpdatePassword(cred Credentials) (string, error)
}
type Credentials struct {
	Email             string    `bson:"email" json:"email,omitempty"`
	CurrentPassword   string    `bson:"currentpassword" json:"currentpassword,omitempty"`
	Password          string    `bson:"password" json:"password,omitempty"`
	CreatedDate       time.Time `bson:"created_date" json:"created_date,omitempty"`
	PhoneNo           string    `bson:"phone_number" json:"phone_number,omitempty"`
	Name              string    `bson:"name" json:"name,omitempty"`
	Status            string    `bson:"status" json:"status,omitempty"`
	Gender            string    `bson:"gender" json:"gender,omitempty"`
	DOB               string    `bson:"dob" json:"dob,omitempty"`
	AccountType       string    `bson:"account_type" json:"account_type,omitempty"`
	VerificationCode  string    `bson:"verification_code" json:"verification_code,omitempty"`
	EmailSentTime     time.Time `bson:"email_sent_time,omitempty" json:"email_sent_time,omitempty"`
	VerifiedTime      time.Time `bson:"verified_time,omitempty" json:"verified_time,omitempty"`
	TermsChecked      bool      `bson:"terms_checked" json:"terms_checked,omitempty"`
	PasswordResetCode string    `bson:"password_reset_code" json:"password_reset_code,omitempty"`
	PasswordResetTime time.Time `bson:"password_reset_time,omitempty" json:"password_reset_time,omitempty"`
	LoggedInUsing     string    `bson:"logged_in_using" json:"logged_in_using,omitempty"`
	Deleted           bool      `bson:"deleted" json:"deleted,omitempty"`
}
