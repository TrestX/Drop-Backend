package api

import (
	entity "Drop/DropPayments/entities"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/aekam27/trestCommon"
	"github.com/spf13/viper"
)

type Response struct {
	Status bool
	Data   entity.AddressDB
}
type UserResponse struct {
	Status bool
	Data   entity.UserDB
}

func GetUserAddress(userAddressId, token string) (entity.AddressDB, error) {
	var address entity.AddressDB
	url := viper.GetString("api.addressbyidurl") + userAddressId

	body, err := trestCommon.GetApi(token, url)
	if err != nil {
		return address, err
	}
	var response Response
	err = json.Unmarshal(body, &response)
	return response.Data, err
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

type ShopResponse struct {
	Status bool
	Data   entity.ShopDB
}

func GetShopDetails(shopId, token string) (entity.ShopDB, error) {
	var shop entity.ShopDB
	url := viper.GetString("api.getshop") + shopId
	body, err := trestCommon.GetApi(token, url)
	if err != nil {
		return shop, err
	}
	var response ShopResponse
	err = json.Unmarshal(body, &response)
	return response.Data, err
}

type PaymentResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Errors  []interface{} `json:"errors"`
	Data    DataS         `json:"data"`
}
type DataS struct {
	Link string `json:"link"`
}

func SendPayemnt(bodyy entity.PaymentIntent) (PaymentResponse, error) {
	var payr PaymentResponse
	url := "https://api.flutterwave.com/v3/payments"
	method := "POST"
	var bearer = "Bearer " + "FLWSECK_TEST-1fc984c8b9de11333ed2c58d1cd1050f-X"
	requestByte, err := json.Marshal(bodyy)
	if err != nil {
		return payr, err
	}
	// resp, err := http.Post("https://httpbin.org/post", "application/json",
	// 	)
	// requestReader := bytes.NewReader(requestByte)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestByte))
	if err != nil {
		return payr, err
	}
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return payr, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return payr, err
	}
	err = json.Unmarshal(body, &payr)
	return payr, err
}

// json_data, err := json.Marshal(values)

// if err != nil {
// 	log.Fatal(err)
// }

// if err != nil {
// 	log.Fatal(err)
// }
