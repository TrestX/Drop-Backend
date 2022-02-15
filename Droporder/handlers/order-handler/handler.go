package orderHandler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aekam27/trestCommon"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"

	controller "Drop/Droporder/controller/order"
	"Drop/Droporder/repository/order"
	notification "Drop/Droporder/repository/order/notificationrepo"
	util "Drop/Droporder/util"
)

var (
	orderService = controller.NewOrderService(order.NewOrderRepository("order"))
)
var (
	notificationService = util.NewNotificationService(notification.NewNotificationRepository("notification"))
)

func SendNotificationWithTopic(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("send notifications", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var message util.Notification
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &message)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to send notification"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to send notification"})
		return
	}
	data, err := notificationService.SendNotificationWithTopic(message.Title, message.Body, message.Topic, message.UserId)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to send notification"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to send notification"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("sendnotification success", logrus.Fields{
		"duration": duration,
	})

}

func Getnotification(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("get notifications", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	userId := ""
	topic := ""
	title := ""
	status := "Active"
	limit := 20
	skip := 0
	var err error
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	userIdS := r.URL.Query().Get("userId")
	topicS := r.URL.Query().Get("topic")
	statusS := r.URL.Query().Get("status")
	titleS := r.URL.Query().Get("title")
	if userIdS != "" {
		userId = userIdS
	}
	if topicS != "" {
		topic = topicS
	}
	if titleS != "" {
		title = titleS
	}
	if statusS != "" {
		status = statusS
	}
	if limitS != "" {
		limit, err = strconv.Atoi(limitS)
		if err != nil {
			limit = 20
		}
	}
	if skipS != "" {
		skip, err = strconv.Atoi(skipS)
		if err != nil {
			skip = 0
		}
	}
	data, err := notificationService.GetNotifications(limit, skip, status, userId, topic, title)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get notification"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get notification"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("get notification success", logrus.Fields{
		"duration": duration,
	})
}

func Deletenotification(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("get notifications", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	userId := ""
	topic := ""
	title := ""
	status := "Active"
	limit := 20
	skip := 0
	var err error
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	userIdS := r.URL.Query().Get("userId")
	topicS := r.URL.Query().Get("topic")
	statusS := r.URL.Query().Get("status")
	titleS := r.URL.Query().Get("title")
	if userIdS != "" {
		userId = userIdS
	}
	if topicS != "" {
		topic = topicS
	}
	if titleS != "" {
		title = titleS
	}
	if statusS != "" {
		status = statusS
	}
	if limitS != "" {
		limit, err = strconv.Atoi(limitS)
		if err != nil {
			limit = 20
		}
	}
	if skipS != "" {
		skip, err = strconv.Atoi(skipS)
		if err != nil {
			skip = 0
		}
	}
	data, err := notificationService.DeleteNotifications(limit, skip, status, userId, topic, title)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get notification"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get notification"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("get notification success", logrus.Fields{
		"duration": duration,
	})
}

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting cart", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")
	if len(tokenString) < 2 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	claims, err := trestCommon.DecodeToken(tokenString[1])
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	var order controller.Order
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &order)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	data, err := orderService.PlaceOrder(claims["userid"].(string), tokenString[1], order)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to place order"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "unable to place order"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("order placed", logrus.Fields{
		"duration": duration,
	})
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("update order status", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")
	if len(tokenString) < 2 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	_, err := trestCommon.DecodeToken(tokenString[1])
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	var orderID = mux.Vars(r)["orderId"]
	if orderID == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set orderID"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set orderID"})
		return
	}
	var order controller.Order
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &order)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	data, err := orderService.UpdateOrder(orderID, order)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to update order status"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": err})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("order status updated", logrus.Fields{
		"duration": duration,
	})
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting orders", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")
	if len(tokenString) < 2 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	claims, err := trestCommon.DecodeToken(tokenString[1])
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	limit := 20
	skip := 0
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	if limitS != "" {
		limit, err = strconv.Atoi(limitS)
		if err != nil {
			limit = 20
		}
	}
	if skipS != "" {
		skip, err = strconv.Atoi(skipS)
		if err != nil {
			skip = 0
		}
	}
	data, err := orderService.GetOrders(claims["userid"].(string), tokenString[1], limit, skip)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get orders"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "unable to get orders"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("orders retrieved", logrus.Fields{
		"duration": duration,
	})
}

func GetLatestOrders(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting orders", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")
	if len(tokenString) < 2 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	claims, err := trestCommon.DecodeToken(tokenString[1])
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	limit := 20
	skip := 0
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	if limitS != "" {
		limit, err = strconv.Atoi(limitS)
		if err != nil {
			limit = 20
		}
	}
	if skipS != "" {
		skip, err = strconv.Atoi(skipS)
		if err != nil {
			skip = 0
		}
	}
	data, err := orderService.GetLatestOrders(claims["userid"].(string), tokenString[1], limit, skip)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get orders"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "unable to get orders"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("orders retrieved", logrus.Fields{
		"duration": duration,
	})
}

func GetOrderDetails(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("get order details", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")
	if len(tokenString) < 2 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	_, err := trestCommon.DecodeToken(tokenString[1])
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	var orderID = mux.Vars(r)["orderId"]
	if orderID == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set orderID"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set orderID"})
		return
	}
	data, err := orderService.GetOrder(orderID)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get order details"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "unable to get order details"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("order details retrieved", logrus.Fields{
		"duration": duration,
	})
}

