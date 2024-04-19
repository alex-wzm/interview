package consumer

import (
	"context"
	"fmt"
	interview "interview-client/internal/api/interview"
	mock_interview "interview-client/internal/api/mock_interview"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestClientHelloWorld(t *testing.T) {

	type TestParameters struct {
		ctx       context.Context
		mockSetup func(*mock_interview.MockInterviewServiceClient)
	}

	type TestResults struct {
		error error
	}

	type TestScenario struct {
		description string
		setupFn     func(t *testing.T) TestParameters
		expectFn    func(t *testing.T, results TestResults)
	}

	testScenarios := []TestScenario{
		{
			description: "Happy path",
			setupFn: func(t *testing.T) TestParameters {
				return TestParameters{
					mockSetup: func(mockClient *mock_interview.MockInterviewServiceClient) {
						mockClient.EXPECT().
							HelloWorld(gomock.Any(), &interview.HelloWorldRequest{}).
							Return(&interview.HelloWorldResponse{Greeting: "Hello , stranger!"}, nil)
					},
				}
			},
			expectFn: func(t *testing.T, results TestResults) {
				if results.error != nil {
					t.Errorf("Expected no error, got %v", results.error)
				}
			},
		},
		{
			description: "unknown error from interview client",
			setupFn: func(t *testing.T) TestParameters {
				return TestParameters{
					mockSetup: func(mockClient *mock_interview.MockInterviewServiceClient) {
						mockClient.EXPECT().
							HelloWorld(gomock.Any(), &interview.HelloWorldRequest{}).
							Return(&interview.HelloWorldResponse{}, fmt.Errorf("error from interview service"))
					},
				}
			},
			expectFn: func(t *testing.T, results TestResults) {
				if results.error == nil {
					t.Errorf("Expected Error, but passed sucesfully")
				}
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, scenario := range testScenarios {
		t.Run(scenario.description, func(t *testing.T) {
			mockClient := mock_interview.NewMockInterviewServiceClient(ctrl)
			consumer := New(mockClient)
			params := scenario.setupFn(t)
			params.mockSetup(mockClient)
			err := consumer.ClientHelloWorld(context.Background())
			scenario.expectFn(t, TestResults{error: err})
		})
	}

}
