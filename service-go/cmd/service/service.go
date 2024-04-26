package main

import (
	"context"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/credentials/insecure"
	"interview-service/internal/config"
	"interview-service/internal/domain/jwt"
	"log/slog"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"interview-service/internal/api"
	"interview-service/internal/api/interview"
)

var Version = "develop"

func main() {
	appCtx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	appCfg, loadErr := config.LoadFromEnv()
	if loadErr != nil {
		slog.Error("failed to load the app config", slog.Any("error", loadErr))
		os.Exit(1)
	}
	// set the default log/slog handler, this will take over so the entire application uses it without having to pass a
	// `logger` around.
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: func() slog.Leveler {
			if appCfg.Debug {
				return slog.LevelDebug
			}
			return slog.LevelInfo
		}(),
	})).With(slog.String("version", Version)))

	// create a dial context for our auth service, this will allow us to validate incoming JWT
	authConn, authConnErr := grpc.DialContext(appCtx,
		appCfg.GRPC.AuthServerAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if authConnErr != nil {
		slog.Error("failed to connect to auth server", slog.Any("error", authConnErr))
		os.Exit(1)
	}
	defer func() { _ = authConn.Close() }()

	// create a new JWT validator, this will call our AuthService to verify
	validator := jwt.New(authConn)

	// create the low level TCP listener for our GRPC service
	lis, listenErr := net.Listen("tcp", appCfg.GRPC.ServerAddr)
	if listenErr != nil {
		slog.Error("failed to listen", slog.Any("error", listenErr), slog.String("address", appCfg.GRPC.ServerAddr))
		os.Exit(1)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(validator.ValidateJWT())),
	)

	interview.RegisterInterviewServiceServer(grpcServer, api.New())
	reflection.Register(grpcServer)

	slog.Info("Starting interview service", slog.String("address", appCfg.GRPC.ServerAddr))
	if e := grpcServer.Serve(lis); e != nil {
		slog.Error("GRPC server failure", slog.Any("error", e))
		os.Exit(1)
	}
}
