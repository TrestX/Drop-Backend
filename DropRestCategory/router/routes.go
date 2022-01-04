package router

import (
	adminHandler "Drop/DropItemCategories/handlers/category-handler"
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
		"/itemcategory",
		adminHandler.AddCategory,
	},
	Route{
		"cat",
		"PUT",
		"/itemcategory/{categoryId}",
		adminHandler.UpdateCategory,
	},
	Route{
		"cat",
		"GET",
		"/itemcategory",
		adminHandler.GetActiveCategory,
	},
	Route{
		"cat",
		"GET",
		"/itemcategory/admin",
		adminHandler.GetAllCategory,
	},
	Route{
		"cat",
		"GET",
		"/itemcategory/itemtags/admin",
		adminHandler.GetAllTags,
	},
	Route{
		"cat",
		"DELETE",
		"/itemcategory/{categoryId}",
		adminHandler.DeleteCategory,
	},
	Route{
		"cat",
		"DELETE",
		"/itemcategory/itemtags/{tagId}",
		adminHandler.DeleteTag,
	},
}
