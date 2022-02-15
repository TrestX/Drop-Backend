package api

import (
	"encoding/json"
	"strings"

	"github.com/aekam27/trestCommon"
	"github.com/spf13/viper"

	entity "Drop/DropItems/entities"
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

type RatingReviewIdsResponse struct {
	Status bool
	Data   []entity.RatingReviewDB
}

func GetRatingDetailsByIDs(ratingIds []string) ([]entity.RatingReviewDB, error) {
	var users []entity.RatingReviewDB
	url := viper.GetString("api.getratingbyidsurl") + strings.Join(ratingIds, ",")
	body, err := trestCommon.GetApi(" ", url)
	if err != nil {
		return users, err
	}
	var response RatingReviewIdsResponse
	err = json.Unmarshal(body, &response)
	return response.Data, err
}
