package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func MakeRefreshToken() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", fmt.Errorf("failed to create random key for reshresh token: %w", err)
	}
	hexString := hex.EncodeToString(key)
	return hexString, nil
}
