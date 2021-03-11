package controllers

import (
	"net/http"

	"github.com/PFadel/golang-restapi-template/app/interfaces"
	"github.com/PFadel/golang-restapi-template/app/models"

	log "github.com/sirupsen/logrus"
)

type controller struct {
	Client *http.Client
}

// NewSampleService creates a new SampleService and is used by dependency injection
func NewSampleService(client *http.Client) interfaces.SampleController {
	return &controller{client}
}

func (c *controller) SampleRequest(params map[string][]string) (*models.SampleResponse, error) {
	url := "https://httpbin.org/get"

	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("Error creating new HTTP request: %s", err.Error())
		return nil, err
	}

	q := r.URL.Query()
	for k, v := range params {
		for _, s := range v {
			q.Add(k, s)
		}
	}
	r.URL.RawQuery = q.Encode()

	var resp models.SampleResponse
	err = request(c.Client, r, http.StatusOK, &resp)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &resp, nil
}
