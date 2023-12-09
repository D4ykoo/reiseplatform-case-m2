package adapter

import (
	"github.com/google/uuid"
	"github.com/mig3177/hotelmanagement/domain/model"
)

type HotelRepositoryImpl struct {
	db ManagementDB
}

func New() HotelRepositoryImpl {
	return HotelRepositoryImpl{
		db: NewDB(),
	}
}

// Save implements ports.HotelRepository.
func (HotelRepositoryImpl) Save(model.Hotel) {
	panic("unimplemented")
}

func (HotelRepositoryImpl) save(model.Hotel) {}
func (HotelRepositoryImpl) Update(model.Hotel) model.Hotel {
	return model.Hotel{}
}
func (HotelRepositoryImpl) Delete(uuid.UUID) bool {
	return false
}
func (HotelRepositoryImpl) FindByID(uuid.UUID) model.Hotel { return model.Hotel{} }
func (HotelRepositoryImpl) Count() int64                   { return 0 }
