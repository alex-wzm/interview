package main

import (
	"context"
	"flag"
	"interview-client/internal/auth"
	"interview-client/internal/config"
	"interview-client/internal/consumer"
	"log/slog"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	username string
	password string
)

func init() {
	flag.StringVar(&username, "u", "", "username")
	flag.StringVar(&password, "p", "", "password")

	flag.Parse()
}

func main() {
	ctx := context.Background()

	cfg, cfgErr := config.LoadFromEnv()
	if cfgErr != nil {
		slog.Error("failed to retrieve config from environment", slog.Any("error", cfgErr))
		os.Exit(1)
	}
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: func() slog.Leveler {
			if cfg.Debug {
				return slog.LevelDebug
			}
			return slog.LevelInfo
		}(),
	})))

	authConn, authDialErr := grpc.DialContext(
		ctx,
		cfg.AuthServerAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if authDialErr != nil {
		slog.Error("failed to connect to auth server", slog.Any("error", authDialErr))
		os.Exit(1)
	}

	authenticator := auth.New(authConn)
	jwt, authErr := authenticator.Auth(ctx, username, password)
	if authErr != nil {
		slog.Error("failed to authenticate", slog.Any("error", authErr))
		os.Exit(1)
	}
	defer func() { _ = authConn.Close() }()

	slog.Info("successfully authenticated", slog.String("username", username))

	ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	conn, dialErr := grpc.DialContext(
		ctxTimeout,
		cfg.ServerAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(jwtRPCCredentials(jwt)),
		grpc.WithBlock(),
	)
	if dialErr != nil {
		slog.Error("failed to connect to service", slog.Any("error", dialErr))
		os.Exit(1)
	}
	defer func() { _ = conn.Close() }()

	cons := consumer.New(conn)

	greeting, err := cons.HelloWorld(ctx)
	if err != nil {
		slog.Error("failed to receive greeting", slog.Any("error", err))
	}
	slog.Info("successfully received greeting", slog.String("greeting", greeting))
}

type jwtAuth struct {
	token string
}

func (j *jwtAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "bearer " + j.token,
	}, nil
}

func (j *jwtAuth) RequireTransportSecurity() bool {
	return false
}

func jwtRPCCredentials(jwt string) *jwtAuth {
	return &jwtAuth{token: jwt}
}
