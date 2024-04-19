package validator

import (
	"context"
	"errors"

	"interview-service/internal/api/authservice"
	auth "interview-service/internal/api/authservice"
	mock_auth "interview-service/internal/api/mock_auth"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"google.golang.org/grpc/metadata"
)

func TestValidateJWTToken(t *testing.T) {
	type TestResults struct {
		ctx   context.Context
		error error
	}

	type TestParameters struct {
		token     string
		mockSetup func(*mock_auth.MockAuthServiceClient)
	}

	type TestScenario struct {
		description string
		setupFn     func(t *testing.T) TestParameters
		expectFn    func(t *testing.T, results TestResults)
	}

	testScenarios := []TestScenario{
		{
			description: "Valid Token",
			setupFn: func(t *testing.T) TestParameters {
				token := "valid-token"
				return TestParameters{
					token: token,
					mockSetup: func(mockClient *mock_auth.MockAuthServiceClient) {
						mockClient.EXPECT().
							ValidateJWT(gomock.Any(), &authservice.ValidateRequest{Token: token}).
							Return(&authservice.ValidateResponse{Valid: true}, nil)
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
			description: "Invalid Token",
			setupFn: func(t *testing.T) TestParameters {
				token := "invalid-token"
				return TestParameters{
					token: token,
					mockSetup: func(mockClient *mock_auth.MockAuthServiceClient) {
						mockClient.EXPECT().
							ValidateJWT(gomock.Any(), &auth.ValidateRequest{Token: token}).
							Return(nil, errors.New("invalid token"))
					},
				}
			},
			expectFn: func(t *testing.T, results TestResults) {
				if results.error == nil || results.error.Error() != "invalid token" {
					t.Errorf("Expected 'invalid token' error, got %v", results.error)
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

			authServiceClient, _ := NewAuthServiceClient(mockAuthClient)
			ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{"authorization": []string{"Bearer " + params.token}})

			newCtx, err := authServiceClient.ValidateJWTToken()(ctx)
			scenario.expectFn(t, TestResults{ctx: newCtx, error: err})
		})
	}
}
