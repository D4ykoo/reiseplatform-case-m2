package dbgorm

import (
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
	con.createTable(&entities.HotelEntity{}, &entities.PictureEntity{}, &entities.TravelEntity{}, &entities.TagEntity{})

	return HotelRepositoryImpl{db: con}
}

//
// CRUD OPERATIONS:
//

func (repo HotelRepositoryImpl) Create(hotel *model.Hotel) (*model.Hotel, error) {
	entity := entities.ToHotelEntity(hotel)

	res := repo.db.Connection.Create(entity)
	return entities.ToHotelModel(entity), res.Error
}

func (repo HotelRepositoryImpl) Update(hotel *model.Hotel) (*model.Hotel, error) {

	// search of existing entry
	var entity entities.HotelEntity
	result := repo.db.Connection.Model(&entities.HotelEntity{}).Preload("Pictures").First(&entity, hotel.Id)

	// Cancel update if entry not exist
	if result.Error != nil {
		return hotel, result.Error
	}

	//repo.db.Connection.Model(&entity).Association("Pictures").Clear()
	//repo.db.Connection.Model(&entity).Association("Hotels").Clear()

	update := entities.ToHotelEntity(hotel)

	// Update entry
	entity.Name = update.Name
	entity.Street = update.Street
	entity.State = update.State
	entity.Land = update.Land
	entity.Pictures = update.Pictures
	entity.VendorRef = update.VendorRef
	entity.Travels = update.Travels

	res := repo.db.Connection.Save(&entity)

	return entities.ToHotelModel(&entity), res.Error
}

func (repo HotelRepositoryImpl) Delete(id uint) error {
	// search of existing entry
	var entity entities.HotelEntity
	result := repo.db.Connection.Model(&entities.HotelEntity{}).Preload("Pictures").Preload("Travels").First(&entity, id)

	// Cancel deletion if entry not exist
	if result.Error != nil {
		return result.Error
	}

	return repo.db.Connection.Select(clause.Associations).Delete(&entity).Error
}

func (repo HotelRepositoryImpl) ListAll() ([]*model.Hotel, error) {
	var entity []entities.HotelEntity
	res := repo.db.Connection.Model(&entities.HotelEntity{}).Preload("Pictures").Preload("Travels").Find(&entity)
	hotels := make([]*model.Hotel, len(entity))

	if res.Error == nil {
		for i, hotel := range entity {
			hotels[i] = entities.ToHotelModel(&hotel)
		}
	}
	return hotels, res.Error
}

func (repo HotelRepositoryImpl) FindByID(id uint) (*model.Hotel, error) {
	var entity entities.HotelEntity
	res := repo.db.Connection.Model(&entities.HotelEntity{}).Preload("Pictures").Preload("Travels").First(&entity, id)
	return entities.ToHotelModel(&entity), res.Error
}

func (repo HotelRepositoryImpl) Count() (int64, error) {
	var count int64
	res := repo.db.Connection.Model(&entities.HotelEntity{}).Count(&count)
	return count, res.Error
}
