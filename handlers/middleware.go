package handlers

import (
	"net/http"
)

func ApiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")
		if apiKey != "expected_api_key" {
			http.Error(w, "API key is invalid", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
