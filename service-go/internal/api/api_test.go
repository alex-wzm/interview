package api

import (
	"context"
	"testing"

	"interview-service/internal/api/interview"
	"interview-service/internal/domain/greeter"
)

func TestServer_HelloWorld(t *testing.T) {
	type TestResults struct {
		response *interview.HelloWorldResponse
		err      error
	}

	type TestParameters struct {
		ctx context.Context
	}

	type TestScenario struct {
		description string
		setupFn     func(t *testing.T) TestParameters
		expectFn    func(t *testing.T, results TestResults)
	}

	testScenarios := []TestScenario{
		{
			description: "formats greeting with given username",
			setupFn: func(t *testing.T) TestParameters {
				username := "Alice"
				ctx := context.WithValue(context.Background(), "username", username)
				return TestParameters{ctx: ctx}
			},
			expectFn: func(t *testing.T, results TestResults) {
				expectedGreeting := greeter.Greet("Alice")
				if results.response.Greeting != expectedGreeting || results.err != nil {
					t.Errorf("Unexpected result: got %v, want %v, error: %v", results.response.Greeting, expectedGreeting, results.err)
				}
			},
		},
		{
			description: "handles missing username context gracefully",
			setupFn: func(t *testing.T) TestParameters {
				ctx := context.Background() // no username set
				return TestParameters{ctx: ctx}
			},
			expectFn: func(t *testing.T, results TestResults) {
				expectedGreeting := greeter.Greet("") // Assuming default behavior is to handle nil or empty username
				if results.response.Greeting != expectedGreeting || results.err != nil {
					t.Errorf("Unexpected result: got %v, want %v, error: %v", results.response.Greeting, expectedGreeting, results.err)
				}
			},
		},
	}

	for _, scenario := range testScenarios {
		t.Run(scenario.description, func(t *testing.T) {
			server := New()
			params := scenario.setupFn(t)

			// Execute the HelloWorld method
			result, err := server.HelloWorld(params.ctx, &interview.HelloWorldRequest{})

			// Pass the results to the expect function
			scenario.expectFn(t, TestResults{
				response: result,
				err:      err,
			})
		})
	}
}
