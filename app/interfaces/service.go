package interfaces

import "github.com/PFadel/golang-restapi-template/app/models"

// SampleService is an interface to use service's funcs
type SampleService interface {
	SampleGet(params map[string][]string) (*models.SampleResponse, error)
}
