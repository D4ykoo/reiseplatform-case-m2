package dbgorm

import (
	"time"

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

func (repo TravelRepositoryImpl) Create(travel *model.Travel, hotelid uint) (*model.Travel, error) {
	entity := entities.ToTravelEntity(travel)
	entity.HotelRef = hotelid
	res := repo.db.Connection.Create(entity)
	return entities.ToTravelModel(entity), res.Error
}

func (repo TravelRepositoryImpl) Update(travel *model.Travel) (*model.Travel, error) {

	// search of existing entry
	var entity entities.TravelEntity
	result := repo.db.Connection.Model(&entities.TravelEntity{}).Preload("Tags").First(&entity, travel.Id)

	// Cancel update if entry not exist
	if result.Error != nil {
		return travel, result.Error
	}

	update := entities.ToTravelEntity(travel)
	entity.From = update.From
	entity.To = update.To
	entity.Price = update.Price
	entity.Vendor = update.Vendor
	entity.Description = update.Description
	entity.Tags = update.Tags

	res := repo.db.Connection.Save(&entity)

	return entities.ToTravelModel(&entity), res.Error
}

func (repo TravelRepositoryImpl) Delete(id uint) error {
	var entity entities.TravelEntity
	result := repo.db.Connection.Model(&entities.TravelEntity{}).Preload("Tags").First(&entity, id)

	// Cancel deletion if entry not exist
	if result.Error != nil {
		return result.Error
	}

	//repo.db.Connection.Model(&entity).Association("Tags").Delete(&entity.Tags)

	result = repo.db.Connection.Delete(&entity)
	return result.Error
}

func (repo TravelRepositoryImpl) ListAll() ([]*model.Travel, error) {
	var entites []entities.TravelEntity
	res := repo.db.Connection.Model(&entities.TravelEntity{}).Preload("Tags").Find(&entites)

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
	err := repo.db.Connection.Model(&entities.TravelEntity{}).Preload("Tags").First(&entity, id).Error
	return entities.ToTravelModel(&entity), err
}

func (repo TravelRepositoryImpl) Count() (int64, error) {
	var count int64
	res := repo.db.Connection.Model(&entities.TravelEntity{}).Count(&count)
	return count, res.Error

}

func (repo TravelRepositoryImpl) FindByName(name string) ([]*model.Travel, error) {
	if len(name) == 0 {
		return []*model.Travel{}, nil
	}
	var entity []entities.TravelEntity

	result := repo.db.Connection.Model(&entities.TravelEntity{}).Preload("Hotels").Where("Hotels.Name LIKE ?", "%"+name+"%").Find(&entity)
	//result := repo.db.Connection.Preload("Hotels", "name IN ?", []string{"groupname"}).Find(&users)
	if result.RowsAffected == 0 || result.Error != nil {
		return []*model.Travel{}, result.Error
	}

	travels := make([]*model.Travel, len(entity))

	for i, travel := range entity {
		travels[i] = entities.ToTravelModel(&travel)
	}

	return travels, result.Error
}

func (repo TravelRepositoryImpl) FindBetween(from time.Time, to time.Time) ([]*model.Travel, error) {

	if from.Unix() > to.Unix() {
		return []*model.Travel{}, nil
	}

	var travel_entity []entities.TravelEntity

	result := repo.db.Connection.Model(&entities.TravelEntity{}).Preload("Tags").Preload("Hotels").
		Preload("Hotels.Pictures").Where("travel_entities.from >= ? AND travel_entities.to <= ?", from, to).Find(&travel_entity)

	if result.RowsAffected == 0 || result.Error != nil {
		return nil, result.Error
	}

	travels := make([]*model.Travel, len(travel_entity))

	for i, travel := range travel_entity {
		travels[i] = entities.ToTravelModel(&travel)
	}
	return travels, result.Error
}
