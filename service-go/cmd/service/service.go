package service

import (
	"context"
	"fmt"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	jwtValidator "interview-service/internal/domain/jwt"
	"log"
	"net"

	"interview-service/internal/api"
	"interview-service/internal/api/interview"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Start() {
	address := fmt.Sprintf("localhost:%d", 8080)
	log.Printf("Starting interview service at %s", address)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	{
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(AuthFunc([]byte("secret"))))
	}

	grpcServer := grpc.NewServer(opts...)
	interview.RegisterInterviewServiceServer(grpcServer, api.New())
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}

const authHeader = "authorization"

func AuthFunc(secret []byte) func(ctx context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, err
		}

		claims, err := jwtValidator.ValidateToken(token, secret)
		if err != nil {
			log.Default().Println(err)
			return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
		}
		ctx = context.WithValue(ctx, authHeader, claims)
		return ctx, nil
	}
}
