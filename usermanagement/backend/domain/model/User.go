package domain

// User Model
//
// May required to add OIDC provider
type User struct {
	Id        uint
	Username  string
	Firstname string
	Lastname  string
	Email     string
	Password  string
}
