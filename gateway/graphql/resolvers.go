package graphql

import (
	"context"
	"fmt"

	"github.com/mtmendonca/teamder-api/common/types"
	"github.com/mtmendonca/teamder-api/gateway/grpc/account"
	"github.com/mtmendonca/teamder-api/gateway/middleware"
)

// Resolver defines graphql resolvers
type Resolver struct {
	GRPCAccountService *account.GRPCAccountService
}

func (_ *Resolver) Hello() string {
	return "Hello, world!"
}
func (_ *Resolver) Sup(ctx context.Context) string {
	return fmt.Sprintf("Sup, world! %s", middleware.GetUserID(ctx))
}

func (r *Resolver) CreateUser(ctx context.Context, args *struct{ Input types.CreateUserInput }) *types.User {
	user, err := r.GRPCAccountService.GetUser(ctx)
	if err != nil {
		fmt.Println("errored", err)
	}
	return user
}
