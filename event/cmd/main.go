package main

import (
	"context"

	"github.com/mtmendonca/teamder-api/common/logger"
	cMongodb "github.com/mtmendonca/teamder-api/common/mongodb"
	"github.com/mtmendonca/teamder-api/event/api"
	"github.com/mtmendonca/teamder-api/event/storage/mongodb"
)

func main() {
	// Load config
	cfg := api.LoadConfig()

	// Initialize logger
	log := logger.New()

	// Initialize mongo connection
	session, err := cMongodb.New(context.Background(), cfg.MongoEndpoint, cfg.MongoDatabase)
	if err != nil {
		panic(err)
	}

	s := &api.Service{
		EventStorage: mongodb.NewEventStorage(session, cfg.MongoDatabase, log),
		Cfg:          cfg,
		Log:          log,
	}

	api.Start(s, cfg.APIPort)
}
