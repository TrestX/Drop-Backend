package favouriteHandler

import (
	controller "Drop/DropFavourite/controller/favourite"

	"Drop/DropFavourite/repository/favourite"

	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aekam27/trestCommon"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	favouriteService = controller.NewFavouriteService(favourite.NewFavouriteRepository("favourite"))
)

func AddFavourite(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting favourite", logrus.Fields{
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
	var itemID = mux.Vars(r)["itemId"]
	if itemID == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set address"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set address"})
		return
	}
	data, err := favouriteService.AddFavourite(itemID, claims["userid"].(string))
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set favourite"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set favourite"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("favourite updated", logrus.Fields{
		"duration": duration,
	})
}

func DeleteFavourite(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting favourite", logrus.Fields{
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
	var itemID = mux.Vars(r)["itemId"]
	if itemID == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set address"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set address"})
		return
	}
	err = favouriteService.DeleteFavourite(itemID, claims["userid"].(string))
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set favourite"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set favourite"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": "Successfully deleted"})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("favourite updated", logrus.Fields{
		"duration": duration,
	})
}

func GetFavourite(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting favourite", logrus.Fields{
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
	data, err := favouriteService.GetFavourite(claims["userid"].(string), limit, skip)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set favourite"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set favourite"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("favourite updated", logrus.Fields{
		"duration": duration,
	})
}

func GetFavouriteShops(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting favourite", logrus.Fields{
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
	data, err := favouriteService.GetFavouriteShop(claims["userid"].(string), limit, skip)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set favourite"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set favourite"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("favourite updated", logrus.Fields{
		"duration": duration,
	})
}
