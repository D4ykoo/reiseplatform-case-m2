package outbound

import (
	"github.com/mig3177/travelmanagement/domain/model"
)

type TravelRepository interface {
	Create(*model.Travel, uint) (*model.Travel, error)
	Update(*model.Travel, uint) (*model.Travel, error)
	Delete(uint) error
	ListAll(uint) ([]*model.Travel, error)
	FindByID(uint) (*model.Travel, error)
	Count() (int64, error)
}
