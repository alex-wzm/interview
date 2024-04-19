package main

import (
	"context"
	"fmt"

	"log"
	"net"

	config "interview-service/config"
	"interview-service/internal/api"
	"interview-service/internal/api/authservice"
	"interview-service/internal/api/interview"
	"interview-service/internal/api/validator"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {

	grpcConfig := config.LoadConfigFromFile(configPath)

	address := fmt.Sprintf("%s:%s", grpcConfig.ServerHost, grpcConfig.UnsecurePort)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ctx := context.Background()

	//Establish Connection to Auth server
	AuthConn, err := grpc.DialContext(
		ctx,
		grpcConfig.AuthServer,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to connect to service"))
	}
	defer AuthConn.Close()

	// Initialize AuthServiceClient
	authClient, err := validator.NewAuthServiceClient(authservice.NewAuthServiceClient(AuthConn))
	if err != nil {
		log.Fatalf("Failed to create auth service client: %v", err)
	}

	//UnaryServerInterceptor is to validate each request before executing the HelloWorld
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			grpc_auth.UnaryServerInterceptor(authClient.ValidateJWTToken()),
		),
	}

	//Create gRPC server
	grpcServer := grpc.NewServer(opts...)
	interviewService := api.New()
	interview.RegisterInterviewServiceServer(grpcServer, interviewService)
	reflection.Register(grpcServer)

	log.Printf("Starting interview service at %s", address)
	grpcServer.Serve(lis)

}

const (
	configPath = "../../config/grpc.json"
)
