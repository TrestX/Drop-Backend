package router

import (
	"net/http"

	orderHandler "Drop/Droporder/handlers/order-handler"

)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains all routes
type Routes []Route

var routes = Routes{

	Route{
		"place order",
		"POST",
		"/order",
		orderHandler.PlaceOrder,
	},
	Route{
		"get all orders",
		"GET",
		"/order",
		orderHandler.GetOrders,
	},
	Route{
		"get single order",
		"GET",
		"/order/{orderId}",
		orderHandler.GetOrderDetails,
	},
	Route{
		"order",
		"PUT",
		"/order/{orderId}",
		orderHandler.UpdateOrder,
	},
	Route{
		"order",
		"GET",
		"/order/admin/orders",
		orderHandler.GetAllOrdersAdmin,
	},
	Route{
		"order",
		"GET",
		"/order/delivery/orders",
		orderHandler.GetNewOrdersDelivery,
	},
	Route{
		"order",
		"GET",
		"/order/delivery/orders/{deliveryID}",
		orderHandler.GetOrdersDeliveryByStatus,
	},
	Route{
		"order",
		"GET",
		"/order/list/{userIds}",
		orderHandler.GetOrderWithIDs,
	},
	Route{
		"order",
		"GET",
		"/order/admin/orders/all",
		orderHandler.GetAdminOrders,
	},
	Route{
		"order",
		"GET",
		"/order/admin/users/all/{deliveryID}",
		orderHandler.GetAllUsers,
	},
	Route{
		"order",
		"GET",
		"/order/latest/order",
		orderHandler.GetLatestOrders,
	},

	Route{
		"ordernotification",
		"GET",
		"/notification",
		orderHandler.Getnotification,
	},
}
