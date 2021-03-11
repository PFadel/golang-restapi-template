package app

import (
	"net/http"

	"github.com/PFadel/golang-restapi-template/app/routes"
)

// Start TODO
func Start() {
	http.HandleFunc("/health-check", routes.HealthCheck())
}
