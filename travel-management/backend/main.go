package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mig3177/travelmanagement/adapter/api"
	dbgorm "github.com/mig3177/travelmanagement/adapter/dbGoRm"
	"github.com/mig3177/travelmanagement/application"
)

func main() {

	// Env

	// Outgoing
	// Repository
	hotelRepo := dbgorm.NewHotelRepository(10, 100)
	travelRepo := dbgorm.NewTravelRepository(10, 100)
	tagRepo := dbgorm.NewTagRepository(10, 100)

	// Application
	// Service
	travelService := application.New(hotelRepo, travelRepo, tagRepo)

	fmt.Println(travelService.FindHotelByTravel(1, 8))

	// Incomming
	// Controller
	service := api.New(travelService)
	// Router
	router := gin.Default()

	// CRUD
	// Create
	router.POST("/api/v1/hotels/", service.CreateHotelRequest)

	// Read
	router.GET("/api/v1/hotels/", service.FindHotels)

	router.GET("/api/v1/hotels/:id/", service.GetHotelById)

	// Update
	router.PUT("/api/v1/hotels/:id/", service.UpdateHotel)
	// Delete
	router.DELETE("/api/v1/hotels/:id/", service.DeleteHotelRequest)

	// Offers
	router.POST("/api/v1/hotels/:id/travels/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
	// Read
	router.GET("/api/v1/hotels/:id/travels/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	router.GET("/api/v1/hotels/:id/travels/:tid/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// Update
	router.PUT("/api/v1/hotels/:id/travels/:tid/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
	// Delete
	router.DELETE("/api/v1/hotels/:id/travels/:tid/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// Tags
	router.POST("/api/v1/tags/", service.CreateTagRequest)
	// Read
	router.GET("/api/v1/tags/", service.ListAllTags)

	router.GET("/api/v1/tags/:id/", service.GetTagById)

	// Update
	router.PUT("/api/v1/tags/:id/", service.UpdateTag)
	// Delete
	router.DELETE("/api/v1/tags/:id/", service.DeleteTagRequest)

	// Run the server on port 8080
	router.Run(":8080")
}
