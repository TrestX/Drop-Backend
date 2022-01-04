package router

import (
	couponHandler "Drop/DropCoupons/handlers/coupon-handler"
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
		"coupon",
		"POST",
		"/coupon",
		couponHandler.AddCoupon,
	},
	Route{
		"coupon",
		"PUT",
		"/coupon/{couponId}",
		couponHandler.UpdateCoupon,
	},
	Route{
		"coupon",
		"GET",
		"/coupon",
		couponHandler.GetCoupon,
	},
	Route{
		"coupon",
		"GET",
		"/coupon/all",
		couponHandler.GetCoupons,
	},
	Route{
		"coupon",
		"GET",
		"/coupon/list/{couponIds}",
		couponHandler.GetCouponsWithIDs,
	},
}
