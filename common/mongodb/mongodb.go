package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// New returns a new mongodb connection
func New(ctx context.Context, endpoint string, database string) (*mongo.Client, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(endpoint)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to mongo ", endpoint, database)
	return client, nil
}
