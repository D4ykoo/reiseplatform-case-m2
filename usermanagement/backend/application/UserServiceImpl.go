package application

import (
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)
import "github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter"

func RunWebServer() {
	utils.LoadFile()

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Authorization", "Origin", "Content-Type", "Accept"}
	router.Use(cors.New(config))

	router.GET("/api/users", adapter.ListUserRequest)
	router.GET("/api/users/:id", adapter.GetUserRequest)
	router.POST("/api/users", adapter.CreateUserRequest)
	router.PUT("/api/users/:id", adapter.UpdateUserRequest)
	router.DELETE("/api/users/:id", adapter.DeleteUserRequest)

	router.POST("/api/login", adapter.LoginRequest)
	router.POST("/api/register", adapter.RegisterRequest)

	router.PUT("/api/reset", adapter.ResetPasswordRequest)

	// start server
	err := router.Run(os.Getenv("DOMAIN"))
	if err != nil {
		log.Fatal(err)
		return
	}
}
