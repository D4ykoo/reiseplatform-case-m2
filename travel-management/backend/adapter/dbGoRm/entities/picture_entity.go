package entities

import (
	"github.com/google/uuid"
)

type PictureEntity struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	Payload     string    `json:"payload" gorm:"not null"`
	HotelRef    uuid.UUID `gorm:"type:uuid;foreignKey:ID"`
	Description string
}
