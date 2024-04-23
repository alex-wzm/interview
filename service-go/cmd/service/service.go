package main

import (
	"fmt"
	"log"
	"net"

	config "interview-service/config"
	"interview-service/internal/api"
	"interview-service/internal/api/auth"
	"interview-service/internal/api/interview"
	"interview-service/internal/logger"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	//grpcConfig := config.LoadConfigFromFile(configPath)
	env := config.LoadEnv()

	address := fmt.Sprintf("%s:%s", env.Host, env.Port)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		logger.LogError(fmt.Errorf("failed to listen: %v", err), true)
	}

	authHelper := auth.New(env)

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			grpc_auth.UnaryServerInterceptor(authHelper.ValidateJWT()),
			grpc_auth.UnaryServerInterceptor(authHelper.AuthorizeRequest()),
		),
	}

	grpcServer := grpc.NewServer(opts...)

	interview.RegisterInterviewServiceServer(grpcServer, api.New())
	reflection.Register(grpcServer)

	log.Printf("Starting interview service at %s", address)
	grpcServer.Serve(lis)

}

const (
	configPath = "./config/grpc.json"
)
