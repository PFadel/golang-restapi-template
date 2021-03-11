package interfaces

import "github.com/PFadel/golang-restapi-template/app/models"

// SampleController is an interface to use controller's funcs
type SampleController interface {
	SampleRequest(params map[string][]string) (*models.SampleResponse, error)
}
