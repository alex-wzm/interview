package jwt

import (
	"context"
	"fmt"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"interview-service/internal/api/auth"
)

type Validator struct {
	auth.UnimplementedAuthServiceServer
	client auth.AuthServiceClient
}

func New(c *grpc.ClientConn) *Validator {
	return &Validator{
		client: auth.NewAuthServiceClient(c),
	}
}

func (v *Validator) ValidateJWT() func(ctx context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, err
		}
		resp, err := v.client.Validate(ctx, &auth.ValidateRequest{
			Token: token,
		})
		if err != nil {
			return nil, err
		}
		if !resp.Valid {
			return nil, fmt.Errorf("provided token is invalid")
		}
		ctx = context.WithValue(ctx, "username", resp.Username)
		return ctx, nil
	}
}
