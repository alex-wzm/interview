package auth

import (
	"context"
	"encoding/base64"
	"fmt"
	"interview-authserver/config"
	jwtValidator "interview-authserver/internal/domain/jwt"
	"interview-authserver/internal/logger"
	"strings"
	"time"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
)

func Authenticate(ctx context.Context, username string, env *config.EnvConfig) (token string, err error) {
	basicAuth, authErr := grpc_auth.AuthFromMD(ctx, "basic")
	if authErr != nil {
		logger.LogError(fmt.Errorf("failed to extract auth header for user %v", username), false)
		err = fmt.Errorf("invalid headers")
		return
	}

	dbytes, dErr := base64.StdEncoding.DecodeString(basicAuth)
	if dErr != nil {
		logger.LogError(fmt.Errorf("failed to decode basic auth header for user %v; %v", username, dErr), false)
		err = fmt.Errorf("invalid headers")
		return
	}

	//validating with user/pass in env for simplicity
	usersecret := strings.Split(string(dbytes), ":")[1]
	logger.LogInfo(env.Username + ":" + env.JwtSecret)
	if username != env.Username || usersecret != env.Usersecret {
		logger.LogError(fmt.Errorf("invalid user %v", username), false)
		err = fmt.Errorf("invalid user")
		return
	}

	logger.LogInfo(fmt.Sprintf("user %v authenticated", username))

	t, tErr := jwtValidator.GenerateToken(username, time.Hour, []byte(env.JwtSecret))
	if tErr != nil {
		logger.LogError(fmt.Errorf("failed to generate token [user: %v] [error: %v]", username, tErr), false)
		err = fmt.Errorf("internal error")
		return
	}

	logger.LogInfo(fmt.Sprintf("successfully generated token for user %v; token: %v", username, t))
	token = t
	return
}
