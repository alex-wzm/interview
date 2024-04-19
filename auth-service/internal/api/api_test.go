package api

import (
	"context"
	"testing"
	"time"

	auth "auth-service/internal/api/authservice"
	jwtValidator "auth-service/internal/domain/jwt"
)

func TestLogin(t *testing.T) {
	server := NewAuthServiceServer("secret-key")

	type TestResults struct {
		response *auth.LoginResponse
		err      error
	}

	type TestParameters struct {
		ctx context.Context
		req *auth.LoginRequest
	}

	type TestScenario struct {
		description string
		setupFn     func(t *testing.T) TestParameters
		expectFn    func(t *testing.T, results TestResults)
	}

	testScenarios := []TestScenario{
		{
			description: "Happy Path",
			setupFn: func(t *testing.T) TestParameters {
				ctx := context.Background()
				req := &auth.LoginRequest{Username: "admin"}
				return TestParameters{ctx: ctx, req: req}
			},
			expectFn: func(t *testing.T, results TestResults) {
				if results.err != nil {
					t.Errorf("Expected no error, got %v", results.err)
				}
				if results.response == nil || results.response.Token == "" {
					t.Errorf("Expected a token, got none")
				}
			},
		},
		{
			description: "login with invalid username",
			setupFn: func(t *testing.T) TestParameters {
				ctx := context.Background()
				req := &auth.LoginRequest{Username: "not-admin"}
				return TestParameters{ctx: ctx, req: req}
			},
			expectFn: func(t *testing.T, results TestResults) {
				if results.err == nil {
					t.Error("Expected an error, got none")
				}
				if results.response != nil {
					t.Error("Expected no response, got a token")
				}
			},
		},
	}

	for _, scenario := range testScenarios {
		t.Run(scenario.description, func(t *testing.T) {
			params := scenario.setupFn(t)
			response, err := server.Login(params.ctx, params.req)
			scenario.expectFn(t, TestResults{
				response: response,
				err:      err,
			})
		})
	}
}

func TestValidateJWT(t *testing.T) {
	server := NewAuthServiceServer("secret-key")

	type TestResults struct {
		response *auth.ValidateResponse
		err      error
	}

	type TestParameters struct {
		ctx context.Context
		req *auth.ValidateRequest
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
				ctx := context.Background()
				token, _ := jwtValidator.GenerateToken("admin", time.Minute*30, []byte("secret-key"))
				req := &auth.ValidateRequest{Token: token}
				return TestParameters{ctx: ctx, req: req}
			},
			expectFn: func(t *testing.T, results TestResults) {
				if results.err != nil {
					t.Errorf("Expected no error, got %v", results.err)
				}
				if results.response == nil || !results.response.Valid {
					t.Error("Expected token to be valid")
				}
			},
		},
		{
			description: "validate invalid JWT",
			setupFn: func(t *testing.T) TestParameters {
				ctx := context.Background()
				req := &auth.ValidateRequest{Token: "invalid-token"}
				return TestParameters{ctx: ctx, req: req}
			},
			expectFn: func(t *testing.T, results TestResults) {
				if results.err == nil {
					t.Error("Expected an error, got none")
				}
				if results.response != nil && results.response.Valid {
					t.Error("Expected token to be invalid")
				}
			},
		},
	}

	for _, scenario := range testScenarios {
		t.Run(scenario.description, func(t *testing.T) {
			params := scenario.setupFn(t)
			response, err := server.ValidateJWT(params.ctx, params.req)
			scenario.expectFn(t, TestResults{
				response: response,
				err:      err,
			})
		})
	}
}
