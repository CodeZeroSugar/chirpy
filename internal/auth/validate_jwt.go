package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to parse token: %w", err)
	}

	stringIssuer, err := token.Claims.GetIssuer()
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to get issuer from token to validate: %w", err)
	}
	if stringIssuer != "chirpy" {
		return uuid.UUID{}, fmt.Errorf("provided issuer '%v' did not match 'chirpy': %w", stringIssuer, err)
	}

	stringID, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to get user id from token: %w", err)
	}
	userID, err := uuid.Parse(stringID)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to parse userID to uuid: %w", err)
	}
	return userID, nil
}
