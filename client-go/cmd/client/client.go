package main

import (
	"context"
	"encoding/json"
	"flag"
	"interview-client/internal/consumer"
	"log"
	"os"

	"github.com/pkg/errors"
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

func main() {
	name := flag.String("name", "Foo Fighters", "Pass arbitrary string to gRPC call request")
	flag.Parse()
	ctx := context.Background()

	config := loadConfig()

	conn, err := grpc.DialContext(
		ctx,
		config.Server,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to connect to service"))
	}
<<<<<<< HEAD:client-go/cmd/client/client.go

=======
	defer conn.Close()
>>>>>>> e3cc5eb (PR feedback changes):client-go/cmd/client/main.go
	consumer := consumer.New(conn)
	consumer.HelloWorld(ctx, *name)
}
