package main

import (
	"context"
	"fmt"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	jwtValidator "interview-service/internal/domain/jwt"
	"log"
	"net"

	log "github.com/sirupsen/logrus"

	"interview-service/internal/api"
	"interview-service/internal/api/interview"

	"flag"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	loglevel := flag.Int("logLevel", 4, "Useful Log levels: Warn = 3; Info = 4; Debug = 5;")
	logGrpc := flag.Bool("logGrpc", false, "Turn ON/OFF grpc middleware logs")
	flag.Parse()

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	if *loglevel >= 0 && *loglevel <= 6 {
		log.SetLevel(log.Level(*loglevel))
	}

	if *logGrpc {
		err := godotenv.Load("logGrpc.env")
		if err != nil {
			log.Infof("Error loading .env file: %+v", err)
		}
		log.Debugf("Loaded .env file")
		log.Debugf("Severity: %s", os.Getenv("GRPC_GO_LOG_SEVERITY_LEVEL"))
		log.Debugf("Verbosity: %s", os.Getenv("GRPC_GO_LOG_VERBOSITY_LEVEL"))
	}

	start()
}

func start() {
	address := fmt.Sprintf("localhost:%d", 8080)
	log.Infof("Starting interview service at %s", address)

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
