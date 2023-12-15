package domain

import (
	model "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
	"github.com/gin-gonic/gin"
)

type (
	IUserService interface {
		CreateUser(user model.User) error
		ChangeUser(id uint, user model.User, oldPassword string) error

		DeleteUserRequest(c *gin.Context)
		GetUserRequest(c *gin.Context)
		ListUserRequest(c *gin.Context)

		// other HTTP requests
		RegisterRequest(c *gin.Context)
		LoginRequest(c *gin.Context)
		LogOutRequest(c *gin.Context)
		ResetPasswordRequest(c *gin.Context)
	}
)
