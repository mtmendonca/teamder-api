package providers

import "github.com/mtmendonca/teamder-api/common/types"

// AuthProvider defines authentication methods for different providers
type AuthProvider interface {
	GetUserInfo(string) *types.User
}
