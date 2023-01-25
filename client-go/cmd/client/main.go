package main

import (
	"context"
	"encoding/json"
	"interview-client/internal/consumer"
	"log"
	"os"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	config configuration
)

type configuration struct {
	Server string `json:"Server"`
}

func init() {
	f, err := os.Open("local.json")
	defer f.Close()
	if err != nil {
		log.Fatalln("failed to open config file: ", err)
	}

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("failed to decode config file: ", err)
	}
}

func main() {
	ctx := context.Background()
	conn, err := grpc.DialContext(
		ctx,
		config.Server,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to connect to service"))
	}
	consumer := consumer.New(conn)
	consumer.HelloWorld(ctx)
}
