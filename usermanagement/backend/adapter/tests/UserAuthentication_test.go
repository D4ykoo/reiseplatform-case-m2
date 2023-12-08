package tests

import (
	adapter "github.com/D4ykoo/usermanagement/backend"
	"regexp"
	"testing"
)

var jwtSecret = "your-fav-secret"

// TestJWTBuilding calls greetings.
// Checking for a valid return value.
func TestJWTBuilding(t *testing.T) {
	username := "test"
	got, err := adapter.CreateJWT(username, jwtSecret, true)

	// generated with: https://jwt.io
	want := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjY4NjYyMDgwMCwidXNlcm5hbWUiOiJ0ZXN0In0.CbmybuROnf_3ClsxXiYiTbK26Dc0e2zSwMeCZZz4guszI-q8LL6HO42HJTeAjQ0gDFRmL4PikQoP8QzdPC03yw"
	valid, exErr := regexp.MatchString(want, got)
	if !valid || exErr != nil {
		t.Fatalf(`createJWT(%q) = %q, %v, want match for %#q, nil`, username, got, err, want)
	}
}

func TestJWTValidating(t *testing.T) {
	token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjY4NjYyMDgwMCwidXNlcm5hbWUiOiJ0ZXN0In0.CbmybuROnf_3ClsxXiYiTbK26Dc0e2zSwMeCZZz4guszI-q8LL6HO42HJTeAjQ0gDFRmL4PikQoP8QzdPC03yw"
	valid, err, claims := adapter.ValidateJWT(token, jwtSecret)

	if !valid || err != nil {
		t.Fatalf(`validateJWT(%q, %q) = %t, %q, %q, is not true, nil, claims`, token, jwtSecret, valid, err, claims)
	}

	username := "test"
	iat := 686620800

	usernameClaim := string(claims["username"].(string))
	iatClaim := int(claims["iat"].(float64))

	uRes, uErr := regexp.MatchString(usernameClaim, username)

	if !uRes || uErr != nil {
		t.Fatalf(`username claim %q does not equal the test value: %q`, usernameClaim, username)
	}

	if iatClaim != 686620800 {
		t.Fatalf(`iat claim %d does not equal the test value: %d`, iatClaim, iat)
	}
}
