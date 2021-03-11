package services

import (
	"github.com/PFadel/golang-restapi-template/app/interfaces"
	"github.com/PFadel/golang-restapi-template/app/models"
)

type service struct {
	interfaces.SampleController
}

// NewSampleService creates a new SampleService and is used by dependency injection
func NewSampleService(cont interfaces.SampleController) interfaces.SampleService {
	return &service{cont}
}

func (s *service) SampleGet(params map[string][]string) (*models.SampleResponse, error) {
	return s.SampleController.SampleRequest(params)
}
