package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/mig3177/hotelmanagement/domain/model"
)

type OfferService interface {
	createOffer(HotelRef model.Hotel, Vendor model.Vendor, From time.Time, To time.Time, price float32, description string) model.Offer
	findOfferByID(ID uuid.UUID) model.Offer
	findOfferByTag(Tag model.Tag) []model.Offer
	findOfferByDate(From time.Time, To time.Time) []model.Offer
	findOfferByVendor(Vendor model.Vendor) []model.Offer
	updateOffer(id uuid.UUID) model.Offer
	deleteOffer(id uuid.UUID) bool
}
