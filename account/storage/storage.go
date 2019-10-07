package storage

import (
	"github.com/mtmendonca/teamder-api/common/types"
)

// Storage provides access to persistence layer
type Storage struct {
	User types.UserStorage
}
