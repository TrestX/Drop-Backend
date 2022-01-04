package router

import (
	shopHandler "Drop/DropShop/handlers/shop-handler"
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
		"shop",
		"POST",
		"/shop",
		shopHandler.AddShop,
	},
	Route{
		"shop",
		"GET",
		"/shop",
		shopHandler.GetShop,
	},
	Route{
		"shop",
		"GET",
		"/shop/{shopId}",
		shopHandler.GetFullShops,
	},
	Route{
		"shop",
		"GET",
		"/shop/select/shopSelected",
		shopHandler.GetPrimaryShop,
	},
	Route{
		"shop",
		"PUT",
		"/shop/{shopId}",
		shopHandler.UpdateShop,
	},
	Route{
		"shop",
		"PUT",
		"/shop/shopSelected/{shopId}",
		shopHandler.UpdatePrimaryShop,
	},
	Route{
		"shop",
		"GET",
		"/shop/featuredShops",
		shopHandler.GetFeaturedShops,
	},
	Route{
		"shop",
		"GET",
		"/shop/shopbytype",
		shopHandler.GetShopsByType,
	},
	Route{
		"shop",
		"POST",
		"/shop/admin/shop",
		shopHandler.AddShopAdmin,
	},
	Route{
		"shop",
		"GET",
		"/shop/admin/shop",
		shopHandler.GetShopAdmin,
	},
	Route{
		"shop",
		"GET",
		"/shop/admin/toprated",
		shopHandler.GetTopRatedShopAdmin,
	},
	Route{
		"shop",
		"GET",
		"/shop/admin/fastestDelivery",
		shopHandler.GetNearestShopAdmin,
	},
}
