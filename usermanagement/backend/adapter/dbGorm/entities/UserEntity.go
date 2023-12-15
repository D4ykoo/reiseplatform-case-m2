package entities

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model
	Username  string
	Firstname string
	Lastname  string
	Email     string
	Password  string
	Salt      string
}
