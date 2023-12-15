package entities

import (
	"time"

	"github.com/google/uuid"
)

type TravelEntity struct {
	ID          uuid.UUID `json:"typ" gorm:"primary_key;"`
	HotelRef    uuid.UUID
	Hotel       HotelEntity `json:"adress" gorm:"foreignKey:HotelRef"`
	Vendor      uuid.UUID   `json:"vendor"`
	From        time.Time
	To          time.Time
	Price       float32
	Description string
	CreatedAt   time.Time   `json:"created" gorm:"autoCreateTime"`
	UpdatedAt   time.Time   `json:"updated" gorm:"autoUpdateTime"`
	Tags        []TagEntity `gorm:"many2many:travel_tags;"`
}
