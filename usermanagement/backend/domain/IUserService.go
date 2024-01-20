package domain

import (
	model "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
)

type (
	IUserService interface {
		CreateUser(user model.User) error
		ChangeUser(id uint, user model.User) error
		FindUser(identifier interface{}) (*model.User, error)
		ListAllUser() (*[]model.User, error)
		DeleteUser(id int) error
	}
)
