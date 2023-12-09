package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mig3177/hotelmanagement/adapter"
	"github.com/mig3177/hotelmanagement/application"
	"github.com/mig3177/hotelmanagement/domain"
	"github.com/mig3177/hotelmanagement/domain/model"
	"github.com/mig3177/hotelmanagement/ports"
)

func abc(domain.HotelService) {

}
func def(ports.HotelRepository) {

}
func main() {

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
	ef := adapter.HotelRepositoryImpl{}
	def(ef)
}
