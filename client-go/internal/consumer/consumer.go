package consumer

import (
	"context"
	"fmt"
	"interview-client/internal/api/interview"
	"interview-client/internal/api/interview/auth"
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

func (s *consumer) HelloWorld(ctx context.Context) {
	resp, err := s.client.HelloWorld(ctx, &interview.HelloWorldRequest{})
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to hello world"))
	}
	fmt.Println(resp)
}

type authorize struct {
	auth.UnimplementedAuthServiceServer
	client auth.AuthServiceClient
}

func NewAuthorize(c *grpc.ClientConn) *authorize {
	return &authorize{
		client: auth.NewAuthServiceClient(c),
	}
}

func (s authorize) Authorize(ctx context.Context, username string, password string, ttl int64) string {
	resp, err := s.client.Authorize(ctx, &auth.AuthRequest{
		Username: username,
		Password: password,
		Ttl: ttl,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return resp.Token

}
