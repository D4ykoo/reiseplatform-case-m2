package inbound

type IUserAuthentication interface {
	createJWT(username string, secret string, test bool) (string, error)
	validateJWT(tokenString string, secret string) (bool, error, any)
}
