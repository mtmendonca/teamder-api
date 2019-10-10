package providers

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/mtmendonca/teamder-api/common/types"
	oauth2 "google.golang.org/api/oauth2/v2"
)

// GoogleAuthProvider encapsulates google log in logic
// https://developers.google.com/identity/sign-in/web/backend-auth
type GoogleAuthProvider struct {
	clientID string
	service  *oauth2.Service
}

// New returns a new provider
func New(ctx context.Context, clientID string) *GoogleAuthProvider {
	s, err := oauth2.NewService(ctx)
	if err != nil {
		fmt.Println("Errored", err)
	}

	return &GoogleAuthProvider{
		service:  s,
		clientID: clientID,
	}
}

// GetUserInfo retrieves user information from google account
func (p *GoogleAuthProvider) GetUserInfo(ctx context.Context, token string) (*types.User, error) {
	// Validate token
	ti, err := p.service.Tokeninfo().IdToken(token).Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	log.Printf("tokeninfo: %#v", ti)

	if ti.Audience != p.clientID {
		return nil, errors.New("Wrong client id " + ti.Audience)
	}

	if !ti.VerifiedEmail {
		return nil, errors.New("Email not verified")
	}

	return &types.User{}, nil
}
