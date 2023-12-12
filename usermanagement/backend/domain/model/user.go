package domain

import (
	"gorm.io/gorm"
)

// User Model
//
// May required to add OIDC provider
type User struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type DBUser struct {
	gorm.Model
	User
	Salt string
}

type ResponseUser struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResetUser struct {
	OldLoginUser LoginUser
	NewPassword  string `json:"newPassword"`
}
