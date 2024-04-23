package main

import (
	"context"
	"fmt"
	"interview-authserver/config"
	"interview-authserver/internal/api"
	"interview-authserver/internal/api/auth"
	"interview-authserver/internal/logger"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	configPath = "./config/grpc.json"
)

func main() {
	//grpcConfig := config.LoadConfigFromFile(configPath)
	env := config.LoadEnv()

	address := fmt.Sprintf("%s:%s", env.Host, env.Port)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		logger.LogError(fmt.Errorf("failed to listen: %v", err), true)
	}

	opts := []grpc.ServerOption{}

	grpcServer := grpc.NewServer(opts...)

	w := WrappedAuth{}
	apiserver := api.New(env, w)

	auth.RegisterAuthServiceServer(grpcServer, apiserver)
	reflection.Register(grpcServer)

	logger.LogInfo(fmt.Sprintf("Starting authserver at %s", address))
	grpcServer.Serve(lis)

}

type WrappedAuth struct {
}

func (w WrappedAuth) Authenticate(ctx context.Context, username string, env *config.EnvConfig) (string, error) {
	return auth.Authenticate(ctx, username, env)
}
