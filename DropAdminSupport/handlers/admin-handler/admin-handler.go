package adminHandler

import (
	controller "Drop/DropAdminSupport/controller/adminsupp"
	"Drop/DropAdminSupport/repository/admin"

	"github.com/aekam27/trestCommon"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	adminacctService = controller.NewAdminService(admin.NewAdminRepository("admin"))
)

func AddAcct(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting ", logrus.Fields{
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
	var acct *controller.Acct
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &acct)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to unmarshal body"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return
	}
	data, err := adminacctService.AddAcct(acct)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set "))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set "})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap(" added", logrus.Fields{
		"duration": duration,
	})
}

func UpdateAcct(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("update ", logrus.Fields{
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
	var acct *controller.Acct
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &acct)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to unmarshal body"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return
	}
	var ID = mux.Vars(r)["Id"]
	if ID == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to update "))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to update "})
		return
	}
	data, err := adminacctService.UpdateAcctStatus(acct, ID)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to update"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to update"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("update  success", logrus.Fields{
		"duration": duration,
	})
}

func GetAllAcct(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("get all", logrus.Fields{
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
	data, err := adminacctService.GetAllAcct(limit, skip)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get all "))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get all "})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("get all  success", logrus.Fields{
		"duration": duration,
	})
}

func GetAllShopAcct(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("get all", logrus.Fields{
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
	data, err := adminacctService.GetAllShopAcct(limit, skip)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get all "))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get all "})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("get all  success", logrus.Fields{
		"duration": duration,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("login", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	user, err := GetCredentials(r)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to parse credentials"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return

	}
	data, err := adminacctService.Login(user)
	if err != nil {
		if err.Error() == "user not verified" {
			trestCommon.ECLog1(errors.Wrapf(err, "unable to login"))
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Email not Verified"})
			return
		}
		trestCommon.ECLog1(errors.Wrapf(err, "unable to login"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "invalid credentials"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "token": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("login successfull", logrus.Fields{"duration": duration})
}

func GetCredentials(r *http.Request) (controller.Credentials, error) {
	var user controller.Credentials

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		return user, err
	}
	user.Email = strings.TrimSpace(user.Email)
	return user, err
}
