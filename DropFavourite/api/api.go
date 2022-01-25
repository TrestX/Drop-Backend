package api

import (
	entity "Drop/DropFavourite/entities"
	"encoding/json"
	"strings"

	"github.com/aekam27/trestCommon"
	"github.com/spf13/viper"
)

type Response struct {
	Status bool
	Data   []entity.ItemDB
}

func GetItems(items []string) ([]entity.ItemDB, error) {

	url := viper.GetString("api.getitembyid") + strings.Join(items, ",")
	body, err := trestCommon.GetApi(" ", url)
	if err != nil {
		return nil, err
	}
	var response Response
	err = json.Unmarshal(body, &response)
	return response.Data, err
}

type ShopIdsResponse struct {
	Status bool
	Data   []entity.ShopDB
}

func GetShopDetailsByIDs(shopIds []string) ([]entity.ShopDB, error) {
	var users []entity.ShopDB
	url := viper.GetString("api.getshopbyidsurl") + strings.Join(shopIds, ",")
	body, err := trestCommon.GetApi(" ", url)
	if err != nil {
		return users, err
	}
	var response ShopIdsResponse
	err = json.Unmarshal(body, &response)
	return response.Data, err
}
