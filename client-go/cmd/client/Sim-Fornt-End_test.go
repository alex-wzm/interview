package main_test

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"testing"

	"interview-client/internal/consumer"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type config struct {
	Server string `json:"Server"`
}

func loadConfig() (c config) {
	f, err := os.Open("./configs/local.json")
	defer f.Close()
	if err != nil {
		log.Fatalln("failed to open config file: ", err)
	}

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		log.Fatalln("failed to decode config file: ", err)
	}

	return c
}

func TestMain(m *testing.M) {
	ctx := context.Background()
	conn, err := setupTestConnection(ctx)
	if err != nil {
		log.Fatalln("failed to connect to test service:", err)
	}
	defer conn.Close()

	consumer.New(conn)

	os.Exit(m.Run())
}

// TODO add JWT auth so that the client can properly run without issues in the unary interceptor.
func setupTestConnection(ctx context.Context) (*grpc.ClientConn, error) {
	config := loadConfig()

	return grpc.DialContext(
		ctx,
		config.Server,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
}
