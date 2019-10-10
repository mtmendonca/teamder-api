package mongodb

import (
	"context"

	"github.com/mtmendonca/teamder-api/common/types"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserMongoStorage abstracts mongo persistence
type UserMongoStorage struct {
	session *mongo.Client
}

// NewUserStorage returns a new userStorage instance
func NewUserStorage(c *mongo.Client) *UserMongoStorage {
	return &UserMongoStorage{
		session: c,
	}
}

// GetByEmail finds user record given unique email
func (s *UserMongoStorage) GetByEmail(c context.Context, email string) (*types.User, error) {
	return &types.User{
		Name:   "foo",
		Email:  "boo",
		Avatar: "avatar",
	}, nil
}
