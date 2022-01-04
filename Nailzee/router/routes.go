package router

import (
	paymentHandler "Nailzee/NailzeePayments/handlers/payment-handler"
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
		"payments",
		"POST",
		"/pay",
		paymentHandler.CreatePaymentIntent,
	},
	Route{
		"payments",
		"POST",
		"/newpayment",
		paymentHandler.CreateNewPaymentIntent,
	},
	Route{
		"payments",
		"POST",
		"/confirmpayment",
		paymentHandler.ConfirmPaymentIntent,
	},
	Route{
		"payments",
		"GET",
		"/publishableKey",
		paymentHandler.GetPublishableKey,
	},
}
