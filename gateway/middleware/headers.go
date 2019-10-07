package middleware

import (
	"net/http"
)

//AddHeaders adds response headers
func AddHeaders(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	next.ServeHTTP(w, r)
}
