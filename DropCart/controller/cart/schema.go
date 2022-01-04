package cart

import (
	entity "Drop/DropCart/entities"
	"time"
)

type CartService interface {
	AddCart(shopID, userId, cartId string, item Items, token string) (string, string, error)
	UpdateCart(userId, status string) (string, error)
	GetCart(userId string) (CartGetSchema, error)
	GetCartWithIDs(cartId []string) ([]CartGetSchema, error)
}
type Items struct {
	Items  []entity.Item   `json:"items"`
	AddOns []entity.AddOns `json:"addons"`
}

type CartGetSchema struct {
	ID          string          `json:"cart_id"`
	UserID      string          `json:"user_id"`
	ShopID      string          `json:"shop_id"`
	Items       []entity.ItemDB `json:"items"`
	AddOns      []entity.AddOns `json:"addons"`
	UpdatedTime time.Time       `json:"updated_time"`
	AddedTime   time.Time       `json:"added_time"`
	Status      string          `json:"status"`
}
type AddOns struct {
	ItemID string `json:"item_id"`
	Name   string `bson:"name" json:"name"`
	Price  int64  `bson:"price" json:"price"`
	Note   string `bson:"note" json:"note"`
}
type Item struct {
	ItemID          string `json:"item_id"`
	ItemName        string `json:"item_name"`
	ItemDescription string `json:"item_description"`
	Status          string `json:"status"`
	Quantity        int64  `json:"quantity"`
	Price           int64  `json:"price"`
	Type            string `json:"type"`
}
