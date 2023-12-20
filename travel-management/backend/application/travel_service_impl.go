package application

import (
	"github.com/google/uuid"
	"github.com/mig3177/travelmanagement/domain/model"
	"github.com/mig3177/travelmanagement/ports/outbound"
)

type TravelServiceImpl struct {
	repo outbound.TravelRepository
}

func NewTravelService(repo outbound.TravelRepository) TravelServiceImpl {
	return TravelServiceImpl{repo: repo}
}

func (service HotelServiceImpl) Create(*model.Travel) error {

}
func (service HotelServiceImpl) Update(*model.Travel) error                 {}
func (service HotelServiceImpl) Delete(*model.Travel) error                 {}
func (service HotelServiceImpl) ListAll() ([]*model.Travel, error)          {}
func (service HotelServiceImpl) FindByID(uuid.UUID) (*model.Travel, error)  {}
func (service HotelServiceImpl) FindByName(string) ([]*model.Travel, error) {}
func (service HotelServiceImpl) FindByTag(string) ([]*model.Travel, error)  {}
func (service HotelServiceImpl) Count() (int, error)                        {}
