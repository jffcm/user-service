package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenGenerator interface {
	Generate(email string) (string, error)
}

type jwtTokenGenerator struct {
	secretKey string
}

func NewJWTTokenGenerator(secretKey string) TokenGenerator {
	return &jwtTokenGenerator{secretKey: secretKey}
}

func (j *jwtTokenGenerator) Generate(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(1 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}
