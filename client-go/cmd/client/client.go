package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/metadata"
	"interview-client/internal/consumer"
	jwt "interview-client/internal/jwt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type config struct {
	Server   string `json:"server"`
	Username string `json:"user"`
	// RequestInterval is the time to wait between sending demo requests to the gRPC service
	RequestIntervalMS int `json:"request_interval_ms"`
	// PublishDelayMaxMS is the maximum time to wait for a message to be published to the pubsub topic
	// The actual time to wait is a random value between 0 and PublishDelayMaxMS
	PublishDelayMaxMS int `json:"publish_delay_max_ms"`
	// ResultProjectID is the GCP project ID where the pubsub topic is located
	ResultProjectID string `json:"result_project_id"`
	// ResultSubID is the pubsub subscription ID to use for receiving publish result messages
	ResultSubID string `json:"result_sub_id"`
	// ResultTopicID is the pubsub topic ID to use for publishing result messages
	ResultTopicID string `json:"result_topic_id"`
	// ResultNumGoroutines is the number of goroutines to use for receiving messages from the pubsub subscription
	ResultNumGoroutines int `json:"result_num_goroutines"`
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
	config := loadConfig()

	// start pubsub subscription for result messages
	go handleResultMessages(config.ResultProjectID, config.ResultTopicID, config.ResultSubID, config.ResultNumGoroutines)

	// generate auth token
	td, err := time.ParseDuration("1h")
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to parse token duration"))
	}
	token, err := jwt.GenerateToken(config.Username, td, []byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to generate token"))
	}
	md := metadata.Pairs("Authorization", "Bearer "+token)
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

	ticker := time.NewTicker(time.Millisecond * time.Duration(config.RequestIntervalMS))
	defer ticker.Stop()
	done := make(chan bool)
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			// create a new client for each request
			consumer := consumer.New(conn)
			// create random delay to simulate processing load on server
			// no need for crypto secure rand with seed
			delay := rand.Intn(config.PublishDelayMaxMS)
			consumer.HelloWorld(ctx, delay)
		}
	}
}

func handleResultMessages(projectID, topicID, subID string, goroutines int) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Println(fmt.Errorf("pubsub.NewClient: %w", err))
	}
	defer client.Close()
	sub := client.Subscription(subID)
	// defaults to 10 if not specified
	if goroutines != 0 {
		sub.ReceiveSettings.NumGoroutines = goroutines
	}
	for {
		err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
			log.Println(fmt.Sprintf("Got message: %q\n", string(msg.Data)))
			msg.Ack()
		})
		if !errors.Is(context.Canceled, err) && err != nil {
			log.Println(fmt.Errorf("pubsub: receive error: %w", err))
		}
	}
}
