package api

import (
	entity "Drop/Droporder/entities"
	"encoding/json"
	"strings"
	"time"

	"github.com/aekam27/trestCommon"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddressResponse struct {
	Status bool
	Data   entity.AddressDB
}
type ShopResponse struct {
	Status bool
	Data   entity.ShopDB
}
type PaymentResponse struct {
	Status bool
	Data   entity.PaymentDB
}

type CartResponse struct {
	Status bool
	Data   entity.CartDB
}
type PaymentIdsResponse struct {
	Status bool
	Data   []entity.PaymentDB
}

type CartIdsResponse struct {
	Status bool
	Data   []entity.CartDB
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

func GetUserAddress(addressId, token string) (entity.AddressDB, error) {
	var address entity.AddressDB
	url := viper.GetString("api.addressbyidurl") + addressId
	body, err := trestCommon.GetApi(token, url)
	if err != nil {
		return address, err
	}
	var response AddressResponse
	err = json.Unmarshal(body, &response)
	return response.Data, err
}

func GetUserPaymentDetails(paymentId, token string) (PaymentResponse, error) {
	var payment PaymentResponse
	url := viper.GetString("api.getpaymentbyid") + paymentId
	body, err := trestCommon.GetApi(token, url)
	if err != nil {
		return payment, err
	}
	var response PaymentResponse
	err = json.Unmarshal(body, &response)
	return response, err
}

func GetUserPaymentDetailsPaymentIds(paymentIds []string) ([]entity.PaymentDB, error) {
	var payment []entity.PaymentDB
	url := viper.GetString("api.getpaymentsbyids") + strings.Join(paymentIds, ",")
	body, err := trestCommon.GetApi(" ", url)
	if err != nil {
		return payment, err
	}
	var response PaymentIdsResponse
	err = json.Unmarshal(body, &response)
	return response.Data, err
}

func GetUserCartDetails(cartId, token string) (CartResponse, error) {
	var cart CartResponse
	url := viper.GetString("api.getcartbyid")
	body, err := trestCommon.GetApi(token, url)
	if err != nil {
		return cart, err
	}
	var response CartResponse
	err = json.Unmarshal(body, &response)
	return response, err
}

func GetCartsDetails(cartIds []string) ([]entity.CartDB, error) {
	var cart []entity.CartDB
	url := viper.GetString("api.getcartbyids") + strings.Join(cartIds, ",")
	body, err := trestCommon.GetApi(" ", url)
	if err != nil {
		return cart, err
	}
	var response CartIdsResponse
	err = json.Unmarshal(body, &response)
	return response.Data, err
}

type UserResponse struct {
	Status bool
	Data   entity.UserDB
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

type AppWalletResponse struct {
	Status bool
	Data   string
}

func PostAppWallet(appWallet interface{}, token string) (string, error) {
	url := viper.GetString("api.getAppWalletURL")
	body, err := trestCommon.PostApi(token, url, appWallet)
	if err != nil {
		return "", err
	}
	var response AppWalletResponse
	err = json.Unmarshal(body, &response)
	return response.Data, err
}

type RatingReviewDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	UserID      string             `bson:"user_id" json:"user_id"`
	EntityID    string             `bson:"entity_id" json:"entity_id"`
	Rating      float64            `bson:"rating" json:"rating"`
	Review      string             `bson:"review" json:"review"`
	Deleted     bool               `bson:"deleted" json:"deleted"`
	UpdatedTime time.Time          `bson:"updated_time" json:"updated_time"`
	AddedTime   time.Time          `bson:"added_time" json:"added_time"`
}

type ReviewResponse struct {
	Data []RatingReviewDB
}

func GetOrderReview(id string, token string) ([]RatingReviewDB, error) {
	url := viper.GetString("api.getOrderreviewURL") + "?entityId=" + id
	body, err := trestCommon.GetApi(" ", url)
	if err != nil {
		return []RatingReviewDB{}, err
	}
	var response ReviewResponse
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
	Data   []RatingReviewDB
}

func GetRatingDetailsByIDs(ratingIds []string) ([]RatingReviewDB, error) {
	var users []RatingReviewDB
	url := viper.GetString("api.getratingbyidsurl") + strings.Join(ratingIds, ",")
	body, err := trestCommon.GetApi(" ", url)
	if err != nil {
		return users, err
	}
	var response RatingReviewIdsResponse
	err = json.Unmarshal(body, &response)
	return response.Data, err
}
