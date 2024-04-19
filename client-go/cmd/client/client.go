package main

import (
	"context"
	"interview-client/internal/consumer"
	"log"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type config struct {
	Server string `json:"Server"`
	AuthServer string `json:"AuthServer"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}


func loadConfig() (c config) {
	viper.SetConfigName("configs/local.json")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Fatalln("failed to open config file: ", err)
	}
	return c
}


func main() {

	config := loadConfig()

	token := getToken(context.Background(), config)
	log.Printf("got token %v", token)

	md := metadata.Pairs("authorization", "bearer "+token)
	ctx := metadata.NewOutgoingContext(context.Background(), md)


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

func getToken(ctx context.Context, c config) string {


	log.Printf("Connecting to auth server %v ", c.AuthServer)
	conn, err := grpc.DialContext(
		context.Background(),
		c.AuthServer,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)

	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to connect to service"))
	}
	log.Printf("Success!")

	authServer := consumer.NewAuthorize(conn)
	return authServer.Authorize(ctx, c.Username, c.Password, 10)


}
