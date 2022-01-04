package api

import (
	entity "Drop/DropChat/entities"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type UserResponse struct {
	Status bool
	Data   entity.UserDB
}

func GetUserDetails(token string) (entity.UserDB, error) {
	var profile entity.UserDB
	url := "https://api.drop-deliveryapp.com/docker1/user/profile"
	method := "GET"
	var bearer = "Bearer " + token

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return profile, err
	}
	req.Header.Add("Authorization", bearer)
	res, err := client.Do(req)
	if err != nil {
		return profile, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
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
	url := "https://api.drop-deliveryapp.com/docker1/user/list/" + strings.Join(userIds, ",")
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return users, err
	}
	res, err := client.Do(req)
	if err != nil {
		return users, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return users, err
	}
	var response UserIdsResponse
	err = json.Unmarshal(body, &response)
	return response.Data, err
}
