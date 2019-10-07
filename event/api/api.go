package api

import (
	"context"
	"errors"
	"net"

	"github.com/mtmendonca/teamder-api/common/grpc/event"
	"github.com/mtmendonca/teamder-api/common/types"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Service defines microservice
type Service struct {
	EventStorage types.EventStorage
	Cfg          *Config
	Log          *logrus.Logger
}

// GetEvents lists events
func (s *Service) GetEvents(ctx context.Context, r *event.GetEventsRequest) (*event.GetEventsResponse, error) {
	s.Log.Debug("Fetching events")
	events, err := s.EventStorage.GetEvents(ctx)
	var eventsMsg = make([]*event.Event, len(events))
	if err != nil {
		s.Log.Warn("Could not get events from database", err)
		return &event.GetEventsResponse{}, err
	}

	for i, e := range events {
		eventsMsg[i] = &event.Event{
			ID:          e.ID,
			Name:        e.Name,
			Description: e.Description,
			Date:        e.Date,
			Venue:       e.Venue,
		}
	}

	s.Log.Debug("Returning the following events", eventsMsg, len(eventsMsg))

	return &event.GetEventsResponse{
		Events: eventsMsg,
	}, nil
}

// CreateEvent creates and returns a new event
func (s *Service) CreateEvent(ctx context.Context, r *event.CreateEventRequest) (*event.CreateEventResponse, error) {
	e := &types.Event{
		Name:        r.GetName(),
		Description: r.GetDescription(),
		Venue:       r.GetVenue(),
		Date:        r.GetDate(),
	}
	err := s.EventStorage.CreateEvent(ctx, e)

	if err != nil {
		s.Log.Warn("Could not create event", err)
		return &event.CreateEventResponse{}, err
	}

	s.Log.Debug("Created event", e)
	return &event.CreateEventResponse{Event: &event.Event{
		Name:        e.Name,
		Description: e.Description,
		Venue:       e.Venue,
		Date:        e.Date,
	}}, nil
}

// GetEventByID returns events by id
func (s *Service) GetEventByID(ctx context.Context, r *event.GetEventByIDRequest) (*event.GetEventByIDResponse, error) {
	ID := r.GetID()

	e, err := s.EventStorage.GetEventByID(ctx, ID)
	if err != nil {
		return &event.GetEventByIDResponse{}, err
	}

	positions := make([]*event.Position, len(e.Positions))
	for i, pos := range e.Positions {
		skills := make([]*event.Skill, len(pos.Skills))
		for j, skill := range pos.Skills {
			skills[j] = &event.Skill{
				Name:  skill.Name,
				Level: skill.Level,
			}
		}
		positions[i] = &event.Position{
			Name:        pos.Name,
			Company:     pos.Company,
			Location:    pos.Location,
			Description: pos.Description,
			Experience:  pos.Experience,
			Education:   pos.Education,
			Skills:      skills,
		}
	}

	return &event.GetEventByIDResponse{
		Event: &event.Event{
			ID:          e.ID,
			Name:        e.Name,
			Description: e.Description,
			Venue:       e.Venue,
			Date:        e.Date,
			Positions:   positions,
		},
	}, nil
}

// CreatePosition adds a position to an event
func (s *Service) CreatePosition(ctx context.Context, r *event.CreatePositionRequest) (*event.CreatePositionResponse, error) {
	res := new(event.CreatePositionResponse)

	skills := make([]*types.Skill, len(r.GetPosition().GetSkills()))
	for i, skill := range r.GetPosition().GetSkills() {
		skills[i] = &types.Skill{
			Name:  skill.GetName(),
			Level: skill.GetLevel(),
		}
	}

	p := &types.Position{
		Name:        r.GetPosition().GetName(),
		Company:     r.GetPosition().GetCompany(),
		Location:    r.GetPosition().GetLocation(),
		Description: r.GetPosition().GetDescription(),
		Experience:  r.GetPosition().GetExperience(),
		Education:   r.GetPosition().GetEducation(),
		Skills:      skills,
	}

	err := s.EventStorage.CreatePosition(ctx, r.GetEventID(), p)

	if err != nil {
		s.Log.Warn("Could not create position", err)
		return nil, errors.New("Could not create position")
	}

	res.Success = true
	return res, nil
}

// Start fires up the service
func Start(s *Service, port string) {
	// Start listener
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	// Attach gRPC server
	g := grpc.NewServer()
	event.RegisterEventServiceServer(g, s)
	if err := g.Serve(lis); err != nil {
		panic(err)
	}
}
