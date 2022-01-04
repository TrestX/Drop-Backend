package api

import (
	entity "Drop/DropDeliveryTracking/entities"
	"encoding/json"
	"strings"

	"github.com/aekam27/trestCommon"
	"github.com/spf13/viper"
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
