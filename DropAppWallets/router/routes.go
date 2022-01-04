package router

import (
	apptransactionHandler "Drop/DropAppWallets/handlers/appwallet-handler"
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
		"apptrans",
		"POST",
		"/apptrans",
		apptransactionHandler.AddAppTransaction,
	},
	Route{
		"apptran",
		"GET",
		"/apptrans",
		apptransactionHandler.GetAppTransaction,
	},
	Route{
		"apptran",
		"PUT",
		"/apptrans/{apptransactionID}",
		apptransactionHandler.Updateapptransaction,
	},
	Route{
		"apptran",
		"GET",
		"/apptrans/all",
		apptransactionHandler.GetAppTransactions,
	},
	Route{
		"apptran",
		"GET",
		"/apptrans/delivery",
		apptransactionHandler.GetDeliveryAppTransactions,
	},
	Route{
		"apptran",
		"GET",
		"/apptrans/seller",
		apptransactionHandler.GetSellerAppTransactions,
	},
	Route{
		"apptran",
		"GET",
		"/apptrans/totalapp",
		apptransactionHandler.GetTotalAppEarning,
	},
	Route{
		"apptran",
		"GET",
		"/apptrans/totaltrans",
		apptransactionHandler.GetTotalTransAmt,
	},
	Route{
		"apptran",
		"GET",
		"/apptrans/admin/seller",
		apptransactionHandler.GetSellerTransactions,
	},
	Route{
		"apptran",
		"GET",
		"/apptrans/admin/delivery",
		apptransactionHandler.GetDeliveryTransactions,
	},
	Route{
		"apptran",
		"GET",
		"/apptrans/admin/seller/update",
		apptransactionHandler.UpdateSellerapptransactionPer,
	},
	Route{
		"apptran",
		"GET",
		"/apptrans/admin/seller/shops",
		apptransactionHandler.GetSellerShopsTransactions,
	},
	Route{
		"apptran",
		"PUT",
		"/apptrans/admin/seller/payment/history",
		apptransactionHandler.UpdateSellerPHTransactions,
	},
	Route{
		"apptran",
		"PUT",
		"/apptrans/admin/shop/payment/history",
		apptransactionHandler.UpdateShopPHTransactions,
	},
	Route{
		"apptran",
		"PUT",
		"/apptrans/admin/delivery/payment/history",
		apptransactionHandler.UpdateDeliveryPHTransactions,
	},
	Route{
		"apptran",
		"GET",
		"/apptrans/admin/payments/history",
		apptransactionHandler.GetSPaymentsHistory,
	},
}
