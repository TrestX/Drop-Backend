package router

import (
	itemHandler "Drop/DropItems/handlers/item-handler"
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
		"item",
		"POST",
		"/item",
		itemHandler.AddItem,
	},
	Route{
		"item",
		"GET",
		"/item",
		itemHandler.GetItem,
	},
	Route{
		"item",
		"GET",
		"/item/{itemId}",
		itemHandler.GetFullItem,
	},
	Route{
		"item",
		"GET",
		"/item/my/item",
		itemHandler.GetMyItem,
	},
	Route{
		"item",
		"GET",
		"/item/list/{itemIds}",
		itemHandler.GetItemWithIDs,
	},
	Route{
		"item",
		"PUT",
		"/item/presigned",
		itemHandler.PreSignedUrl,
	},
	Route{
		"item",
		"PUT",
		"/item/presigned",
		itemHandler.PreSignedUrl,
	},
	Route{
		"item",
		"PUT",
		"/item/{itemID}",
		itemHandler.UpdateItem,
	},
	Route{
		"item",
		"GET",
		"/item/category/struct",
		itemHandler.GetItemStruc,
	},
	Route{
		"item",
		"GET",
		"/item/featured/items",
		itemHandler.GetFeaturedItem,
	},
	Route{
		"item",
		"GET",
		"/item/toprated/items",
		itemHandler.GetTopRatedItems,
	},
	Route{
		"item",
		"GET",
		"/item/featured/shop/items",
		itemHandler.GetShopFeaturedItem,
	},
	Route{
		"item",
		"GET",
		"/item/popular/items",
		itemHandler.GetPopularItem,
	},
}
