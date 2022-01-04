package api

import (
	entity "Drop/DropUserAccount/entities"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/aekam27/trestCommon"
	"github.com/spf13/viper"
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
