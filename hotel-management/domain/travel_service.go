package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/mig3177/hotelmanagement/domain/model"
)

type TravelService interface {
	CreateOffer(HotelRef model.Hotel, Vendor model.Vendor, From time.Time, To time.Time, price float32, description string) model.Travel
	FindOfferByID(ID uuid.UUID) model.Travel
	FindOfferByTag(Tag model.Tag) []model.Travel
	FindOfferByDate(From time.Time, To time.Time) []model.Travel
	FindOfferByVendor(Vendor model.Vendor) []model.Travel
	UpdateOffer(id uuid.UUID) model.Travel
	DeleteOffer(id uuid.UUID) bool
}
