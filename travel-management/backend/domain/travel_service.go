package domain

import (
	"time"

	"github.com/mig3177/travelmanagement/domain/model"
)

type TravelService interface {
	NewHotel(string, model.Address, uint, string, []*model.Picture) (*model.Hotel, error)
	NewTravel(hotelRef uint, vendor uint, from time.Time, to time.Time, price float32, description string) (*model.Travel, error)

	GetHotel(uint) (*model.Hotel, error)
	GetTravel(uint) (*model.Travel, error)

	FindHotelTravel(string, string, time.Time, time.Time, ...uint) ([]*model.Hotel, error)

	UpdateHotel(*model.Hotel) (*model.Hotel, error)
	UpdateTravel(*model.Travel, uint) (*model.Travel, error)

	RemoveHotel(uint) error
	RemoveTravel(uint) error
}
