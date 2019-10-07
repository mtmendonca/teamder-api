package api

import (
	"encoding/json"
	"net/http"
)

var (
	// ErrInvalid returns http 40x
	ErrInvalid = Error{StatusCode: http.StatusBadRequest, Message: "Invalid request"}
	// ErrInternalServer returns http 50x
	ErrInternalServer = Error{StatusCode: http.StatusInternalServerError, Message: "Internal server Error"}
	// ErrUnauthorized returns http 401
	ErrUnauthorized = Error{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"}
	// ErrForbidden returns http 403
	ErrForbidden = Error{StatusCode: http.StatusForbidden, Message: "Forbidden"}
)

// Error defines error fields
type Error struct {
	StatusCode int    `json:"status,omitempty"`
	Message    string `json:"message,omitempty"`
}

// Send adds error to response and sends it back to client
func (e Error) Send(w http.ResponseWriter) error {
	statusCode := e.StatusCode
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(e)
}
