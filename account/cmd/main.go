package main

import (
	"context"

	"github.com/mtmendonca/teamder-api/account/api"
	providers "github.com/mtmendonca/teamder-api/account/providers/auth"
	"github.com/mtmendonca/teamder-api/account/storage/mongodb"
	"github.com/mtmendonca/teamder-api/common/logger"
	cMongodb "github.com/mtmendonca/teamder-api/common/mongodb"
)

func main() {
	// Load config
	cfg := api.LoadConfig()

	// Initialize Logger
	log := logger.New()

	// Initialize mongo connection
	session, err := cMongodb.New(context.Background(), cfg.MongoEndpoint, cfg.MongoDatabase)
	if err != nil {
		panic(err)
	}

	// Initialize GoogleProvider
	g, err := providers.NewGoogleProvider(context.Background(), cfg.GoogleAPIClientID)
	if err != nil {
		panic(err)
	}

	s := &api.Service{
		UserStorage:    mongodb.NewUserStorage(session, cfg.MongoDatabase, log),
		GoogleProvider: g,
		Cfg:            cfg,
		Log:            log,
	}

	api.Start(s, cfg.APIPort)
}
