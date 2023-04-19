package main

import (
	"context"
	"encoding/json"
	"flag"
	"interview-client/internal/consumer"
	"os"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type config struct {
	Server string `json:"Server"`
}

func loadConfig() (*config, error) {
	f, err := os.Open("./configs/local.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var c config
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func main() {
	name := flag.String("name", "Foo Fighters", "Pass arbitrary string to gRPC call request")
	loglevel := flag.Int("logLevel", 4, "Useful Log levels: Warn = 3; Info = 4; Debug = 5;")
	flag.Parse()
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	if *loglevel >= 0 && *loglevel <= 6 {
		log.SetLevel(log.Level(*loglevel))
	}

	config, err := loadConfig()
	if err != nil {
		log.WithError(err).Fatal("failed to load client config")
	}
	ctx := context.Background()
	conn, err := grpc.DialContext(
		ctx,
		config.Server,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.WithError(err).Fatalln("connection to service failed")
	}
	defer conn.Close()
	consumer := consumer.New(conn)
	consumer.HelloWorld(ctx, *name)
}
