package validator

import (
	"context"
	"errors"
	"fmt"
	"interview-service/internal/api/authservice"
	auth "interview-service/internal/api/authservice"
	"log"
	"os"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
)

type AuthServiceClient struct {
	client auth.AuthServiceClient
}

// NewAuthServiceClient initializes a new AuthService client
func NewAuthServiceClient(client authservice.AuthServiceClient) (*AuthServiceClient, error) {
	return &AuthServiceClient{
		client: client,
	}, nil
}

// Calls the ValidateJWT function in Auth service
func (a *AuthServiceClient) ValidateJWTToken() func(ctx context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		log.SetOutput(os.Stdout)
		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		fmt.Println()
		if err != nil {
			return nil, err
		}

		resp, err := a.client.ValidateJWT(ctx, &auth.ValidateRequest{Token: token})
		if err != nil {
			return nil, err
		}

		if !resp.Valid {
			return nil, errors.New("invalid token")
		}

		ctx = context.WithValue(ctx, username, resp.Username)

		return ctx, nil
	}
}

const username = "username"
