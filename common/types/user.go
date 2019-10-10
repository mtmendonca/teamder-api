package types

import "context"

// User defines a user record
type User struct {
	Name   string
	Email  string
	Avatar string
}

// UserStorage saves users to a persistence layer
type UserStorage interface {
	GetByEmail(context.Context, string) (*User, error)
}
