package main

import (
	"context"
	"flag"
	"interview-client/internal/consumer"
	"log"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	name := flag.String("name", "Foo Fighters", "Pass arbitrary string to gRPC call request")
	flag.Parse()
	ctx := context.Background()
	serviceAddress := "127.0.0.1:8080"
	conn, err := grpc.DialContext(
		ctx,
		serviceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to connect to service"))
	}
	defer conn.Close()
	consumer := consumer.New(conn)
	consumer.HelloWorld(ctx, *name)
}
