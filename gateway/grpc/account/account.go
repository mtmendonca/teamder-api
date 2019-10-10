package account

import (
	"context"
	"fmt"
	"time"

	"github.com/mtmendonca/teamder-api/common/grpc/account"
	"google.golang.org/grpc"
)

// Service provides an api to gRPC
type Service interface {
	Login(context.Context, account.LoginRequest) (string, error)
}

// GRPCAccountService has a gRPC client
type GRPCAccountService struct {
	Client account.AccountClient
}

// New instantiates Service with a given gRPC connection
func New(conn *grpc.ClientConn) *GRPCAccountService {
	c := account.NewAccountClient(conn)
	return &GRPCAccountService{Client: c}
}

// Login logs user in with provider and creates a local record if none exists
func (s *GRPCAccountService) Login(ctx context.Context, r account.LoginRequest) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	userResponse, err := s.Client.Login(ctx, &r)
	if err != nil {
		fmt.Println(userResponse)
		return "", err
	}

	return "calculated token", nil
	// return userResponse.GetToken(), nil
}
