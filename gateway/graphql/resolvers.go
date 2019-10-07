package graphql

import (
	"context"
	"fmt"

	"github.com/mtmendonca/backend/common/types"
	"github.com/mtmendonca/backend/gateway/middleware"
)

// Resolver defines graphql resolvers
type Resolver struct{}

func (_ *Resolver) Hello() string {
	return "Hello, world!"
}
func (_ *Resolver) Sup(ctx context.Context) string {
	return fmt.Sprintf("Sup, world! %s", middleware.GetUserID(ctx))
}

func (_ *Resolver) CreateUser(ctx context.Context, args *struct{ Input types.CreateUserInput }) *types.User {
	return &types.User{Name: middleware.GetUserID(ctx), Email: args.Input.Email, Avatar: args.Input.Avatar}
}
