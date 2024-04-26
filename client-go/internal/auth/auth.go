package auth

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"interview-client/internal/api/auth"
)

type Auth struct {
	auth.UnimplementedAuthServiceServer
	client auth.AuthServiceClient
}

func New(c *grpc.ClientConn) *Auth {
	return &Auth{client: auth.NewAuthServiceClient(c)}
}

func (a *Auth) Auth(ctx context.Context, username, password string) (string, error) {
	resp, err := a.client.Auth(ctx, &auth.AuthRequest{Username: username, Password: password})
	if err != nil {
		return "", fmt.Errorf("login failed: %w", err)
	}
	return resp.Token, nil
}
