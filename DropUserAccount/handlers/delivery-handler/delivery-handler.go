package deliveryHandler

import (
	controller "Drop/DropUserAccount/controller/delivery-profile"
	"Drop/DropUserAccount/repository/user"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aekam27/trestCommon"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	deliveryService = controller.NewDeliveryRegisterationService(user.NewProfileRepository("users"))
)

func SetDeliveryProfile(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("set delivery profile", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var creds *controller.Credentials
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &creds)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to unmarshal body"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return
	}
	data, err := deliveryService.RegisterDeliveryPerson(*creds)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set delivery profile"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set delivery profile"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("delivery profile added", logrus.Fields{
		"duration": duration,
	})
}
