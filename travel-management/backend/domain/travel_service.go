package domain

import (
	"time"

	"github.com/mig3177/travelmanagement/domain/model"
)

type TravelService interface {
	NewHotel(string, model.Address, model.Vendor, string, []*model.Picture, []*model.Tag) (*model.Hotel, error)
	NewTravel(hotelRef uint, vendor uint, from time.Time, to time.Time, price float32, description string) (*model.Travel, error)
	NewTag(string) (*model.Tag, error)

	GetHotel(uint) (*model.Hotel, error)
	GetTravel(uint) (*model.Travel, error)
	GetTag(uint) (*model.Tag, error)

	ListHotelTravel() ([]*model.Hotel, error)
	FindHotelTravel(string, string, *time.Time, *time.Time, []uint) ([]*model.Hotel, error)
	FindHotelByTravel(uint, uint) (*model.Hotel, error)
	ListTags() ([]*model.Tag, error)

	UpdateHotel(*model.Hotel) (*model.Hotel, error)
	UpdateTravel(*model.Travel, uint) (*model.Travel, error)
	UpdateTags(*model.Tag) (*model.Tag, error)

	RemoveHotel(uint) error
	RemoveTravel(uint) error
	RemoveTag(uint) error
}
