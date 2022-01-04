package main

import (
	"Nailzee/NailzeePayments/router"
	"log"
	"net/http"

	"github.com/aekam27/trestCommon"
	"github.com/rs/cors"
)

func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.AllowAll().Handler
	return handleCORS(handler)
}
func main() {
	trestCommon.LoadConfig()
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":6004", setupGlobalMiddleware(router)))
}
