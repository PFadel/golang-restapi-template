package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"

	"github.com/PFadel/golang-restapi-template/app/models"
)

// HealthCheck TODO
func HealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			var resp models.HealthCheckResponse
			resp.ApplicationVersion = "0.1.0"
			resp.Application.Goroutines = runtime.NumGoroutine()

			var mem runtime.MemStats
			runtime.ReadMemStats(&mem)

			resp.Application.HeapAlloc = fmt.Sprintf("%d mb", mem.Sys/1000000) // bytes

			b, _ := json.Marshal(resp)

			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}
