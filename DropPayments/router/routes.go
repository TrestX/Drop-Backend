package router

import (
	paymentHandler "Drop/DropPayments/handlers/payment-handler"
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
		"/payments/create/paymentintent",
		paymentHandler.CreatePaymentIntent,
	},
	Route{
		"payments",
		"PUT",
		"/payments/{paymentID}",
		paymentHandler.UpdatePaymentStatus,
	},
	Route{
		"payments",
		"GET",
		"/payments",
		paymentHandler.GetPaymentSDetails,
	},
	Route{
		"payments",
		"GET",
		"/payments/success/{paymentID}",
		paymentHandler.GetPaymentSuccessDetails,
	},
	Route{
		"payments",
		"GET",
		"/payments/{paymentID}",
		paymentHandler.GetPaymentDetails,
	},
	Route{
		"payments",
		"GET",
		"/payments/get/publishableKey",
		paymentHandler.GetPublishableKey,
	},
	Route{
		"payments",
		"GET",
		"/payments/list/{paymentIds}",
		paymentHandler.GetPaymentsWithIDs,
	},
	Route{
		"payments",
		"GET",
		"/payments/admin/payments",
		paymentHandler.GetAdminPaymentSDetails,
	},
}
