package graphql

import (
	"context"

	"github.com/mtmendonca/teamder-api/common/types"
	"github.com/sirupsen/logrus"
)

// Resolver defines graphql resolvers
type Resolver struct {
	Log *logrus.Logger
}

// Events returns a list of upcoming events
func (r *Resolver) Events(ctx context.Context) ([]*types.Event, error) {
	events := []*types.Event{
		&types.Event{
			ID:          "id-01",
			Name:        "GDG Devfest Cerrado 2019",
			Description: "O maior evento sobre tecnologias Google do Cerrado!",
			Venue:       "UniAlfa Perimetral",
			Date:        "1571455000601",
		},
	}

	return events, nil
}

// CreateEvent creates and returns a new event
func (r *Resolver) CreateEvent(ctx context.Context, args *struct{ Input types.Event }) (*types.Event, error) {
	r.Log.Info("Creating event")
	// aditional magic will go here
	r.Log.Info("Created event", args.Input)
	return &types.Event{
		ID:          "id-" + args.Input.Name,
		Name:        args.Input.Name,
		Description: args.Input.Description,
		Venue:       args.Input.Venue,
		Date:        args.Input.Date,
	}, nil
}
