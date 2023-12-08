package application

import "github.com/gin-gonic/gin"
import "github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter"

func main() {
	router := gin.Default()

	router.GET("/users", adapter.ListUserRequest)
	router.GET("/users/:id", adapter.GetUserRequest)
	router.POST("/users", adapter.CreateUserRequest)
	router.PUT("/users/:id", adapter.UpdateUserRequest)
	router.DELETE("/users/:id", adapter.DeleteUserRequest)

	router.POST("/login", adapter.LoginRequest)
	router.POST("/reset", adapter.ResetPasswordRequest)

	// start server
	router.Run() // TODO: Address from env file
}
