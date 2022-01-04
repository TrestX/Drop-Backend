package paymentHandler

import (
	controller "Nailzee/NailzeePayments/controller/payment"
	"Nailzee/NailzeePayments/repository/payment"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/aekam27/trestCommon"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	paymentService = controller.NewPaymentService(payment.NewPaymentRepository("nailzeepayments"))
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
	data, data2, err := paymentService.CreatePaymentIntent(claims["userid"].(string), tokenString[1], paymentParams)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to create payment intent"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": err, "clientSecret": data, "intent": data2})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "clientSecret": data, "intent": data2})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("payment intent created", logrus.Fields{
		"duration": duration,
	})
}

func CreateNewPaymentIntent(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("creating new payment intent", logrus.Fields{
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
	data, data2, err := paymentService.CreateNewPaymentIntent(claims["userid"].(string), tokenString[1], paymentParams)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to create new payment intent"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": err, "clientSecret": data, "intent": data2})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "clientSecret": data, "intent": data2})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("new payment intent created", logrus.Fields{
		"duration": duration,
	})
}

func ConfirmPaymentIntent(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("creating new payment intent", logrus.Fields{
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
	data, err := paymentService.ConfirmPaymentIntent(claims["userid"].(string), tokenString[1], paymentParams)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to create new payment intent"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": err, "data": data})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("new payment intent created", logrus.Fields{
		"duration": duration,
	})
}
