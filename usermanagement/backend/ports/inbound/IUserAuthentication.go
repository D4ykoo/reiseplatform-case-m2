package inbound

import "github.com/golang-jwt/jwt/v5"

type IUserAuthentication interface {
	CreateJWT(username string, userId *uint, secret string, test bool) (string, error)
	ValidateJWT(tokenString string, secret string) (bool, error, jwt.MapClaims)
}
