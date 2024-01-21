package test

import (
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter"
	"regexp"
	"testing"
)

var jwtSecret = "your-jwt-secret"

// TestJWTBuilding calls greetings.
// Checking for a valid return value.
func TestJWTBuilding(t *testing.T) {
	auth := adapter.InitAuth()
	username := "test"
	var userId uint = 1

	got, err := auth.CreateJWT(username, &userId, jwtSecret, true)

	// generated with: https://jwt.io
	want := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjY4NjYyMDgwMCwidXNlcl9pZCI6MSwidXNlcm5hbWUiOiJ0ZXN0In0.j1McK-wmsGqNWQ9ir_DqT90uFFvIt7zjYbABe1Lf0iJ8wbZeaqedbTr5v2UB43ZLCySEwTx3QzROTXltzTIXoA"
	valid, exErr := regexp.MatchString(want, got)
	if !valid || exErr != nil {
		t.Fatalf(`createJWT(%q) = %q, %v, want match for %#q, nil`, username, got, err, want)
	}
}

func TestJWTValidating(t *testing.T) {
	auth := adapter.InitAuth()

	token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjY4NjYyMDgwMCwidXNlcl9pZCI6MSwidXNlcm5hbWUiOiJ0ZXN0In0.j1McK-wmsGqNWQ9ir_DqT90uFFvIt7zjYbABe1Lf0iJ8wbZeaqedbTr5v2UB43ZLCySEwTx3QzROTXltzTIXoA"
	valid, err, claims := auth.ValidateJWT(token, jwtSecret)

	if !valid || err != nil {
		t.Fatalf(`validateJWT(%q, %q) = %t, %q, %q, is not true, nil, claims`, token, jwtSecret, valid, err, claims)
	}

	username := "test"
	var userId uint = 1
	iat := 686620800

	usernameClaim := string(claims["username"].(string))
	userIdClaim := uint(claims["user_id"].(float64))
	iatClaim := int(claims["iat"].(float64))

	uRes, uErr := regexp.MatchString(usernameClaim, username)

	if !uRes || uErr != nil {
		t.Fatalf(`username claim %q does not equal the test value: %q`, usernameClaim, username)
	}

	if userIdClaim != 1 {
		t.Fatalf(`iat claim %d does not equal the test value: %d`, userIdClaim, userId)
	}

	if iatClaim != 686620800 {
		t.Fatalf(`iat claim %d does not equal the test value: %d`, iatClaim, iat)
	}
}
