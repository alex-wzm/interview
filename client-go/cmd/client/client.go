package main

import (
	"context"

	"interview-client/config"
	"interview-client/internal/consumer"
	"interview-client/internal/logger"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func makeAuthCall(env *config.EnvConfig) *grpc.ClientConn {
	ctx := context.Background()
	authconn, err := grpc.DialContext(
		ctx,
		env.Authserver,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)

	if err != nil {
		logger.LogError(errors.Wrap(err, "failed to connect to authserver"), true)
	}

	return authconn
}

func main() {
	ctx := context.Background()
	env := config.LoadEnv()

	apiconn, err := grpc.DialContext(
		ctx,
		env.Apiserver,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		logger.LogError(errors.Wrap(err, "failed to connect to service"), true)
	}

	authconn := makeAuthCall(env)
	defer func() {
		apiconn.Close()
		authconn.Close()
	}()

	consumer := consumer.New(apiconn, authconn)
	logger.LogInfo("authenticating and dialing service")
	consumer.Authenticate(ctx, env)
	consumer.HelloWorld(ctx)
}
