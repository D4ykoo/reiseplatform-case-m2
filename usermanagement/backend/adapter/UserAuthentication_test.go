package adapter

import (
	"regexp"
	"testing"
)

var jwtSecret = "your-fav-secret"

// TestJWTBuilding calls greetings.
// Checking for a valid return value.
func TestJWTBuilding(t *testing.T) {
	username := "test"
	got, err := createJWT(username, jwtSecret, true)

	// generated with: https://jwt.io
	want := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjY4NjYyMDgwMCwidXNlcm5hbWUiOiJ0ZXN0In0.CbmybuROnf_3ClsxXiYiTbK26Dc0e2zSwMeCZZz4guszI-q8LL6HO42HJTeAjQ0gDFRmL4PikQoP8QzdPC03yw"
	valid, exErr := regexp.MatchString(want, got)
	if !valid || exErr != nil {
		t.Fatalf(`createJWT(%q) = %q, %v, want match for %#q, nil`, username, got, err, want)
	}
}

func TestJWTValidating(t *testing.T) {

}
