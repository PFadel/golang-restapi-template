package app

import (
	"net/http"

	"github.com/PFadel/golang-restapi-template/app/routes"
)

// Start faz o registro das handle funcs no server HTTP
func Start() {
	http.HandleFunc("/health-check", routes.HealthCheck())
}
