package app

import (
	"net/http"

	"github.com/PFadel/golang-restapi-template/app/routes"
)

// Start register the handle funcs for the patterns
func Start() {
	http.HandleFunc("/health-check", routes.HealthCheck())
}
