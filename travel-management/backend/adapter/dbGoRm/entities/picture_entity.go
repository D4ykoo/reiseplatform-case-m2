package entities

import (
	"encoding/json"

	"github.com/mig3177/travelmanagement/domain/model"
	"gorm.io/gorm"
)

type PictureEntity struct {
	gorm.Model
	Payload     string `json:"payload" gorm:"not null"`
	HotelRef    uint   `gorm:"type:uint;foreignKey:ID"`
	Description string `json:"description"`
}

func (p *PictureEntity) String() string {
	res, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(res)
}

func ToPicEntity(pic *model.Picture) *PictureEntity {
	return &PictureEntity{Payload: pic.Payload, Description: pic.Description, Model: gorm.Model{ID: pic.Id}}
}

func ToPicModel(pic *PictureEntity) *model.Picture {
	return &model.Picture{Id: pic.ID, Payload: pic.Payload, Description: pic.Description}
}
