package api

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"authservice-go/internal/api/auth"
	jwtValidator "authservice-go/internal/domain/jwt"
)

type server struct {
	auth.UnimplementedAuthServiceServer
	users  map[string]string
	secret []byte
}

func New(users map[string]string, secret []byte) *server {
	return &server{users: users, secret: secret}
}

func (s *server) Auth(ctx context.Context, in *auth.AuthRequest) (*auth.AuthResponse, error) {
	if pass, ok := s.users[in.GetUsername()]; ok {
		if pass == in.GetPassword() {
			token, err := jwtValidator.GenerateToken(in.GetUsername(), time.Hour, s.secret)
			if err != nil {
				slog.Error("failed to generate token", slog.Any("error", err))
				return &auth.AuthResponse{}, err
			}
			// successfully passed authentication,return our generated token so the user can use it against other services
			return &auth.AuthResponse{Token: token}, nil
		}
	}
	return &auth.AuthResponse{}, fmt.Errorf("unable to authenticate user")
}

func (s *server) Validate(ctx context.Context, in *auth.ValidateRequest) (*auth.ValidateResponse, error) {
	claims, err := jwtValidator.ValidateToken(in.Token, s.secret)
	if err != nil {
		slog.Error("failed to validate token", slog.Any("error", err))
		return &auth.ValidateResponse{}, err
	}
	return &auth.ValidateResponse{Valid: true, Username: claims.Username}, nil
}
