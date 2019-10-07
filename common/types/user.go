package types

// User defines a user record
type User struct {
	Name   string
	Email  string
	Avatar *string
}

// CreateUserInput defines the required fields for creating a User
type CreateUserInput struct {
	Name   string
	Email  string
	Avatar *string
}
