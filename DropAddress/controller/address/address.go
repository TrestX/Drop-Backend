package address

import (
	entity "Drop/DropAddress/entities"
	"Drop/DropAddress/repository/address"

	"github.com/aekam27/trestCommon"

	"errors"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = address.NewAddressRepository("address")
)

type addressService struct{}

func NewAddressService(repository address.UserRepository) AddressService {
	repo = repository
	return &addressService{}
}

func (add *addressService) AddAddress(address *Address, userId string) (string, error) {
	var addressEntity entity.AddressDB
	if address == nil {
		err := errors.New("address id missing")
		trestCommon.ECLog2(
			"add address section",
			err,
		)
		return "", err
	}
	data, err := add.GetPrimaryAddress(userId)
	if err == nil && address.Primary {
		id := data.ID
		setParameters := bson.M{"primary": false}

		setParameters["updated_time"] = time.Now()
		filter := bson.M{"_id": id}
		set := bson.M{
			"$set": setParameters,
		}
		_, err = repo.UpdateOne(filter, set)
		if err != nil {
			trestCommon.ECLog3(
				"update address section",
				err,
				logrus.Fields{
					"address_id": data.ID,
				})
		}

	}
	addressEntity.ID = primitive.NewObjectID()
	addressEntity.UserID = userId
	addressEntity.Address = address.Address
	addressEntity.City = address.City
	addressEntity.Country = address.Country
	addressEntity.State = address.State
	addressEntity.Pin = address.Pin
	geoLocation := []float64{address.Longitude, address.Latitude}
	addressEntity.GeoLocation = bson.M{"type": "Point", "coordinates": geoLocation}
	addressEntity.Primary = address.Primary
	addressEntity.CreatedTime = time.Now()
	result, err := repo.InsertOne(addressEntity)
	if err != nil {
		return "", err
	}
	return result, nil
}
func createPreSignedDownloadUrl(url string) string {
	s := strings.Split(url, "?")
	o := strings.Split(s[0], "/")
	fileName := o[4]
	path := o[3]
	downUrl, _ := trestCommon.PreSignedDownloadUrl(fileName, path)
	return downUrl
}
func (*addressService) UpdateAddress(address *Address, addressid string) (string, error) {
	if addressid == "" {
		err := errors.New("address id missing")
		trestCommon.ECLog2(
			"update address section",
			err,
		)
		return "", err
	}
	id, _ := primitive.ObjectIDFromHex(addressid)
	setParameters := bson.M{}

	if address.Address != "" {
		setParameters["address"] = address.Address
	}
	if address.State != "" {
		setParameters["state"] = address.State
	}
	if address.City != "" {
		setParameters["city"] = address.City
	}
	if address.Country != "" {
		setParameters["country"] = address.Country
	}
	if address.Pin != "" {
		setParameters["pin"] = address.Pin
	}
	if address.Latitude != 0 && address.Longitude != 0 {
		geoLocation := []float64{address.Longitude, address.Latitude}
		setParameters["geo_location"] = bson.M{"type": "Point", "coordinates": geoLocation}
	}
	setParameters["updated_time"] = time.Now()
	filter := bson.M{"_id": id}
	set := bson.M{
		"$set": setParameters,
	}
	result, err := repo.UpdateOne(filter, set)
	if err != nil {
		trestCommon.ECLog3(
			"update address section",
			err,
			logrus.Fields{
				"address_id": addressid,
				"address":    address,
			})
		return "", err
	}

	return result, nil
}

func (add *addressService) PrimaryAddress(addressid string, userId string) (string, error) {
	if addressid == "" {
		err := errors.New("address id missing")
		trestCommon.ECLog2(
			"update address section",
			err,
		)
		return "", err
	}
	data, err := add.GetPrimaryAddress(userId)
	if err == nil {
		id := data.ID
		setParameters := bson.M{"primary": false}

		setParameters["updated_time"] = time.Now()
		filter := bson.M{"_id": id}
		set := bson.M{
			"$set": setParameters,
		}
		_, err = repo.UpdateOne(filter, set)
		if err != nil {
			trestCommon.ECLog3(
				"update address section",
				err,
				logrus.Fields{
					"address_id": addressid,
				})
		}

	}
	id, _ := primitive.ObjectIDFromHex(addressid)

	setParameters := bson.M{"primary": true}

	setParameters["updated_time"] = time.Now()
	filter := bson.M{"_id": id}
	set := bson.M{
		"$set": setParameters,
	}
	return repo.UpdateOne(filter, set)

}

func (*addressService) GetAddress(userId string, limit, skip int) ([]entity.AddressDB, error) {

	address, err := repo.Find(bson.M{"user_id": userId}, bson.M{}, limit, skip)
	if err != nil {
		trestCommon.ECLog2(
			"GetAddress section",
			err,
		)
		return address, err
	}
	return address, nil
}
func (*addressService) GetAddressUsingID(addressId string, userID string) (entity.AddressDB, error) {
	if addressId == "" {
		err := errors.New("address id missing")
		trestCommon.ECLog2(
			"update address section",
			err,
		)
		return entity.AddressDB{}, err
	}
	id, _ := primitive.ObjectIDFromHex(addressId)
	return repo.FindOne(bson.M{"_id": id}, bson.M{})

}
func (*addressService) GetPrimaryAddress(userID string) (entity.AddressDB, error) {
	return repo.FindOne(bson.M{"user_id": userID, "primary": true}, bson.M{})

}

func (*addressService) DeleteAddress(addressID string) error {
	id, _ := primitive.ObjectIDFromHex(addressID)
	return repo.DeleteOne(bson.M{"_id": id})
}
