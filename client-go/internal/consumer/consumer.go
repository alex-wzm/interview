package consumer

import (
	"context"
	"fmt"
	"interview-client/internal/api/interview"
	"log"
	"sync"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type consumer struct {
	interview.UnimplementedInterviewServiceServer
	Client interview.InterviewServiceClient
}

var (
	once     sync.Once
	instance *consumer
)

func New(c *grpc.ClientConn) *consumer {
	once.Do(func() {
		instance = &consumer{
			Client: interview.NewInterviewServiceClient(c),
		}
	})
	return instance
}

func GetConsumer() *consumer {
	return instance
}

func (s *consumer) HelloWorld(ctx context.Context, req *interview.HelloWorldRequest) (*interview.HelloWorldResponse, error) {
	resp, err := s.Client.HelloWorld(ctx, req)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to execute HelloWorld"))
		return nil, err
	}
	fmt.Println(resp)
	return resp, nil
}
