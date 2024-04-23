package consumer

import (
	"context"
	"encoding/base64"
	"fmt"
	"interview-client/config"
	"interview-client/internal/api/auth"
	"interview-client/internal/api/interview"
	"interview-client/internal/logger"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type consumer struct {
	interview.UnimplementedInterviewServiceServer
	client     interview.InterviewServiceClient
	authClient auth.AuthServiceClient
	token      string
}

func New(apiconn *grpc.ClientConn, authconn *grpc.ClientConn) *consumer {
	return &consumer{
		client:     interview.NewInterviewServiceClient(apiconn),
		authClient: auth.NewAuthServiceClient(authconn),
	}
}

func (s *consumer) HelloWorld(ctx context.Context) {
	md := metadata.Pairs("authorization", "Bearer "+s.token)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp, err := s.client.HelloWorld(ctx, &interview.HelloWorldRequest{})
	if err != nil {
		logger.LogError(errors.Wrap(err, "failed to hello world"), true)
	}
	fmt.Println(resp.Greeting)
	logger.LogInfo(resp.Greeting)
}

func (s *consumer) Authenticate(ctx context.Context, env *config.EnvConfig) {
	user, pass := env.User, env.Pass
	secretbytes := []byte(fmt.Sprintf("%s:%s", user, pass))
	encodedstr := base64.StdEncoding.EncodeToString(secretbytes)
	md := metadata.Pairs("authorization", "Basic "+encodedstr)
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, authErr := s.authClient.Authenticate(ctx, &auth.AuthRequest{Username: user})
	if authErr != nil {
		logger.LogError(errors.Wrap(authErr, "failed to authenticate"), true)
	}

	logger.LogInfo("successful client authentication; storing token")
	s.token = res.Token
}
