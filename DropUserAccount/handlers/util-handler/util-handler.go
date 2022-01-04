package utilHandler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/aekam27/trestCommon"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
)

type Req struct {
	Type string `json:"type"`
	Data S3File `json:"data"`
}
type S3File struct {
	Name []string `json:"name"`
	Path string   `json:"path"`
}

func GetPreSignedUrl(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("get presigned url", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var s3file Req
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &s3file)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to unmarshal body"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return
	}
	preSignedUrl, err := createPreSignedUrl(s3file)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set profile"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set profile"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "preSignedUrl": preSignedUrl})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("profile updated", logrus.Fields{
		"duration": duration,
	})
}

func createPreSignedUrl(s3file Req) ([]string, error) {
	l := []string{}
	for i := 0; i < len(s3file.Data.Name); i++ {
		fileName := strconv.Itoa(int(time.Now().Unix())) + s3file.Data.Name[i]
		url, _ := preSignedUrl(fileName, s3file.Data.Path)
		l = append(l, url)
	}
	return l, nil
}

func preSignedUrl(filename, path string) (string, error) {
	filename = strings.ReplaceAll(filename, " ", "")
	filename = strconv.Itoa(int(time.Now().Unix())) + filename
	opts := &storage.SignedURLOptions{
		Scheme: storage.SigningSchemeV4,
		Method: "PUT",
		Headers: []string{
			"Content-Type:image/jpeg",
		},
		GoogleAccessID: viper.GetString("gcp.email"),
		PrivateKey:     []byte(viper.GetString("gcp.private_key")),
		Expires:        time.Now().Add(15 * time.Minute),
	}
	str, err := storage.SignedURL(viper.GetString("gcp.bucket"), filename, opts)
	if err != nil {
		trestCommon.ECLog2("failed to create presigned url", err)
		return "", err
	}
	return str, nil
}
