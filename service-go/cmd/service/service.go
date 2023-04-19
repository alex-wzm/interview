package main

import (
	"context"
	"fmt"
	"net"
	"os"

	config "interview-service/config"

	log "github.com/sirupsen/logrus"

	"interview-service/internal/api"
	"interview-service/internal/api/interview"
	jwt "interview-service/internal/domain/jwt"

	"flag"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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
	grpcConfig := config.LoadConfigFromFile(configPath)
	address := fmt.Sprintf("%s:%s", grpcConfig.ServerHost, grpcConfig.UnsecurePort)
	log.Infof("Starting interview service at %s", address)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var jwtSecret = os.Getenv("JWT_SECRET")

	if jwtSecret == "" {
		log.Fatalf("error loading secret from envoirnment")
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			grpc_auth.UnaryServerInterceptor(validateJWT([]byte(jwtSecret))),
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
	configPath = "./config/grpc.json"
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
<<<<<<< HEAD
			//log.Default().Println(err)
			log.WithError(err).Debug("JWT validation failed")
=======
			// log.Default().Println(err)
			log.Debugf("Error %v", err)
>>>>>>> 94725c2 (PR feedback changes)
			return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
		}

		ctx = context.WithValue(ctx, authHeader, claims)

		return ctx, nil
	}
}
