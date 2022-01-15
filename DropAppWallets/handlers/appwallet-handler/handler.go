package apptransactionHandler

import (
	controller "Drop/DropAppWallets/controller/appwallet"
	"Drop/DropAppWallets/repository/appwallet"
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
	apptransactionService = controller.NewAppWalletService(appwallet.NewAppWalletRepository("apptransactions"))
)

func AddAppTransaction(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("adding to apptransaction", logrus.Fields{
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
	var appWallet controller.AppWalletSchema
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &appWallet)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	data, err := apptransactionService.AddTransaction(appWallet, tokenString[1])
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to add to apptransaction"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to add to apptransaction"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("apptransaction updated", logrus.Fields{
		"duration": duration,
	})
}

func Updateapptransaction(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("updating apptransaction", logrus.Fields{
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
	var apptransactionID = mux.Vars(r)["apptransactionID"]
	if apptransactionID == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set apptransactionID"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set apptransactionID"})
		return
	}
	var appWallet controller.AppWalletSchema
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &appWallet)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	data, err := apptransactionService.UpdateTrans(apptransactionID, appWallet.Status)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set apptransaction"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set apptransaction"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("apptransaction updated", logrus.Fields{
		"duration": duration,
	})
}

func GetAppTransaction(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("retrieving apptransaction", logrus.Fields{
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
	status := ""
	statusS := r.URL.Query().Get("status")
	if statusS != "" {
		status = statusS
	}
	entity := ""
	entityS := r.URL.Query().Get("entity")
	if entityS != "" {
		entity = entityS
	}
	entityid := ""
	entityIdS := r.URL.Query().Get("entityId")
	if entityIdS != "" {
		entityid = entityIdS
	}
	orderid := ""
	orderidS := r.URL.Query().Get("orderid")
	if orderidS != "" {
		orderid = orderidS
	}
	transactionId := ""
	transactionIdS := r.URL.Query().Get("transactionId")
	if transactionIdS != "" {
		transactionId = transactionIdS
	}
	data, err := apptransactionService.GetTransaction(transactionId, status, entity, entityid, orderid)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to retrieve apptransaction"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to retrieve apptransaction"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("apptransaction retrieved", logrus.Fields{
		"duration": duration,
	})
}

func GetAppTransactions(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("retrieving apptransaction", logrus.Fields{
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
	status := ""
	statusS := r.URL.Query().Get("status")
	if statusS != "" {
		status = statusS
	}
	entity := ""
	entityS := r.URL.Query().Get("entity")
	if entityS != "" {
		entity = entityS
	}
	entityid := ""
	entityIdS := r.URL.Query().Get("entityId")
	if entityIdS != "" {
		entityid = entityIdS
	}
	orderid := ""
	orderidS := r.URL.Query().Get("orderid")
	if orderidS != "" {
		orderid = orderidS
	}
	transactionId := ""
	transactionIdS := r.URL.Query().Get("transactionId")
	if transactionIdS != "" {
		transactionId = transactionIdS
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
	data, err := apptransactionService.GetTransactions(transactionId, status, entity, entityid, orderid, limit, skip)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to retrieve apptransaction"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to retrieve apptransaction"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("apptransaction retrieved", logrus.Fields{
		"duration": duration,
	})
}

func GetDeliveryAppTransactions(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("retrieving apptransaction", logrus.Fields{
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
	status := ""
	statusS := r.URL.Query().Get("status")
	if statusS != "" {
		status = statusS
	}
	entityid := ""
	entityIdS := r.URL.Query().Get("entityId")
	if entityIdS != "" {
		entityid = entityIdS
	}
	orderid := ""
	orderidS := r.URL.Query().Get("orderid")
	if orderidS != "" {
		orderid = orderidS
	}
	transactionId := ""
	transactionIdS := r.URL.Query().Get("transactionId")
	if transactionIdS != "" {
		transactionId = transactionIdS
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
	fromD := ""
	endD := ""
	fromDS := r.URL.Query().Get("from")
	endDS := r.URL.Query().Get("to")
	if fromDS != "" {
		fromD = fromDS
	}
	if endDS != "" {
		endD = endDS
	}
	typeD := r.URL.Query().Get("daytype")

	data, total, settled, unsettled, totaltip, settledtip, unsettledtip, earningdiff, difftyp, err := apptransactionService.GetDeliveryPersonBalance(transactionId, status, entityid, orderid, tokenString[1], fromD, endD, typeD, limit, skip)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to retrieve apptransaction"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to retrieve apptransaction"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data, "earningdiff": earningdiff, "difftyp": difftyp, "total": total, "settled": settled, "unsettled": unsettled, "totaltip": totaltip, "settledtip": settledtip, "unsettledtip": unsettledtip})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("apptransaction retrieved", logrus.Fields{
		"duration": duration,
	})
}

func GetSellerAppTransactions(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("retrieving apptransaction", logrus.Fields{
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
	status := ""
	statusS := r.URL.Query().Get("status")
	if statusS != "" {
		status = statusS
	}
	entityid := ""
	entityIdS := r.URL.Query().Get("entityId")
	if entityIdS != "" {
		entityid = entityIdS
	}
	orderid := ""
	orderidS := r.URL.Query().Get("orderid")
	if orderidS != "" {
		orderid = orderidS
	}
	transactionId := ""
	transactionIdS := r.URL.Query().Get("transactionId")
	if transactionIdS != "" {
		transactionId = transactionIdS
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
	fromD := ""
	endD := ""
	fromDS := r.URL.Query().Get("from")
	endDS := r.URL.Query().Get("to")
	if fromDS != "" {
		fromD = fromDS
	}
	if endDS != "" {
		endD = endDS
	}
	data, total, settled, unsettled, err := apptransactionService.GetSellerPersonBalance(transactionId, status, entityid, orderid, fromD, endD, limit, skip)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to retrieve apptransaction"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to retrieve apptransaction"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data, "total": total, "settled": settled, "unsettled": unsettled})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("apptransaction retrieved", logrus.Fields{
		"duration": duration,
	})
}

func GetTotalAppEarning(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("retrieving apptransaction", logrus.Fields{
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
	data, err := apptransactionService.GetAppEarning()
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to retrieve apptransaction"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to retrieve apptransaction"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("apptransaction retrieved", logrus.Fields{
		"duration": duration,
	})
}

func GetTotalTransAmt(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("retrieving apptransaction", logrus.Fields{
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
	data, err := apptransactionService.GetTotalTransactions()
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to retrieve apptransaction"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to retrieve apptransaction"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("apptransaction retrieved", logrus.Fields{
		"duration": duration,
	})
}

func GetSellerTransactions(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("retrieving apptransaction", logrus.Fields{
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
	fromD := ""
	endD := ""
	fromDS := r.URL.Query().Get("from")
	endDS := r.URL.Query().Get("to")
	if fromDS != "" {
		fromD = fromDS
	}
	if endDS != "" {
		endD = endDS
	}
	data, err := apptransactionService.GetSellerPersonS(tokenString[1], fromD, endD, limit, skip)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to retrieve apptransaction"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to retrieve apptransaction"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("apptransaction retrieved", logrus.Fields{
		"duration": duration,
	})
}

func UpdateSellerapptransactionPer(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("updating apptransaction per", logrus.Fields{
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
	per := ""
	sId := ""
	c := ""
	perS := r.URL.Query().Get("per")
	sIdS := r.URL.Query().Get("sId")
	cS := r.URL.Query().Get("c")
	if sIdS != "" {
		sId = sIdS
	}
	if perS != "" {
		per = perS
	}
	if cS != "" {
		c = cS
	}

	data, err := apptransactionService.UpdateSellerTransPer(sId, per, c)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set apptransaction per"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set apptransaction per"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("apptransaction updated per", logrus.Fields{
		"duration": duration,
	})
}

func GetDeliveryTransactions(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("retrieving apptransaction", logrus.Fields{
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
	fromD := ""
	endD := ""
	fromDS := r.URL.Query().Get("from")
	endDS := r.URL.Query().Get("to")
	if fromDS != "" {
		fromD = fromDS
	}
	if endDS != "" {
		endD = endDS
	}
	data, err := apptransactionService.GetDeliveryPersonS(tokenString[1], fromD, endD, limit, skip)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to retrieve apptransaction"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to retrieve apptransaction"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("apptransaction retrieved", logrus.Fields{
		"duration": duration,
	})
}

func GetSellerShopsTransactions(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("retrieving apptransaction", logrus.Fields{
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
	sellerId := ""
	sellerIdS := r.URL.Query().Get("sellerId")
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	if limitS != "" {
		limit, err = strconv.Atoi(limitS)
		if err != nil {
			limit = 20
		}
	}
	if sellerId != "" {
		sellerId = sellerIdS
	}
	if skipS != "" {
		skip, err = strconv.Atoi(skipS)
		if err != nil {
			skip = 0
		}
	}
	fromD := ""
	endD := ""
	fromDS := r.URL.Query().Get("from")
	endDS := r.URL.Query().Get("to")
	if fromDS != "" {
		fromD = fromDS
	}
	if endDS != "" {
		endD = endDS
	}
	data, err := apptransactionService.GetSellerPersonShops(tokenString[1], sellerId, fromD, endD, limit, skip)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to retrieve apptransaction"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to retrieve apptransaction"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("apptransaction retrieved", logrus.Fields{
		"duration": duration,
	})
}

func UpdateShopPHTransactions(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("retrieving apptransaction", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")
	var err error
	var appWallet controller.SettingPaymentHistory
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &appWallet)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	data, err := apptransactionService.UpdateShopPaymentHistory(appWallet.ID, appWallet.Name, appWallet.Email, appWallet.PhoneNo, appWallet.DoneBy, tokenString[1], appWallet.Amount)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to retrieve apptransaction"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to retrieve apptransaction"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("apptransaction retrieved", logrus.Fields{
		"duration": duration,
	})
}

func UpdateSellerPHTransactions(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("retrieving apptransaction", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")
	var err error
	var appWallet controller.SettingPaymentHistory
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &appWallet)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	data, err := apptransactionService.UpdateSellerPaymentHistory(appWallet.ID, appWallet.Name, appWallet.Email, appWallet.PhoneNo, appWallet.DoneBy, tokenString[1], appWallet.Amount)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to retrieve apptransaction"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to retrieve apptransaction"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("apptransaction retrieved", logrus.Fields{
		"duration": duration,
	})
}
func UpdateDeliveryPHTransactions(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("retrieving apptransaction", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")
	var err error
	var appWallet controller.SettingPaymentHistory
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &appWallet)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	data, err := apptransactionService.UpdateDeliveryPaymentHistory("", appWallet.Name, appWallet.Email, appWallet.PhoneNo, appWallet.DoneBy, appWallet.Type, tokenString[1], appWallet.Amount)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to retrieve apptransaction"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to retrieve apptransaction"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("apptransaction retrieved", logrus.Fields{
		"duration": duration,
	})
}
func GetSPaymentsHistory(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("retrieving apptransaction", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_ = strings.Split(r.Header.Get("Authorization"), " ")
	limit := 20
	skip := 0
	var err error
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
	data, err := apptransactionService.GetPAymentsHistory(limit, skip)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to retrieve apptransaction"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to retrieve apptransaction"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("apptransaction retrieved", logrus.Fields{
		"duration": duration,
	})
}
