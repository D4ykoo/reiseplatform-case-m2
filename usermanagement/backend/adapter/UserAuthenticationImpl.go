package adapter

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func loadJWTEnv(key string) {
	// TODO: Load env file and get value
}

func createJWT(username string, secret string, test bool) (string, error) {
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

func validateJWT(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("my_secret_key"), nil
	})

	if err != nil {
		return false, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok {
		return true, nil
	}

	return false, nil
}
