package api

import (
	entity "Drop/DropWallet/entities"
	"encoding/json"

	"github.com/aekam27/trestCommon"
	"github.com/spf13/viper"
)

type Response struct {
	Status bool
	Data   interface{}
}
type Profile struct {
	Wallet string `bson:"wallet" json:"wallet,omitempty"`
}

func UpdateUserWallet(wallet Profile, userId, token string) (string, error) {
	url := viper.GetString("api.updateprofileurl")
	body, err := trestCommon.PostApi(token, url, wallet)
	if err != nil {
		return "", err
	}
	var response Response
	err = json.Unmarshal(body, &response)
	return "success", err
}

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
