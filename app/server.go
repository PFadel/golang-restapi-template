package app

import (
	"net/http"

	"github.com/PFadel/golang-restapi-template/app/container"
	"github.com/PFadel/golang-restapi-template/app/interfaces"
	"github.com/PFadel/golang-restapi-template/app/routes"

	"go.uber.org/fx"
)

// Start register the handle funcs for the patterns
func Start() {
	fx.New(
		container.SampleService(),
		container.SampleController(),
		container.HTTPClient(),
		fx.Invoke((func(svc interfaces.SampleService) {
			http.HandleFunc("/get", routes.SampleRoute(svc))
		})),
	)
	http.HandleFunc("/health-check", routes.HealthCheck())
}
