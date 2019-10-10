package main

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	cmnGrpc "github.com/mtmendonca/teamder-api/common/grpc"
	"github.com/mtmendonca/teamder-api/gateway/api"
	gql "github.com/mtmendonca/teamder-api/gateway/graphql"
	"github.com/mtmendonca/teamder-api/gateway/grpc/account"
)

func main() {
	cfg := api.LoadConfig()

	// Connect to and instantiate Acccount Service
	conn, err := cmnGrpc.GetConn(cfg.AccountEndpoint)
	if err != nil {
		panic(err)
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
	api.Start(service, ":"+cfg.APIPort)
}
