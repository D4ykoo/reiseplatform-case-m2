package kafka

import (
	"log"
	"os"
	"time"

	"github.com/mig3177/travelmanagement/adapter/kafka/dto"
	"github.com/mig3177/travelmanagement/domain/model"
)

type TravelEventImpl struct {
	MsgService MessageService
}

func New() TravelEventImpl {
	service, err := NewMsgService(os.Getenv("BROKERS"), os.Getenv("TOPIC"))

	// Stop Service
	if err != nil {
		log.Panic(err.Error())
	}

	return TravelEventImpl{
		MsgService: service,
	}
}

func (service TravelEventImpl) HotelAdded(obj *model.Hotel) {
	event := dto.HotelEvent{
		Type: "Create",
		Log:  "Create new hotel (" + obj.Name + ")",
		Time: time.Now().UTC(),
	}
	service.MsgService.PublishAsJSON(event)
}

func (service TravelEventImpl) HotelRemoved(obj *model.Hotel) {
	event := dto.HotelEvent{
		Type: "Remove",
		Log:  "Removed hotel (" + string(rune(obj.Id)) + ")",
		Time: time.Now().UTC(),
	}
	service.MsgService.PublishAsJSON(event)
}
func (service TravelEventImpl) HotelVisited(obj *model.Hotel) {
	event := dto.HotelEvent{
		Type: "Visit",
		Log:  "Visited offers from hotel (" + obj.Name + ")",
		Time: time.Now().UTC(),
	}
	service.MsgService.PublishAsJSON(event)
}
func (service TravelEventImpl) HotelUpdated(obj *model.Hotel) {
	event := dto.HotelEvent{
		Type: "Update",
		Log:  "Update offers from hotel (" + obj.Name + ")",
		Time: time.Now().UTC(),
	}
	service.MsgService.PublishAsJSON(event)
}
