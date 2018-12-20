package middlewares

import (
	"net/http"

	"github.com/gorilla/mux"
)

func setHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

// SetCommonMiddleware sets common middlewares
func SetCommonMiddleware(r *mux.Router) {
	r.Use(setHeader)
}
