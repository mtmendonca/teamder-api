package types

import "context"

// User defines a user record
type User struct {
	ID     string
	Name   string
	Email  string
	Avatar string
}

// Skill that a user may have
type Skill struct {
	Name  string
	Level string
}

// SetProfileInput defines the input type updating a user profile
type SetProfileInput struct {
	Education  string
	Experience string
	Role       string
	Skills     []*Skill
}

// UserStorage defines persistence operations
type UserStorage interface {
	Create(context.Context, *User) error
	GetByEmail(context.Context, string) (*User, error)
	SetProfile(context.Context, string, *SetProfileInput) error
}
