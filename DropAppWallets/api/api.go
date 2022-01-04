package api

import (
	entity "Drop/DropAppWallets/entities"
	"encoding/json"
	"strings"

	"github.com/aekam27/trestCommon"
	"github.com/spf13/viper"
)

type Response struct {
	Status bool
	Data   entity.OrderDB
}
type PayResponse struct {
	Status bool
	Data   []entity.PaymentEntityDB
}

func GetOrder(orderid, token string) (entity.OrderDB, error) {
	url := viper.GetString("api.getorderbyid") + orderid
	body, err := trestCommon.GetApi(token, url)
	if err != nil {
		return entity.OrderDB{}, err
	}
	var response Response
	err = json.Unmarshal(body, &response)
	return response.Data, err
}
func GetPaymentByIds(paymentid []string) ([]entity.PaymentEntityDB, error) {
	url := viper.GetString("api.getpaymentsbyids") + strings.Join(paymentid, ",")
	body, err := trestCommon.GetApi(" ", url)
	if err != nil {
		return []entity.PaymentEntityDB{}, err
	}
	var response PayResponse
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

type SettingResponse struct {
	Status bool
	Data   []entity.SettingDB
}

func GetSettings() ([]entity.SettingDB, error) {
	var users []entity.SettingDB
	url := viper.GetString("api.getperct")
	body, err := trestCommon.GetApi(" ", url)
	if err != nil {
		return users, err
	}
	var response SettingResponse
	err = json.Unmarshal(body, &response)
	return response.Data, err
}

func GetSellers(token string) ([]entity.UserDB, error) {
	var users []entity.UserDB
	url := viper.GetString("api.getSellers")
	body, err := trestCommon.GetApi(token, url)
	if err != nil {
		return users, err
	}
	var response UserIdsResponse
	err = json.Unmarshal(body, &response)
	return response.Data, err
}

func GetDelivery(token string) ([]entity.UserDB, error) {
	var users []entity.UserDB
	url := viper.GetString("api.getDelivery")
	body, err := trestCommon.GetApi(token, url)
	if err != nil {
		return users, err
	}
	var response UserIdsResponse
	err = json.Unmarshal(body, &response)
	return response.Data, err
}

type ShopResponse struct {
	Status bool
	Data   []entity.ShopDB
}

func GetShopBySId(sid, token string) ([]entity.ShopDB, error) {
	var shops []entity.ShopDB
	url := viper.GetString("api.getshopbysid") + "?sellerId=" + sid
	body, err := trestCommon.GetApi(token, url)
	if err != nil {
		return shops, err
	}
	var response ShopResponse
	err = json.Unmarshal(body, &response)
	return response.Data, err
}
