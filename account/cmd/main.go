package main

import (
	"context"
	"log"
	"net"

	"github.com/mtmendonca/teamder-api/common/grpc/account"
	"google.golang.org/grpc"
)

type Service struct{}

// GetUser returns a user
func (s *Service) GetUser(context.Context, *account.GetUserRequest) (*account.UserResponse, error) {
	return &account.UserResponse{Name: "foo", Email: "bar", Avatar: "boo"}, nil
}

func main() {
	// service := &api.Service{}
	// api.Start(service)
	s := &Service{}
	lis, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	g := grpc.NewServer()
	account.RegisterAccountServer(g, s)
	if err := g.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
