package domain

import "github.com/gin-gonic/gin"

type IUserService interface {
	// CRUD REST operations
	createUserRequest(c *gin.Context)
	updateUserRequest(c *gin.Context)
	deleteUserRequest(c *gin.Context)
	getUserRequest(c *gin.Context)

	// other HTTP requests
	loginRequest(c *gin.Context)
	resetPasswordRequest(c *gin.Context)
}
