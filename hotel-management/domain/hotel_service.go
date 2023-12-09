package domain

import (
	"github.com/google/uuid"
	"github.com/mig3177/hotelmanagement/domain/model"
)

type HotelService interface {
	CreateHotel(Name string, Address model.Address, description string) model.Hotel
	FindHotelByID(uuid.UUID) model.Hotel
	FindHotelByName(Name string) []model.Hotel
	UpdateHotel(HotelRef *model.Hotel) model.Hotel
	DeleteHotel(ID uuid.UUID) bool
}
