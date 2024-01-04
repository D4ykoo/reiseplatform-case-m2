package outbound

import (
	"github.com/mig3177/travelmanagement/domain/model"
)

type TravelEvents interface {
	HotelAdded(event *model.Hotel)
	HotelRemoved(event *model.Hotel)
	HotelVisited(event *model.Hotel)
	HotelUpdated(event *model.Hotel)
}
