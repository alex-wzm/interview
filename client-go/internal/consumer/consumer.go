package consumer

import (
	"context"
	"google.golang.org/grpc"
	"interview-client/internal/api/interview"
)

type consumer struct {
	interview.UnimplementedInterviewServiceServer
	client interview.InterviewServiceClient
}

func New(c *grpc.ClientConn) *consumer {
	return &consumer{
		client: interview.NewInterviewServiceClient(c),
	}
}

func (s *consumer) HelloWorld(ctx context.Context) (string, error) {
	resp, err := s.client.HelloWorld(ctx, &interview.HelloWorldRequest{})
	if err != nil {
		return "", err
	}
	return resp.Greeting, nil
}
