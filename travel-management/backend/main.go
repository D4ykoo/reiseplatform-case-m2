package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mig3177/travelmanagement/adapter/api"
	dbgorm "github.com/mig3177/travelmanagement/adapter/dbGoRm"
	"github.com/mig3177/travelmanagement/adapter/kafka"
	"github.com/mig3177/travelmanagement/application"
	"github.com/mig3177/travelmanagement/utils"
)

func main() {

	// Env
	utils.LoadFile()
	isDebug, errBool := strconv.ParseBool(os.Getenv("DEBUG"))

	if errBool != nil {
		log.Fatal(errBool, "Try to change the DEBUG field in the .env file")
	}

	if !isDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	// Outgoing
	// Repository
	hotelRepo := dbgorm.NewHotelRepository(10, 100)
	travelRepo := dbgorm.NewTravelRepository(10, 100)
	tagRepo := dbgorm.NewTagRepository(10, 100)
	// Message broker
	eventService := kafka.New()
	// Application
	// Service
	travelService := application.New(hotelRepo, travelRepo, tagRepo, eventService)

	// Incomming
	// Controller
	service := api.New(travelService)
	// Router
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://travelmngt-web:8085"}
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Authorization", "Origin", "Content-Type", "Accept"}

	// TODO CORS
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{""})
	//router.Use(cors.New(config))

	router.GET("/api/v1/loginstatus", service.CheckLoginStatus)

	// CRUD
	// Create
	router.POST("/api/v1/hotels", service.CreateHotelRequest)

	// Read
	router.GET("/api/v1/hotels", func(ctx *gin.Context) {
		fmt.Println(ctx.Query("name"))
		service.FindHotels(ctx)
	})

	router.GET("/api/v1/hotels/:id", service.GetHotelById)

	// Update
	router.PUT("/api/v1/hotels/:id", service.UpdateHotel)
	// Delete
	router.DELETE("/api/v1/hotels/:id", service.DeleteHotelRequest)

	// Offers
	router.POST("/api/v1/hotels/:id/travels", service.CreateTravelRequest)
	// Read
	/*router.GET("/api/v1/hotels/:id/travels", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})*/

	router.GET("/api/v1/hotels/:id/travels/:tid", service.GetTravelById)

	// Update
	router.PUT("/api/v1/hotels/:id/travels/:tid", service.UpdateTravel)
	// Delete
	router.DELETE("/api/v1/hotels/:id/travels/:tid", service.DeleteTravel)

	// Tags
	router.POST("/api/v1/tags", service.CreateTagRequest)
	// Read
	router.GET("/api/v1/tags", service.ListAllTags)

	router.GET("/api/v1/tags/:id", service.GetTagById)

	// Update
	router.PUT("/api/v1/tags/:id", service.UpdateTag)
	// Delete
	router.DELETE("/api/v1/tags/:id", service.DeleteTagRequest)

	err := router.Run(os.Getenv("API_URL"))
	if err != nil {
		log.Fatal(err)
		return
	}
}
