package account

import (
	"context"
	"time"

	"github.com/mtmendonca/teamder-api/common/grpc/account"
	"github.com/mtmendonca/teamder-api/common/types"
	"google.golang.org/grpc"
)

// Service provides an api to gRPC
type Service interface {
	GetUserByEmail(context.Context, string) (*types.User, error)
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

// GetUserByEmail finds user based on email
func (s *GRPCAccountService) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	userResponse, err := s.Client.GetUserByEmail(ctx, &account.GetUserByEmailRequest{Email: email})
	if err != nil {
		return nil, err
	}

	return &types.User{
		Name:   userResponse.Name,
		Email:  userResponse.Email,
		Avatar: userResponse.Avatar,
	}, nil
}
