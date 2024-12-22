package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strings"
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
			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					break
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}
