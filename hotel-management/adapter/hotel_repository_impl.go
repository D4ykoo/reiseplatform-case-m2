package adapter

import (
	"github.com/google/uuid"
	"github.com/mig3177/hotelmanagement/adapter/entities"
	"github.com/mig3177/hotelmanagement/domain/model"
	"gorm.io/gorm/clause"
)

type HotelRepositoryImpl struct {
	db pgRepository
}

func NewHotelRepository() HotelRepositoryImpl {
	con := initPGConnection(10, 100)
	con.createTable(&entities.HotelEntity{}, &entities.PictureEnitiy{})
	return HotelRepositoryImpl{db: con}
}

func (repo HotelRepositoryImpl) Save(hotel model.Hotel) {
	pictures := make([]entities.PictureEnitiy, len(hotel.Pictures))
	for i, pic := range hotel.Pictures {
		pictures[i] = entities.PictureEnitiy{ID: pic.ID, Payload: pic.Payload, HotelRef: hotel.ID, Description: pic.Description}
	}
	entity := entities.HotelEntity{ID: hotel.ID, Name: hotel.Name,
		Street: hotel.Address.Street, State: hotel.Address.State, Land: hotel.Address.Land,
		VendorRef: hotel.Vendor.ID, Description: hotel.Description, Pictures: pictures}
	repo.db.Connection.Create(&entity)
	//if len(pictures) < 0 {
	//		repo.db.Connection.Create(&pictures)
	//	}
}

func (repo HotelRepositoryImpl) Update(hotel model.Hotel) {
	pictures := make([]entities.PictureEnitiy, len(hotel.Pictures))
	for i, pic := range hotel.Pictures {
		pictures[i] = entities.PictureEnitiy{ID: pic.ID, Payload: pic.Payload, HotelRef: hotel.ID, Description: pic.Description}
	}
	entity := entities.HotelEntity{ID: hotel.ID, Name: hotel.Name,
		Street: hotel.Address.Street, State: hotel.Address.State, Land: hotel.Address.Land,
		VendorRef: hotel.Vendor.ID, Description: hotel.Description, Pictures: pictures}
	repo.db.Connection.Save(&entity)
}

func (repo HotelRepositoryImpl) Delete(hotel model.Hotel) bool {
	pictures := make([]entities.PictureEnitiy, len(hotel.Pictures))
	for i, pic := range hotel.Pictures {
		pictures[i] = entities.PictureEnitiy{ID: pic.ID, Payload: pic.Payload, HotelRef: hotel.ID, Description: pic.Description}
	}
	entity := entities.HotelEntity{ID: hotel.ID, Name: hotel.Name,
		Street: hotel.Address.Street, State: hotel.Address.State, Land: hotel.Address.Land,
		VendorRef: hotel.Vendor.ID, Description: hotel.Description, Pictures: pictures}
	repo.db.Connection.Select(clause.Associations).Delete(&entity, entity.ID)
	return false
}

func (repo HotelRepositoryImpl) GetAll() ([]model.Hotel, error) {
	var entity []entities.HotelEntity
	err := repo.db.Connection.Model(&entities.HotelEntity{}).Preload("Pictures").Find(&entity).Error
	hotels := make([]model.Hotel, len(entity))

	if err == nil {
		for i, hotel := range entity {
			pictures := make([]model.Picture, len(hotel.Pictures))
			for j, pic := range hotel.Pictures {
				pictures[j] = model.Picture{ID: pic.ID, Payload: pic.Payload, Description: pic.Description}
			}
			hotels[i] = model.Hotel{ID: hotel.ID, Name: hotel.Name,
				Address: model.Address{Street: hotel.Street, State: hotel.State, Land: hotel.Land},
				Vendor:  model.Vendor{ID: hotel.ID}, Description: hotel.Description, Pictures: pictures}
		}
	}
	return hotels, err
}

func (repo HotelRepositoryImpl) Count() int {
	hotels, err := repo.GetAll()
	if err != nil {
		return 0
	} else {
		return len(hotels)
	}
}

func (repo HotelRepositoryImpl) FindByID(id uuid.UUID) model.Hotel {
	var entity entities.HotelEntity
	err := repo.db.Connection.Model(&entities.HotelEntity{}).Preload("Pictures").First(&entity, id).Error
	if err != nil {
		return model.Hotel{}
	}
	pictures := make([]model.Picture, len(entity.Pictures))
	for i, pic := range entity.Pictures {
		pictures[i] = model.Picture{ID: pic.ID, Payload: pic.Payload, Description: pic.Description}
	}
	return model.Hotel{ID: entity.ID, Name: entity.Name,
		Address: model.Address{Street: entity.Street, State: entity.State, Land: entity.Land},
		Vendor:  model.Vendor{ID: entity.VendorRef}, Description: entity.Description, Pictures: pictures}
}

/*
func (repo HotelRepositoryImpl) FindByID(id uuid.UUID) model.Hotel {

	entity := repo.db.getBy("ID = ?", id.String())[0]
	return model.Hotel{ID: entity.ID, Name: entity.Name, Address: model.Address{Street: entity.Street, State: entity.State, Land: entity.Land}, Description: entity.Description}
}
users
func (repo HotelRepositoryImpl) Count() int64 {
	res := repo.db.getBy("ID != ?", "")
	return int64(len(res))
}
*/
