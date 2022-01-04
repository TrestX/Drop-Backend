package router

import (
	trackingHandler "Drop/DropDeliveryTracking/handlers/tracking-handler"
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
		"tracking",
		"POST",
		"/tracking/{deliveryId}",
		trackingHandler.AddTracking,
	},
	Route{
		"tracking",
		"PUT",
		"/tracking/{trackingId}",
		trackingHandler.UpdateTracking,
	},
	Route{
		"tracking",
		"GET",
		"/tracking/{trackingId}",
		trackingHandler.GetByTrackingID,
	},
	Route{
		"tracking",
		"GET",
		"/tracking",
		trackingHandler.GetTracking,
	},
	Route{
		"tracking",
		"GET",
		"/tracking/admin/deliveryperson/tracking",
		trackingHandler.GetNearDeliveryPerson,
	},
}
