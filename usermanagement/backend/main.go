package main

import (
	"log"
	"os"
	"strconv"

	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter/api"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunWebServer() {
	utils.LoadFile()
	isDebug, errBool := strconv.ParseBool(os.Getenv("DEBUG"))

	if errBool != nil {
		log.Fatal(errBool, "Try to change the DEBUG field in the .env file")
	}

	if !isDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("DOMAIN")}
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Authorization", "Origin", "Content-Type", "Accept"}
	router.Use(cors.New(config))

	v1 := router.Group("/v1")
	{
		v1.GET("/api/users", api.ListUserRequest)
		v1.GET("/api/users/:id", api.GetUserRequest)
		v1.POST("/api/users", api.CreateUserRequest)
		v1.PUT("/api/users/:id", api.UpdateUserRequest)
		v1.DELETE("/api/users/:id", api.DeleteUserRequest)

		v1.POST("/api/login", api.LoginRequest)
		v1.POST("/api/register", api.RegisterRequest)

		v1.PUT("/api/reset", api.ResetPasswordRequest)
		v1.GET("/api/logout", api.LogoutRequest)
	}
	// start server
	err := router.Run(os.Getenv("API_URL"))
	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	RunWebServer()
}
