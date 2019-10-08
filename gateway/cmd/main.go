package main

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/mtmendonca/teamder-api/gateway/api"
	gql "github.com/mtmendonca/teamder-api/gateway/graphql"
)

func main() {
	// Configure graphql
	s := gql.GetSchema()
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(s, &gql.Resolver{}, opts...)

	// Configure service
	service := &api.Service{
		GraphqlHandler: &relay.Handler{Schema: schema},
	}

	// Start api
	api.Start(service)
}
