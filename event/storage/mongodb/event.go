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
	Positions   []*interface{}     `json:"positions,omitempty" bson:"positions,omitempty"`
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
