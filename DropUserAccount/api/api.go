package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/aekam27/trestCommon"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"

	entity "Drop/DropUserAccount/entities"
)

type AddressResponse struct {
	Status       bool
	Data         string
	PresignedUrl string
}

func AddUserAddress(address entity.Address, token string) (AddressResponse, error) {
	var response AddressResponse
	url := viper.GetString("api.addressurl")
	body, err := trestCommon.PostApi(token, url, address)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(body, &response)
	return response, err
}

type OrdersResponse struct {
	Status bool
	Data   []entity.OrderOutput
}
type OrdersIdsResponse struct {
	Status bool
	Data   OrderInteface
}
type OrderInteface struct {
	OrderList   []entity.OrderDB   `json:"order_list,omitempty"`
	PaymentList []entity.PaymentDB `json:"payment_list,omitempty"`
	CartList    []entity.CartDB    `json:"carts_list,omitempty"`
}
type Resp struct {
	Data []interface{} `json:"data"`
}
type RespCount struct {
	Data map[string]int `json:"data"`
}

func GetUserOrders(orderIds []string) (OrderInteface, error) {
	var response OrdersIdsResponse
	url := viper.GetString("api.orderIdsurl") + strings.Join(orderIds, ",")
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return response.Data, err
	}
	res, err := client.Do(req)
	if err != nil {
		return response.Data, err
	}
	defer res.Body.Close()

	body, err := trestCommon.GetApi(" ", url)
	if err != nil {
		return response.Data, err
	}
	err = json.Unmarshal(body, &response)
	return response.Data, err
}

func GetOrders(uidId, did, token, status string) (int, error) {
	var resp Resp
	url := ""
	if uidId != "" {
		url = viper.GetString("api.orders") + "?status=" + status + "&shopID=&deliveryId=&userId=" + uidId
	}
	if did != "" {
		url = viper.GetString("api.orders") + "?status=" + status + "&shopID=&deliveryId=" + did + "&userId="
	}
	body, err := trestCommon.GetApi(token, url)
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal(body, &resp)
	return len(resp.Data), err
}
func GetOrdersCounts(uidId []string) (map[string]int, error) {
	var resp RespCount
	url := viper.GetString("api.getordercount") + strings.Join(uidId, ",")

	body, err := trestCommon.GetApi("", url)
	if err != nil {
		return resp.Data, err
	}
	err = json.Unmarshal(body, &resp)
	return resp.Data, err
}

func SendEmail(name, email, link string) error {
	url := "https://api.emailjs.com/api/v1.0/email/send"
	sendbody := bson.M{
		"service_id":   "service_yxroydb",
		"template_id":  "template_tkjgk33",
		"user_id":      "user_IDQkJAjDozb8YEWlPkLZ8",
		"access_token": "23564f14c7bd59b8021623d723b2ca49",
		"template_params": bson.M{
			"to_name":  name,
			"to_email": email,
			"message":  link,
		},
	}
	method := "POST"
	requestByte, err := json.Marshal(sendbody)
	if err != nil {
		return err
	}
	requestReader := bytes.NewReader(requestByte)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, requestReader)

	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))
	return nil
}
