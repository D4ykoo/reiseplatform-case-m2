package application

import (
	"github.com/google/uuid"
	"github.com/mig3177/travelmanagement/domain/model"
)

type HotelServiceImpl struct {
}

func (HotelServiceImpl) CreateHotel(Name string, Address model.Address, description string) model.Hotel {
	return model.Hotel{}
}
func (HotelServiceImpl) FindHotelByID(uuid.UUID) model.Hotel {
	return model.Hotel{}
}
func (HotelServiceImpl) FindHotelByName(Name string) []model.Hotel {
	return []model.Hotel{}
}
func (HotelServiceImpl) UpdateHotel(HotelRef *model.Hotel) model.Hotel { return model.Hotel{} }
func (HotelServiceImpl) DeleteHotel(ID uuid.UUID) bool                 { return false }
