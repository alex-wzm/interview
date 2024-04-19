package consumer

import (
	"context"
	"fmt"
	"interview-client/internal/api/interview"
)

type consumer struct {
	interview.UnimplementedInterviewServiceServer
	client interview.InterviewServiceClient
}

// Create a new instance for consumer
func New(client interview.InterviewServiceClient) *consumer {
	return &consumer{
		client: client,
	}
}

// Calls HelloWorld funciton from interview-service
func (s *consumer) ClientHelloWorld(ctx context.Context) error {
	resp, err := s.client.HelloWorld(context.Background(), &interview.HelloWorldRequest{})
	if err != nil {
		return err
	}
	fmt.Println(resp)
	return nil
}
