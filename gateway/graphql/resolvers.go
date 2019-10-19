package graphql

import (
	"context"

	"github.com/mtmendonca/teamder-api/common/grpc/event"
	"github.com/mtmendonca/teamder-api/common/types"
	"github.com/sirupsen/logrus"
)

// Resolver defines graphql resolvers
type Resolver struct {
	Log             *logrus.Logger
	GRPCEventClient event.EventServiceClient
}

// Events returns a list of upcoming events
func (r *Resolver) Events(ctx context.Context) ([]*types.Event, error) {
	res, err := r.GRPCEventClient.GetEvents(ctx, &event.GetEventsRequest{})
	if err != nil {
		r.Log.Warn("Could not get events", err)
		return nil, err
	}

	events := make([]*types.Event, len(res.Events))
	for i, e := range res.Events {
		events[i] = &types.Event{
			ID:          e.GetID(),
			Name:        e.GetName(),
			Description: e.GetDescription(),
			Venue:       e.GetVenue(),
			Date:        e.GetDate(),
		}
	}

	return events, nil
}

// CreateEvent creates and returns a new event
func (r *Resolver) CreateEvent(ctx context.Context, args *struct{ Input types.Event }) (*types.Event, error) {
	r.Log.Info("Creating event")
	response, err := r.GRPCEventClient.CreateEvent(ctx, &event.CreateEventRequest{
		Name:        args.Input.Name,
		Description: args.Input.Description,
		Venue:       args.Input.Venue,
		Date:        args.Input.Date,
	})

	if err != nil {
		r.Log.Warn("Could not create event", err)
		return nil, err
	}

	event := response.GetEvent()
	r.Log.Info("Created event", event)
	return &types.Event{
		ID:          event.GetID(),
		Name:        event.GetName(),
		Description: event.GetDescription(),
		Venue:       event.GetVenue(),
		Date:        event.GetDate(),
	}, nil
}
