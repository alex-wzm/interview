package consumer

import (
	"context"
	"interview-client/internal/api/interview"

	log "github.com/sirupsen/logrus"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
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

func (s *consumer) HelloWorld(ctx context.Context, name string) {
	req := &interview.HelloWorldRequest{
		Name: name,
	}
	resp, err := s.client.HelloWorld(context.Background(), req)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to hello world"))
	}
	log.Infof("Client request: %v", req)
	log.Infof("Service response: %v", resp)
}
