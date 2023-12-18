package dbgorm

import (
	"github.com/google/uuid"
	"github.com/mig3177/travelmanagement/adapter/dbGoRm/entities"
	"github.com/mig3177/travelmanagement/domain/model"
	"gorm.io/gorm/clause"
)

type HotelRepositoryImpl struct {
	db pgRepository
}

// Open connection to the database
// Creates tables if not exist
func NewHotelRepository(min int, max int) HotelRepositoryImpl {
	con := initPGConnection(min, max)
	con.createTable(&entities.HotelEntity{}, &entities.PictureEntity{})

	return HotelRepositoryImpl{db: con}
}

// Function for converting from model to entity
func ToHotelEntity(hotel *model.Hotel) *entities.HotelEntity {
	pictures := make([]*entities.PictureEntity, len(hotel.Pictures))
	for i, pic := range hotel.Pictures {
		pictures[i] = &entities.PictureEntity{ID: pic.ID, Payload: pic.Payload, HotelRef: hotel.ID, Description: pic.Description}
	}
	return &entities.HotelEntity{ID: hotel.ID, Name: hotel.Name,
		Street: hotel.Address.Street, State: hotel.Address.State, Land: hotel.Address.Land,
		VendorRef: hotel.Vendor.ID, Description: hotel.Description, Pictures: pictures}
}

// Function for converting from entity to model
func ToHotelModel(entity *entities.HotelEntity) *model.Hotel {
	pictures := make([]*model.Picture, len(entity.Pictures))
	for i, pic := range entity.Pictures {
		pictures[i] = &model.Picture{ID: pic.ID, Payload: pic.Payload, Description: pic.Description}
	}
	return &model.Hotel{ID: entity.ID, Name: entity.Name,
		Address: model.Address{Street: entity.Street, State: entity.State, Land: entity.Land},
		Vendor:  model.Vendor{ID: entity.VendorRef}, Description: entity.Description, Pictures: pictures}
}

//
// CRUD OPERATIONS:
//

func (repo HotelRepositoryImpl) Create(hotel *model.Hotel) error {
	entity := ToHotelEntity(hotel)
	return repo.db.Connection.Create(entity).Error
}

func (repo HotelRepositoryImpl) Update(hotel *model.Hotel) error {
	entity := ToHotelEntity(hotel)
	return repo.db.Connection.Save(entity).Error
}

func (repo HotelRepositoryImpl) Delete(hotel *model.Hotel) error {
	entity := ToHotelEntity(hotel)
	return repo.db.Connection.Select(clause.Associations).Delete(entity, entity.ID).Error
}

func (repo HotelRepositoryImpl) ListAll() ([]*model.Hotel, error) {
	var entity []entities.HotelEntity
	err := repo.db.Connection.Model(&entities.HotelEntity{}).Preload("Pictures").Find(&entity).Error
	hotels := make([]*model.Hotel, len(entity))

	if err == nil {
		for i, hotel := range entity {
			hotels[i] = ToHotelModel(&hotel)
		}
	}
	return hotels, err
}

func (repo HotelRepositoryImpl) FindByID(id uuid.UUID) (*model.Hotel, error) {
	var entity entities.HotelEntity
	err := repo.db.Connection.Model(&entities.HotelEntity{}).Preload("Pictures").First(&entity, id).Error
	return ToHotelModel(&entity), err
}

func (repo HotelRepositoryImpl) Count() (int, error) {
	hotels, err := repo.ListAll()
	if err != nil {
		return 0, err
	} else {
		return len(hotels), err
	}
}
