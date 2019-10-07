package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/mtmendonca/teamder-api/common/grpc/account"
)

func (s *Service) login(w http.ResponseWriter, r *http.Request) {
	var form struct {
		Token    string
		Provider string
		Name     string
		Email    string
		Avatar   string
	}
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		ErrInvalid.Send(w)
		return
	}

	s.Log.Debug("Loggin in with ", form.Token)
	input := &account.LoginRequest{
		Token:    form.Token,
		Provider: form.Provider,
		Name:     form.Name,
		Email:    form.Email,
		Avatar:   form.Avatar,
	}
	userResponse, err := s.GRPCAccountClient.Login(ctx, input)
	if err != nil {
		s.Log.Warn(err)
		ErrUnauthorized.Send(w)
		return
	}

	s.Log.Debug("User token:", userResponse.GetToken())

	Success(userResponse.GetToken(), http.StatusOK).Send(w)
}
