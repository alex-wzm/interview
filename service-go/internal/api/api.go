package api

import (
	"context"
	"fmt"
	"interview-service/internal/api/interview"
	"interview-service/internal/domain/greeter"
)

type server struct {
	interview.UnimplementedInterviewServiceServer
}

func New() *server {
	return &server{}
}

func (s *server) HelloWorld(ctx context.Context, r *interview.HelloWorldRequest) (*interview.HelloWorldResponse, error) {
	username, ok := ctx.Value("username").(string)
	// ensure the value from the context is our expected data type
	if !ok {
		return &interview.HelloWorldResponse{}, fmt.Errorf("unexpected error occured when retrieving username from context")
	}

	return &interview.HelloWorldResponse{
		Greeting: greeter.Greet(username),
	}, nil
}
