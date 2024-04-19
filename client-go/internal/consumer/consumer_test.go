package consumer_test

import (
	"context"
	"testing"

	"interview-client/internal/api/interview"
	"interview-client/internal/consumer"
)

type testCase struct {
	name         string
	request      *interview.HelloWorldRequest
	expectedResp *interview.HelloWorldResponse
	expectedErr  error
}

// TestHelloWorld uses a table-driven test to check the behavior of the HelloWorld function.
func TestHelloWorld(t *testing.T) {
	c := consumer.GetConsumer() // This assumes that you have a method to get the consumer instance.

	tests := []testCase{
		{
			name:         "Empty Request",
			request:      &interview.HelloWorldRequest{Name: ""},
			expectedResp: &interview.HelloWorldResponse{Greeting: "Hello, World!"},
			expectedErr:  nil,
		},
		{
			name:         "Non-Empty Request",
			request:      &interview.HelloWorldRequest{Name: "Alice"},
			expectedResp: &interview.HelloWorldResponse{Greeting: "Hello, Alice!"},
			expectedErr:  nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := c.HelloWorld(context.Background(), tc.request)
			if err != tc.expectedErr {
				t.Errorf("Test %s failed: expected error %v, got %v", tc.name, tc.expectedErr, err)
			}
			if err == nil && resp.GetGreeting() != tc.expectedResp.GetGreeting() {
				t.Errorf("Test %s failed: expected response %v, got %v", tc.name, tc.expectedResp.GetGreeting(), resp.GetGreeting())
			}
		})
	}
}
