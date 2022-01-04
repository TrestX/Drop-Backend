package router

import (
	adminHandler "Drop/DropAdminSupport/handlers/admin-handler"
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
		"/admin/acct",
		adminHandler.AddAcct,
	},
	Route{
		"cat",
		"POST",
		"/admin/acct/login",
		adminHandler.Login,
	},
	Route{
		"cat",
		"PUT",
		"/admin/acct/{Id}",
		adminHandler.UpdateAcct,
	},
	Route{
		"cat",
		"GET",
		"/admin/accts",
		adminHandler.GetAllAcct,
	},
	Route{
		"cat",
		"GET",
		"/admin/accts/sellers",
		adminHandler.GetAllShopAcct,
	},
}
