package api

import (
	entity "Drop/DropCart/entities"
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

type ItemResponse struct {
	Status bool
	Data   entity.ItemDB
}

func GetItem(itemId string) (entity.ItemDB, error) {
	url := viper.GetString("api.getitem") + itemId
	body, err := trestCommon.GetApi(" ", url)
	if err != nil {
		return entity.ItemDB{}, err
	}
	var response ItemResponse
	err = json.Unmarshal(body, &response)
	return response.Data, err
}
