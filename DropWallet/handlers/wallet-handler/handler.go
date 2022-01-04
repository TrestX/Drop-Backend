package walletHandler

import (
	controller "Drop/DropWallet/controller/wallet"
	"Drop/DropWallet/repository/wallet"
	"io/ioutil"

	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/aekam27/trestCommon"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	walletService = controller.NewWalletService(wallet.NewWalletRepository("wallet"))
)

func UpdateWallet(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("cahnging to wallet", logrus.Fields{
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
	var wallet controller.Wallet
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &wallet)
	if wallet.UserID == "" {
		wallet.UserID = claims["userid"].(string)
	}
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	data, err := walletService.AddWalletTransaction(wallet, tokenString[1])
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to change to wallet"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to change to wallet"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("wallet updated", logrus.Fields{
		"duration": duration,
	})
}

func GetWallet(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("retrieving wallet", logrus.Fields{
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
	status := ""
	statusS := r.URL.Query().Get("status")
	if statusS != "" {
		status = statusS
	}
	data, err := walletService.GetWallet(claims["userid"].(string), status)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to retrieve wallet"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to retrieve wallet"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("wallet retrieved", logrus.Fields{
		"duration": duration,
	})
}

func GetWalletByUserId(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("retrieving wallet", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	var userid = mux.Vars(r)["userId"]
	if userid == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get wallet"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get wallet"})
		return
	}
	data, err := walletService.GetWalletByUserId(userid)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to retrieve wallet"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to retrieve wallet"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("wallet retrieved", logrus.Fields{
		"duration": duration,
	})
}

func GetWalletWithIDs(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting wallet", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	var wallet = mux.Vars(r)["walletIds"]
	if wallet == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get wallet"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get wallet"})
		return
	}
	wallets := strings.Split(wallet, ",")
	data, err := walletService.GetWalletWithIDs(wallets)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get wallet"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get wallet"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("wallets", logrus.Fields{
		"duration": duration,
	})
}
