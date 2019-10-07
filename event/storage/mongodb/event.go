package mongodb

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mtmendonca/teamder-api/common/types"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collectionName = "events"
)

type eventDoc struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Venue       string             `json:"venue,omitempty" bson:"venue,omitempty"`
	Date        string             `json:"date,omitempty" bson:"date,omitempty"`
	Positions   []*positionDoc     `json:"positions,omitempty" bson:"positions,omitempty"`
}

type skillDoc struct {
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Level string `json:"level,omitempty" bson:"level,omitempty"`
}

type positionDoc struct {
	Name        string      `json:"name,omitempty" bson:"name,omitempty"`
	Company     string      `json:"company,omitempty" bson:"company,omitempty"`
	Location    string      `json:"location,omitempty" bson:"location,omitempty"`
	Description string      `json:"description,omitempty" bson:"description,omitempty"`
	Experience  string      `json:"experience,omitempty" bson:"experience,omitempty"`
	Education   string      `json:"education,omitempty" bson:"education,omitempty"`
	Skills      []*skillDoc `json:"skills,omitempty" bson:"skills,omitempty"`
}

// EventMongoStorage abstracts mongo persistence
type EventMongoStorage struct {
	session      *mongo.Client
	databaseName string
	log          *logrus.Logger
}

// NewEventStorage returns a new userStorage instance
func NewEventStorage(c *mongo.Client, d string, l *logrus.Logger) *EventMongoStorage {
	return &EventMongoStorage{
		session:      c,
		databaseName: d,
		log:          l,
	}
}

func (s *EventMongoStorage) collection() *mongo.Collection {
	c := s.session.Database(s.databaseName).Collection(collectionName)
	return c
}

// GetEvents finds user record given unique email
func (s *EventMongoStorage) GetEvents(ctx context.Context) ([]*types.Event, error) {
	c := s.collection()
	filter := bson.M{}

	cursor, err := c.Find(ctx, filter)
	if err != nil {
		s.log.Warn("Could not find events", err)
		return nil, errors.New("Could not find events")
	}
	defer cursor.Close(ctx)

	var events []*types.Event
	for cursor.Next(ctx) {
		e := eventDoc{}

		err := cursor.Decode(&e)
		if err != nil {
			s.log.Warn("Could not parse event", err)
			return nil, err
		}

		events = append(events, &types.Event{
			ID:          e.ID.Hex(),
			Name:        e.Name,
			Description: e.Description,
			Venue:       e.Venue,
			Date:        e.Date,
		})

	}

	return events, nil
}

// CreateEvent saves and returns an Event
func (s *EventMongoStorage) CreateEvent(ctx context.Context, event *types.Event) error {
	c := s.collection()

	input := eventDoc{
		Name:        event.Name,
		Description: event.Description,
		Venue:       event.Venue,
		Date:        event.Date,
	}

	res, err := c.InsertOne(ctx, input)
	if err != nil {
		s.log.Warn("Could not save event to database", err)
		return err
	}

	event.ID = res.InsertedID.(primitive.ObjectID).Hex()
	s.log.Debug("Saved new event", event)
	return nil
}

// GetEventByID returns an event based on its ID
func (s *EventMongoStorage) GetEventByID(ctx context.Context, ID string) (*types.Event, error) {
	c := s.collection()
	var event eventDoc

	oID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		s.log.Warn("Invalid event id", err)
		return nil, err
	}

	filter := bson.M{"_id": oID}
	err = c.FindOne(ctx, filter).Decode(&event)
	if err != nil {
		s.log.Warn("Could not find event for id ", ID)
		return nil, err
	}

	positions := make([]*types.Position, len(event.Positions))
	for i, pos := range event.Positions {
		skills := make([]*types.Skill, len(pos.Skills))
		for j, skill := range pos.Skills {
			skills[j] = &types.Skill{
				Name:  skill.Name,
				Level: skill.Level,
			}
		}
		positions[i] = &types.Position{
			Name:        pos.Name,
			Company:     pos.Company,
			Location:    pos.Location,
			Description: pos.Description,
			Experience:  pos.Experience,
			Education:   pos.Education,
			Skills:      skills,
		}
	}

	return &types.Event{
		ID:          event.ID.Hex(),
		Name:        event.Name,
		Description: event.Description,
		Venue:       event.Venue,
		Date:        event.Date,
		Positions:   positions,
	}, nil
}

// CreatePosition adds a position do the event document
func (s *EventMongoStorage) CreatePosition(ctx context.Context, ID string, p *types.Position) error {
	c := s.collection()

	oID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		s.log.Warn("Could not derive object id from", ID)
		return err
	}
	filter := bson.M{"_id": oID}

	skills := make([]*skillDoc, len(p.Skills))
	for i, skill := range p.Skills {
		skills[i] = &skillDoc{
			Name:  skill.Name,
			Level: skill.Level,
		}
	}

	pDoc := &positionDoc{
		Name:        p.Name,
		Company:     p.Company,
		Location:    p.Location,
		Description: p.Description,
		Experience:  p.Experience,
		Education:   p.Education,
		Skills:      skills,
	}

	_, err = c.UpdateOne(
		ctx,
		filter,
		bson.M{"$push": bson.M{"positions": pDoc}},
	)
	if err != nil {
		return err
	}

	return nil

}
