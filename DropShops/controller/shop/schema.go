package shop

import (
	entity "Drop/DropShop/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShopService interface {
	AddShopAdmin(shop *Shop) (string, error)
	AddShop(shop *Shop, sellerId string) (string, error)
	UpdateShop(shop *Shop, sellerId string) (string, error)
	PrimaryShop(shopid string, sellerId string) (string, error)
	GetPrimaryShop(userId string) (entity.ShopDB, error)
	GetShop(sellerId string, limit, skip int) ([]entity.ShopDB, error)
	GetShopUsingID(shopId string, sellerId string) (entity.ShopDB, error)
	GetFeaturedShop(limit, skip int) ([]entity.ShopDB, error)
	SearchShopByType(typ string, limit, skip int) ([]entity.ShopDB, error)
	GetShopAdmin(limit, skip int, sellerId, sType, status, featured, deal, rating, priceu, pricel, lowest string, lat, long float64) ([]OpSchema, error)
	GetNearestShopAdmin(limit, skip int, sellerId, sType, status string, lat, long float64) ([]OpSchema, error)
	GetTopRatedShopAdmin(limit, skip int, sellerId, sType, status string) ([]OpSchema, error)
	GetAdminUsersWithIDs(userIds []string) ([]entity.ShopDB, error)
}

type Shop struct {
	SellerID        string   `bson:"seller_id" json:"seller_id,omitempty"`
	Address         string   `bson:"address" json:"address,omitempty"`
	Country         string   `bson:"country" json:"country,omitempty"`
	State           string   `bson:"state,omitempty" json:"state,omitempty"`
	City            string   `bson:"city,omitempty" json:"city,omitempty"`
	Pin             string   `bson:"pin" json:"pin,omitempty"`
	Primary         bool     `bson:"primary" json:"primary,omitempty"`
	Timing          string   `bson:"timing" json:"timing,omitempty"`
	ShopName        string   `bson:"shop_name" json:"shop_name,omitempty"`
	ShopPhotos      []string `bson:"shop_photos" json:"shop_photos,omitempty"`
	ShopStatus      string   `bson:"shop_status" json:"shop_status,omitempty"`
	ShopDescription string   `bson:"shop_description" json:"shop_description,omitempty"`
	ShopLogo        string   `bson:"shop_logo" json:"shop_logo,omitempty"`
	ShopBanner      string   `bson:"shop_banner" json:"shop_banner,omitempty"`
	Longitude       float64  `bson:"longitude" json:"longitude,omitempty"`
	Latitude        float64  `bson:"latitude" json:"latitude,omitempty"`
	Type            string   `bson:"type" json:"type,omitempty"`
	Featured        bool     `bson:"featured" json:"featured,omitempty"`
	Deal            string   `bson:"deal" json:"deal,omitempty"`
	DeliveryType    string   `bson:"delivery" json:"delivery,omitempty"`
	Cuisine         string   `bson:"cuisine" json:"cuisine,omitempty"`
	Tags            string   `bson:"tags" json:"tags,omitempty"`
	Rating          float64  `bson:"rating" json:"rating,omitempty"`
	NumbofRating    int64    `bson:"nrating" json:"nrating,omitempty"`
	MinOrderAmount  int64    `bson:"minorderamount" json:"minorderamount"`
}

type OpSchema struct {
	ID              primitive.ObjectID `json:"shop_id,omitempty"`
	SellerID        string             `json:"seller_id,omitempty"`
	Address         string             `json:"address,omitempty"`
	Country         string             `json:"country,omitempty"`
	State           string             `json:"state,omitempty"`
	City            string             `json:"city,omitempty"`
	Pin             string             `json:"pin,omitempty"`
	Primary         bool               `json:"primary,omitempty"`
	Type            string             `json:"type,omitempty"`
	Timing          string             `json:"timing,omitempty"`
	ShopName        string             `json:"shop_name,omitempty"`
	ShopLogo        string             `json:"shop_logo,omitempty"`
	ShopBanner      string             `json:"shop_banner,omitempty"`
	ShopPhotos      []string           `json:"shop_photos,omitempty"`
	ShopStatus      string             `json:"shop_status,omitempty"`
	Featured        bool               `json:"featured,omitempty"`
	ShopDescription string             `json:"shop_description,omitempty"`
	GeoLocation     bson.M             `json:"geo_location,omitempty"`
	CreatedTime     time.Time          `json:"created_time,omitempty"`
	UpdatedTime     time.Time          `json:"updated_time,omitempty"`
	Tags            string             `bson:"tags" json:"tags,omitempty"`
	Rating          float64            `bson:"rating" json:"rating,omitempty"`
	SellerEmail     string             `json:"sellerEmail,omitempty"`
	Deal            string             `bson:"deal" json:"deal,omitempty"`
	DeliveryType    string             `bson:"delivery" json:"delivery,omitempty"`
	Cuisine         string             `bson:"cuisine" json:"cuisine,omitempty"`
	MinOrderAmount  int64              `bson:"minorderamount" json:"minorderamount"`
}
