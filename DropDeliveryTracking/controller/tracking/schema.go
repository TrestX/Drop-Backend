package tracking

import entity "Drop/DropDeliveryTracking/entities"

type TrackingService interface {
	AddLocation(tracking *Tracking, deliveryBoyId string) (string, error)
	UpdateLocation(tracking *Tracking, trackingId string) (string, error)
	GetAllTrackingDetails(token string, limit, skip int) ([]entity.TrackingDB, error)
	GetTrackingByTrackingID(trackingId string) (entity.TrackingDB, error)
	GetAllDeliveryPersobOfAddressDetails(token string, limit, skip int, latitude, longitude float64) ([]entity.UserDB, error)
}

type Tracking struct {
	Longitude float64 `bson:"longitude" json:"longitude,omitempty"`
	Latitude  float64 `bson:"latitude" json:"latitude,omitempty"`
}
