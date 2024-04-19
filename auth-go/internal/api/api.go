package api

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	"context"
	"interview-auth/internal/api/interview/auth"
)

// JWTClaims represents the custom claims for our JWT
type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string, expiresIn time.Duration, secret []byte) (string, error) {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims for the token
	claims := JWTClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiresIn).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token.Claims = claims

	// Sign the token with the provided secret
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

type server struct {
	auth.UnimplementedAuthServiceServer
}

func New() *server {
	return &server{}
}

func (s *server) Authorize(ctx context.Context, r *auth.AuthRequest) (*auth.JWTResponse, error) {
	// TODO: authenticate username/password
	token, err  := GenerateToken(r.Username, time.Second * time.Duration(r.Ttl), []byte("secret"))

	return &auth.JWTResponse{
		Token: token,
		
	}, err
}