package couponHandler

import (
	controller "Drop/DropCoupons/controller/coupon"

	"Drop/DropCoupons/repository/coupon"
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
	couponService = controller.NewCouponService(coupon.NewCouponRepository("coupon"))
)

func AddCoupon(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("adding coupon", logrus.Fields{
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
	var coupon controller.Coupon
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &coupon)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	data, err := couponService.AddCoupon(coupon, tokenString[1])
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to add coupon"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to add coupon"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("coupon added", logrus.Fields{
		"duration": duration,
	})
}

func GetCoupon(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("retrieving coupon", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	status := ""
	statusS := r.URL.Query().Get("status")
	if statusS != "" {
		status = statusS
	}
	code := ""
	codeS := r.URL.Query().Get("code")
	if codeS != "" {
		code = codeS
	}
	validAmt := ""
	validAmtS := r.URL.Query().Get("validamount")
	if validAmtS != "" {
		validAmt = validAmtS
	}
	maxDis := ""
	maxDisS := r.URL.Query().Get("maxdiscount")
	if maxDisS != "" {
		maxDis = maxDisS
	}
	usagePD := ""
	usagePDS := r.URL.Query().Get("usageperday")
	if usagePDS != "" {
		usagePD = usagePDS
	}
	maxUsage := ""
	maxUsageS := r.URL.Query().Get("maxusage")
	if maxUsageS != "" {
		maxUsage = maxUsageS
	}
	disPer := ""
	disPerS := r.URL.Query().Get("discount")
	if disPerS != "" {
		disPer = disPerS
	}
	cId := ""
	cIdS := r.URL.Query().Get("couponId")
	if cIdS != "" {
		cId = cIdS
	}
	data, err := couponService.GetCoupon(code, validAmt, maxDis, usagePD, maxUsage, disPer, status, cId)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to retrieve coupon"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to retrieve coupon"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("coupon retrieved", logrus.Fields{
		"duration": duration,
	})
}

func GetCoupons(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("retrieving coupons", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	status := ""
	statusS := r.URL.Query().Get("status")
	if statusS != "" {
		status = statusS
	}
	code := ""
	codeS := r.URL.Query().Get("code")
	if codeS != "" {
		code = codeS
	}
	validAmt := ""
	validAmtS := r.URL.Query().Get("validamount")
	if validAmtS != "" {
		validAmt = validAmtS
	}
	maxDis := ""
	maxDisS := r.URL.Query().Get("maxdiscount")
	if maxDisS != "" {
		maxDis = maxDisS
	}
	usagePD := ""
	usagePDS := r.URL.Query().Get("usageperday")
	if usagePDS != "" {
		usagePD = usagePDS
	}
	maxUsage := ""
	maxUsageS := r.URL.Query().Get("maxusage")
	if maxUsageS != "" {
		maxUsage = maxUsageS
	}
	disPer := ""
	disPerS := r.URL.Query().Get("discount")
	if disPerS != "" {
		disPer = disPerS
	}
	cId := ""
	cIdS := r.URL.Query().Get("couponId")
	if cIdS != "" {
		cId = cIdS
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
	data, err := couponService.GetCoupons(code, validAmt, maxDis, usagePD, maxUsage, disPer, status, cId, limit, skip)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to retrieve coupons"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to retrieve coupons"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("coupons retrieved", logrus.Fields{
		"duration": duration,
	})
}

func GetCouponsWithIDs(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting coupon", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	var coupon = mux.Vars(r)["couponIds"]
	if coupon == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get coupon"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get coupon"})
		return
	}
	coupons := strings.Split(coupon, ",")
	data, err := couponService.GetCouponWithIDs(coupons)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get coupon"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get coupon"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("coupons", logrus.Fields{
		"duration": duration,
	})
}

func UpdateCoupon(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("update coupon ", logrus.Fields{
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
	var coupon controller.Coupon
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &coupon)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to unmarshal body"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return
	}
	var couponID = mux.Vars(r)["couponId"]
	if couponID == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to update coupon "))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to update coupon "})
		return
	}
	data, err := couponService.UpdateCoupon(coupon, couponID)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to update coupon "))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to update coupon "})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("update coupon  success", logrus.Fields{
		"duration": duration,
	})
}
