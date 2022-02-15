package order

import (
	"time"

	"Drop/Droporder/api"
	entity "Drop/Droporder/entities"
)

type OrderService interface {
	GetOrderWithIDs(userId []string) (interface{}, error)
	PlaceOrder(userId, token string, order Order) (AdminOrderOutput, error)
	UpdateOrder(orderId string, order Order) (string, error)
	GetOrder(orderId string) (entity.OrderDB, error)
	GetOrders(userId, token string, limit, skip int) ([]OrderOutput, error)
	GetAllOrdersAdmin(token, status string, limit, skip int) ([]OrderOutput, error)
	GetNewDeliveryOrders(token string, limit, skip int, latitude, longitude float64) ([]OrderOutput, error)
	GetActiveDeliveryOrders(token string, limit, skip int, deliveryID, status string) ([]OrderOutput, error)
	GetAdminOrders(token string, limit, skip int, shopID, sellerID, deliveryId, userId, status, fromD, endD string) ([]AdminOrderOutput, error)
	GetLatestOrders(userId, token string, limit, skip int) (AdminOrderOutput, error)
	GetAllUsers(token string, limit, skip int, deliveryID, status string) ([]interface{}, error)
	GetOrdersCount(userId []string) (map[string]int, error)
}
type Order struct {
	PaymentID            string               `bson:"payment_id" json:"payment_id"`
	ShopID               string               `bson:"shop_id" json:"shop_id"`
	CartID               string               `bson:"cart_id" json:"cart_id"`
	DeliveryID           string               `bson:"delivery_id" json:"delivery_id"`
	ShopAddressID        string               `bson:"shop_address_id" json:"shop_address_id"`
	UserAddressID        string               `bson:"user_address_id" json:"user_address_id"`
	TrackingID           string               `bson:"tracking_id" json:"tracking_id"`
	Status               string               `bson:"status" json:"status"`
	TipAmount            int64                `bson:"tip_amount" json:"tip_amount"`
	SellerAmount         int64                `bson:"seller_amount" json:"seller_amount"`
	DropAmount           int64                `bson:"drop_amount" json:"drop_amount"`
	OrderStatusList      []entity.OrderStatus `bson:"order_status_list" json:"order_status_list"`
	DeliveryCharge       int64                `bson:"delivery_charge" json:"delivery_charge"`
	DeliveryPersonCut    int64                `bson:"delivery_person_cut" json:"delivery_person_cut"`
	DeliveryPersonCutPer int64                `bson:"delivery_person_cut_per" json:"delivery_person_cut_per"`
	DeliveryCode         int64                `bson:"delivery_code" json:"delivery_code"`
	Refunded             bool                 `bson:"refunded" json:"refunded"`
	RefundedUsing        string               `bson:"refunded_using" json:"refunded_using"`
}
type OrderInteface struct {
	OrderList   []entity.OrderDB   `json:"order_list"`
	PaymentList []entity.PaymentDB `json:"payment_list"`
	CartList    []entity.CartDB    `json:"carts_list"`
}
type OrderOutput struct {
	ID               string                 `json:"order_Id"`
	OrderStatus      string                 `json:"orderStatus"`
	OrderedAt        time.Time              `json:"orderedAt"`
	AcceptedAt       time.Time              `json:"accepted_at"`
	ItemsOrdered     []entity.ItemDB        `json:"itemsOrdered"`
	PaymentAmount    int64                  `json:"paymentAmount"`
	OrderStatusList  []entity.OrderStatus   `bson:"order_status_list" json:"order_status_list"`
	PaymentCurrency  string                 `json:"paymentCurrency"`
	OrderDescription string                 `json:"orderDescription"`
	DeliveryDetails  entity.ShippingDetails `json:"deliveryDetails"`
	ShopDetails      entity.ShopDB          `json:"shopDetails"`
	SellerDetails    entity.UserDB          `json:"sellerDetails"`
	RefundedUsing    string                 `bson:"refunded_using" json:"refunded_using"`
	Refunded         bool                   `bson:"refunded" json:"refunded"`
}
type GEO struct {
	Data []float64 `bson:"data" json:"data"`
}
type AdminOrderOutput struct {
	ID                    string                 `json:"order_Id"`
	OrderStatus           string                 `json:"orderStatus"`
	OrderedAt             time.Time              `json:"orderedAt"`
	AcceptedAt            time.Time              `json:"accepted_at"`
	DeliveredTime         time.Time              `json:"delivered_time"`
	OrderPickeUpTime      time.Time              `json:"order_picke_up"`
	ItemsOrdered          []entity.ItemDB        `json:"itemsOrdered"`
	PaymentAmount         int64                  `json:"paymentAmount"`
	DeliveryPersonDetails entity.UserDB          `json:"deliveryPersonDetails"`
	DeliveryID            string                 `bson:"delivery_id" json:"delivery_id"`
	PaymentCurrency       string                 `json:"paymentCurrency"`
	OrderDescription      string                 `json:"orderDescription"`
	PaymentStatus         string                 `json:"paymentStatus"`
	CouponCode            string                 `json:"coupon_code"`
	PaymentMethodTypes    string                 `json:"payment_method"`
	DeliveryDetails       entity.ShippingDetails `json:"deliveryDetails"`
	ShopDetails           entity.ShopDB          `json:"shopDetails"`
	SellerDetails         entity.UserDB          `json:"sellerDetails"`
	OrderReview           api.RatingReviewDB     `json:"orderReview"`
	OrderStatusList       []entity.OrderStatus   `bson:"order_status_list" json:"order_status_list"`
	Distance              float64                `json:"distance"`
	TipAmount             int64                  `bson:"tip_amount" json:"tip_amount"`
	UserDetails           entity.UserDB          `json:"userDetails"`
	RefundedUsing         string                 `bson:"refunded_using" json:"refunded_using"`
	Refunded              bool                   `bson:"refunded" json:"refunded"`
}

type AppWallet struct {
	OrderId           string `bson:"order_id" json:"order_id"`
	ShopId            string `bson:"shop_id" json:"shop_id"`
	PaymentId         string `bson:"payment_id" json:"payment_id"`
	SellerID          string `bson:"seller_id" json:"seller_id"`
	DeliveryID        string `bson:"delivery_id" json:"delivery_id"`
	Status            string `bson:"status" json:"status"`
	SellerAmount      int64  `bson:"seller_amount" json:"seller_amount"`
	DropAmount        int64  `bson:"drop_amount" json:"drop_amount"`
	DeliveryCharge    int64  `bson:"delivery_charge" json:"delivery_charge"`
	DeliveryPersonCut int64  `bson:"delivery_person_cut" json:"delivery_person_cut"`
	TipAmount         int64  `bson:"tip_amount" json:"tip_amount"`
}
