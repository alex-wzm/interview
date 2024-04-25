package consumer

import (
	"context"
	"fmt"
	"interview-client/internal/api/interview"
	"log"

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

func (s *consumer) HelloWorld(ctx context.Context, pdms int) {
	resp, err := s.client.HelloWorld(ctx, &interview.HelloWorldRequest{
		PublishDelayMS: int32(pdms),
	})
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to hello world"))
	}
	fmt.Println(resp)
}
