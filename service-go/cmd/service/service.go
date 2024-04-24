package main

import (
	"context"
	"fmt"

	"log"
	"net"

	config "interview-service/config"
	"interview-service/internal/api"
	"interview-service/internal/api/interview"
	jwt "interview-service/internal/domain/jwt"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func main() {
	appCfg, loadErr := config.Load()
	if loadErr != nil {
		log.Fatalf("failed to load the app config: %v", loadErr)
	}

	address := fmt.Sprintf("%s:%s", appCfg.GRPC.ServerHost, appCfg.GRPC.UnsecurePort)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			grpc_auth.UnaryServerInterceptor(validateJWT([]byte(appCfg.JWTSecret))),
		),
	}

	grpcServer := grpc.NewServer(opts...)

	interview.RegisterInterviewServiceServer(grpcServer, api.New())
	reflection.Register(grpcServer)

	log.Printf("Starting interview service at %s", address)
	grpcServer.Serve(lis)

}

const (
	authHeader = "authorization"
)

// validateJWT parses and validates a bearer jwt
//
// TODO: move to own package (in ./internal/api/auth) using a constructor that privately sets the secret
func validateJWT(secret []byte) func(ctx context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, err
		}

		claims, err := jwt.ValidateToken(token, secret)
		if err != nil {
			log.Default().Println(err)
			return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
		}

		ctx = context.WithValue(ctx, authHeader, claims)

		return ctx, nil
	}
}
