package graphql

import (
	"context"
	"time"

	"github.com/mtmendonca/teamder-api/common/grpc/account"
	"github.com/mtmendonca/teamder-api/common/grpc/event"
	"github.com/mtmendonca/teamder-api/common/types"
	"github.com/sirupsen/logrus"
)

const (
	contextKeyUserID = types.ContextKey("userID")
)

// Resolver defines graphql resolvers
type Resolver struct {
	GRPCAccountClient account.AccountServiceClient
	GRPCEventClient   event.EventServiceClient
	Log               *logrus.Logger
}

// SetProfile saves a user profile
func (r *Resolver) SetProfile(ctx context.Context, args *struct{ Input types.SetProfileInput }) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	skills := make([]*account.Skill, len(args.Input.Skills))
	for i, skill := range args.Input.Skills {
		skills[i] = &account.Skill{
			Name:  skill.Name,
			Level: skill.Level,
		}
	}

	response, err := r.GRPCAccountClient.SetProfile(ctx, &account.SetProfileRequest{
		UserID:     ctx.Value(contextKeyUserID).(string),
		Education:  args.Input.Education,
		Experience: args.Input.Experience,
		Role:       args.Input.Role,
		Skills:     skills,
	})

	if err != nil {
		r.Log.Warn(response)
		return false, err
	}

	return true, nil
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

// Event returns an event by id
func (r *Resolver) Event(ctx context.Context, args *struct{ ID string }) (*types.Event, error) {
	res, err := r.GRPCEventClient.GetEventByID(ctx, &event.GetEventByIDRequest{ID: args.ID})
	if err != nil {
		r.Log.Warn("Could not get event with id", args.ID, err)
		return nil, err
	}

	positions := make([]*types.Position, len(res.Event.GetPositions()))
	for i, p := range res.Event.GetPositions() {
		skills := make([]*types.Skill, len(p.GetSkills()))
		for j, s := range p.GetSkills() {
			skills[j] = &types.Skill{
				Name:  s.GetName(),
				Level: s.GetLevel(),
			}
		}

		positions[i] = &types.Position{
			Name:        p.GetName(),
			Company:     p.GetCompany(),
			Location:    p.GetLocation(),
			Description: p.GetDescription(),
			Experience:  p.GetExperience(),
			Education:   p.GetEducation(),
			Skills:      skills,
		}
	}

	return &types.Event{
		ID:          res.Event.GetID(),
		Name:        res.Event.GetName(),
		Description: res.Event.GetDescription(),
		Venue:       res.Event.GetVenue(),
		Date:        res.Event.GetDate(),
		Positions:   positions,
	}, nil
}

// CreatePosition adds an open position to an event
func (r *Resolver) CreatePosition(ctx context.Context, args *struct {
	EventID string
	Input   types.Position
}) (bool, error) {
	r.Log.Info("Creating position", args)

	skills := make([]*event.Skill, len(args.Input.Skills))
	for i, s := range args.Input.Skills {
		skills[i] = &event.Skill{
			Name:  s.Name,
			Level: s.Level,
		}
	}

	position := &event.Position{
		Name:        args.Input.Name,
		Company:     args.Input.Company,
		Location:    args.Input.Location,
		Description: args.Input.Description,
		Experience:  args.Input.Experience,
		Education:   args.Input.Education,
		Skills:      skills,
	}

	res, err := r.GRPCEventClient.CreatePosition(ctx, &event.CreatePositionRequest{Position: position, EventID: args.EventID})
	if err != nil || res.GetSuccess() == false {
		r.Log.Warn("Could not create Position", err)
		return false, err
	}

	return true, nil
}
