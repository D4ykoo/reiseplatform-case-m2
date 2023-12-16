package ports

import (
	"github.com/google/uuid"
	"github.com/mig3177/travelmanagement/domain/model"
)

type TravelRepository interface {
	Save(model.Travel) error
	Update(model.Travel) error
	Delete(model.Travel) error
	ListAll() ([]*model.Hotel, error)
	FindByID(uuid.UUID) (model.Travel, error)
	Count() (int, error)
}
