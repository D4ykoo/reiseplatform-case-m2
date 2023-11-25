package domain

import (
	"github.com/google/uuid"
	"github.com/mig3177/hotelmanagement/domain/model"
)

type HotelService interface {
	createHotel(Name string, Address model.Address, description string) model.Hotel
	findHotelByID(uuid.UUID) model.Hotel
	findHotelByName(Name string) []model.Hotel
	updateHotel(HotelRef *model.Hotel) model.Hotel
	deleteHotel(ID uuid.UUID) bool
}
