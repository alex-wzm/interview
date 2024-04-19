package authorizer

import (
	"context"
	auth "interview-client/internal/api/authservice" // Assuming this is the generated package for AuthService
)

type AuthClient struct {
	client auth.AuthServiceClient
}

// NewAuthServiceClient initializes a new AuthService client
func New(client auth.AuthServiceClient) (*AuthClient, error) {
	return &AuthClient{
		client: client,
	}, nil
}

// ClientLogin calls login from auth server and returns a JWT token
func (a *AuthClient) ClientLogin(ctx context.Context, username string) (string, error) {
	resp, err := a.client.Login(ctx, &auth.LoginRequest{
		Username: username,
	})
	if err != nil {
		return "", err
	}
	return resp.Token, nil
}
