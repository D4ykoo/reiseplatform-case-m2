package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mig3177/travelmanagement/adapter/api"
	dbgorm "github.com/mig3177/travelmanagement/adapter/dbGoRm"
	"github.com/mig3177/travelmanagement/application"
)

func main() {
	// Env

	repo := dbgorm.NewHotelRepository(10, 100)
	service := application.NewHotelService(repo)
	// Service
	controller := api.NewRestHotelController(service)
	// Router
	router := gin.Default()

	// CRUD
	router.POST("/api/hotels", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	router.GET("/api/hotels/:id", controller.GetHotelRequest)

	router.GET("/api/hotels/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	router.POST("/api/hotels/filter", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	router.PUT("/api/hotels:id", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	router.DELETE("/api/hotels:id", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// Run the server on port 8080
	router.Run(":8080")

	/*
				fmt.Println("HUHU")
				p := &model.Address{}
				p.Land = "LAND"
				uu := uuid.New()
				hotel := model.Hotel{
					ID: uu,
				}
				fmt.Println(p.Land)
				fmt.Println(hotel.ID)
				ade := application.HotelServiceImpl{}
				abc(ade)
				///

			repo := adapter.NewHotelRepository()
			newHotel := model.Hotel{ID: uuid.New(), Name: "Im weißen Rößl", Address: model.Address{Street: "Markt 47", State: "St. Wolfgang im Salzkammergut", Land: "Österreich"}, Pictures: []*model.Picture{{ID: uuid.New(), Payload: "kandfmanfadfafaf", Description: "ABC"}}}
			repo.Create(&newHotel)
			len := repo.Count()
			fmt.Println(len)
			uu := repo.FindByID(newHotel.ID)
			fmt.Println(uu)
			fmt.Println((uu.Pictures[0]))

			newHotel.Address.Street = "Markt 74"
			repo.Save(&newHotel)
			fmt.Println(repo.FindByID(newHotel.ID))
			repo.Delete(&newHotel)
			len = repo.Count()
			fmt.Println(len)

			hotels, err := repo.GetAll()
				fmt.Println(hotels)
			fmt.Println(err)

		newHotel := model.Hotel{ID: uuid.New(), Name: "Im weißen Rößl", Address: model.Address{Street: "Markt 47", State: "St. Wolfgang im Salzkammergut", Land: "Österreich"}, Pictures: []*model.Picture{{ID: uuid.New(), Payload: "kandfmanfadfafaf", Description: "ABC"}}}
	*/
	/*
		fmt.Println("-------------------------------------------------")
		travel := model.Travel{ID: uuid.New(), Hotel: []*model.Hotel{&newHotel}, Vendor: model.Vendor{ID: uuid.New(), Username: "Herbert"}, From: time.Now(), To: time.Now(), Price: 320.50, Description: "DES", Tags: []*model.Tag{{Typ: 55, Name: "Strand"}, {Typ: 40, Name: "Wandern"}}}
		fmt.Println(travel.Hotel)
		repo2 := dbgorm.NewTravelRepository(10, 100)
		repo2.Create(&travel)
		fmt.Println("####")
		nn, _ := repo2.FindByID(travel.ID)
		fmt.Println(nn.Tags)
		repo2.Delete(&travel)*/
}
