package main

import (
	"log/slog"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"authservice-go/internal/api"
	"authservice-go/internal/api/auth"
	"authservice-go/internal/config"
)

var Version = "develop"

func main() {
	cfg, cfgErr := config.LoadFromEnv()
	if cfgErr != nil {
		slog.Error("failed to load config from environment", slog.Any("error", cfgErr))
		os.Exit(1)
	}
	// set up the default log/slog configuration using JSON output for centralized logging services
	// enable Debug output if DEBUG set in environment
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: func() slog.Leveler {
			if cfg.Debug {
				return slog.LevelDebug
			}
			return slog.LevelInfo
		}(),
	})).With(slog.String("version", Version)))

	lis, listenErr := net.Listen("tcp", cfg.GRPC.ListenAddr)
	if listenErr != nil {
		slog.Error("failed to listen", slog.Any("error", listenErr), slog.String("address", cfg.GRPC.ListenAddr))
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()
	auth.RegisterAuthServiceServer(grpcServer, api.New(cfg.Users, []byte(cfg.JWTSecret)))
	reflection.Register(grpcServer)

	slog.Info("Starting authentication service", slog.String("address", cfg.GRPC.ListenAddr))
	if e := grpcServer.Serve(lis); e != nil {
		slog.Error("GRPC server failure", slog.Any("error", e))
		os.Exit(1)
	}
}
