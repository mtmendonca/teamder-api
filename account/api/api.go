package api

import (
	"context"
	"log"
	"net"

	"github.com/mtmendonca/teamder-api/common/user"
	"google.golang.org/grpc"
)

// Service defines microservice
type Service struct{}

// GetUser returns a user
func (s *Service) GetUser(context.Context, *user.GetUserRequest) (*user.UserResponse, error) {
	return &user.UserResponse{Name: "foo", Email: "bar", Avatar: "boo"}, nil
}

// Start fires up the service
func Start(s *Service) {
	lis, err := net.Listen("tcp", "3001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	g := grpc.NewServer()
	user.RegisterUserServiceServer(g, s)
	if err := g.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
