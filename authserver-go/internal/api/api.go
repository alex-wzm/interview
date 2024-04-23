package api

import (
	"context"
	"interview-authserver/config"
	"interview-authserver/internal/api/auth"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Authenticator interface {
	Authenticate(context.Context, string, *config.EnvConfig) (string, error)
}

type server struct {
	auth.UnimplementedAuthServiceServer
	env           *config.EnvConfig
	authenticator Authenticator
}

func New(env *config.EnvConfig, authenticator Authenticator) *server {
	return &server{env: env, authenticator: authenticator}
}

func (s *server) Authenticate(ctx context.Context, req *auth.AuthRequest) (*auth.AuthResponse, error) {
	token, authErr := s.authenticator.Authenticate(ctx, req.Username, s.env)
	if authErr != nil {
		return nil, status.Error(codes.Unauthenticated, authErr.Error())
	}

	return &auth.AuthResponse{Token: token}, nil
}
