package types

import "context"

// Event that can be attended
type Event struct {
	ID          string
	Name        string
	Description string
	Venue       string
	Date        string
	Positions   []*interface{}
}

// EventStorage defines persistence operations
type EventStorage interface {
	GetEvents(context.Context) ([]*Event, error)
	CreateEvent(context.Context, *Event) error
}
