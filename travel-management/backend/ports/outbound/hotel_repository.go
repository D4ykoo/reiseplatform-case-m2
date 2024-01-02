package outbound

import (
	"github.com/mig3177/travelmanagement/domain/model"
)

type HotelRepository interface {
	Create(*model.Hotel) (*model.Hotel, error)
	Update(*model.Hotel) (*model.Hotel, error)
	Delete(uint) error
	ListAll() ([]*model.Hotel, error)
	FindByID(uint) (*model.Hotel, error)
	Count() (int64, error)
}
