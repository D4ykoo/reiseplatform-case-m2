package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/mig3177/travelmanagement/domain/model"
)

type TravelService interface {
	CreateTravel(HotelRef model.Hotel, Vendor model.Vendor, From time.Time, To time.Time, price float32, description string) model.Travel
	FindTravelByID(ID uuid.UUID) model.Travel
	FindTravelByTag(Tag model.Tag) []model.Travel
	FindTravelByDate(From time.Time, To time.Time) []model.Travel
	FindTravelByVendor(Vendor model.Vendor) []model.Travel
	UpdateOffer(id uuid.UUID) model.Travel
	DeleteOffer(id uuid.UUID) bool
}
