package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/aekam27/trestCommon"
	"github.com/rs/cors"
)

func putReq(token, url string, doc interface{}) ([]byte, error) {
	method := "PUT"
	var bearer = "Bearer " + token
	requestByte, err := json.Marshal(doc)
	if err != nil {
		return []byte{}, err
	}
	requestReader := bytes.NewReader(requestByte)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, requestReader)
	if err != nil {
		return []byte{}, err
	}
	req.Header.Add("Authorization", bearer)
	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
func delReq(token, url string, doc interface{}) ([]byte, error) {
	method := "DELETE"
	var bearer = "Bearer " + token
	requestByte, err := json.Marshal(doc)
	if err != nil {
		return []byte{}, err
	}
	requestReader := bytes.NewReader(requestByte)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, requestReader)
	if err != nil {
		return []byte{}, err
	}
	req.Header.Add("Authorization", bearer)
	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func sendReq(token, url string, w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		result, err := trestCommon.GetApi(token, url)
		if err != nil {
			w.WriteHeader(http.StatusOK)
			w.Write(result)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
	if r.Method == "POST" {
		var data interface{}
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, &data)
		result, err := trestCommon.PostApi(token, url, data)
		if err != nil {
			w.WriteHeader(http.StatusOK)
			w.Write(result)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
	if r.Method == "PUT" {
		var data interface{}
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, &data)
		result, err := putReq(token, url, data)
		if err != nil {
			w.WriteHeader(http.StatusOK)
			w.Write(result)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
	if r.Method == "DELETE" {
		var data interface{}
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, &data)
		result, err := delReq(token, url, data)
		if err != nil {
			w.WriteHeader(http.StatusOK)
			w.Write(result)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}

func main() {
	log.Println("Started proxy")
	mux := http.NewServeMux()
	mux.HandleFunc("/coupon", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers,Authorization")
		token := strings.Split(r.Header.Get("Authorization"), " ")
		url := "http://localhost:6033/api/v1" + r.URL.Path + "?" + r.URL.RawQuery
		if len(token) > 1 {
			sendReq(token[1], url, w, r)
		} else {
			sendReq(" ", url, w, r)
		}
	})
	mux.HandleFunc("/coupon/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers,Authorization")
		token := strings.Split(r.Header.Get("Authorization"), " ")
		url := "http://localhost:6033/api/v1" + r.URL.Path + "?" + r.URL.RawQuery
		if len(token) > 1 {
			sendReq(token[1], url, w, r)
		} else {
			sendReq(" ", url, w, r)
		}
	})
	mux.HandleFunc("/favourite", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers,Authorization")
		token := strings.Split(r.Header.Get("Authorization"), " ")
		url := "http://localhost:6012/api/v1" + r.URL.Path + "?" + r.URL.RawQuery
		if len(token) > 1 {
			sendReq(token[1], url, w, r)
		} else {
			sendReq(" ", url, w, r)
		}
	})
	mux.HandleFunc("/favourite/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers,Authorization")
		token := strings.Split(r.Header.Get("Authorization"), " ")
		url := "http://localhost:6012/api/v1" + r.URL.Path + "?" + r.URL.RawQuery
		if len(token) > 1 {
			sendReq(token[1], url, w, r)
		} else {
			sendReq(" ", url, w, r)
		}
	})
	mux.HandleFunc("/item", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers,Authorization")
		token := strings.Split(r.Header.Get("Authorization"), " ")
		url := "http://localhost:6011/api/v1" + r.URL.Path + "?" + r.URL.RawQuery
		if len(token) > 1 {
			sendReq(token[1], url, w, r)
		} else {
			sendReq(" ", url, w, r)
		}
	})
	mux.HandleFunc("/item/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers,Authorization")
		token := strings.Split(r.Header.Get("Authorization"), " ")
		url := "http://localhost:6011/api/v1" + r.URL.Path + "?" + r.URL.RawQuery
		if len(token) > 1 {
			sendReq(token[1], url, w, r)
		} else {
			sendReq(" ", url, w, r)
		}
	})
	mux.HandleFunc("/ratingsReviews", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers,Authorization")
		token := strings.Split(r.Header.Get("Authorization"), " ")
		url := "http://localhost:6018/api/v1" + r.URL.Path + "?" + r.URL.RawQuery
		if len(token) > 1 {
			sendReq(token[1], url, w, r)
		} else {
			sendReq(" ", url, w, r)
		}
	})
	mux.HandleFunc("/ratingsReviews/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers,Authorization")
		token := strings.Split(r.Header.Get("Authorization"), " ")
		url := "http://localhost:6018/api/v1" + r.URL.Path + "?" + r.URL.RawQuery
		if len(token) > 1 {
			sendReq(token[1], url, w, r)
		} else {
			sendReq(" ", url, w, r)
		}
	})

	mux.HandleFunc("/itemcategory", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers,Authorization")
		token := strings.Split(r.Header.Get("Authorization"), " ")
		url := "http://localhost:6058/api/v1" + r.URL.Path + "?" + r.URL.RawQuery
		if len(token) > 1 {
			sendReq(token[1], url, w, r)
		} else {
			sendReq(" ", url, w, r)
		}
	})
	mux.HandleFunc("/itemcategory/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers,Authorization")
		token := strings.Split(r.Header.Get("Authorization"), " ")
		url := "http://localhost:6058/api/v1" + r.URL.Path + "?" + r.URL.RawQuery
		if len(token) > 1 {
			sendReq(token[1], url, w, r)
		} else {
			sendReq(" ", url, w, r)
		}
	})
	handler := cors.AllowAll().Handler(mux)
	log.Fatal(http.ListenAndServe(":6002", handler))
}
