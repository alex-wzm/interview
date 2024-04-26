package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type (
	// Claims represents the custom claims for our JWT
	Claims struct {
		Username string `json:"username"`
		jwt.StandardClaims
	}

	Tokenizer interface {
		GenerateToken(username string, expiresIn time.Duration, secret []byte) (string, error)
		ValidateToken(tokenString string, secret []byte) (*Claims, error)
	}
)

var ErrInvalidToken = errors.New("invalid token")

// GenerateToken generates a JWT token with the provided username and expiration time
func GenerateToken(username string, expiresIn time.Duration, secret []byte) (string, error) {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims for the token
	claims := Claims{
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
func ValidateToken(tokenString string, secret []byte) (*Claims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secret, nil
	})

	// Check if the token is valid
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, ErrInvalidToken
		}
		return nil, err
	}

	// Extract the claims from the token
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
