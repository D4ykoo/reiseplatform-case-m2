package entities

import "github.com/google/uuid"

type PictureEnitiy struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	Payload     string    `json:"payload" gorm:"not null"`
	HotelRef    uuid.UUID
	Description string
}
