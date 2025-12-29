package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	a := headers.Get("Authorization")
	if len(a) == 0 {
		return a, fmt.Errorf("'Authorization' does not exist in header")
	}
	tokenString := strings.TrimSpace(strings.ReplaceAll(a, "ApiKey", ""))
	return tokenString, nil
}
