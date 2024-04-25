package auth

import (
	"context"
	jwt "interview-service/internal/domain/jwt"
	"log"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth struct {
	authHeader string
	secret     []byte
}

func NewAuth(authHeader string, secret []byte) *Auth {
	return &Auth{authHeader: authHeader, secret: secret}
}

func (a *Auth) ValidateJWT() func(ctx context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			log.Println("ValidateJWT: AuthFromMD: " + err.Error())
			return nil, err
		}

		claims, err := jwt.ValidateToken(token, a.secret)
		if err != nil {
			log.Default().Println("ValidateJWT: ValidateToken: " + err.Error())
			return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
		}

		ctx = context.WithValue(ctx, a.authHeader, claims)

		return ctx, nil
	}
}
