package tracking

import (
	"errors"
	"strings"
	"time"

	"github.com/aekam27/trestCommon"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"Drop/DropDeliveryTracking/api"
	entity "Drop/DropDeliveryTracking/entities"
	"Drop/DropDeliveryTracking/repository/tracking"

)

var (
	repo = tracking.NewTrackingRepository("tracking")
)

type trackingService struct{}

func NewTrackingService(repository tracking.TrackingRepository) TrackingService {
	repo = repository
	return &trackingService{}
}

func (add *trackingService) AddLocation(tracking *Tracking, deliveryId string) (string, error) {
	if deliveryId == "" {
		err := errors.New("delivery boy id missing")
		trestCommon.ECLog2(
			"add tracking location",
			err,
		)
		return "", err
	}
	_, err := checkByTrackingID(deliveryId)
	if err != nil {
		var trackingEntity entity.TrackingDB
	trackingEntity.DeliveryID = deliveryId
	trackingEntity.CreatedTime = time.Now()
	trackingEntity.ID = primitive.NewObjectID()
	geoLocation := []float64{tracking.Longitude, tracking.Latitude}
	trackingEntity.GeoLocation = bson.M{"type": "Point", "coordinates": geoLocation}
	return repo.InsertOne(trackingEntity)
	}
	setParameters := bson.M{}
	if tracking.Latitude != 0 && tracking.Longitude != 0 {
		geoLocation := []float64{tracking.Longitude, tracking.Latitude}
		setParameters["geo_location"] = bson.M{"type": "Point", "coordinates": geoLocation}
	}
	setParameters["updated_time"] = time.Now()
	filter := bson.M{"delivery_id": deliveryId}
	set := bson.M{
		"$set": setParameters,
	}
	result, err := repo.UpdateOne(filter, set)
	if err != nil {
		trestCommon.ECLog3(
			"update tracking location",
			err,
			logrus.Fields{
				"tracking_id": deliveryId,
			})
		return "", err
	}

	return result, nil
	
}

func (*trackingService) UpdateLocation(tracking *Tracking, trackingid string) (string, error) {
	if trackingid == "" {
		err := errors.New("tracking id missing")
		trestCommon.ECLog2(
			"update tracking location",
			err,
		)
		return "", err
	}
	_, err := checkByTrackingID(trackingid)
	if err != nil {
		var trackingEntity entity.TrackingDB
		trackingEntity.DeliveryID = trackingid
		trackingEntity.CreatedTime = time.Now()
		trackingEntity.ID = primitive.NewObjectID()
		geoLocation := []float64{tracking.Longitude, tracking.Latitude}
		trackingEntity.GeoLocation = bson.M{"type": "Point", "coordinates": geoLocation}
		return repo.InsertOne(trackingEntity)
		}
	setParameters := bson.M{}
	if tracking.Latitude != 0 && tracking.Longitude != 0 {
		geoLocation := []float64{tracking.Longitude, tracking.Latitude}
		setParameters["geo_location"] = bson.M{"type": "Point", "coordinates": geoLocation}
	}
	setParameters["updated_time"] = time.Now()
	filter := bson.M{"delivery_id": trackingid}
	set := bson.M{
		"$set": setParameters,
	}
	result, err := repo.UpdateOne(filter, set)
	if err != nil {
		trestCommon.ECLog3(
			"update tracking location",
			err,
			logrus.Fields{
				"tracking_id": trackingid,
			})
		return "", err
	}

	return result, nil
}

func checkByTrackingID(id string) (entity.TrackingDB, error) {
	tracking, err := repo.FindOne(bson.M{"delivery_id": id}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get Tracking Details section",
			err,
		)
		return tracking, err
	}
	return tracking, nil
}

func getUserDetails(token string) error {
	user, err := api.GetUserDetails(token)
	if err != nil {
		return err
	}
	if strings.ToLower(user.AccountType) != "admin" {
		return errors.New("user doesnot have admin privilages")
	}
	return nil
}

func (*trackingService) GetAllTrackingDetails(token string, limit, skip int) ([]entity.TrackingDB, error) {
	err := getUserDetails(token)
	if err != nil {
		return []entity.TrackingDB{}, err
	}
	tracking, err := repo.Find(bson.M{}, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"Get Tracking section",
			err,
		)
		return tracking, err
	}
	return tracking, nil
}

func (*trackingService) GetTrackingByTrackingID(trackingId string) (entity.TrackingDB, error) {
	id, _ := primitive.ObjectIDFromHex(trackingId)
	tracking, err := repo.FindOne(bson.M{"delivery_id": id}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get Tracking Details section",
			err,
		)
		return tracking, err
	}
	return tracking, nil
}

func (*trackingService) GetAllDeliveryPersobOfAddressDetails(token string, limit, skip int, latitude, longitude float64) ([]entity.UserDB, error) {
	err := getUserDetails(token)
	if err != nil {
		return []entity.UserDB{}, err
	}
	tracking, err := repo.Find(bson.M{"geo_location": bson.M{
		"$near": bson.M{
			"$geometry": bson.M{
				"type":        "Point",
				"coordinates": []float64{latitude, longitude},
			},
			"$maxDistance": 10000,
		},
	}}, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"Get Tracking section",
			err,
		)
		return []entity.UserDB{}, err
	}
	l := []string{}
	for i := 0; i < len(tracking); i++ {
		l = append(l, tracking[i].DeliveryID)
	}
	listOfDeliveryBoys, err := api.GetUsersDetailsByIDs(l)
	if err != nil {
		trestCommon.ECLog2(
			"Get Tracking section",
			err,
		)
		return []entity.UserDB{}, err
	}
	return listOfDeliveryBoys, nil
}
