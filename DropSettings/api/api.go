package api

import (
	"encoding/json"

	"github.com/aekam27/trestCommon"
	"github.com/spf13/viper"
)

type Response struct {
	Status bool
	Data   int
}

func UpdateSellerPer(sellerid, per, c, token string) (int, error) {
	url := viper.GetString("api.updatetransac") + "?sId=" + sellerid + "&per=" + per + "&c=" + c
	body, err := trestCommon.GetApi(token, url)
	if err != nil {
		return 0, err
	}
	var response Response
	err = json.Unmarshal(body, &response)
	return response.Data, err
}
