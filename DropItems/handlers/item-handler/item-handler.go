package itemHandler

import (
	controller "Drop/DropItems/controller/item"

	"Drop/DropItems/repository/item"

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
	itemService = controller.NewItemService(item.NewItemRepository("item"))
)

func AddItem(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting item", logrus.Fields{
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
	var item *controller.Item
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &item)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to unmarshal body"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return
	}
	data, err := itemService.AddItem(item, claims["userid"].(string))
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

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting item", logrus.Fields{
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
	var itemID = mux.Vars(r)["itemID"]
	if itemID == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set item"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set item"})
		return
	}
	var item *controller.Item
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &item)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to unmarshal body"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return
	}
	data, err := itemService.UpdateItem(item, itemID)
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

func GetItem(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting item", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	limit := 20
	skip := 0
	var err error
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	category := ""
	categoryS := r.URL.Query().Get("category")
	if categoryS != "" {
		category = categoryS
	}
	name := ""
	nameS := r.URL.Query().Get("name")
	if nameS != "" {
		name = nameS
	}
	typee := ""
	typeeS := r.URL.Query().Get("type")
	if typeeS != "" {
		typee = typeeS
	}
	sellerId := ""
	sellerIdS := r.URL.Query().Get("sellerId")
	if sellerIdS != "" {
		sellerId = sellerIdS
	}
	featured := ""
	featuredS := r.URL.Query().Get("featured")
	if featuredS != "" {
		featured = featuredS
	}
	search := ""
	searchS := r.URL.Query().Get("search")
	if searchS != "" {
		search = searchS
	}
	shopID := ""
	shopIDS := r.URL.Query().Get("shopID")
	if shopIDS != "" {
		shopID = shopIDS
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
	data, err := itemService.GetItem(shopID, category, name, typee, sellerId, featured, search, limit, skip)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get item"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get item"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("item retrieved", logrus.Fields{
		"duration": duration,
	})
}

func GetMyItem(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting item", logrus.Fields{
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
	shopID := r.URL.Query().Get("shopID")

	data, err := itemService.GetSellerItem(claims["userid"].(string), shopID, limit, skip)
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
func GetFullItem(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting item", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	var item = mux.Vars(r)["itemId"]
	if item == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set item"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set item"})
		return
	}
	data, err := itemService.GetItemUsingID(item)
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

func GetItemWithIDs(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting item", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	var item = mux.Vars(r)["itemIds"]
	if item == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set item"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set item"})
		return
	}
	items := strings.Split(item, ",")
	data, err := itemService.GetItemWithIDs(items)
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
func PreSignedUrl(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("Account Presigned url", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	name := mux.Vars(r)["filename"]
	if name == "" {
		trestCommon.ECLog1(errors.New("failed to get file name and path"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "unauthorized"})
		return
	}
	_, err := trestCommon.DecodeToken(strings.Split(r.Header.Get("Authorization"), " ")[1])
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "unauthorized"})
		return
	}
	file_path := make([]string, 0)
	preSignedURLs := make([]string, 0)
	names := strings.Split(name, ",")
	for i := range names {
		currentTime := strconv.FormatInt(time.Now().Unix(), 10)
		newName := currentTime + names[i]
		preSignedURL, _ := trestCommon.PreSignedUrl(newName, "images")
		preSignedURLs = append(preSignedURLs, preSignedURL)
		file := "images/" + newName
		file_path = append(file_path, file)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "preSignedURLs": preSignedURLs, "file_paths": file_path})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("Account Presigned url", logrus.Fields{"duration": duration})
}

func GetFeaturedItem(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting item", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	limit := 20
	skip := 0
	var err error
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	category := r.URL.Query().Get("category")
	subCategory := r.URL.Query().Get("shopType")
	shopID := r.URL.Query().Get("shopID")
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
	data, err := itemService.GetFeaturedItem(shopID, category, subCategory, limit, skip)
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
func GetPopularItem(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting item", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	limit := 20
	skip := 0
	var err error
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	category := r.URL.Query().Get("category")
	subCategory := r.URL.Query().Get("shopType")
	shopID := r.URL.Query().Get("shopID")
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
	data, err := itemService.GetPopularItem(shopID, category, subCategory, limit, skip)
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
func GetShopFeaturedItem(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting item", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	limit := 20
	skip := 0
	var err error
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	category := r.URL.Query().Get("category")
	subCategory := r.URL.Query().Get("shopType")
	shopID := r.URL.Query().Get("shopID")
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
	data, err := itemService.GetShopFeaturedItem(shopID, category, subCategory, limit, skip)
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

func GetTopRatedItems(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting item", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	limit := 20
	skip := 0
	var err error
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	category := r.URL.Query().Get("category")
	subCategory := r.URL.Query().Get("shopType")
	shopID := r.URL.Query().Get("shopID")
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
	data, err := itemService.GetFeaturedItem(shopID, category, subCategory, limit, skip)
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

func GetItemStruc(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting item", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	limit := 20
	skip := 0
	var err error
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	category := ""
	categoryS := r.URL.Query().Get("category")
	if categoryS != "" {
		category = categoryS
	}
	name := ""
	nameS := r.URL.Query().Get("name")
	if nameS != "" {
		name = nameS
	}
	typee := ""
	typeeS := r.URL.Query().Get("foodtype")
	if typeeS != "" {
		typee = typeeS
	}
	sellerId := ""
	sellerIdS := r.URL.Query().Get("sellerId")
	if sellerIdS != "" {
		sellerId = sellerIdS
	}
	featured := ""
	featuredS := r.URL.Query().Get("featured")
	if featuredS != "" {
		featured = featuredS
	}
	search := ""
	searchS := r.URL.Query().Get("search")
	if searchS != "" {
		search = searchS
	}
	shopID := ""
	shopIDS := r.URL.Query().Get("shopID")
	if shopIDS != "" {
		shopID = shopIDS
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
	stypee := ""
	stypeeS := r.URL.Query().Get("type")
	if stypeeS != "" {
		stypee = stypeeS
	}
	deal := r.URL.Query().Get("deal")
	data, err := itemService.GetItemCategoryStructured(shopID, category, deal, name, typee, sellerId, search, featured, stypee, limit, skip)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get item"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get item"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("item retrieved", logrus.Fields{
		"duration": duration,
	})
}
