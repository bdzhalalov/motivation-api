package middleware

import (
	"motivations-api/config"
	"net/http"
)

func CheckApiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("API-Key") != config.Cfg.APIKey {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Access Denied"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
