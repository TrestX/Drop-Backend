package router

import (
	adminHandler "Drop/DropSettings/handlers/admin-handler"
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
		"setting",
		"POST",
		"/setting",
		adminHandler.AddSetting,
	},
	Route{
		"setting",
		"GET",
		"/setting",
		adminHandler.GetAllSettings,
	},
	Route{
		"setting",
		"GET",
		"/setting/current",
		adminHandler.GetSettingsCurrent,
	},
}
