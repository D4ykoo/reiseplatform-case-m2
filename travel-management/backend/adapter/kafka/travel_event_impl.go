package kafka

import (
	"fmt"
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
		log.Default().Printf(err.Error())
	}

	return TravelEventImpl{
		MsgService: service,
	}
}

func (service TravelEventImpl) HotelAdded(obj *model.Hotel, err error) {
	var msg string
	var typ string
	if err != nil {
		typ = "CreateError"
		msg = "Error creating a new hotel listing"
	} else {
		typ = "Create"
		msg = fmt.Sprintf("User '%s' creates a new hotel listing (ID: %d, Name: '%s')", obj.Vendor.Username, obj.Id, obj.Name)
	}
	event := dto.HotelTravelEvent{
		Type: typ,
		Log:  msg,
		Time: time.Now().UTC(),
	}
	service.MsgService.PublishAsJSON(event)
}

func (service TravelEventImpl) HotelRemoved(id uint, user string, err error) {
	var msg string
	var typ string
	if err != nil {
		typ = "RemoveError"
		msg = "Error removing a hotel listing"
	} else {
		typ = "Remove"
		msg = fmt.Sprintf("User '%s' removed a hotel listing (ID: %d)", user, id)
	}
	event := dto.HotelTravelEvent{
		Type: typ,
		Log:  msg,
		Time: time.Now().UTC(),
	}
	service.MsgService.PublishAsJSON(event)
}

func (service TravelEventImpl) HotelUpdated(obj *model.Hotel, err error) {
	var msg string
	var typ string
	if err != nil {
		typ = "UpdateError"
		msg = "Error updating a hotel listing"
	} else {
		typ = "Update"
		msg = fmt.Sprintf("User '%s' updated a hotel listing (ID: %d, Name: '%s')", obj.Vendor.Username, obj.Id, obj.Name)
	}
	event := dto.HotelTravelEvent{
		Type: typ,
		Log:  msg,
		Time: time.Now().UTC(),
	}
	service.MsgService.PublishAsJSON(event)
}

func (service TravelEventImpl) TravelAdded(obj *model.Travel, err error) {
	var msg string
	var typ string
	if err != nil {
		typ = "CreateError"
		msg = "Error creating a new travel offer"
	} else {
		typ = "Create"
		msg = fmt.Sprintf("User '%s' creates a new travel offer (ID: %d)", obj.Vendor.Username, obj.Id)
	}
	event := dto.HotelTravelEvent{
		Type: typ,
		Log:  msg,
		Time: time.Now().UTC(),
	}
	service.MsgService.PublishAsJSON(event)
}

func (service TravelEventImpl) TravelUpdated(obj *model.Travel, err error) {
	var msg string
	var typ string
	if err != nil {
		typ = "UpdateError"
		msg = "Error updating a new travel offer"
	} else {
		typ = "Update"
		msg = fmt.Sprintf("User '%s' updated a travel offer (ID: %d)", obj.Vendor.Username, obj.Id)
	}
	event := dto.HotelTravelEvent{
		Type: typ,
		Log:  msg,
		Time: time.Now().UTC(),
	}
	service.MsgService.PublishAsJSON(event)
}

func (service TravelEventImpl) TravelRemoved(id uint, user string, err error) {
	var msg string
	var typ string
	if err != nil {
		typ = "RemoveError"
		msg = "Error removing a new travel offer"
	} else {
		typ = "Remove"
		msg = fmt.Sprintf("User '%s' removed travel offer (ID: %d)", user, id)
	}
	event := dto.HotelTravelEvent{
		Type: typ,
		Log:  msg,
		Time: time.Now().UTC(),
	}
	service.MsgService.PublishAsJSON(event)
}
