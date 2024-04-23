package auth

import (
	"context"
	"encoding/base64"
	"interview-authserver/config"
	"testing"

	"google.golang.org/grpc/metadata"
)

func Test_Authenticate_Success(t *testing.T) {
	mockenv := config.EnvConfig{
		JwtSecret:  "notasecret",
		Username:   "testuser",
		Usersecret: "testsecret",
	}

	raw := mockenv.Username + ":" + mockenv.Usersecret
	encodedstr := base64.StdEncoding.EncodeToString([]byte(raw))
	md := metadata.Pairs("Authorization", "Basic "+encodedstr)
	mockctx := metadata.NewIncomingContext(context.Background(), md)

	token, err := Authenticate(mockctx, mockenv.Username, &mockenv)
	if err != nil {
		t.Error("unexpected authentication error occurred")
	}

	if token == "" {
		t.Error("unexpected empty token string")
	}
}

func TestAuthenticate_MissingHeader(t *testing.T) {
	mockenv := config.EnvConfig{
		JwtSecret:  "notasecret",
		Username:   "testuser",
		Usersecret: "testsecret",
	}

	mockctx := context.Background()

	_, err := Authenticate(mockctx, mockenv.Username, &mockenv)
	if err == nil {
		t.Error("expected authentication error 'invalid headers' to occur")
	}

	if err.Error() != "invalid headers" {
		t.Error("expected invalid headers error")
	}
}

func TestAuthenticate_DecodeError(t *testing.T) {
	mockenv := config.EnvConfig{
		JwtSecret:  "notasecret",
		Username:   "testuser",
		Usersecret: "testsecret",
	}

	raw := mockenv.Username + ":" + mockenv.Usersecret
	md := metadata.Pairs("Authorization", "Basic "+raw)
	mockctx := metadata.NewIncomingContext(context.Background(), md)

	_, err := Authenticate(mockctx, mockenv.Username, &mockenv)
	if err == nil {
		t.Error("expected authentication error 'invalid headers' to occur")
	}

	if err.Error() != "invalid headers" {
		t.Error("expected invalid headers error")
	}
}

func TestAuthenticate_InvalidUserName(t *testing.T) {
	mockenv := config.EnvConfig{
		JwtSecret:  "notasecret",
		Username:   "testuser",
		Usersecret: "testsecret",
	}

	raw := "invalid_user" + ":" + mockenv.Usersecret
	encodedstr := base64.StdEncoding.EncodeToString([]byte(raw))
	md := metadata.Pairs("Authorization", "Basic "+encodedstr)
	mockctx := metadata.NewIncomingContext(context.Background(), md)

	_, err := Authenticate(mockctx, "invalid_user", &mockenv)
	if err == nil {
		t.Error("expected authentication error 'invalid user' to occur")
	}

	if err.Error() != "invalid user" {
		t.Error("expected invalid user error")
	}
}

func TestAuthenticate_InvalidUserSecret(t *testing.T) {
	mockenv := config.EnvConfig{
		JwtSecret:  "notasecret",
		Username:   "testuser",
		Usersecret: "testsecret",
	}

	raw := mockenv.Username + ":" + "invalid_secret"
	encodedstr := base64.StdEncoding.EncodeToString([]byte(raw))
	md := metadata.Pairs("Authorization", "Basic "+encodedstr)
	mockctx := metadata.NewIncomingContext(context.Background(), md)

	_, err := Authenticate(mockctx, mockenv.Username, &mockenv)
	if err == nil {
		t.Error("expected authentication error 'invalid user' to occur")
	}

	if err.Error() != "invalid user" {
		t.Error("expected invalid user error")
	}
}
