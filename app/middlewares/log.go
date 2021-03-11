package middlewares

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

var forbbidenHeaders = []string{
	"Authorization",
}

type respWriter struct {
	http.ResponseWriter
	status int
	body   []byte
}

func generateTraceKey() string {
	uuidV4, err := uuid.NewRandom()
	if err != nil {
		return ""
	}
	return uuidV4.String()
}

func in(k string, l []string) bool {
	k = strings.ToLower(k)
	for _, s := range l {
		if strings.ToLower(s) == k {
			return true
		}
	}

	return false
}

func (w *respWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *respWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = http.StatusOK
	}
	w.body = b
	n, err := w.ResponseWriter.Write(b)
	return n, err
}

// Log middleware
func Log(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		tKey := generateTraceKey()
		route := r.Method + " " + r.URL.String()
		r.Header.Set("TraceKey", tKey)

		rb, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.WithFields(log.Fields{
				"TraceKey": tKey,
				"Route":    route,
			}).Errorf("error reading request body :%s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		reader := ioutil.NopCloser(bytes.NewBuffer(rb))
		cleanReader := ioutil.NopCloser(bytes.NewBuffer(rb))
		r.Body = cleanReader

		buffer := new(bytes.Buffer)
		buffer.ReadFrom(reader)

		logHeaders := make(http.Header)
		for k, v := range r.Header {
			if !in(k, forbbidenHeaders) {
				logHeaders[k] = v
			}
		}

		log.WithFields(log.Fields{
			"TraceKey": tKey,
			"Route":    route,
			"Body":     buffer.String(),
			"Headers":  logHeaders,
			"Params":   r.URL.Query(),
		}).Infof("Request to %s", route)

		rec := respWriter{ResponseWriter: w}

		start := time.Now()
		h(&rec, r)
		elapsed := time.Since(start)

		log.WithFields(log.Fields{
			"Body":                      string(rec.body),
			"Headers":                   rec.Header(),
			"Route":                     route,
			"StatusCode":                rec.status,
			"TraceKey":                  tKey,
			"ElapsedTimeInMilliseconds": elapsed.Nanoseconds() / 1000000,
		}).Infof("Response from %s", route)
	}
}
