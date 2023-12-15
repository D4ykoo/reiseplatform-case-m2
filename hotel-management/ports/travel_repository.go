package ports

import (
	"github.com/google/uuid"
	"github.com/mig3177/hotelmanagement/domain/model"
)

type TravelRepository interface {
	Save(model.Travel)
	Update(model.Travel)
	Delete(model.Travel) bool
	FindByID(uuid.UUID) model.Travel
	Count() int
}
