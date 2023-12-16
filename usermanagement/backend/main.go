package main

import (
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter/api"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
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

	router.GET("/api/users", api.ListUserRequest)
	router.GET("/api/users/:id", api.GetUserRequest)
	router.POST("/api/users", api.CreateUserRequest)
	router.PUT("/api/users/:id", api.UpdateUserRequest)
	router.DELETE("/api/users/:id", api.DeleteUserRequest)

	router.POST("/api/login", api.LoginRequest)
	router.POST("/api/register", api.RegisterRequest)

	router.PUT("/api/reset", api.ResetPasswordRequest)
	router.GET("/api/logout", api.LogoutRequest)

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
