package api

import (
	"context"
	"net"

	"github.com/mtmendonca/teamder-api/common/grpc/account"
	"github.com/mtmendonca/teamder-api/common/types"
	"google.golang.org/grpc"
)

// Service defines microservice
type Service struct {
	UserStorage types.UserStorage
}

// Login logs user in with provider, creates local record and returns token
func (s *Service) Login(c context.Context, r *account.LoginRequest) (*string, error) {
	u, err := s.UserStorage.GetByEmail(c, r.Email)
	if err != nil {
		panic(err)
	}

	if u == nil {
		return nil, nil
	}

	return &account.UserResponse{Name: u.Name, Email: u.Email, Avatar: u.Avatar}, nil
}

// Start fires up the service
func Start(s *Service, port string) {
	// Start listener
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	// Attach gRPC server
	g := grpc.NewServer()
	account.RegisterAccountServer(g, s)
	if err := g.Serve(lis); err != nil {
		panic(err)
	}
}
