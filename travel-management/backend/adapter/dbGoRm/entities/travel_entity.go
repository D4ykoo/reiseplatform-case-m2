package entities

import (
	"encoding/json"
	"time"

	"github.com/mig3177/travelmanagement/domain/model"
	"gorm.io/gorm"
)

type TravelEntity struct {
	gorm.Model
	HotelRef    uint      `gorm:"type:uint;foreignKey:ID"`
	Vendor      uint      `json:"vendor"`
	From        time.Time `json:"from"`
	To          time.Time `json:"to"`
	Price       float32   `json:"price"`
	Description string    `json:"description"`
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

	return &TravelEntity{Vendor: travel.Vendor.Id, From: travel.From, To: travel.To,
		Price: travel.Price, Description: travel.Description, Model: gorm.Model{ID: travel.Id}}
}

// Function for converting from entity to model
func ToTravelModel(entity *TravelEntity) *model.Travel {

	return &model.Travel{Id: entity.ID, Vendor: model.Vendor{Id: entity.Vendor}, From: entity.From, To: entity.To,
		Price: entity.Price, Description: entity.Description, CreatedAt: entity.CreatedAt, UpdatedAt: entity.UpdatedAt}
}
