package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mig3177/hotelmanagement/adapter"
	"github.com/mig3177/hotelmanagement/domain"
	"github.com/mig3177/hotelmanagement/domain/model"
	"github.com/mig3177/hotelmanagement/ports"
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

		repo := adapter.NewHotelRepository()*/
	newHotel := model.Hotel{ID: uuid.New(), Name: "Im weißen Rößl", Address: model.Address{Street: "Markt 47", State: "St. Wolfgang im Salzkammergut", Land: "Österreich"}, Pictures: []model.Picture{{ID: uuid.New(), Payload: "kandfmanfadfafaf", Description: "ABC"}}}
	/*repo.Save(newHotel)
	len := repo.Count()
	fmt.Println(len)
	fmt.Println(repo.FindByID(newHotel.ID))
	newHotel.Address.Street = "Markt 74"
	repo.Update(newHotel)
	fmt.Println(repo.FindByID(newHotel.ID))
	repo.Delete(newHotel)
	len = repo.Count()
	fmt.Println(len)

	hotels, err := repo.GetAll()

	fmt.Println(hotels)
	fmt.Println(err)
	*/
	fmt.Println("-------------------------------------------------")
	travel := model.Travel{ID: uuid.New(), Hotel: newHotel, Vendor: model.Vendor{ID: uuid.New(), Username: "Herbert"}, From: time.Now(), To: time.Now(), Price: 320.50, Description: "DES", Tags: []model.Tag{{Typ: 55, Name: "Strand"}, {Typ: 40, Name: "Wandern"}}}
	fmt.Println(travel.Tags)
	repo2 := adapter.NewTravelRepository()
	repo2.Save(travel)
	repo2.FindByID(travel.ID)
	repo2.Delete(travel)
}
