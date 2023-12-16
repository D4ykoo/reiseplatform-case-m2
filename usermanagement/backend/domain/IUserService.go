package domain

import (
	model "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
)

type (
	IUserService interface {
		CreateUser(user model.User) error
		ChangeUser(id uint, user model.User, oldPassword string) error
		FindUser(identifier interface{}) (*model.User, error)
		ListAllUser() (*[]model.User, error)

		//DeleteUserRequest(c *gin.Context)
		//GetUserRequest(c *gin.Context)
		//ListUserRequest(c *gin.Context)

		// other HTTP requests
		RegisterUser(user model.User)
		LoginUser(username string, password string) error
		ResetPassword(username string, newPassword string) error

		// TODO what about logout?
	}
)
