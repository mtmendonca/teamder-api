package account

import (
	"context"
	"time"

	"github.com/mtmendonca/teamder-api/common/grpc/account"
	"github.com/mtmendonca/teamder-api/common/types"
	"google.golang.org/grpc"
)

type AccountService interface {
	GetUser() types.User
}

type GRPCAccountService struct {
	Client account.AccountClient
}

func (s *GRPCAccountService) GetUser(ctx context.Context) (*types.User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	userResponse, err := s.Client.GetUser(ctx, &account.GetUserRequest{})
	if err != nil {
		return nil, err
	}

	return &types.User{
		Name:   userResponse.Name,
		Email:  userResponse.Email,
		Avatar: &userResponse.Avatar,
	}, nil
}

func New(conn *grpc.ClientConn) *GRPCAccountService {
	c := account.NewAccountClient(conn)
	return &GRPCAccountService{Client: c}
}
