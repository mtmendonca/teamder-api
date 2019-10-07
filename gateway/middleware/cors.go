package middleware

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Cors applies cords rules to router
func Cors(r *mux.Router) http.Handler {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Origin", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	return handlers.CORS(originsOk, headersOk, methodsOk)(r)
}
