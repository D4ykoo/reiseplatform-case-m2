package entities

import (
	"encoding/json"
	"time"

	"github.com/mig3177/travelmanagement/domain/model"
	"gorm.io/gorm"
)

type TravelEntity struct {
	gorm.Model
	HotelRef    uint         `gorm:"type:uint;foreignKey:ID"`
	Vendor      uint         `json:"vendor"`
	From        time.Time    `json:"from"`
	To          time.Time    `json:"to"`
	Price       float32      `json:"price"`
	Description string       `json:"description"`
	Tags        []*TagEntity `json:"tags" gorm:"many2many:travel_tags;"`
}

func (t *TravelEntity) String() string {
	res, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	return string(res)
}

// Function for converting from model to entity
func ToTravelEntity(travel *model.Travel) *TravelEntity {

	tags := make([]*TagEntity, len(travel.Tags))
	for i, tag := range travel.Tags {
		tags[i] = ToTagEntity(tag)
	}

	return &TravelEntity{Vendor: travel.Vendor.Id, From: travel.From, To: travel.To,
		Price: travel.Price, Description: travel.Description, Tags: tags, Model: gorm.Model{ID: travel.Id}}
}

// Function for converting from entity to model
func ToTravelModel(entity *TravelEntity) *model.Travel {

	tags := make([]*model.Tag, len(entity.Tags))
	for i, tag := range entity.Tags {
		tags[i] = ToTagModel(tag)
	}

	return &model.Travel{Id: entity.ID, Vendor: model.Vendor{Id: entity.Vendor}, From: entity.From, To: entity.To,
		Price: entity.Price, Description: entity.Description, Tags: tags, CreatedAt: entity.CreatedAt, UpdatedAt: entity.UpdatedAt}
}
