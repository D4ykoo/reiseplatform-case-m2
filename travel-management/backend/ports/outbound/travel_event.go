package outbound

import (
	"github.com/mig3177/travelmanagement/domain/model"
)

type TravelEvents interface {
	HotelAdded(*model.Hotel, error)
	HotelRemoved(uint, string, error)
	HotelUpdated(*model.Hotel, error)
	TravelAdded(*model.Travel, error)
	TravelRemoved(uint, string, error)
	TravelUpdated(*model.Travel, error)
}
