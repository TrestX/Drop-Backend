package addressHandler

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

	controller "Drop/DropDeliveryTracking/controller/tracking"
	"Drop/DropDeliveryTracking/repository/tracking"

)

var (
	trackingService = controller.NewTrackingService(tracking.NewTrackingRepository("tracking"))
)

func AddTracking(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting delivery tracking", logrus.Fields{
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
	var tracking *controller.Tracking
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &tracking)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to unmarshal body"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return
	}
	var deliveryID = mux.Vars(r)["deliveryId"]
	if deliveryID == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set tracking"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set tracking"})
		return
	}
	data, err := trackingService.AddLocation(tracking, claims["userid"].(string))
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set tracking"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set tracking"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("tracking added", logrus.Fields{
		"duration": duration,
	})
}

func UpdateTracking(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("update tracking location", logrus.Fields{
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
	var tracking *controller.Tracking
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &tracking)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to unmarshal body"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return
	}
	var trackingID = mux.Vars(r)["trackingId"]
	if trackingID == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to update tracking location"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to update tracking location"})
		return
	}
	data, err := trackingService.UpdateLocation(tracking, claims["userid"].(string))
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to update tracking location"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to update tracking location"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("update tracking location success", logrus.Fields{
		"duration": duration,
	})
}

func GetTracking(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("get all trackings", logrus.Fields{
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
	data, err := trackingService.GetAllTrackingDetails(tokenString[1], limit, skip)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get all trackings"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get all trackings"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("get all trackings success", logrus.Fields{
		"duration": duration,
	})
}

func GetByTrackingID(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("get tracking by tracking id", logrus.Fields{
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
	var trackingId = mux.Vars(r)["trackingId"]
	if trackingId == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get tracking by tracking id"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get tracking by tracking id"})
		return
	}
	data, err := trackingService.GetTrackingByTrackingID(trackingId)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get tracking by tracking id"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get tracking by tracking id"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("get tracking by tracking id success", logrus.Fields{
		"duration": duration,
	})
}

func GetNearDeliveryPerson(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting delivery person", logrus.Fields{
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
	data, err := trackingService.GetAllDeliveryPersobOfAddressDetails(tokenString[1], limit, skip, latitude, longitude)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get delivery persons"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "unable to get delivery person"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("delivery person retrieved", logrus.Fields{
		"duration": duration,
	})
}
