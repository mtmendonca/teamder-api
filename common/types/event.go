package types

import "context"

// Event that can be attended
type Event struct {
	ID          string
	Name        string
	Description string
	Venue       string
	Date        string
	Positions   []*Position
}

// Position defines a job offering
type Position struct {
	Name        string
	Company     string
	Location    string
	Description string
	Experience  string
	Education   string
	Skills      []*Skill
}

// EventStorage defines persistence operations
type EventStorage interface {
	GetEvents(context.Context) ([]*Event, error)
	CreateEvent(context.Context, *Event) error
	GetEventByID(context.Context, string) (*Event, error)
	CreatePosition(context.Context, string, *Position) error
}
