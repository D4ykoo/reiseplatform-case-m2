package outbound

import (
	"time"

	"github.com/mig3177/travelmanagement/domain/model"
)

type TravelRepository interface {
	Create(*model.Travel) (*model.Travel, error)
	Update(*model.Travel) (*model.Travel, error)
	Delete(uint) error
	ListAll() ([]*model.Travel, error)
	FindByID(uint) (*model.Travel, error)
	FindByName(string) ([]*model.Travel, error)
	//FindByTag(string) ([]*model.Travel, error)
	FindBetween(time.Time, time.Time) ([]*model.Travel, error)
	Count() (int64, error)
}
