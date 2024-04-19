package api

import (
	"context"

	"interview-service/internal/api/interview"
	"interview-service/internal/domain/greeter"
)

type server struct {
	interview.UnimplementedInterviewServiceServer
}

func New() *server {
	return &server{}
}

// HelloWorld Implementation - calls Greet function to send greeting message
func (s *server) HelloWorld(ctx context.Context, r *interview.HelloWorldRequest) (*interview.HelloWorldResponse, error) {
	username, _ := ctx.Value("username").(string)
	greeting := greeter.Greet(username)
	return &interview.HelloWorldResponse{
		Greeting: greeting,
	}, nil
}
