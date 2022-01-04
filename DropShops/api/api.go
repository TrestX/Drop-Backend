package api

import (
	entity "Drop/DropShop/entities"
	"encoding/json"
	"strings"
	"time"

	"github.com/aekam27/trestCommon"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserResponse struct {
	Status bool
	Data   entity.UserDB
}

func GetUserDetails(token string) (entity.UserDB, error) {
	var profile entity.UserDB
	url := viper.GetString("api.getprofileurl")
	body, err := trestCommon.GetApi(token, url)
	if err != nil {
		return profile, err
	}
	var response UserResponse
	err = json.Unmarshal(body, &response)
	return response.Data, err
}

type UserIdsResponse struct {
	Status bool
	Data   []entity.UserDB
}

func GetUsersDetailsByIDs(userIds []string) ([]entity.UserDB, error) {
	var users []entity.UserDB
	url := viper.GetString("api.getprofilebyidsurl") + strings.Join(userIds, ",")
	body, err := trestCommon.GetApi(" ", url)
	if err != nil {
		return users, err
	}
	var response UserIdsResponse
	err = json.Unmarshal(body, &response)
	return response.Data, err
}

type RatingReviewDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID      string             `bson:"user_id" json:"user_id,omitempty"`
	EntityID    string             `bson:"entity_id" json:"entity_id,omitempty"`
	Rating      float64            `bson:"rating" json:"rating,omitempty"`
	Review      string             `bson:"review" json:"review,omitempty"`
	Deleted     bool               `bson:"deleted" json:"deleted,omitempty"`
	UpdatedTime time.Time          `bson:"updated_time" json:"updated_time,omitempty"`
	AddedTime   time.Time          `bson:"added_time" json:"added_time,omitempty"`
}

type ReviewResponse struct {
	Data []RatingReviewDB
}

func GetOrderReview(id string, token string) ([]RatingReviewDB, error) {
	url := viper.GetString("api.getOrderreviewURL") + "?entityId=" + id
	body, err := trestCommon.GetApi(" ", url)
	if err != nil {
		return []RatingReviewDB{}, err
	}
	var response ReviewResponse
	err = json.Unmarshal(body, &response)
	return response.Data, err
}
