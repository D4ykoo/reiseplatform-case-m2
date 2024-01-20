package application

import (
	"time"

	"github.com/mig3177/travelmanagement/domain/model"
	"github.com/mig3177/travelmanagement/ports/outbound"
)

type TravelServiceImpl struct {
	travels outbound.TravelRepository
	events  outbound.TravelEvents
}

func NewTravelService(trepo outbound.TravelRepository, events outbound.TravelEvents) TravelServiceImpl {
	return TravelServiceImpl{travels: trepo, events: events}
}

func (service TravelServiceImpl) NewTravel(hotelRef uint, vendor uint, from time.Time, to time.Time, price float32, description string) (*model.Travel, error) {
	offer := &model.Travel{Vendor: model.Vendor{Id: vendor}, From: from, To: to, Price: price, Description: description}
	return service.travels.Create(offer, hotelRef)
}

func (service TravelServiceImpl) GetTravel(id uint) (*model.Travel, error) {
	return service.travels.FindByID(id)
}

func (service TravelServiceImpl) UpdateTravel(offerUpdate *model.Travel, id uint, hotelRef uint) (*model.Travel, error) {
	return service.travels.Update(offerUpdate, id, hotelRef)
}

func (service TravelServiceImpl) RemoveTravel(id uint) error {
	return service.travels.Delete(id)
}
