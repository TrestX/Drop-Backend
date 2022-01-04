package api

import (
	entity "Drop/DropCategories/entities"
	"encoding/json"

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
