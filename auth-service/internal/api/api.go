package api

import (
	"context"
	"fmt"
	"log"
	"time"

	auth "auth-service/internal/api/authservice"
	jwtValidator "auth-service/internal/domain/jwt"
)

type AuthServiceServer struct {
	jwtKey []byte
	auth.UnimplementedAuthServiceServer
}

func NewAuthServiceServer(secret string) *AuthServiceServer {
	return &AuthServiceServer{jwtKey: []byte(secret)}
}

// Login function - called by interview-client
func (s *AuthServiceServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	//For now the logic is set to check for the admin user
	//Possible extension - seed some users and give access based on their names/roles.
	if req.Username == "admin" {
		token, err := jwtValidator.GenerateToken(req.Username, time.Minute*30, s.jwtKey)
		if err != nil {
			return nil, err
		}
		return &auth.LoginResponse{Token: token}, nil
	}
	return nil, fmt.Errorf("invalid user")
}

// To validate the jwt token - called by interview-service
func (s *AuthServiceServer) ValidateJWT(ctx context.Context, req *auth.ValidateRequest) (*auth.ValidateResponse, error) {
	//Get token from request and validate
	token := req.Token
	claims, err := jwtValidator.ValidateToken(token, s.jwtKey)
	if err != nil {
		log.Default().Println(err)
		return &auth.ValidateResponse{Valid: false, Username: ""}, err
	}

	return &auth.ValidateResponse{Valid: true, Username: claims.Username}, nil
}
