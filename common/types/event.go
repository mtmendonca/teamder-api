package types

// Event that can be attended
type Event struct {
	ID          string
	Name        string
	Description string
	Venue       string
	Date        string
	Positions   []*interface{}
}
