package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	a := headers.Get("Authorization")
	if len(a) == 0 {
		return a, fmt.Errorf("'Authorization' does not exist in header")
	}
	tokenString := strings.TrimSpace(strings.ReplaceAll(a, "Bearer", ""))
	return tokenString, nil
}
