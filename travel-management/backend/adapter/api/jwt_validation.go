package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWT(tokenString string, secret string) (bool, error, jwt.MapClaims) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		if secret == "" {
			secret = os.Getenv("JWT_SECRET")
		}
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

func ValidateLoginStatus(c *gin.Context) (jwt.MapClaims, error) {

	cookie, cookieErr := c.Cookie("authTravel")

	fmt.Println(cookie)
	if cookieErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": cookieErr.Error()})
		return nil, cookieErr
	}

	_, valErr, claim := ValidateJWT(cookie, os.Getenv("JWT_SECRET"))

	if valErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": valErr.Error()})
		return nil, valErr
	}

	return claim, nil
}
