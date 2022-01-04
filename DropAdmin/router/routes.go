package router

import (
	adminHandler "Drop/DropAdmin/handlers/admin-handler"
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
		"banner",
		"POST",
		"/banner",
		adminHandler.AddBanner,
	},
	Route{
		"banner",
		"PUT",
		"/banner/{bannerId}",
		adminHandler.UpdateBanner,
	},
	Route{
		"banner",
		"GET",
		"/banner",
		adminHandler.GetActiveBanners,
	},
	Route{
		"banner",
		"GET",
		"/banner/admin/banner",
		adminHandler.GetAllBanners,
	},
}
