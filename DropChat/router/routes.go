package router

import (
	chatHandler "Drop/DropChat/handlers/chat-handler"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"chat",
		"POST",
		"/chat",
		chatHandler.AddChat,
	},
	Route{
		"chat",
		"GET",
		"/chat/{chatId}",
		chatHandler.GetChat,
	},
	Route{
		"chat",
		"GET",
		"/chat/user/{userId}",
		chatHandler.GetChatWithUserID,
	},
	Route{
		"chat",
		"PUT",
		"/chat",
		chatHandler.UpdateChat,
	},
	Route{
		"chat",
		"GET",
		"/chats",
		chatHandler.GetChats,
	},
	//Route{
	//	"cart",
	//	"Delete",
	//	"/cart/{itemId}",
	//	cartHandler.DeleteCart,
	//},
}
