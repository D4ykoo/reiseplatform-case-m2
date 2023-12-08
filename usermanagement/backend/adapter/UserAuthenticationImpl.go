package adapter

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func loadJWTEnv(key string) {
	// TODO: Load env file and get value
}

func CreateJWT(username string, secret string, test bool) (string, error) {
	var token *jwt.Token

	if test {
		token = jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
			"username": username,
			"iat":      time.Date(1991, 10, 5, 0, 0, 0, 0, time.UTC).Unix(),
		})
	} else {
		token = jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
			"username": username,
			"iat":      time.Now().Add(time.Hour * 1).Unix(),
		})
	}

	if secret == "" {
		secret = "your-fav-secret" // TODO: Env value
	}
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, err
}

// last return value is used for testing purposes
func ValidateJWT(tokenString string, secret string) (bool, error, jwt.MapClaims) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		if secret == "" {
			secret = "your-fav-secret" // TODO: Env value
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})

	if err != nil {
		return false, err, nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return true, nil, claims
	}

	return false, nil, nil
}
