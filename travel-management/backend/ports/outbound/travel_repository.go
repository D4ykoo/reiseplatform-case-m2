package outbound

import (
	"time"

	"github.com/google/uuid"
	"github.com/mig3177/travelmanagement/domain/model"
)

type TravelRepository interface {
	Create(*model.Travel) error
	Update(*model.Travel) error
	Delete(*model.Travel) error
	ListAll() ([]*model.Travel, error)
	FindByID(uuid.UUID) (*model.Travel, error)
	FindByName(string) ([]*model.Travel, error)
	//FindByTag(string) ([]*model.Travel, error)
	FindBetween(time.Time, time.Time) ([]*model.Travel, error)
	Count() (int, error)
}
