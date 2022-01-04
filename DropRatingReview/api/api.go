package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

type ShopResponse struct {
	Status bool
	Data   string
}

func UpdateShopRating(id string, doc interface{}, token string) (string, error) {
	url := viper.GetString("api.updateShopReview") + id
	method := "PUT"
	var bearer = "Bearer " + token
	requestByte, err := json.Marshal(doc)
	if err != nil {
		return "", err
	}
	requestReader := bytes.NewReader(requestByte)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, requestReader)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", bearer)
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var response ShopResponse
	err = json.Unmarshal(body, &response)
	return response.Data, err
}
