package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mig3177/hotelmanagement/domain/model"
)

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

}
