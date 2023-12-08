package domain

import "github.com/gin-gonic/gin"

type IUserService interface {
	// CRUD REST operations
	CreateUserRequest(c *gin.Context)
	UpdateUserRequest(c *gin.Context)
	DeleteUserRequest(c *gin.Context)
	GetUserRequest(c *gin.Context)

	// other HTTP requests
	LoginRequest(c *gin.Context)
	ResetPasswordRequest(c *gin.Context)
}
