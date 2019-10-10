package graphql

import (
	"context"
	"fmt"

	cmnAccount "github.com/mtmendonca/teamder-api/common/grpc/account"
	"github.com/mtmendonca/teamder-api/gateway/grpc/account"
)

// Resolver defines graphql resolvers
type Resolver struct {
	GRPCAccountService account.Service
}

// Login authenticates user with google. Creates local user record if none exists
func (r *Resolver) Login(ctx context.Context, args *struct{ Input cmnAccount.LoginRequest }) string {
	token, err := r.GRPCAccountService.Login(ctx, args.Input)
	if err != nil {
		fmt.Println(err)
	}
	return token
}
