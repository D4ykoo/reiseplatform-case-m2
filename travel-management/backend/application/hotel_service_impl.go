package application

import (
	"github.com/google/uuid"
	"github.com/mig3177/travelmanagement/domain/model"
	"github.com/mig3177/travelmanagement/ports/outbound"
)

type HotelServiceImpl struct {
	hotels outbound.HotelRepository
}

func NewHotelService(repo outbound.HotelRepository) HotelServiceImpl {
	return HotelServiceImpl{hotels: repo}
}

func (service HotelServiceImpl) CreateHotel(name string, address model.Address, userid uuid.UUID, description string, pics []*model.Picture) (*model.Hotel, error) {
	// TODO check user is valid
	// TODO check hotel already exist
	hotel := &model.Hotel{ID: uuid.New(), Address: address, Name: name, Description: description, Vendor: model.Vendor{ID: userid, Username: "Walter"}, Pictures: pics}
	err := service.hotels.Save(hotel)
	return hotel, err
}

func (service HotelServiceImpl) FindHotelByID(id uuid.UUID) (*model.Hotel, error) {
	return service.FindHotelByID(id)
}

func (service HotelServiceImpl) FindHotelByName(Name string) ([]*model.Hotel, error) {
	return nil, nil
}

func (service HotelServiceImpl) UpdateHotel(id uuid.UUID, name string, address model.Address, userid uuid.UUID, description string, pics []*model.Picture) error {
	hotel, err := service.hotels.FindByID(id)
	if err != nil {
		return err
	}
	hotel.Address = address
	hotel.Name = name
	hotel.Vendor = model.Vendor{ID: userid, Username: "Maria"}
	hotel.Description = description
	hotel.Pictures = pics
	err = service.hotels.Update(hotel)
	return err
}

func (service HotelServiceImpl) DeleteHotel(id uuid.UUID) error {
	// TODO check user is valid
	hotel, err := service.hotels.FindByID(id)
	if err != nil {
		return err
	}
	return service.hotels.Delete(hotel)
}
