package authorization

import (
	"time"
	"github.com/dgrijalva/jwt-go"
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

