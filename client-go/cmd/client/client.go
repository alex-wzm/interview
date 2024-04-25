package main

import (
	"context"
	"encoding/json"
	"interview-client/internal/consumer"
	"interview-client/internal/jwt"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
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

	username := "someuser"

	duration, _ := time.ParseDuration("10h")
	jwtToken, err := jwt.GenerateToken(username, duration, []byte("secret"))
	if err != nil {
		log.Default().Println("GenerateToken error: " + err.Error())
	}
	md := metadata.Pairs("Authorization", "Bearer "+jwtToken)
	ctx = metadata.NewOutgoingContext(ctx, md)

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
