package dbgorm

import (
	"time"

	"github.com/google/uuid"
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

// Function for converting from model to entity
func ToTravelEntity(travel *model.Travel) *entities.TravelEntity {

	hotel_entities := make([]*entities.HotelEntity, len(travel.Hotel))
	for i, hotel := range travel.Hotel {
		hotel_entities[i] = ToHotelEntity(hotel)
	}

	tags := make([]*entities.TagEntity, len(travel.Tags))
	for i, tag := range travel.Tags {
		tags[i] = &entities.TagEntity{Typ: tag.Typ, Name: tag.Name}
	}

	return &entities.TravelEntity{ID: travel.ID, Hotels: hotel_entities, Vendor: travel.Vendor.ID, From: travel.From, To: travel.To,
		Price: travel.Price, Description: travel.Description, CreatedAt: travel.CreatedAt, UpdatedAt: travel.UpdatedAt, Tags: tags}
}

// Function for converting from entity to model
func ToTravelModel(entity *entities.TravelEntity) *model.Travel {
	hotels := make([]*model.Hotel, len(entity.Hotels))
	for i, hotel := range entity.Hotels {
		hotels[i] = ToHotelModel(hotel)
	}

	tags := make([]*model.Tag, len(entity.Tags))
	for i, tag := range entity.Tags {
		tags[i] = &model.Tag{Typ: tag.Typ, Name: tag.Name}
	}

	return &model.Travel{ID: entity.ID, Hotel: hotels, Vendor: model.Vendor{ID: entity.Vendor}, From: entity.From, To: entity.To,
		Price: entity.Price, Description: entity.Description, Tags: tags}
}

//
// CRUD OPERATIONS:
//

func (repo TravelRepositoryImpl) Create(travel *model.Travel) error {
	travel_entity := ToTravelEntity(travel)
	return repo.db.Connection.Create(travel_entity).Error
}

func (repo TravelRepositoryImpl) Update(travel *model.Travel) error {
	travel_entity := ToTravelEntity(travel)
	return repo.db.Connection.Save(travel_entity).Error
}

func (repo TravelRepositoryImpl) Delete(travel *model.Travel) error {
	travel_entity := ToTravelEntity(travel)
	repo.db.Connection.Model(travel_entity).Association("Tags").Delete(travel_entity.Tags)
	repo.db.Connection.Model(travel_entity).Association("Hotels").Delete(travel_entity.Hotels)
	err := repo.db.Connection.Delete(travel_entity).Error
	return err
}

func (repo TravelRepositoryImpl) ListAll() ([]*model.Travel, error) {
	var travel_entity []entities.TravelEntity
	err := repo.db.Connection.Model(&entities.TravelEntity{}).Preload("Tags").Preload("Hotels").
		Preload("Hotels.Pictures").Find(travel_entity).Error

	travels := make([]*model.Travel, len(travel_entity))

	if err == nil {
		for i, travel := range travel_entity {
			travels[i] = ToTravelModel(&travel)
		}
	}
	return travels, err
}

func (repo TravelRepositoryImpl) FindByID(id uuid.UUID) (*model.Travel, error) {
	travel_entity := &entities.TravelEntity{}
	err := repo.db.Connection.Model(&entities.TravelEntity{}).Preload("Tags").Preload("Hotels").
		Preload("Hotels.Pictures").First(travel_entity, id).Error
	return ToTravelModel(travel_entity), err
}

func (repo TravelRepositoryImpl) Count() (int, error) {
	hotels, err := repo.ListAll()
	if err != nil {
		return 0, err
	} else {
		return len(hotels), err
	}
}

func (repo TravelRepositoryImpl) FindByName(name string) ([]*model.Travel, error) {
	if len(name) == 0 {
		return nil, nil
	}
	var travel_entity []entities.TravelEntity

	result := repo.db.Connection.Where("name LIKE ?", "%"+name+"%").Find(&travel_entity)
	if result.RowsAffected == 0 || result.Error != nil {
		return []*model.Travel{}, result.Error
	}

	travels := make([]*model.Travel, len(travel_entity))

	for i, travel := range travel_entity {
		travels[i] = ToTravelModel(&travel)
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
		travels[i] = ToTravelModel(&travel)
	}
	return travels, nil
}
