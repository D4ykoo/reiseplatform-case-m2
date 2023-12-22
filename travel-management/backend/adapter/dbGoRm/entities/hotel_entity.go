package entities

import (
	"encoding/json"

	"github.com/mig3177/travelmanagement/domain/model"
	"gorm.io/gorm"
)

type HotelEntity struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Street      string `json:"street" gorm:"not null"`
	State       string `json:"state" gorm:"not null"`
	Land        string `json:"land" gorm:"not null"`
	VendorRef   uint   `json:"vendorid" gorm:"not null"`
	Description string
	Pictures    []*PictureEntity `json:"pictures" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:HotelRef"`
	Travels     []*TravelEntity  `json:"travels" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:HotelRef"`
	Tags        []*TagEntity     `json:"tags" gorm:"many2many:tag_hotel;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (h *HotelEntity) String() string {
	res, err := json.Marshal(h)
	if err != nil {
		return ""
	}
	return string(res)
}

func ToHotelEntity(hotel *model.Hotel) *HotelEntity {
	pictures := make([]*PictureEntity, len(hotel.Pictures))
	for i, pic := range hotel.Pictures {
		pictures[i] = ToPicEntity(pic)
	}

	travels := make([]*TravelEntity, len(hotel.Travels))
	for i, travel := range hotel.Travels {
		travels[i] = ToTravelEntity(travel)
	}

	tags := make([]*TagEntity, len(hotel.Tags))
	for i, tag := range hotel.Tags {
		tags[i] = ToTagEntity(tag)
	}

	return &HotelEntity{Name: hotel.Name,
		Street: hotel.Address.Street, State: hotel.Address.State, Land: hotel.Address.Land,
		VendorRef: hotel.Vendor.Id, Description: hotel.Description, Pictures: pictures, Model: gorm.Model{ID: hotel.Id}, Travels: travels, Tags: tags}
}

func ToHotelModel(entity *HotelEntity) *model.Hotel {
	pictures := make([]*model.Picture, len(entity.Pictures))
	for i, pic := range entity.Pictures {
		pictures[i] = ToPicModel(pic)
	}

	travels := make([]*model.Travel, len(entity.Travels))
	for i, travel := range entity.Travels {
		travels[i] = ToTravelModel(travel)
	}

	tags := make([]*model.Tag, len(entity.Tags))
	for i, tag := range entity.Tags {
		tags[i] = ToTagModel(tag)
	}

	return &model.Hotel{Id: entity.ID, Name: entity.Name,
		Address: model.Address{Street: entity.Street, State: entity.State, Land: entity.Land},
		Vendor:  model.Vendor{Id: entity.VendorRef}, Description: entity.Description, Pictures: pictures, Travels: travels, Tags: tags}
}
