package domain

import (
	"time"

	"github.com/mig3177/travelmanagement/domain/model"
)

type HotelService interface {
	NewHotel(string, model.Address, model.Vendor, string, []*model.Picture, []*model.Tag) (*model.Hotel, error)
	GetHotel(uint) (*model.Hotel, error)
	ListHotelTravel() ([]*model.Hotel, error)
	FindHotelTravel(string, string, *time.Time, *time.Time, []uint) ([]*model.Hotel, error)
	FindHotelByTravel(uint, uint) (*model.Hotel, error)
	UpdateHotel(*model.Hotel) (*model.Hotel, error)
	RemoveHotel(uint, string) error
}
