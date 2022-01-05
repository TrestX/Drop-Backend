package chatHandler

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"

	controller "Drop/DropChat/controller/chat"
	"Drop/DropChat/log"
	"Drop/DropChat/repository/chat"
	util "Drop/DropChat/utils"

)

var (
	chatService = controller.NewChatService(chat.NewChatRepository("chat"))
)

func AddChat(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	log.DLogMap("adding to chat", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")
	if len(tokenString) < 2 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	claims, err := util.DecodeToken(tokenString[1])
	if err != nil {
		log.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	var chats controller.Chats
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &chats)
	if err != nil {
		log.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	data, err := chatService.AddChat(claims["userid"].(string), chats)
	if err != nil {
		log.ECLog1(errors.Wrapf(err, "unable to add to chat"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to add to chat"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	log.DLogMap("chat updated", logrus.Fields{
		"duration": duration,
	})
}

func UpdateChat(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	log.DLogMap("updating to chat", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")
	if len(tokenString) < 2 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	claims, err := util.DecodeToken(tokenString[1])
	if err != nil {
		log.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	var chats controller.Chats
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &chats)
	if err != nil {
		log.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	data, err := chatService.UpdateChat(claims["userid"].(string), chats)
	if err != nil {
		log.ECLog1(errors.Wrapf(err, "unable to add to chat"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to updating to chat"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	log.DLogMap("chat updated", logrus.Fields{
		"duration": duration,
	})
}

func GetChat(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	log.DLogMap("retrieving chat", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")
	if len(tokenString) < 2 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	_, err := util.DecodeToken(tokenString[1])
	if err != nil {
		log.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	var chatID = mux.Vars(r)["chatId"]
	if chatID == "" {
		log.ECLog1(errors.Wrapf(err, "unable to update "))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to update "})
		return
	}
	data, err := chatService.GetChat(chatID)
	if err != nil {
		log.ECLog1(errors.Wrapf(err, "unable to retrieve chat"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to retrieve chat"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	log.DLogMap("chat retrieved", logrus.Fields{
		"duration": duration,
	})
}

func GetChatWithUserID(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	log.DLogMap("retrieving chat", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")
	if len(tokenString) < 2 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	claim, err := util.DecodeToken(tokenString[1])
	if err != nil {
		log.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	var userId = mux.Vars(r)["userId"]
	status := r.URL.Query().Get("status")
	
	if userId == "" {
		log.ECLog1(errors.Wrapf(err, "unable to update "))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to update "})
		return
	}
	data, err := chatService.GetChatForUser(userId, claim["userid"].(string),status)
	if err != nil {
		log.ECLog1(errors.Wrapf(err, "unable to retrieve chat"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to retrieve chat"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	log.DLogMap("chat retrieved", logrus.Fields{
		"duration": duration,
	})
}
func GetChats(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	log.DLogMap("get all ", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")
	if len(tokenString) < 2 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	_, err := util.DecodeToken(tokenString[1])
	if err != nil {
		log.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	user := ""
	chat := ""
	status := ""
	userS := r.URL.Query().Get("userId")
	chatS := r.URL.Query().Get("chatId")
	statusS := r.URL.Query().Get("status")
	if userS != "" {
		user = userS
	}
	if chatS != "" {
		chat = chatS
	}
	if statusS != "" {
		status = statusS
	}
	data, err := chatService.GetChats(user, chat, status)
	if err != nil {
		log.ECLog1(errors.Wrapf(err, "unable to get all banners"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get all banners"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	log.DLogMap("get all banners success", logrus.Fields{
		"duration": duration,
	})
}
