package main

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/mtmendonca/teamder-api/common/logger"
	"github.com/mtmendonca/teamder-api/gateway/api"
	gql "github.com/mtmendonca/teamder-api/gateway/graphql"
)

func main() {
	// Load Config
	cfg := api.LoadConfig()

	// Initialize logger
	log := logger.New()

	// Configure graphql
	s := gql.GetSchema()
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(s, &gql.Resolver{
		Log: log,
	}, opts...)

	// Configure service
	service := &api.Service{
		Cfg:            cfg,
		GraphqlHandler: &relay.Handler{Schema: schema},
		Log:            log,
	}

	// Start api
	api.Start(service, ":"+cfg.APIPort)
}
