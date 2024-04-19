package main

import (
	"fmt"
	"os"

	"log"
	"net"

	config "auth-service/config"
	"auth-service/internal/api"
	auth "auth-service/internal/api/authservice"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	//load the configuarations
	grpcConfig := config.LoadConfigFromFile(configPath)

	address := fmt.Sprintf("%s:%s", grpcConfig.ServerHost, grpcConfig.UnsecurePort)

	//listen to the ip address
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//retrieve secrets from environment
	var jwtSecret = os.Getenv("JWT_SECRET")

	//create a new grpcServer
	grpcServer := grpc.NewServer()

	authService := api.NewAuthServiceServer(jwtSecret)
	auth.RegisterAuthServiceServer(grpcServer, authService)
	reflection.Register(grpcServer)

	log.Printf("Starting auth service at %s", address)
	grpcServer.Serve(lis)

}

const (
	configPath = "../../config/grpc.json"
)
