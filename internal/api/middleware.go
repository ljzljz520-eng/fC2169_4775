package api

import "net/http"

func WithHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Campus-Service", "awards")
		next.ServeHTTP(w, r)
	})
}
func Chain(h http.Handler) http.Handler { return WithHeaders(MethodAllowed(h)) }
