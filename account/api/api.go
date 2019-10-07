package api

import (
	"context"
	"errors"
	"net"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mtmendonca/teamder-api/account/providers"
	"github.com/mtmendonca/teamder-api/common/grpc/account"
	"github.com/mtmendonca/teamder-api/common/types"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	google = "google"
)

// Service defines microservice
type Service struct {
	UserStorage    types.UserStorage
	GoogleProvider providers.AuthProvider
	Cfg            *Config
	Log            *logrus.Logger
}

func (s *Service) isGoogleTokenValid(ctx context.Context, token string) bool {
	// Get user info from google service using signed token
	s.Log.Debug("Getting user account info for token ", token)

	_, err := s.GoogleProvider.GetUserInfo(ctx, token)
	if err != nil {
		s.Log.Warn("api error: ", err)
		return false
	}

	s.Log.Debug("user token is valid", token)
	return true
}

// isTokenValid validates token according to provider
func (s *Service) isTokenValid(ctx context.Context, token string, provider string) bool {
	s.Log.Debug("validating token from provider: ", provider)
	if provider == google {
		return s.isGoogleTokenValid(ctx, token)
	}
	return false
}

func (s *Service) generateJWTToken(u *types.User) (string, error) {
	s.Log.Debug("Generating JTW Token for user ", u)
	// generate access token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub":    u.ID,
		"exp":    time.Now().Add(time.Hour * time.Duration(72)).Unix(),
		"name":   u.Name,
		"email":  u.Email,
		"avatar": u.Avatar,
	})

	return token.SignedString([]byte(s.Cfg.JWTSecret))
}

// Login logs user in with provider, creates local record and returns token
func (s *Service) Login(ctx context.Context, r *account.LoginRequest) (*account.LoginResponse, error) {
	var user *types.User
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	// Validate token
	isValid := s.isTokenValid(ctx, r.GetToken(), r.GetProvider())
	if !isValid {
		s.Log.Warn("Invalid token")
		return nil, errors.New("Invalid token")
	}

	// Find or create user record
	user, err := s.UserStorage.GetByEmail(ctx, r.GetEmail())
	if err != nil {
		user = &types.User{
			Name:   r.GetName(),
			Email:  r.GetEmail(),
			Avatar: r.GetAvatar(),
		}

		cerr := s.UserStorage.Create(ctx, user)
		if cerr != nil {
			return nil, errors.New("Could not load user")
		}
	}

	// Generate a long-living token for the user
	token, err := s.generateJWTToken(user)

	if err != nil {
		s.Log.Warn("Could not generate token", err)
		return nil, err
	}

	return &account.LoginResponse{Token: token}, nil
}

// SetProfile saves a user profile
func (s *Service) SetProfile(ctx context.Context, r *account.SetProfileRequest) (*account.SetProfileResponse, error) {
	skills := make([]*types.Skill, len(r.GetSkills()))
	for i, skill := range r.GetSkills() {
		skills[i] = &types.Skill{
			Name:  skill.GetName(),
			Level: skill.GetLevel(),
		}
	}

	err := s.UserStorage.SetProfile(ctx, r.GetUserID(), &types.SetProfileInput{
		Education:  r.GetEducation(),
		Experience: r.GetExperience(),
		Role:       r.GetRole(),
		Skills:     skills,
	})
	if err != nil {
		return &account.SetProfileResponse{Success: false}, err
	}

	return &account.SetProfileResponse{Success: true}, nil
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
	account.RegisterAccountServiceServer(g, s)
	if err := g.Serve(lis); err != nil {
		panic(err)
	}
}
