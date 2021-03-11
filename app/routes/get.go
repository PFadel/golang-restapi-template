package routes

import (
	"encoding/json"
	"net/http"

	"github.com/PFadel/golang-restapi-template/app/errors"
	"github.com/PFadel/golang-restapi-template/app/interfaces"
	"github.com/PFadel/golang-restapi-template/app/models"
)

// SampleRoute is an wrapper to the HandlerFunc that checks the application's health
func SampleRoute(service interfaces.SampleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")

			data, err := service.SampleGet(r.URL.Query())
			if err != nil {
				respErr := models.SampleErrorResponse{
					ErrorDescription: err.Error(),
				}
				w.WriteHeader(http.StatusInternalServerError)

				if _, ok := err.(errors.ErrorConnectingToAPI); ok {
					w.WriteHeader(http.StatusServiceUnavailable)
					respErr.Error = "ERROR_CONNECTING_TO_API"
				} else if _, ok := err.(errors.ErrorParsingResponse); ok {
					respErr.Error = "ERROR_PARSING_RESPONSE"
				} else if newErr, ok := err.(errors.UnexpectedStatusCode); ok {
					respErr.Error = "UNEXPECTED_STATUSCODE"
					w.WriteHeader(newErr.Actual)
				}

				output, _ := json.Marshal(respErr)
				w.Write(output)
				return
			}

			output, _ := json.Marshal(data)
			w.Write(output)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}
