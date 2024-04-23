package auth

import (
	"context"
	"fmt"
	"interview-service/config"
	jwtValidator "interview-service/internal/domain/jwt"
	"interview-service/internal/logger"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthHelper struct {
	env *config.EnvConfig
}

const (
	authHeader = "authorization"
)

func New(env *config.EnvConfig) *AuthHelper {
	return &AuthHelper{env: env}
}

func (a *AuthHelper) ValidateJWT() func(ctx context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, err
		}

		claims, err := jwtValidator.ValidateToken(token, []byte(a.env.JwtSecret))
		if err != nil {
			logger.LogError(err, false)
			return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
		}

		logger.LogInfo("jwt validated")
		ctx = context.WithValue(ctx, authHeader, claims)
		return ctx, nil
	}
}

func (a *AuthHelper) AuthorizeRequest() func(ctx context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		authHeader := ctx.Value(authHeader)
		claims, ok := authHeader.(*jwtValidator.JWTClaims)
		if !ok {
			logger.LogError(fmt.Errorf("failed to process claims during authorization; %v", authHeader), false)
			return nil, status.Error(codes.PermissionDenied, "authorization failed")
		}

		if vErr := claims.Valid(); vErr != nil {
			logger.LogError(fmt.Errorf("expired claims %v; %v", authHeader, vErr), false)
			return nil, status.Error(codes.PermissionDenied, "invalid claims")
		}

		//authorizing user based on env user creds for simplicity
		if claims.Username != a.env.Admin {
			logger.LogError(fmt.Errorf("unauthorized user %v", claims.Username), false)
			return nil, status.Error(codes.PermissionDenied, "unauthorized")
		}

		logger.LogInfo(fmt.Sprintf("authorized api access for user %v", claims.Username))
		return ctx, nil
	}
}
