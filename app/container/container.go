package container

import (
	"net/http"
	"time"

	"github.com/PFadel/golang-restapi-template/app/controllers"
	"github.com/PFadel/golang-restapi-template/app/services"

	"go.uber.org/fx"
)

func SampleService() fx.Option {
	return fx.Provide(services.NewSampleService)
}

func SampleController() fx.Option {
	return fx.Provide(controllers.NewSampleService)
}

func HTTPClient() fx.Option {
	return fx.Provide(func() *http.Client { return &http.Client{Timeout: time.Second * time.Duration(30)} })
}
