package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"interview-client/internal/consumer"
	"log"
	"os"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

func main() {
	ctx := context.Background()

	config := loadConfig()
	// Verify the server cert is in prod.
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	creds := credentials.NewTLS(tlsConfig)

	conn, err := grpc.DialContext(
		ctx,
		config.Server,
		grpc.WithTransportCredentials(creds),
		grpc.WithBlock(),
	)

	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to connect to service"))
	}

	consumer := consumer.New(conn)

	consumer.HelloWorld(ctx)
}
