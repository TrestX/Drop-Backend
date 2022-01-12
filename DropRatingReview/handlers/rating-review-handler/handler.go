package ratingReviewHandler

import (
	controller "Drop/DropRatingReview/controller/ratingreview"
	"Drop/DropRatingReview/repository/ratingReview"
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
)

var (
	ratingReviewService = controller.NewRatingReviewService(ratingReview.NewRatingReviewRepository("ratingReview"))
)

func AddReviewRating(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("add ratings/review", logrus.Fields{
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
	var ratinReviewParams controller.RatingReviewSchema
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &ratinReviewParams)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	var entityID = mux.Vars(r)["entityID"]
	if entityID == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set entityID"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set entityID"})
		return
	}
	data, err := ratingReviewService.AddRatingReview(claims["userid"].(string), entityID, tokenString[1], ratinReviewParams)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to add ratings/review"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "unable to add ratings/review"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("add ratings/review added", logrus.Fields{
		"duration": duration,
	})
}

func UpdateReviewRating(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("add ratings/review", logrus.Fields{
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
	var ratinReviewParams controller.RatingReviewSchema
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &ratinReviewParams)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	var ratingReviewID = mux.Vars(r)["ratingReviewID"]
	if ratingReviewID == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set ratingReviewID"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set ratingReviewID"})
		return
	}
	data, err := ratingReviewService.UpdateRatingReview(ratingReviewID, ratinReviewParams)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to add ratings/review"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "unable to add ratings/review"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("add ratings/review added", logrus.Fields{
		"duration": duration,
	})
}

func GetReviewsAndRatings(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting rating/reviews Details", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	limit := 20
	skip := 0
	entityId := ""
	userId := ""
	var err error
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	entityIdS := r.URL.Query().Get("entityId")
	userIdS := r.URL.Query().Get("userId")
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
	if entityIdS != "" {
		entityId = entityIdS
	}
	if userIdS != "" {
		userId = userIdS
	}
	data, err := ratingReviewService.GetRatingsReview(userId, entityId, limit, skip)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get rating/reviews"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "unable to get rating/reviews"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("successfully retrieved rating/reviews details", logrus.Fields{
		"duration": duration,
	})
}

func GetReviewRatingsWithIDs(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting users", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	var users = mux.Vars(r)["reviewIds"]
	if users == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get users"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get users"})
		return
	}
	user := strings.Split(users, ",")
	data, err := ratingReviewService.GetReviewRatingsWithIDs(user)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get users"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get users"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("users retrieved", logrus.Fields{
		"duration": duration,
	})
}
