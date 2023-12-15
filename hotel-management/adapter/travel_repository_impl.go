package adapter

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mig3177/hotelmanagement/adapter/entities"
	"github.com/mig3177/hotelmanagement/domain/model"
)

type TravelRepositoryImpl struct {
	db pgRepository
}

func NewTravelRepository() TravelRepositoryImpl {
	con := initPGConnection(10, 100)
	con.createTable(&entities.HotelEntity{}, &entities.PictureEnitiy{}, &entities.TagEntity{}, &entities.TravelEntity{})
	return TravelRepositoryImpl{db: con}
}

func (repo TravelRepositoryImpl) Save(travel model.Travel) {
	hotel := travel.Hotel
	pictures := make([]entities.PictureEnitiy, len(hotel.Pictures))
	for i, pic := range hotel.Pictures {
		pictures[i] = entities.PictureEnitiy{ID: pic.ID, Payload: pic.Payload, HotelRef: hotel.ID, Description: pic.Description}
	}
	tags := make([]entities.TagEntity, len(travel.Tags))
	for i, tag := range travel.Tags {
		tags[i] = entities.TagEntity{Typ: tag.Typ, Name: tag.Name}
	}
	hotel_entity := entities.HotelEntity{ID: hotel.ID, Name: hotel.Name,
		Street: hotel.Address.Street, State: hotel.Address.State, Land: hotel.Address.Land,
		VendorRef: hotel.Vendor.ID, Description: hotel.Description, Pictures: pictures}
	travel_entity := entities.TravelEntity{ID: travel.ID, Hotel: hotel_entity, Vendor: travel.Vendor.ID, From: travel.From, To: travel.To,
		Price: travel.Price, Description: travel.Description, CreatedAt: travel.CreatedAt, UpdatedAt: travel.UpdatedAt, Tags: tags}
	repo.db.Connection.Create(&travel_entity)

}
func (repo TravelRepositoryImpl) Update(travel model.Travel) {
	hotel := travel.Hotel
	pictures := make([]entities.PictureEnitiy, len(hotel.Pictures))
	for i, pic := range hotel.Pictures {
		pictures[i] = entities.PictureEnitiy{ID: pic.ID, Payload: pic.Payload, HotelRef: hotel.ID, Description: pic.Description}
	}
	tags := make([]entities.TagEntity, len(travel.Tags))
	for i, tag := range travel.Tags {
		tags[i] = entities.TagEntity{Typ: tag.Typ, Name: tag.Name}
	}
	hotel_entity := entities.HotelEntity{ID: hotel.ID, Name: hotel.Name,
		Street: hotel.Address.Street, State: hotel.Address.State, Land: hotel.Address.Land,
		VendorRef: hotel.Vendor.ID, Description: hotel.Description, Pictures: pictures}
	travel_entity := entities.TravelEntity{ID: travel.ID, Hotel: hotel_entity, Vendor: travel.Vendor.ID, From: travel.From, To: travel.To,
		Price: travel.Price, Description: travel.Description, CreatedAt: travel.CreatedAt, UpdatedAt: travel.UpdatedAt, Tags: tags}
	repo.db.Connection.Save(&travel_entity)

}
func (repo TravelRepositoryImpl) Delete(travel model.Travel) bool {
	hotel := travel.Hotel
	pictures := make([]entities.PictureEnitiy, len(hotel.Pictures))
	for i, pic := range hotel.Pictures {
		pictures[i] = entities.PictureEnitiy{ID: pic.ID, Payload: pic.Payload, HotelRef: hotel.ID, Description: pic.Description}
	}
	tags := make([]entities.TagEntity, len(travel.Tags))
	for i, tag := range travel.Tags {
		tags[i] = entities.TagEntity{Typ: tag.Typ, Name: tag.Name}
	}
	hotel_entity := entities.HotelEntity{ID: hotel.ID, Name: hotel.Name,
		Street: hotel.Address.Street, State: hotel.Address.State, Land: hotel.Address.Land,
		VendorRef: hotel.Vendor.ID, Description: hotel.Description, Pictures: pictures}
	travel_entity := entities.TravelEntity{ID: travel.ID, Hotel: hotel_entity, Vendor: travel.Vendor.ID, From: travel.From, To: travel.To,
		Price: travel.Price, Description: travel.Description, CreatedAt: travel.CreatedAt, UpdatedAt: travel.UpdatedAt, Tags: tags}
	repo.db.Connection.Model(&travel_entity).Association("Tags").Delete(travel_entity.Tags)
	repo.db.Connection.Delete(&travel_entity)
	return false
}
func (repo TravelRepositoryImpl) FindByID(id uuid.UUID) model.Travel {
	var travel_entity []entities.TravelEntity
	_ = repo.db.Connection.Model(&entities.TravelEntity{}).Preload("Tags").Preload("Hotel").Preload("Hotel.Pictures").First(&travel_entity, id).Error

	fmt.Println(travel_entity)
	return model.Travel{}
}
func (repo TravelRepositoryImpl) Count() int { return 0 }

/*
func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Model(&User{}).Preload("Languages").Find(&users).Error
	return users, err
  }*/
