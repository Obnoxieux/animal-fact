package server

import (
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/gorilla/mux"
)

func CORSMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origEnv := os.Getenv("ALLOWED_ORIGINS")
			if origEnv == "" {
				next.ServeHTTP(w, r)
			}
			allowedOrigins := strings.Split(origEnv, ",")

			origin := r.Header.Get("Origin")
			if slices.Contains(allowedOrigins, origin) {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}
			next.ServeHTTP(w, r)
		})
	}
}
