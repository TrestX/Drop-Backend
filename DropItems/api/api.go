package api

import (
	entity "Drop/DropItems/entities"
	"encoding/json"

	"github.com/aekam27/trestCommon"
	"github.com/spf13/viper"
)

type ReviewResponse struct {
	Data []entity.RatingReviewDB
}

func GetOrderReview(id string, token string) ([]entity.RatingReviewDB, error) {
	url := viper.GetString("api.getOrderreviewURL") + "?entityId=" + id
	body, err := trestCommon.GetApi(" ", url)
	if err != nil {
		return []entity.RatingReviewDB{}, err
	}
	var response ReviewResponse
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
