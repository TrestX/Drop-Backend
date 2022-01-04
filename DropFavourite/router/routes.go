package router

import (
	favouriteHandler "Drop/DropFavourite/handlers/favourite-handler"
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
		"favourite",
		"POST",
		"/favourite/{itemId}",
		favouriteHandler.AddFavourite,
	},
	Route{
		"favourite",
		"GET",
		"/favourite",
		favouriteHandler.GetFavourite,
	},
	Route{
		"favourite",
		"Delete",
		"/favourite/{itemId}",
		favouriteHandler.DeleteFavourite,
	},
}
