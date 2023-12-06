package domain

import "github.com/gin-gonic/gin"

type IUserService interface {
	// CRUD REST operations
	createUser(c *gin.Context)
	updateUser(c *gin.Context)
	deleteUser(c *gin.Context)
	getUser(c *gin.Context)

	// other HTTP requests
	loginRequest(c *gin.Context)
	resetPasswordRequest(c *gin.Context)
}
