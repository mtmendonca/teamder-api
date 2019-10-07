package api

// Forbidden indicates lack of authorization context
type ErrForbidden struct {
	Message string `json:"message"`
}
