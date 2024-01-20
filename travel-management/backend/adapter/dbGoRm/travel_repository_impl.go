package dbgorm

import (
	"github.com/mig3177/travelmanagement/adapter/dbGoRm/entities"
	"github.com/mig3177/travelmanagement/domain/model"
)

type TravelRepositoryImpl struct {
	db pgRepository
}

// Open connection to the database
// Creates tables if not exist
func NewTravelRepository(min, max int) TravelRepositoryImpl {
	con := initPGConnection(min, max)
	con.createTable(&entities.HotelEntity{}, &entities.PictureEntity{}, &entities.TagEntity{}, &entities.TravelEntity{})
	return TravelRepositoryImpl{db: con}
}

//
// CRUD OPERATIONS:
//

func (repo TravelRepositoryImpl) Create(travel *model.Travel, hotelref uint) (*model.Travel, error) {
	entity := entities.ToTravelEntity(travel)
	entity.HotelRef = hotelref
	res := repo.db.Connection.Create(entity)
	return entities.ToTravelModel(entity), res.Error
}

func (repo TravelRepositoryImpl) Update(travel *model.Travel, id uint, hotelref uint) (*model.Travel, error) {

	// search of existing entry
	var entity entities.TravelEntity
	result := repo.db.Connection.Model(&entities.TravelEntity{}).Where("hotel_ref = ?", hotelref).First(&entity, id)

	// Cancel update if entry not exist
	if result.Error != nil {
		return travel, result.Error
	}

	update := entities.ToTravelEntity(travel)
	entity.From = update.From
	entity.To = update.To
	entity.Price = update.Price
	entity.VendorRef = update.VendorRef
	entity.VendorName = update.VendorName
	entity.Description = update.Description

	res := repo.db.Connection.Save(&entity)

	return entities.ToTravelModel(&entity), res.Error
}

func (repo TravelRepositoryImpl) Delete(id uint) error {
	var entity entities.TravelEntity
	result := repo.db.Connection.Model(&entities.TravelEntity{}).First(&entity, id)

	// Cancel deletion if entry not exist
	if result.Error != nil {
		return result.Error
	}

	result = repo.db.Connection.Delete(&entity)
	return result.Error
}

func (repo TravelRepositoryImpl) ListAll(hotelref uint) ([]*model.Travel, error) {
	var entites []entities.TravelEntity
	res := repo.db.Connection.Model(&entities.TravelEntity{}).Where("HotelRef = ?", hotelref).Find(&entites)

	travels := make([]*model.Travel, len(entites))

	if res.Error == nil {
		for i, travel := range entites {
			travels[i] = entities.ToTravelModel(&travel)
		}
	}

	return travels, res.Error
}

func (repo TravelRepositoryImpl) FindByID(id uint) (*model.Travel, error) {
	entity := entities.TravelEntity{}
	res := repo.db.Connection.Model(&entities.TravelEntity{}).First(&entity, id)
	if res.Error != nil {
		return &model.Travel{}, res.Error
	}
	travel := entities.ToTravelModel(&entity)
	return travel, nil
}

func (repo TravelRepositoryImpl) Count() (int64, error) {
	var count int64
	res := repo.db.Connection.Model(&entities.TravelEntity{}).Count(&count)
	return count, res.Error
}
