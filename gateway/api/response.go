package api

import (
	"encoding/json"
	"net/http"
)

// Response defines a successful http response
type Response struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"status,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

// Success builds a Success struct based on an input and http status
func Success(result interface{}, status int) *Response {
	return &Response{
		Success:    true,
		StatusCode: status,
		Data:       result,
	}
}

// Send writes a Response to an http response
func (r *Response) Send(w http.ResponseWriter) error {
	w.WriteHeader(r.StatusCode)
	return json.NewEncoder(w).Encode(r)
}
