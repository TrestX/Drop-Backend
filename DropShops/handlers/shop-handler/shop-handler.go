package shopHandler

import (
	controller "Drop/DropShop/controller/shop"
	"Drop/DropShop/repository/shop"

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
	shopService = controller.NewShopService(shop.NewShopRepository("shop"))
)

func AddShop(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting shop", logrus.Fields{
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
	var shop *controller.Shop
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &shop)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to unmarshal body"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return
	}
	data, err := shopService.AddShop(shop, claims["userid"].(string))
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set address"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set address"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("address updated", logrus.Fields{
		"duration": duration,
	})
}

func UpdateShop(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting shop", logrus.Fields{
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
	var shop *controller.Shop
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &shop)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to unmarshal body"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return
	}
	var shopID = mux.Vars(r)["shopId"]
	if shopID == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set shop"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set shop"})
		return
	}
	data, err := shopService.UpdateShop(shop, shopID)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set shop"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set shop"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("shop updated", logrus.Fields{
		"duration": duration,
	})
}

func UpdatePrimaryShop(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting shop", logrus.Fields{
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
	var shop = mux.Vars(r)["shopId"]
	if shop == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set shop"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set address"})
		return
	}
	data, err := shopService.PrimaryShop(shop, claims["userid"].(string))
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set shop"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set shop"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("shop updated", logrus.Fields{
		"duration": duration,
	})
}
func GetPrimaryShop(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting shop", logrus.Fields{
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

	data, err := shopService.GetPrimaryShop(claims["userid"].(string))
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get shop"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get shop"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("shop", logrus.Fields{
		"duration": duration,
	})
}

func GetShop(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting shop", logrus.Fields{
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
	data, err := shopService.GetShop(claims["userid"].(string), limit, skip)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get shop"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get shop"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("shop", logrus.Fields{
		"duration": duration,
	})
}

func GetFullShops(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting address", logrus.Fields{
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
	var shop = mux.Vars(r)["shopId"]
	if shop == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set shop"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set address"})
		return
	}
	data, err := shopService.GetShopUsingID(shop, claims["userid"].(string))
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set address"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set address"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("address updated", logrus.Fields{
		"duration": duration,
	})
}

func GetFeaturedShops(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting shop", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	limit := 20
	skip := 0
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	var err error
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
	data, err := shopService.GetFeaturedShop(limit, skip)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get shop"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get shop"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("shop", logrus.Fields{
		"duration": duration,
	})
}

func GetShopsByType(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting shop", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	limit := 20
	skip := 0
	typ := ""
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	typeS := r.URL.Query().Get("type")
	var err error
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
	if typeS != "" {
		typ = typeS
	}
	data, err := shopService.SearchShopByType(typ, limit, skip)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get shop"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get shop"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("shop", logrus.Fields{
		"duration": duration,
	})
}

func AddShopAdmin(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting shop", logrus.Fields{
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
	var shop *controller.Shop
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &shop)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to unmarshal body"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return
	}
	data, err := shopService.AddShopAdmin(shop)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set address"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set address"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("address updated", logrus.Fields{
		"duration": duration,
	})
}

func GetShopAdmin(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting shop", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	limit := 20
	skip := 0
	sellerId := ""
	sType := ""
	status := ""
	featured := ""
	deal := ""
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	sellerIdS := r.URL.Query().Get("sellerId")
	sTypeS := r.URL.Query().Get("type")
	statusS := r.URL.Query().Get("status")
	featuredS := r.URL.Query().Get("featured")
	dealS := r.URL.Query().Get("deal")
	rating := r.URL.Query().Get("rating")
	priceu := r.URL.Query().Get("lowerprice")
	pricel := r.URL.Query().Get("upperprice")
	lowest := r.URL.Query().Get("lowestDP")
	lat := 0.0
	long := 0.0

	if statusS != "" {
		status = statusS
	}
	if r.URL.Query().Get("lat") != "" {
		latt, err := strconv.Atoi(r.URL.Query().Get("lat"))
		lat = float64(latt)
		if err != nil {
			lat = 0.0
		}
	}
	if r.URL.Query().Get("long") != "" {
		longg, err := strconv.Atoi(r.URL.Query().Get("long"))
		long = float64(longg)
		if err != nil {
			long = 0.0
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
	if sellerIdS != "" {
		sellerId = sellerIdS
	}
	if dealS != "" {
		deal = dealS
	}
	if sTypeS != "" {
		sType = sTypeS
	}
	if featuredS != "" {
		featured = featuredS
	}
	pickup := r.URL.Query().Get("pickup")
	data, err := shopService.GetShopAdmin(limit, skip, sellerId, sType, status, featured, deal, rating, priceu, pricel, lowest, pickup, lat, long)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get shop"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get shop"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("shop", logrus.Fields{
		"duration": duration,
	})
}

func GetNearestShopAdmin(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting shop", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	limit := 20
	skip := 0
	sellerId := ""
	sType := ""
	status := ""
	latitude := 0.00
	longitude := 0.00
	latitudeS := r.URL.Query().Get("latitude")
	longitudeS := r.URL.Query().Get("longitude")
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	sellerIdS := r.URL.Query().Get("sellerId")
	sTypeS := r.URL.Query().Get("type")
	statusS := r.URL.Query().Get("status")
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
	if sellerIdS != "" {
		sellerId = sellerIdS
	}
	if sTypeS != "" {
		sType = sTypeS
	}
	data, err := shopService.GetNearestShopAdmin(limit, skip, sellerId, sType, status, latitude, longitude)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get shop"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get shop"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("shop", logrus.Fields{
		"duration": duration,
	})
}

func GetTopRatedShopAdmin(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting shop", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	limit := 20
	skip := 0
	sellerId := ""
	sType := ""
	status := ""
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	sellerIdS := r.URL.Query().Get("sellerId")
	sTypeS := r.URL.Query().Get("type")
	statusS := r.URL.Query().Get("status")
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
	if sellerIdS != "" {
		sellerId = sellerIdS
	}
	if sTypeS != "" {
		sType = sTypeS
	}
	data, err := shopService.GetTopRatedShopAdmin(limit, skip, sellerId, sType, status)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get shop"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get shop"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("shop", logrus.Fields{
		"duration": duration,
	})
}
func GetShopsWithIDs(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting users", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	var users = mux.Vars(r)["shopIds"]
	if users == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get users"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get users"})
		return
	}
	user := strings.Split(users, ",")
	data, err := shopService.GetAdminUsersWithIDs(user)
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
