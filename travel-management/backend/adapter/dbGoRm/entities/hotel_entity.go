package entities

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type HotelEntity struct {
	ID          uuid.UUID        `json:"id" gorm:"type:uuid;primary_key"`
	Name        string           `json:"name" gorm:"not null"`
	Street      string           `json:"street"`
	State       string           `json:"state"`
	Land        string           `json:"land"`
	VendorRef   uuid.UUID        `json:"vendorId" gorm:"type:uuid"`
	Description string           `json:"description"`
	CreatedAt   time.Time        `json:"created" gorm:"autoCreateTime"`
	UpdatedAt   time.Time        `json:"updated" gorm:"autoUpdateTime"`
	Pictures    []*PictureEntity `gorm:"foreignKey:HotelRef"`
	Travels     []*TravelEntity  `gorm:"many2many:travel_hotel;"`
}

func (b *HotelEntity) String() string {
	return fmt.Sprintf("%s", b.ID)
}
