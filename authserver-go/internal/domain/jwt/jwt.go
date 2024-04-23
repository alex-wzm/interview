package jwtValidator

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

// JWTClaims represents the custom claims for our JWT
type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type JWTlib interface {
	GenerateToken(username string, expiresIn time.Duration, secret []byte) (string, error)
	ValidateToken(tokenString string, secret []byte) (*JWTClaims, error)
}

var ErrInvalidToken = errors.New("invalid token")

// GenerateToken generates a JWT token with the provided username and expiration time
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

// ValidateToken validates a JWT token with the provided secret and returns the claims
func ValidateToken(tokenString string, secret []byte) (*JWTClaims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secret, nil
	})

	// Check if the token is valid
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, ErrInvalidToken
		}
		return nil, err
	}

	// Extract the claims from the token
	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
