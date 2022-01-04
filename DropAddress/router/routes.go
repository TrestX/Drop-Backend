package router

import (
	addressHandler "Drop/DropAddress/handlers/address-handler"
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
		"address",
		"POST",
		"/address",
		addressHandler.AddAddress,
	},
	Route{
		"address",
		"GET",
		"/address",
		addressHandler.GetAddress,
	},
	Route{
		"address",
		"GET",
		"/address/{addressId}",
		addressHandler.GetFullAddress,
	},
	Route{
		"address",
		"GET",
		"/address/select/addressSelected",
		addressHandler.GetPrimaryAddress,
	},
	Route{
		"address",
		"PUT",
		"/address/{addressId}",
		addressHandler.UpdateAddress,
	},
	Route{
		"address",
		"PUT",
		"/address/addressSelected/{addressId}",
		addressHandler.UpdatePrimaryAddress,
	},
	Route{
		"address",
		"GET",
		"/address/admin/address",
		addressHandler.GetAdminAddress,
	},
	Route{
		"address",
		"DELETE",
		"/address/delete/{addressId}",
		addressHandler.DeleteAddress,
	},
}
