package graphql

import (
	"context"
	"fmt"

	"github.com/mtmendonca/teamder-api/common/types"
	"github.com/mtmendonca/teamder-api/gateway/grpc/account"
)

// Resolver defines graphql resolvers
type Resolver struct {
	GRPCAccountService account.Service
}

type getUserByEmailArgs struct {
	Email string
}

// GetUserByEmail fetches user from gprc service
func (r *Resolver) GetUserByEmail(ctx context.Context, args *getUserByEmailArgs) *types.User {
	user, err := r.GRPCAccountService.GetUserByEmail(ctx, args.Email)
	if err != nil {
		fmt.Println("errored", err)
	}
	return user
}
