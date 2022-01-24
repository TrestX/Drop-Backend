package address

import entity "Drop/DropAddress/entities"

type AddressService interface {
	AddAddress(address *Address, userId string) (string, error)
	UpdateAddress(address *Address, userId string) (string, error)
	PrimaryAddress(addressid string, userId string) (string, error)
	GetPrimaryAddress(userId string) (entity.AddressDB, error)
	DeleteAddress(addressID string) error
	GetAddress(userId string, limit, skip int) ([]entity.AddressDB, error)
	GetAddressUsingID(addressId string, userID string) (entity.AddressDB, error)
}

type Address struct {
	UserID    string  `bson:"user_id" json:"user_id,omitempty"`
	Address   string  `bson:"address" json:"address,omitempty"`
	Country   string  `bson:"country" json:"country,omitempty"`
	State     string  `bson:"state,omitempty" json:"state,omitempty"`
	City      string  `bson:"city,omitempty" json:"city,omitempty"`
	Pin       string  `bson:"pin" json:"pin,omitempty"`
	Primary   bool    `bson:"primary" json:"primary,omitempty"`
	Timing    string  `bson:"timing" json:"timing,omitempty"`
	Longitude float64 `bson:"longitude" json:"longitude,omitempty"`
	Latitude  float64 `bson:"latitude" json:"latitude,omitempty"`
	Type      string  `bson:"type" json:"type,omitempty"`
	Note      string  `bson:"note" json:"note,omitempty"`
}

type OP struct {
	UserID    string  `bson:"user_id" json:"user_id,omitempty"`
	Name      string  `bson:"name" json:"name,omitempty"`
	Address   string  `bson:"address" json:"address,omitempty"`
	Country   string  `bson:"country" json:"country,omitempty"`
	State     string  `bson:"state,omitempty" json:"state,omitempty"`
	City      string  `bson:"city,omitempty" json:"city,omitempty"`
	Pin       string  `bson:"pin" json:"pin,omitempty"`
	Primary   bool    `bson:"primary" json:"primary,omitempty"`
	Timing    string  `bson:"timing" json:"timing,omitempty"`
	Longitude float64 `bson:"longitude" json:"longitude,omitempty"`
	Latitude  float64 `bson:"latitude" json:"latitude,omitempty"`
	Type      string  `bson:"type" json:"type,omitempty"`
	Note      string  `bson:"note" json:"note,omitempty"`
}
