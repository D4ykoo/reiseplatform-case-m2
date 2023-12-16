package ports

import (
	"github.com/google/uuid"
	"github.com/mig3177/hotelmanagement/domain/model"
)

type HotelRepository interface {
	Save(model.Hotel) error
	Update(model.Hotel) error
	Delete(model.Hotel) error
	ListAll() ([]*model.Hotel, error)
	FindByID(uuid.UUID) (model.Hotel, error)
	Count() (int, error)
}
