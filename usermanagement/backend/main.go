package main

import (
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter/dbGorm"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter/kafka"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/application"
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
	//isDemo, errBoolDemo := strconv.ParseBool(os.Getenv("DEMO"))

	if errBool != nil {
		log.Fatal(errBool, "Try to change the DEBUG field in the .env file")
	}

	if !isDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	//if errBoolDemo != nil {
	//	log.Fatal(errBool, "Try to change the DEMO field in the .env file")
	//}
	// add user -> find a place were to include this, since it is a hexagonal architecture
	//user := domain.User{
	//	Username:  "demoUser",
	//	Firstname: "demo",
	//	Lastname:  "user",
	//	Email:     "user@demo.demo",
	//	Password:  "demo",
	//}
	//user.Password = HashPassword(user.Password, []byte(os.Getenv("SALT")))
	//
	//_ = dbGorm.Save(user)

	// outgoing
	userRepo := dbGorm.Init()
	messageService := kafka.Init()
	auth := adapter.InitAuth()
	email := adapter.InitEmail(messageService)
	userService := application.InitUserService(userRepo)
	httpUserService := application.InitHttpService(userRepo)

	userServiceHttpController := api.InitHttpUserController(httpUserService, messageService, email, auth)
	userServiceController := api.Init(userService, messageService, auth)

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("DOMAIN"), "http://usermanagement-frontend:8081"}
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Authorization", "Origin", "Content-Type", "Accept"}

	router.Use(cors.New(config))

	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", userServiceController.ListUserRequest)
		v1.GET("/users/:id", userServiceController.GetUserRequest)
		v1.POST("/users", userServiceController.CreateUserRequest)
		v1.PUT("/users/:id", userServiceController.UpdateUserRequest)
		v1.DELETE("/users/:id", userServiceController.DeleteUserRequest)

		v1.POST("/login", userServiceHttpController.LoginRequest)
		v1.POST("/register", userServiceHttpController.RegisterRequest)

		v1.PUT("/reset", userServiceHttpController.ResetPasswordRequest)
		v1.GET("/logout", userServiceHttpController.LogoutRequest)
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
