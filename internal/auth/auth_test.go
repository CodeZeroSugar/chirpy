package auth

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeJWT(t *testing.T) {
	type MakeJWTTest struct {
		userID      uuid.UUID
		tokenSecret string
		expiresIn   time.Duration
		want        bool
	}

	tests := []MakeJWTTest{
		{userID: uuid.New(), tokenSecret: "testsecret", expiresIn: 24 * time.Hour, want: true},
	}

	for _, tc := range tests {
		result, err := MakeJWT(tc.userID, tc.tokenSecret, tc.expiresIn)
		if err != nil {
			t.Fatalf("error was not nil: %v", err)
		}
		got := false
		if result != "" {
			got = true
		}
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected %v, got: %v", tc.want, got)
		}
	}
}
