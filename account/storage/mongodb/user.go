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
	collectionName = "users"
)

type userDoc struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty" bson:"name,omitempty"`
	Email  string             `json:"email,omitempty" bson:"email,omitempty"`
	Avatar string             `json:"avatar,omitempty" bson:"avatar,omitempty"`
}

// UserMongoStorage abstracts mongo persistence
type UserMongoStorage struct {
	session      *mongo.Client
	databaseName string
	log          *logrus.Logger
}

// NewUserStorage returns a new userStorage instance
func NewUserStorage(c *mongo.Client, d string, l *logrus.Logger) *UserMongoStorage {
	return &UserMongoStorage{
		session:      c,
		databaseName: d,
		log:          l,
	}
}

func (s *UserMongoStorage) collection() *mongo.Collection {
	c := s.session.Database(s.databaseName).Collection(collectionName)
	return c
}

// GetByEmail finds user record given unique email
func (s *UserMongoStorage) GetByEmail(ctx context.Context, email string) (*types.User, error) {
	var user userDoc

	c := s.collection()
	filter := bson.M{"email": email}

	err := c.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		s.log.Warn("Could not find user", err)
		return nil, errors.New("Could not find user")
	}

	return &types.User{
		ID:     user.ID.Hex(),
		Name:   user.Name,
		Email:  user.Email,
		Avatar: user.Avatar,
	}, nil
}

// Create inserts a user record into the database
func (s *UserMongoStorage) Create(ctx context.Context, u *types.User) error {
	c := s.collection()
	input := userDoc{
		Name:   u.Name,
		Email:  u.Email,
		Avatar: u.Avatar,
	}

	res, err := c.InsertOne(ctx, input)
	if err != nil {
		return err
	}

	u.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

// SetProfile saves a user profile
func (s *UserMongoStorage) SetProfile(ctx context.Context, userID string, profile *types.SetProfileInput) error {
	c := s.collection()
	oID, _ := primitive.ObjectIDFromHex(userID)
	filter := bson.M{"_id": oID}

	_, err := c.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"profile": profile}})

	if err != nil {
		return err
	}

	return nil
}
