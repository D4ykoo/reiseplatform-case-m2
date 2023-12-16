package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	dbgorm "github.com/mig3177/travelmanagement/adapter/dbGoRm"
	"github.com/mig3177/travelmanagement/domain"
	"github.com/mig3177/travelmanagement/domain/model"
	"github.com/mig3177/travelmanagement/ports"
)

func abc(domain.HotelService) {

}
func def(ports.HotelRepository) {

}
func main() {
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
	*/
	newHotel := model.Hotel{ID: uuid.New(), Name: "Im weißen Rößl", Address: model.Address{Street: "Markt 47", State: "St. Wolfgang im Salzkammergut", Land: "Österreich"}, Pictures: []*model.Picture{{ID: uuid.New(), Payload: "kandfmanfadfafaf", Description: "ABC"}}}

	/**/
	fmt.Println("-------------------------------------------------")
	travel := model.Travel{ID: uuid.New(), Hotel: []*model.Hotel{&newHotel}, Vendor: model.Vendor{ID: uuid.New(), Username: "Herbert"}, From: time.Now(), To: time.Now(), Price: 320.50, Description: "DES", Tags: []*model.Tag{{Typ: 55, Name: "Strand"}, {Typ: 40, Name: "Wandern"}}}
	fmt.Println(travel.Hotel)
	repo2 := dbgorm.NewTravelRepository(10, 100)
	repo2.Create(&travel)
	fmt.Println("####")
	nn, _ := repo2.FindByID(travel.ID)
	fmt.Println(nn.Tags)
	repo2.Delete(&travel)
}
