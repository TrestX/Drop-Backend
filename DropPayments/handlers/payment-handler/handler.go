package paymentHandler

import (
	controller "Drop/DropPayments/controller/payment"
	"Drop/DropPayments/repository/payment"
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
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	paymentService = controller.NewPaymentService(payment.NewPaymentRepository("payments"))
)

type PaymentStatus struct {
	Status string `json:"status"`
}

func GetPublishableKey(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting Publishable Key", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	publishableKey := viper.GetString("stripe.publishablekey")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": publishableKey})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("publishabe key retrieved", logrus.Fields{
		"duration": duration,
	})
}

func CreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("creating payment intent", logrus.Fields{
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
	var paymentParams controller.PaymentIntentSchema
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &paymentParams)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	data, paymentId, err := paymentService.CreatePaymentIntent(claims["userid"].(string), tokenString[1], paymentParams)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to create payment intent"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "unable to create payment intent"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data, "paymentId": paymentId})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("payment intent created", logrus.Fields{
		"duration": duration,
	})
}

func UpdatePaymentStatus(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("updating payment status", logrus.Fields{
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
	var paymentID = mux.Vars(r)["paymentID"]
	if paymentID == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set paymentID"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set paymentID"})
		return
	}
	var paymentStatus PaymentStatus
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &paymentStatus)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	data, err := paymentService.UpdatePaymentStatus(claims["userid"].(string), paymentID, paymentStatus.Status)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "Unable to update payment status"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to update payment status"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("payment status updated successfully", logrus.Fields{
		"duration": duration,
	})
}

func GetPaymentSDetails(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting Payments Details", logrus.Fields{
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
	data, err := paymentService.GetPaymentsDetails(claims["userid"].(string), status, limit, skip)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get payments details for user"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "unable to get payments details for user"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("successfully retrieved payments details", logrus.Fields{
		"duration": duration,
	})
}

func GetPaymentDetails(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting Payment Details", logrus.Fields{
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
	var paymentID = mux.Vars(r)["paymentID"]
	if paymentID == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set paymentID"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set paymentID"})
		return
	}
	data, err := paymentService.GetPaymentDetails(claims["userid"].(string), paymentID)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get payment details for user"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "unable to get payment details for user"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("successfully retrieved payment details", logrus.Fields{
		"duration": duration,
	})
}

func GetPaymentsWithIDs(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting item", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	var payment = mux.Vars(r)["paymentIds"]
	if payment == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set item"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set item"})
		return
	}
	paymentIds := strings.Split(payment, ",")
	data, err := paymentService.GetPaymentWithIDs(paymentIds)
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
func GetAdminPaymentSDetails(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting Payments Details", logrus.Fields{
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
	user := ""
	seller := ""
	shop := ""
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	statusS := r.URL.Query().Get("status")
	userS := r.URL.Query().Get("user")
	sellerS := r.URL.Query().Get("seller")
	shopS := r.URL.Query().Get("shop")
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
	if sellerS != "" {
		seller = sellerS
	}
	if shopS != "" {
		shop = shopS
	}
	if userS != "" {
		user = userS
	}
	data, err := paymentService.GetAdminPaymentDetails(status, user, tokenString[1], seller, shop, limit, skip)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get payments details for user"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "unable to get payments details for user"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("successfully retrieved payments details", logrus.Fields{
		"duration": duration,
	})
}

func GetPaymentSuccessDetails(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting Payment Details", logrus.Fields{
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
	var paymentID = mux.Vars(r)["paymentID"]
	if paymentID == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set paymentID"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set paymentID"})
		return
	}
	data, err := paymentService.UpdatePaymentStatusSuccess(claims["userid"].(string), paymentID)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get payment details for user"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "unable to get payment details for user"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("successfully retrieved payment details", logrus.Fields{
		"duration": duration,
	})
}
