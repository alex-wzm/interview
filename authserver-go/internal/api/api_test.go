package api

import (
	"context"
	"fmt"
	"interview-authserver/config"
	"interview-authserver/internal/api/auth"
	"testing"
)

type mockAuthenticator struct{}

func (m mockAuthenticator) Authenticate(ctx context.Context, username string, env *config.EnvConfig) (string, error) {
	if env.Username == username {
		return "sometoken", nil
	}
	return "", fmt.Errorf("user not found")
}

func Test_Authenticate_Success(t *testing.T) {
	mockenv := config.EnvConfig{
		JwtSecret:  "fakesecret",
		Username:   "fakeuser",
		Usersecret: "fakeusersecret",
	}

	req := &auth.AuthRequest{
		Username: "fakeuser",
	}

	mockAuthenticator := mockAuthenticator{}
	s := New(&mockenv, mockAuthenticator)

	res, err := s.Authenticate(context.Background(), req)
	if err != nil {
		t.Error("unexpected error")
	}

	if res.Token != "sometoken" {
		t.Error("expected token from mockauthenticator")
	}
}

func Test_Authenticate_Fail(t *testing.T) {
	mockenv := config.EnvConfig{
		JwtSecret:  "fakesecret",
		Username:   "fakeuser",
		Usersecret: "fakeusersecret",
	}

	req := &auth.AuthRequest{
		Username: "invaliduser",
	}

	mockAuthenticator := mockAuthenticator{}
	s := New(&mockenv, mockAuthenticator)

	res, err := s.Authenticate(context.Background(), req)
	if err == nil {
		t.Error("expected an error")
	}

	if res != nil {
		t.Error("unexpected token")
	}
}
