package auth

import (
	"fmt"
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
		{userID: uuid.New(), tokenSecret: "12345", expiresIn: 24 * time.Hour, want: true},
		{userID: uuid.New(), tokenSecret: "helloworld", expiresIn: 24 * time.Hour, want: true},
		{userID: uuid.New(), tokenSecret: "$up3r$3cr3t!!", expiresIn: 24 * time.Hour, want: true},
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
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestValidateJWTUser(t *testing.T) {
	type ValidateJWTTest struct {
		tokenString, tokenSecret string
		want                     uuid.UUID
	}

	testUser := uuid.New()
	testSecret := "supersecretpassword"
	testExpires := 24 * time.Hour
	testToken, err := MakeJWT(testUser, testSecret, testExpires)
	if err != nil {
		t.Fatalf("error creating test token: %v", err)
	}

	tests := []ValidateJWTTest{
		{tokenString: testToken, tokenSecret: testSecret, want: testUser},
	}
	for _, tc := range tests {
		got, err := ValidateJWT(tc.tokenString, tc.tokenSecret)
		if err != nil {
			t.Fatalf("error validating jwt: %v", err)
		}
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestValidateJWTExpired(t *testing.T) {
	type ValidateJWTTest struct {
		tokenString, tokenSecret string
		want                     uuid.UUID
	}
	testUser := uuid.New()
	testSecret := "supersecretpassword"
	testExpires := -24 * time.Hour
	testToken, err := MakeJWT(testUser, testSecret, testExpires)
	if err != nil {
		t.Fatalf("error creating test token: %v", err)
	}

	tests := []ValidateJWTTest{
		{tokenString: testToken, tokenSecret: testSecret, want: uuid.UUID{}},
	}
	for _, tc := range tests {
		got, err := ValidateJWT(tc.tokenString, tc.tokenSecret)
		if err == nil {
			t.Fatalf("Expected expired token error, got: %v", err)
		}
		fmt.Printf("%v\n", err)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestValidateJWTWrongIssuer(t *testing.T) {
	type ValidateJWTTest struct {
		tokenString, tokenSecret string
		want                     uuid.UUID
	}
	testUser := uuid.New()
	testSecret := "supersecretpassword"
	testExpires := 24 * time.Hour
	testToken, err := MakeJWTBadIssuer(testUser, testSecret, testExpires)
	if err != nil {
		t.Fatalf("error creating test token: %v", err)
	}

	tests := []ValidateJWTTest{
		{tokenString: testToken, tokenSecret: testSecret, want: uuid.UUID{}},
	}
	for _, tc := range tests {
		got, err := ValidateJWT(tc.tokenString, tc.tokenSecret)
		if err == nil {
			t.Fatalf("Expected incorrect issuer error, got: %v", err)
		}
		fmt.Printf("%v\n", err)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
