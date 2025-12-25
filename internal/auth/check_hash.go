package auth

import (
	"fmt"

	"github.com/alexedwards/argon2id"
)

func CheckPasswordHash(password, hash string) (bool, error) {
	result, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return false, fmt.Errorf("password hash comparison failed: %w", err)
	}
	return result, nil
}
