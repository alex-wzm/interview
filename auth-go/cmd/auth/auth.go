package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"interview-auth/internal/api"
	"interview-auth/internal/api/interview/auth"
	"google.golang.org/grpc/reflection"
)
func main() {

	address := "127.0.0.1:8081"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}


	opts := []grpc.ServerOption{}

	grpcServer := grpc.NewServer(opts...)

	auth.RegisterAuthServiceServer(grpcServer, api.New())
	reflection.Register(grpcServer)

	log.Printf("Starting auth service at %s", address)
	grpcServer.Serve(lis)

}