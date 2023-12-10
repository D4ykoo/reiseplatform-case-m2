package adapter

import (
	"github.com/google/uuid"
	"github.com/mig3177/hotelmanagement/adapter/entities"
	"github.com/mig3177/hotelmanagement/domain/model"
)

type HotelRepositoryImpl struct {
	db CrudRepository[entities.HotelEntity]
}

func NewHotelRepository() HotelRepositoryImpl {
	return HotelRepositoryImpl{
		db: initPGConnection[entities.HotelEntity](),
	}
}

func (repo HotelRepositoryImpl) Save(hotel model.Hotel) {
	entity := entities.HotelEntity{ID: hotel.ID, Name: hotel.Name, Street: hotel.Address.Street, State: hotel.Address.State, Land: hotel.Address.Land, Description: hotel.Description}
	repo.db.create(entity)
}

func (repo HotelRepositoryImpl) Update(hotel model.Hotel) {
	entity := entities.HotelEntity{ID: hotel.ID, Name: hotel.Name, Street: hotel.Address.Street, State: hotel.Address.State, Land: hotel.Address.Land, Description: hotel.Description}
	repo.db.update(entity)
}

func (repo HotelRepositoryImpl) Delete(hotel model.Hotel) bool {
	entity := entities.HotelEntity{ID: hotel.ID, Name: hotel.Name, Street: hotel.Address.Street, State: hotel.Address.State, Land: hotel.Address.Land, Description: hotel.Description}
	return repo.db.delete(entity)
}
func (repo HotelRepositoryImpl) FindByID(id uuid.UUID) model.Hotel {
	entity := repo.db.getBy("ID = ?", id.String())[0]
	return model.Hotel{ID: entity.ID, Name: entity.Name, Address: model.Address{Street: entity.Street, State: entity.State, Land: entity.Land}, Description: entity.Description}
}

func (repo HotelRepositoryImpl) Count() int64 {
	res := repo.db.getBy("ID != ?", "")
	return int64(len(res))
}
