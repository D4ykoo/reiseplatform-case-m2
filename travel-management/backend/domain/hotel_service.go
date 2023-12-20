package domain

import (
	"github.com/google/uuid"
	"github.com/mig3177/travelmanagement/domain/model"
)

type HotelService interface {
	CreateHotel(name string, address model.Address, userid uuid.UUID, description string, pics []*model.Picture) (*model.Hotel, error)
	FindHotelByID(uuid.UUID) (*model.Hotel, error)
	FindHotelByName(Name string) ([]*model.Hotel, error)
	UpdateHotel(*model.Hotel) (*model.Hotel, error)
	DeleteHotel(ID uuid.UUID) error
}
