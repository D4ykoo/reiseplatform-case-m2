package kafka

import (
	"log"
	"os"

	"github.com/mig3177/travelmanagement/adapter/kafka/dto"
	"github.com/mig3177/travelmanagement/domain/model"
)

type TravelEventImpl struct {
	msgService MessageService
}

func New() TravelEventImpl {

	service, err := NewMsgService(os.Getenv("BROKERS"), os.Getenv("TOPIC"))
	log.Default().Print(err)
	return TravelEventImpl{
		msgService: service,
	}
}

func (service TravelEventImpl) HotelAdded(obj model.Hotel) {
	event := dto.HotelEvent{
		ObjId: obj.Id,
		Type_: "Create",
		Event: "Create new hotel (" + obj.Name + ")",
	}
	service.msgService.PublishAsJSON(event)
}

func (service TravelEventImpl) HotelRemoved(obj model.Hotel) {
	event := dto.HotelEvent{
		ObjId: obj.Id,
		Type_: "Remove",
		Event: "Removed hotel (" + obj.Name + ")",
	}
	service.msgService.PublishAsJSON(event)
}
func (service TravelEventImpl) HotelVisited(obj model.Hotel) {
	event := dto.HotelEvent{
		ObjId: obj.Id,
		Type_: "Visit",
		Event: "Visited offers from hotel (" + obj.Name + ")",
	}
	service.msgService.PublishAsJSON(event)
}
func (service TravelEventImpl) HotelUpdated(obj model.Hotel) {
	event := dto.HotelEvent{
		ObjId: obj.Id,
		Type_: "Update",
		Event: "Update offers from hotel (" + obj.Name + ")",
	}
	service.msgService.PublishAsJSON(event)
}
