package router

import (
	walletHandler "Drop/DropWallet/handlers/wallet-handler"
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
		"wallet",
		"POST",
		"/wallet",
		walletHandler.UpdateWallet,
	},
	Route{
		"wallet",
		"GET",
		"/wallet",
		walletHandler.GetWallet,
	},
	Route{
		"wallet",
		"GET",
		"/wallet/{userId}",
		walletHandler.GetWalletByUserId,
	},
	Route{
		"wallet",
		"GET",
		"/wallet/list/{walletIds}",
		walletHandler.GetWalletWithIDs,
	},
}