func GetAllOrdersAdmin(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting orders", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")
	if len(tokenString) < 2 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	_, err := trestCommon.DecodeToken(tokenString[1])
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	limit := 20
	skip := 0
	status := ""
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	statusS := r.URL.Query().Get("status")
	if limitS != "" {
		limit, err = strconv.Atoi(limitS)
		if err != nil {
			limit = 20
		}
	}
	if skipS != "" {
		skip, err = strconv.Atoi(skipS)
		if err != nil {
			skip = 0
		}
	}
	if statusS != "" {
		status = statusS
	}
	data, err := orderService.GetAllOrdersAdmin(tokenString[1], status, limit, skip)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get orders"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "unable to get orders"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("orders retrieved", logrus.Fields{
		"duration": duration,
	})
}
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting orders", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")
	if len(tokenString) < 2 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	claim, err := trestCommon.DecodeToken(tokenString[1])
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	limit := 100
	skip := 0
	status := "Ordered"
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	statusS := r.URL.Query().Get("status")
	if limitS != "" {
		limit, err = strconv.Atoi(limitS)
		if err != nil {
			limit = 100
		}
	}
	if skipS != "" {
		skip, err = strconv.Atoi(skipS)
		if err != nil {
			skip = 0
		}
	}
	if statusS != "" {
		status = statusS
	}
	data, err := orderService.GetAllUsers(tokenString[1], limit, skip, claim["userid"].(string), status)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get orders"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "unable to get orders"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("orders retrieved", logrus.Fields{
		"duration": duration,
	})
}

func GetNewOrdersDelivery(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting orders", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")
	if len(tokenString) < 2 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	_, err := trestCommon.DecodeToken(tokenString[1])
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	limit := 20
	skip := 0
	latitude := 0.00
	longitude := 0.00
	latitudeS := r.URL.Query().Get("latitude")
	longitudeS := r.URL.Query().Get("longitude")
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	if latitudeS != "" {
		latitude, err = strconv.ParseFloat(latitudeS, 64)
		if err != nil {
			latitude = 0.00
		}
	}
	if longitudeS != "" {
		longitude, err = strconv.ParseFloat(longitudeS, 64)
		if err != nil {
			longitude = 0.00
		}
	}
	if limitS != "" {
		limit, err = strconv.Atoi(limitS)
		if err != nil {
			limit = 20
		}
	}
	if skipS != "" {
		skip, err = strconv.Atoi(skipS)
		if err != nil {
			skip = 0
		}
	}
	data, err := orderService.GetNewDeliveryOrders(tokenString[1], limit, skip, latitude, longitude)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get orders"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "unable to get orders"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("orders retrieved", logrus.Fields{
		"duration": duration,
	})
}

func GetOrdersDeliveryByStatus(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting orders", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")
	if len(tokenString) < 2 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	_, err := trestCommon.DecodeToken(tokenString[1])
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	limit := 20
	skip := 0
	status := "Ordered"
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	statusS := r.URL.Query().Get("status")
	if limitS != "" {
		limit, err = strconv.Atoi(limitS)
		if err != nil {
			limit = 20
		}
	}
	if skipS != "" {
		skip, err = strconv.Atoi(skipS)
		if err != nil {
			skip = 0
		}
	}
	if statusS != "" {
		status = statusS
	}
	var deliveryID = mux.Vars(r)["deliveryID"]
	if deliveryID == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get deliveryID"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get deliveryID"})
		return
	}
	data, err := orderService.GetActiveDeliveryOrders(tokenString[1], limit, skip, deliveryID, status)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get orders"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "unable to get orders"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("orders retrieved", logrus.Fields{
		"duration": duration,
	})
}

func GetOrderWithIDs(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting item", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	var order = mux.Vars(r)["userIds"]
	if order == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set item"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set item"})
		return
	}
	orders := strings.Split(order, ",")
	data, err := orderService.GetOrderWithIDs(orders)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set item"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set item"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("item updated", logrus.Fields{
		"duration": duration,
	})
}
func GetOrdersCount(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting item", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	var order = mux.Vars(r)["userIds"]
	if order == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set item"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set item"})
		return
	}
	orders := strings.Split(order, ",")
	data, err := orderService.GetOrdersCount(orders)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set item"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set item"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("item updated", logrus.Fields{
		"duration": duration,
	})
}

func GetAdminOrders(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting orders", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")
	if len(tokenString) < 2 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	_, err := trestCommon.DecodeToken(tokenString[1])
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	limit := 20
	skip := 0
	status := ""
	shopID := ""
	deliveryId := ""
	userId := ""
	sellerID := ""
	fromD := ""
	endD := ""
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	statusS := r.URL.Query().Get("status")
	shopIDS := r.URL.Query().Get("shopID")
	sellerIDS := r.URL.Query().Get("sellerID")
	deliveryIdS := r.URL.Query().Get("deliveryId")
	userIdS := r.URL.Query().Get("userId")
	fromDS := r.URL.Query().Get("from")
	endDS := r.URL.Query().Get("to")
	if limitS != "" {
		limit, err = strconv.Atoi(limitS)
		if err != nil {
			limit = 20
		}
	}
	if skipS != "" {
		skip, err = strconv.Atoi(skipS)
		if err != nil {
			skip = 0
		}
	}
	if statusS != "" {
		status = statusS
	}
	if shopIDS != "" {
		shopID = shopIDS
	}
	if sellerIDS != "" {
		sellerID = sellerIDS
	}
	if deliveryIdS != "" {
		deliveryId = deliveryIdS
	}
	if userIdS != "" {
		userId = userIdS
	}
	if fromDS != "" {
		fromD = fromDS
	}
	if endDS != "" {
		endD = endDS
	}
	data, err := orderService.GetAdminOrders(tokenString[1], limit, skip, shopID, sellerID, deliveryId, userId, status, fromD, endD)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get orders"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "unable to get orders"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("orders retrieved", logrus.Fields{
		"duration": duration,
	})
}
