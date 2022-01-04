package profile

import entity "Drop/DropUserAccount/entities"

type ProfileService interface {
	UpdateProfile(profile *Profile, userid string) (string, error)
	GetProfile(userID, token string) (entity.UserDB, int, int, int, int, error)
	CheckPhoneNumber(phoneNumber string) (entity.UserDB, error)
	GetAllUsers(userID, accountType, token string) ([]AdminOutput, error)
	ChangeProfileStatus(userID, status string) (string, error)
	GetAdminProfile(typ, status, stype string) ([]entity.UserDB, error)
	GetAdminUsersWithIDs(userIds []string) ([]entity.UserDB, error)
}

type Profile struct {
	Name                string `bson:"name" json:"name,omitempty"`
	AlternateEmail      string `bson:"alternate_email,omitempty" json:"alternate_email,omitempty"`
	PhoneNo             string `bson:"phone_number" json:"phone_number,omitempty"`
	AccountType         string `bson:"account_type" json:"account_type,omitempty"`
	LoggedInUsing       string `bson:"logged_in_using" json:"logged_in_using,omitempty"`
	Gender              string `bson:"gender" json:"gender,omitempty"`
	DOB                 string `bson:"dob" json:"dob,omitempty"`
	ProfilePhoto        string `bson:"profile_photo" json:"profile_photo,omitempty"`
	ApprovalStatus      string `bson:"approval_status" json:"approval,omitempty"`
	Availability        string `bson:"availability" json:"availability,omitempty"`
	Deleted             bool   `bson:"deleted" json:"deleted,omitempty"`
	Wallet              string `bson:"wallet" json:"wallet,omitempty"`
	NationalID          string `bson:"national_id" json:"national_id,omitempty"`
	PictureID           string `bson:"picture_id" json:"picture_id,omitempty"`
	VehiclePhoto        string `bson:"vehicle_photo" json:"vehicle_photo,omitempty"`
	VehicleRegistration string `bson:"vehicle_registration_document" json:"vehicle_registration_document,omitempty"`
	VehicleNumber       string `bson:"vehicle_number" json:"vehicle_number,omitempty"`
	VehicleType         string `bson:"vehicle_type" json:"vehicle_type,omitempty"`
}

type OutputInterface struct {
	Users  []entity.UserDB `json:"users,omitempty"`
	Orders interface{}     `json:"orders,omitempty"`
}

type AdminOutput struct {
	ID          string     `json:"id,omitempty"`
	Email       string     `json:"email,omitempty"`
	Name        string     `json:"name,omitempty"`
	Gender      string     `json:"gender,omitempty"`
	DOB         string     `json:"dob,omitempty"`
	AccountType string     `json:"account_type,omitempty"`
	Wallet      string     `json:"wallet,omitempty"`
	Orders      []OrdersOP `json:"orders"`
}

type OrdersOP struct {
	ID                  string        `json:"id,omitempty"`
	PaymentAmount       string        `json:"payment_amount,omitempty"`
	PaymentMethod       string        `json:"payment_method,omitempty"`
	Status              string        `json:"status,omitempty"`
	ItemsOrderd         []entity.Item `json:"items_orderd,omitempty"`
	DeliveryAddress     string        `json:"delivery_address,omitempty"`
	DeliveryGeoLocation interface{}   `json:"delivery_geo_location,omitempty"`
}
