package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/PFadel/golang-restapi-template/app/errors"

	log "github.com/sirupsen/logrus"
)

func retrieveRequestBody(r *http.Request) string {
	var bodyBytes []byte

	// Read RequestBody
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
		// Restore the io.ReadCloser to its original state
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	return string(bodyBytes)
}

func retrieveResponseBody(r *http.Response) []byte {
	var bodyBytes []byte

	// Read Response
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
		// Restore the io.ReadCloser to its original state
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	return bodyBytes
}

func request(cli *http.Client, req *http.Request, expectedStatusCode int, dataModel ...interface{}) error {
	log.WithFields(log.Fields{
		"Headers":     req.Header,
		"RequestBody": retrieveRequestBody(req),
		"URL":         req.URL.String(),
		"Params":      req.URL.Query(),
	}).Infof("Executing request to: %s", req.URL.String())

	start := time.Now()
	response, err := cli.Do(req)
	elapsed := time.Since(start)
	if err != nil {
		return errors.ErrorConnectingToAPI{BaseError: err}
	}
	defer response.Body.Close()

	// Read Response Body
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	// reset the response body to the original unread state
	response.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	log.WithFields(log.Fields{
		"ElapsedTimeInMilliseconds": elapsed.Nanoseconds() / 1000000,
		"Headers":                   response.Header,
		"URL":                       req.URL.String(),
		"Code":                      response.StatusCode,
		"ResponseBody":              string(bodyBytes),
	}).Infof("Response from %s", req.URL.String())

	if response.StatusCode != expectedStatusCode {
		return errors.UnexpectedStatusCode{
			Actual:   response.StatusCode,
			Expected: expectedStatusCode,
		}
	}

	if dataModel != nil {
		rBody := retrieveResponseBody(response)
		for d := range dataModel {
			err = json.Unmarshal(rBody, &dataModel[d])
			if err == nil {
				return nil
			}
		}
		return errors.ErrorParsingResponse{BaseError: err}
	}

	return nil
}
