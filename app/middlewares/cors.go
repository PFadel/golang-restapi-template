package middlewares

import (
	"net/http"

	"github.com/rs/cors"
)

// Cors middleware
func Cors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			cors.Default().HandlerFunc(w, r)
		} else {
			h(w, r)
		}
	}
}
