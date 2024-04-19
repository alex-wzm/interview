package main

import (
	"context"
	"encoding/json"
	"interview-client/internal/api/authservice"
	"interview-client/internal/api/interview"
	"interview-client/internal/authorizer"
	"interview-client/internal/consumer"

	"log"
	"os"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type config struct {
	Server     string `json:"Server"`
	AuthServer string `json:"AuthServer"`
}

type jwtCredentials struct {
	token string
}

func attachToken(token string) credentials.PerRPCCredentials {
	return &jwtCredentials{token: token}
}

// PerRPCCredentials is an interface that abstracts GetRequestMetadata and RequireTransportSecurity
func (c *jwtCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer " + c.token,
	}, nil
}

func (c *jwtCredentials) RequireTransportSecurity() bool {
	return false // set to false as we are not using TLS
}

// to load the configurations
func loadConfig() (c config) {
	f, err := os.Open("../../configs/local.json")
	if err != nil {
		log.Fatalln("failed to open config file: ", err)
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		log.Fatalln("failed to decode config file: ", err)
	}

	return c
}

const username = "admin"

func main() {
	ctx := context.Background()
	config := loadConfig()

	//Establish Connection to Auth server
	AuthConn, err := grpc.DialContext(
		ctx,
		config.AuthServer,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to connect to service"))
	}
	defer AuthConn.Close()

	// Initialize AuthServiceClient
	authClient, err := authorizer.New(authservice.NewAuthServiceClient(AuthConn))
	if err != nil {
		log.Fatalf("Failed to create auth service client: %v", err)
	}

	// Perform login to obtain JWT
	jwtToken, err := authClient.ClientLogin(ctx, username)
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}

	//Establish connection to interview-service
	conn, err := grpc.DialContext(
		ctx,
		config.Server,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(attachToken(jwtToken)),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to connect to service"))
	}
	defer conn.Close()

	consumer := consumer.New(interview.NewInterviewServiceClient(conn))

	err = consumer.ClientHelloWorld(ctx)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "hello world failed"))
	}
}
