package main

import (
	"crypto/tls"
	"fmt"

	"log"
	"net"

	config "interview-service/config"
	"interview-service/internal/api"
	"interview-service/internal/api/interview"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

const (
	configPath = "./config/grpc.json"
)

func main() {

	grpcConfig := config.LoadConfigFromFile(configPath)

	address := fmt.Sprintf("%s:%s", grpcConfig.ServerHost, grpcConfig.UnsecurePort)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	tlsCert, err := tls.LoadX509KeyPair("../certs/server-crt.pem", "../certs/server-key.pem")
	if err != nil {
		log.Fatalf("failed to load server TLS certificate: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&tlsCert)),
	}

	grpcServer := grpc.NewServer(opts...)

	interview.RegisterInterviewServiceServer(grpcServer, api.New())
	reflection.Register(grpcServer)

	log.Printf("Starting interview service at %s", address)
	grpcServer.Serve(lis)

}
