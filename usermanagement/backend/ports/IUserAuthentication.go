package ports

type IUserAuthentication interface {
	createJWT()
	validateJWT()
}
