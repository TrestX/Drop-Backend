package main

import (
	"Drop/DropUserAccount/router"
	"log"
	"net/http"

	"github.com/aekam27/trestCommon"
	"github.com/rs/cors"
)

// setupGlobalMiddleware will setup CORS
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.AllowAll().Handler
	return handleCORS(handler)
}

// our main function
func main() {

	trestCommon.LoadConfig()
	// create router and start listen on port 8000
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":6009", setupGlobalMiddleware(router)))
}
