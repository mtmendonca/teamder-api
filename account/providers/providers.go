package providers

import (
	"context"

	"github.com/mtmendonca/teamder-api/common/types"
)

// AuthProvider defines authentication methods for different providers
type AuthProvider interface {
	GetUserInfo(context.Context, string) (*types.User, error)
}
