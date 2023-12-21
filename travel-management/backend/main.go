package main

import (
	"fmt"
	"time"

	dbgorm "github.com/mig3177/travelmanagement/adapter/dbGoRm"
	"github.com/mig3177/travelmanagement/domain/model"
	"github.com/mig3177/travelmanagement/ports/outbound"
)

func main() {
	/*
		// Env

		// Outgoing
		// Repository
		hotelRepo := dbgorm.NewHotelRepository(10, 100)
		//travelRepo := dbgorm.NewTravelRepository(10, 100)

		// Application
		// Service
		hotelService := application.NewHotelService(hotelRepo)
		//travelService := application.New(travelRepo)

		// Incomming
		// Controller
		hotelController := api.NewRestHotelController(hotelService)
		// Router
		router := gin.Default()

		// CRUD
		// Create
		router.POST("/api/hotels", func(c *gin.Context) {
			c.String(200, "Hello, World!")
		})
		// Read
		router.GET("/api/hotels/:id", hotelController.GetHotelRequest)

		router.GET("/api/hotels/", hotelController.GetHotelsByNameRequest)

		// Update
		router.PUT("/api/hotels:id", hotelController.UpdateHotelRequest)
		// Delete
		router.DELETE("/api/hotels:id", func(c *gin.Context) {
			c.String(200, "Hello, World!")
		})

		// Run the server on port 8080
		router.Run(":8080")
	*/
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

		fmt.Println("-------------------------------------------------")
		travel := model.Travel{ID: uuid.New(), Hotel: []*model.Hotel{&newHotel}, Vendor: model.Vendor{ID: uuid.New(), Username: "Herbert"}, From: time.Now(), To: time.Now(), Price: 320.50, Description: "DES", Tags: []*model.Tag{{Typ: 55, Name: "Strand"}, {Typ: 40, Name: "Wandern"}}}
		fmt.Println(travel.Hotel)
		repo2 := dbgorm.NewTravelRepository(10, 100)
		repo2.Create(&travel)
		fmt.Println("####")
		nn, _ := repo2.FindByID(travel.ID)
		fmt.Println(nn.From)
		dd, ee := repo2.FindBetween(travel.From, travel.To)
		fmt.Println(dd[0].Hotel[0].Address, ee)
		repo2.Delete(&travel)
	*/
	/*
		abc := &entities.PictureEntity{Model: gorm.Model{ID: 1253}, Payload: "DJSKJDKS", Description: "adfadf"}
		ar := []*entities.PictureEntity{abc}
		def := &entities.HotelEntity{Model: gorm.Model{ID: 5262222}, Name: "HUHUH", Street: "STEET", State: "STAST", Land: "§LNAN", VendorRef: 25314, Description: "ööö", Pictures: ar}
		fmt.Println(abc)
		fmt.Println(def)
	*/

	var repo outbound.HotelRepository
	repo = dbgorm.NewHotelRepository(10, 100)

	pic := model.Picture{Payload: "PAYLOAD", Description: "DES"}
	pics := []*model.Picture{&pic}

	hotel := model.Hotel{Name: "waterloo", Address: model.Address{Street: "STREET", State: "STATE", Land: "LAND"}, Description: "DES2", Vendor: model.Vendor{Id: 1253, Username: "user"}, Pictures: pics}

	newhotel, _ := repo.Create(&hotel)

	hotelByID, _ := repo.FindByID(newhotel.Id)
	from := time.Date(

		2009, 11, 17, 0, 0, 0, 0, time.UTC)

	to := time.Date(

		2009, 11, 25, 0, 0, 0, 0, time.UTC)
	travel := model.Travel{Vendor: model.Vendor{Id: 12358, Username: "Herbert"}, From: from, To: to, Price: 8884.51, Description: "DES22", Tags: []*model.Tag{{Typ: 58, Name: "Wandern"}}}

	//new_Pic := model.Picture{Payload: "Kuwukiland-PAYLOAD", Description: "Kuwukiland-DES"}
	pic.Id = newhotel.Pictures[0].Id
	//pics2 := []*model.Picture{&new_Pic, &pic}
	//travels := []*model.Travel{&travel}

	//newhotel.Address.Land = "Kuwukiland"
	//newhotel.Pictures = pics2
	//newhotel.Travels = travels

	repo2 := dbgorm.NewTravelRepository(10, 100)
	repo2.Create(&travel, newhotel.Id)

	//updated_hotel, _ := repo.Update(newhotel)
	//repo.Delete(newhotel.Id)
	currentHotels, _ := repo.ListAll()

	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println(currentHotels)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println(hotelByID)
	fmt.Println("--------------------------------------------------------------------------------")
	hotelByID, _ = repo.FindByID(newhotel.Id)
	fmt.Println(hotelByID)

	//fmt.Println(updated_hotel)

	/*
	   var repo outbound.TravelRepository
	   repo = dbgorm.NewTravelRepository(10, 100)

	   pic := model.Picture{Payload: "PAYLOAD", Description: "DES"}
	   pics := []*model.Picture{&pic}
	   from := time.Date(

	   	2009, 11, 17, 0, 0, 0, 0, time.UTC)

	   to := time.Date(

	   	2009, 11, 25, 0, 0, 0, 0, time.UTC)

	   hotel := model.Hotel{Name: "waterloo", Address: model.Address{Street: "STREET", State: "STATE", Land: "LAND"}, Description: "DES2", Vendor: model.Vendor{Id: 1253, Username: "user"}, Pictures: pics}
	   travel := model.Travel{Hotel: []*model.Hotel{&hotel}, Vendor: model.Vendor{Id: 12358, Username: "Herbert"}, From: from, To: to, Price: 8884.51, Description: "DES22", Tags: []*model.Tag{&model.Tag{Typ: 58, Name: "Wandern"}}}
	   current_travel, _ := repo.Create(&travel)
	   //current_travels, _ := repo.ListAll()
	   travel_ID, _ := repo.FindByID(current_travel.Id)

	   new_Pic := model.Picture{Payload: "Kuwukiland-PAYLOAD", Description: "Kuwukiland-DES"}
	   pics2 := []*model.Picture{&new_Pic, &pic}

	   hotel2 := model.Hotel{Name: "ELBA", Address: model.Address{Street: "STREET2", State: "STATE2", Land: "LAND2"}, Description: "DES2", Vendor: model.Vendor{Id: 1253, Username: "user"}, Pictures: pics2}

	   current_travel.Hotel[0] = &hotel2
	   current_travel.Description = " DAS WIRD NICHTS"

	   updated_hotel, _ := repo.Update(current_travel)

	   hotel_name, _ := repo.FindByName(hotel.Name)
	   fmt.Println(hotel_name)
	   repo.Delete(updated_hotel.Id)

	   fmt.Println("--------------------------------------------------------------------------------")
	   //fmt.Println(current_travel)
	   fmt.Println("--------------------------------------------------------------------------------")
	   //fmt.Println(current_travels)
	   fmt.Println("--------------------------------------------------------------------------------")
	   fmt.Println(travel_ID)
	   fmt.Println("--------------------------------------------------------------------------------")
	   fmt.Println(updated_hotel)
	*/
}
