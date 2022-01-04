package router

import (
	adminHandler "Drop/DropCategories/handlers/category-handler"
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
		"cat",
		"POST",
		"/category",
		adminHandler.AddCategory,
	},
	Route{
		"cat",
		"PUT",
		"/category/{categoryId}",
		adminHandler.UpdateCategory,
	},
	Route{
		"cat",
		"GET",
		"/category",
		adminHandler.GetActiveCategory,
	},
	Route{
		"cat",
		"GET",
		"/category/admin/category",
		adminHandler.GetAllCategory,
	},
}
