package main

import (
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
	ginModeSetup()

	router := gin.Default()
	corsSetup(router)
	routingSetup(router)

	// Run
	err := router.Run(os.Getenv("API_URL"))
	if err != nil {
		log.Fatal(err)
		return
	}
}

func ginModeSetup() {
	isDebug, errBool := strconv.ParseBool(os.Getenv("DEBUG"))

	if errBool != nil {
		log.Fatal(errBool, "Try to change the DEBUG field in the .env file")
	}

	if !isDebug {
		gin.SetMode(gin.ReleaseMode)
	}
}

func corsSetup(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8085", "http://localhost:12345"}
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Authorization", "Origin", "Content-Type", "Accept"}

	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{""})
	router.Use(cors.New(config))
}

func routingSetup(router *gin.Engine) {

	// Outgoing
	// Repository
	hotelRepo := dbgorm.NewHotelRepository(10, 100)
	travelRepo := dbgorm.NewTravelRepository(10, 100)
	tagRepo := dbgorm.NewTagRepository(10, 100)
	// Message broker
	eventService := kafka.New()
	// Application
	// Service
	travelService := application.NewTravelService(travelRepo, eventService)
	tagService := application.NewTagService(tagRepo)
	hotelService := application.NewHotelService(hotelRepo, eventService)

	// Incomming
	// Controller
	controller := api.New(travelService, tagService, hotelService)
	api := router.Group("/api/v1")

	{
		api.GET("/loginstatus", controller.CheckLoginStatus)

		// CRUD
		// Create
		api.POST("/hotels", controller.CreateHotelRequest)

		// Read
		api.GET("/hotels", controller.FindHotels)

		api.GET("/hotels/:id", controller.GetHotelById)

		// Update
		api.PUT("/hotels/:id", controller.UpdateHotel)
		// Delete
		api.DELETE("/hotels/:id", controller.DeleteHotelRequest)

		// Offers
		api.POST("/hotels/:id/travels", controller.CreateTravelRequest)

		api.GET("/hotels/:id/travels/:tid", controller.GetTravelById)

		// Update
		api.PUT("/hotels/:id/travels/:tid", controller.UpdateTravel)
		// Delete
		api.DELETE("/hotels/:id/travels/:tid", controller.DeleteTravel)

		// Tags
		api.POST("/tags", controller.CreateTagRequest)
		// Read
		api.GET("/tags", controller.ListAllTags)

		api.GET("/tags/:id", controller.GetTagById)

		// Update
		api.PUT("/tags/:id", controller.UpdateTag)
		// Delete
		api.DELETE("/tags/:id", controller.DeleteTagRequest)
	}

}
