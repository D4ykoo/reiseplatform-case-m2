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
	result := repo.db.Connection.Model(&entities.HotelEntity{}).Preload(clause.Associations).First(&entity, hotel.Id)

	// Cancel update if entry not exist
	if result.Error != nil {
		return hotel, result.Error
	}

	update := entities.ToHotelEntity(hotel)

	rPicIds := findRemovedPics(update.Pictures, entity.Pictures)
	for _, id := range rPicIds {
		repo.DeletePic(id)
	}

	tags := findRemovedTags(update.Tags, entity.Tags)
	repo.db.Connection.Model(&entity).Association("Tags").Delete(tags)

	// Update entry
	entity.Name = update.Name
	entity.Street = update.Street
	entity.State = update.State
	entity.Land = update.Land
	entity.Pictures = update.Pictures
	entity.VendorRef = update.VendorRef
	entity.VendorName = update.VendorName
	entity.Travels = update.Travels
	entity.Description = update.Description
	entity.Tags = update.Tags

	res := repo.db.Connection.Save(&entity)

	return entities.ToHotelModel(&entity), res.Error
}

func (repo HotelRepositoryImpl) Delete(id uint) error {
	// search of existing entry
	var entity entities.HotelEntity
	result := repo.db.Connection.Model(&entities.HotelEntity{}).First(&entity, id)

	// Cancel deletion if entry not exist
	if result.Error != nil {
		return result.Error
	}

	// Delte all
	return repo.db.Connection.Select(clause.Associations).Delete(&entity).Error
}

func (repo HotelRepositoryImpl) DeletePic(id uint) error {
	// search of existing entry
	var entity entities.PictureEntity
	result := repo.db.Connection.Model(&entities.PictureEntity{}).First(&entity, id)

	// Cancel deletion if entry not exist
	if result.Error != nil {
		return result.Error
	}

	// Delte all
	return repo.db.Connection.Select(clause.Associations).Delete(&entity).Error
}

func (repo HotelRepositoryImpl) ListAll() ([]*model.Hotel, error) {
	var entity []entities.HotelEntity
	res := repo.db.Connection.Model(&entities.HotelEntity{}).Preload(clause.Associations).Find(&entity)
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
	res := repo.db.Connection.Model(&entities.HotelEntity{}).Preload(clause.Associations).First(&entity, id)
	return entities.ToHotelModel(&entity), res.Error
}

func (repo HotelRepositoryImpl) Count() (int64, error) {
	var count int64
	res := repo.db.Connection.Model(&entities.HotelEntity{}).Count(&count)
	return count, res.Error
}

func findRemovedPics(updatePics, currentPics []*entities.PictureEntity) []uint {
	missingEntries := make([]uint, 0)

	set := make(map[int]bool)

	// Create a set from the first array
	for _, pic := range updatePics {
		set[int(pic.ID)] = true // setting the initial value to true
	}

	// Check elements in the second array against the set
	for _, pic := range currentPics {
		if !set[int(pic.ID)] {
			missingEntries = append(missingEntries, pic.ID)
		}
	}
	return missingEntries
}

func findRemovedTags(updateTags, currentTags []*entities.TagEntity) []*entities.TagEntity {
	missingEntries := make([]*entities.TagEntity, 0)

	set := make(map[int]bool)

	// Create a set from the first array
	for _, tag := range updateTags {
		set[int(tag.ID)] = true // setting the initial value to true
	}

	// Check elements in the second array against the set
	for _, tag := range currentTags {
		if !set[int(tag.ID)] {
			missingEntries = append(missingEntries, tag)
		}
	}
	return missingEntries
}
