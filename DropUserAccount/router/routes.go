package router

import (
	deliveryHandler "Drop/DropUserAccount/handlers/delivery-handler"
	profileHandler "Drop/DropUserAccount/handlers/profile-handler"
	sellerHandler "Drop/DropUserAccount/handlers/seller-handler"
	settingHandler "Drop/DropUserAccount/handlers/setting-handler"
	userHandler "Drop/DropUserAccount/handlers/user-handler"
	utilHandler "Drop/DropUserAccount/handlers/util-handler"
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
		"signup",
		"POST",
		"/user/signup",
		userHandler.SignUp,
	},
	Route{
		"login",
		"POST",
		"/user/login",
		userHandler.Login,
	},
	Route{
		"signup",
		"POST",
		"/user/google/login",
		userHandler.GLogin,
	},
	Route{
		"signup",
		"POST",
		"/user/google/signup",
		userHandler.GSignUp,
	},
	Route{
		"profile",
		"POST",
		"/user/profile",
		profileHandler.SetProfile,
	},
	Route{
		"profile",
		"GET",
		"/user/profile",
		profileHandler.Profile,
	},
	Route{
		"profile",
		"PUT",
		"/user/profile/{uid}",
		profileHandler.ChangeProfileStatus,
	},
	Route{
		"profile",
		"GET",
		"/user/admin/profile",
		profileHandler.GetAllProfile,
	},
	Route{
		"checkphone",
		"GET",
		"/user/checkphone/{phone_number}",
		profileHandler.CheckPhoneNumber,
	},
	Route{
		"setting",
		"GET",
		"/user/setting",
		settingHandler.Setting,
	},
	Route{
		"setting",
		"PUT",
		"/user/setting",
		settingHandler.SetSetting,
	},
	Route{
		"delivery",
		"POST",
		"/user/delivery/register",
		deliveryHandler.SetDeliveryProfile,
	},
	Route{
		"seller",
		"POST",
		"/user/seller/register",
		sellerHandler.SetSellerProfile,
	},
	Route{
		"util",
		"POST",
		"/user/util/presignedurl",
		utilHandler.GetPreSignedUrl,
	},
	Route{
		"profile",
		"GET",
		"/user/admin/all/profile",
		profileHandler.GetAdminProfile,
	},
	Route{
		"profile",
		"GET",
		"/user/list/{usersIds}",
		profileHandler.GetUsersWithIDs,
	},
}
