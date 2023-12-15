package ports

import (
	"github.com/google/uuid"
	"github.com/mig3177/hotelmanagement/domain/model"
)

type HotelRepository interface {
	Save(model.Hotel)
	Update(model.Hotel)
	Delete(model.Hotel) bool
	FindByID(uuid.UUID) model.Hotel
	Count() int
}
