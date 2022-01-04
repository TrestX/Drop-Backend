package router

import (
	cartHandler "Drop/DropCart/handlers/cart-handler"
	"net/http"
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
		"cart",
		"POST",
		"/cart/{shopId}",
		cartHandler.AddCart,
	},
	Route{
		"cart",
		"GET",
		"/cart",
		cartHandler.GetCart,
	},
	Route{
		"cart",
		"PUT",
		"/cart/{cartID}",
		cartHandler.UpdateCart,
	},
	Route{
		"item",
		"GET",
		"/cart/list/{cartIds}",
		cartHandler.GetCartsWithIDs,
	},
	//Route{
	//	"cart",
	//	"Delete",
	//	"/cart/{itemId}",
	//	cartHandler.DeleteCart,
	//},
}
