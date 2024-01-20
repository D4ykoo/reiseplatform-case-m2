package domain

import (
	"time"

	"github.com/mig3177/travelmanagement/domain/model"
)

type TravelService interface {
	NewTravel(hotelRef uint, vendor uint, from time.Time, to time.Time, price float32, description string) (*model.Travel, error)
	GetTravel(uint) (*model.Travel, error)
	UpdateTravel(*model.Travel, uint, uint) (*model.Travel, error)
	RemoveTravel(uint) error
}
