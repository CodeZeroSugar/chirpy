package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to parse token: %w", err)
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
