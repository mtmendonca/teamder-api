package main

import (
	"log"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/mtmendonca/teamder-api/gateway/api"
	gql "github.com/mtmendonca/teamder-api/gateway/graphql"
	"github.com/mtmendonca/teamder-api/gateway/grpc/account"
	"google.golang.org/grpc"
)

func main() {
	// Connect to and instantiate Acccount Service
	conn, err := grpc.Dial("localhost:3001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	as := account.New(conn)

	// Configure graphql
	s := gql.GetSchema()
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(s, &gql.Resolver{
		GRPCAccountService: as,
	}, opts...)

	// Configure service
	service := &api.Service{
		GraphqlHandler: &relay.Handler{Schema: schema},
	}

	// Start api
	api.Start(service)
}
