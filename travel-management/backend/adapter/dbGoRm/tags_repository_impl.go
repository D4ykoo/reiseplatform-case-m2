package dbgorm

import (
	"github.com/mig3177/travelmanagement/adapter/dbGoRm/entities"
	"github.com/mig3177/travelmanagement/domain/model"
	"gorm.io/gorm/clause"
)

type TagRepositoryImpl struct {
	db pgRepository
}

// Open connection to the database
// Creates tables if not exist
func NewTagRepository(min int, max int) TagRepositoryImpl {
	con := initPGConnection(min, max)
	con.createTable(&entities.HotelEntity{}, &entities.PictureEntity{}, &entities.TravelEntity{}, &entities.TagEntity{})

	return TagRepositoryImpl{db: con}
}

//
// CRUD OPERATIONS:
//

func (repo TagRepositoryImpl) Create(tag *model.Tag) (*model.Tag, error) {
	entity := entities.ToTagEntity(tag)

	res := repo.db.Connection.Create(tag)
	return entities.ToTagModel(entity), res.Error
}

func (repo TagRepositoryImpl) Update(tag *model.Tag) (*model.Tag, error) {

	// search of existing entry
	var entity entities.TagEntity
	result := repo.db.Connection.Model(&entities.TagEntity{}).First(&entity, tag.Id)

	// Cancel update if entry not exist
	if result.Error != nil {
		return tag, result.Error
	}

	update := entities.ToTagEntity(tag)

	// Update entry
	entity.Name = update.Name

	res := repo.db.Connection.Save(&entity)

	return entities.ToTagModel(&entity), res.Error
}

func (repo TagRepositoryImpl) Delete(id uint) error {
	// search of existing entry
	var entity entities.TagEntity
	result := repo.db.Connection.Model(&entities.TagEntity{}).First(&entity, id)

	// Cancel deletion if entry not exist
	if result.Error != nil {
		return result.Error
	}

	// Delte all
	return repo.db.Connection.Select(clause.Associations).Delete(&entity).Error
}

func (repo TagRepositoryImpl) ListAll() ([]*model.Tag, error) {
	var entity []entities.TagEntity
	res := repo.db.Connection.Model(&entities.TagEntity{}).Find(&entity)
	tags := make([]*model.Tag, len(entity))

	if res.Error == nil {
		for i, tag := range entity {
			tags[i] = entities.ToTagModel(&tag)
		}
	}
	return tags, res.Error
}

func (repo TagRepositoryImpl) FindByID(id uint) (*model.Tag, error) {
	var entity entities.TagEntity
	res := repo.db.Connection.Model(&entities.TagEntity{}).First(&entity, id)
	return entities.ToTagModel(&entity), res.Error
}

func (repo TagRepositoryImpl) Count() (int64, error) {
	var count int64
	res := repo.db.Connection.Model(&entities.TagEntity{}).Count(&count)
	return count, res.Error
}
