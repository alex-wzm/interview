package authorizer

import (
	"context"
	"fmt"
	"interview-client/internal/api/authservice"
	mock_auth "interview-client/internal/api/mock_auth"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestClientLogin(t *testing.T) {
	type TestParameters struct {
		ctx       context.Context
		username  string
		mockSetup func(*mock_auth.MockAuthServiceClient)
	}

	type TestResults struct {
		token string
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
					username: "test-admin",
					mockSetup: func(mockClient *mock_auth.MockAuthServiceClient) {
						mockClient.EXPECT().
							Login(gomock.Any(), gomock.Any()).
							Return(&authservice.LoginResponse{}, nil)
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
			description: "Error returned from auth service",
			setupFn: func(t *testing.T) TestParameters {
				return TestParameters{
					mockSetup: func(mockClient *mock_auth.MockAuthServiceClient) {
						mockClient.EXPECT().
							Login(gomock.Any(), gomock.Any()).
							Return(nil, fmt.Errorf("some error"))
					},
				}
			},
			expectFn: func(t *testing.T, results TestResults) {
				if results.error == nil {
					t.Errorf("Expected error, got %v", results.error)
				}
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuthClient := mock_auth.NewMockAuthServiceClient(ctrl)

	for _, scenario := range testScenarios {
		t.Run(scenario.description, func(t *testing.T) {
			params := scenario.setupFn(t)
			params.mockSetup(mockAuthClient)

			authServiceClient, _ := New(mockAuthClient)
			ctx := context.Background()

			resp, err := authServiceClient.ClientLogin(ctx, "username")
			scenario.expectFn(t, TestResults{token: resp, error: err})
		})
	}
}
