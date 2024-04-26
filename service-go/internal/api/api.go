package api

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"interview-service/internal/api/interview"
	"interview-service/internal/domain/greeter"
)

type server struct {
	interview.UnimplementedInterviewServiceServer
	resultProjectID string
	resultTopicID   string
}

func New(projectID, topicID string) *server {
	return &server{
		resultProjectID: projectID,
		resultTopicID:   topicID,
	}
}

func (s *server) HelloWorld(ctx context.Context, r *interview.HelloWorldRequest) (*interview.HelloWorldResponse, error) {
	greeting := greeter.Greet(r.GetName())

	// simulate some async work with pubsub here
	go func() {
		pubCtx := context.Background()
		client, err := pubsub.NewClient(pubCtx, s.resultProjectID)
		if err != nil {
			log.Println(fmt.Errorf("pubsub: NewClient: %w", err))
		}
		defer client.Close()
		// TODO - randomize with max value from request data
		publishDelay := r.PublishDelayMS
		time.Sleep(time.Millisecond * time.Duration(publishDelay))
		helloResult := struct {
			Greeting       string
			PublishDelayMS int
		}{
			Greeting:       greeting,
			PublishDelayMS: int(publishDelay),
		}

		resultBytes, err := json.Marshal(helloResult)
		if err != nil {
			log.Println(fmt.Errorf("json: Marshal: %w", err))
		}
		t := client.Topic(s.resultTopicID)
		result := t.Publish(pubCtx, &pubsub.Message{
			Data: resultBytes,
		})
		id, err := result.Get(pubCtx)
		if err != nil {
			log.Println(fmt.Errorf("pubsub: Publish: %w", err))
		}
		log.Println(fmt.Sprintf("Published a message; msg ID: %v\n", id))
	}()

	return &interview.HelloWorldResponse{
		Greeting: greeting,
	}, nil
}
