package jwtValidator

import (
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	type testResult struct {
		token string
		err   error
	}
	tests := []struct {
		description string
		username    string
		expiresIn   time.Duration
		secret      []byte
		expectFn    func(*testing.T, testResult)
	}{
		{
			description: "Generate token successfully",
			username:    "testuser",
			expiresIn:   time.Minute * 5,
			secret:      []byte("mysecret"),
			expectFn: func(t *testing.T, result testResult) {
				if result.err != nil {
					t.Errorf("Expected no error, got %v", result.err)
				}
				if result.token == "" {
					t.Error("Expected a token, got empty string")
				}
			},
		},
		{
			description: "Generate token with empty username",
			username:    "",
			expiresIn:   time.Minute * 5,
			secret:      []byte("mysecret"),
			expectFn: func(t *testing.T, result testResult) {
				if result.err != nil {
					t.Errorf("Expected no error, got %v", result.err)
				}
				if result.token == "" {
					t.Error("Expected a token, got empty string")
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			token, err := GenerateToken(test.username, test.expiresIn, test.secret)
			test.expectFn(t, testResult{token: token, err: err})
		})
	}
}

func TestValidateToken(t *testing.T) {
	type testResult struct {
		claims *JWTClaims
		err    error
	}
	tests := []struct {
		description string
		token       string
		secret      []byte
		expectFn    func(*testing.T, testResult)
	}{
		{
			description: "Validate valid token",
			token: func() string {
				token, _ := GenerateToken("testuser", time.Minute*5, []byte("mysecret"))
				return token
			}(),
			secret: []byte("mysecret"),
			expectFn: func(t *testing.T, result testResult) {
				if result.err != nil {
					t.Errorf("Expected no error, got %v", result.err)
				}
				if result.claims.Username != "testuser" {
					t.Errorf("Expected username 'testuser', got '%v'", result.claims.Username)
				}
			},
		},
		{
			description: "Validate token with wrong secret",
			token: func() string {
				token, _ := GenerateToken("testuser", time.Minute*5, []byte("mysecret"))
				return token
			}(),
			secret: []byte("wrongsecret"),
			expectFn: func(t *testing.T, result testResult) {
				if result.err == nil {
					t.Error("Expected error, got none")
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			claims, err := ValidateToken(test.token, test.secret)
			test.expectFn(t, testResult{claims: claims, err: err})
		})
	}
}
