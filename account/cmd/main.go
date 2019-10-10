package main

import (
	"context"

	"github.com/mtmendonca/teamder-api/account/api"
	"github.com/mtmendonca/teamder-api/account/storage/mongodb"
	cMongodb "github.com/mtmendonca/teamder-api/common/mongodb"
)

func main() {
	// Load config
	cfg := api.LoadConfig()

	// Initialize mongo connection
	db, err := cMongodb.New(context.Background(), cfg.MongoEndpoint, cfg.MongoDatabase)
	if err != nil {
		panic(err)
	}

	s := &api.Service{
		UserStorage: mongodb.NewUserStorage(db),
	}

	api.Start(s, cfg.APIPort)
}
