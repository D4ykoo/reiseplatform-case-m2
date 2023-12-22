package application

import (
	"strings"
	"time"

	"github.com/mig3177/travelmanagement/domain/model"
	"github.com/mig3177/travelmanagement/ports/outbound"
)

type TravelServiceImpl struct {
	hotels  outbound.HotelRepository
	travels outbound.TravelRepository
}

func New(hrepo outbound.HotelRepository, trepo outbound.TravelRepository) TravelServiceImpl {
	return TravelServiceImpl{hotels: hrepo, travels: trepo}
}

func (service TravelServiceImpl) NewHotel(name string, address model.Address, vendorId uint, description string, pics []*model.Picture) (*model.Hotel, error) {
	// TODO check user is valid
	// TODO check hotel already exist
	hotel := &model.Hotel{Address: address, Name: name, Description: description, Vendor: model.Vendor{Id: vendorId, Username: "Walter"}, Pictures: pics}
	return service.hotels.Create(hotel)
}

func (service TravelServiceImpl) NewTravel(hotelRef uint, vendor uint, from time.Time, to time.Time, price float32, description string) (*model.Travel, error) {
	offer := &model.Travel{Vendor: model.Vendor{Id: vendor}, From: from, To: to, Price: price, Description: description}
	return service.travels.Create(offer, hotelRef)
}

func (service TravelServiceImpl) FindHotel(name, land string, from, to time.Time, tags ...uint) ([]*model.Hotel, error) {
	hotels, error := service.hotels.ListAll()
	result := make([]*model.Hotel, len(hotels))
	for i, hotel := range hotels {
		if len(name) != 0 && !strings.Contains(strings.ToLower(name), strings.ToLower(hotel.Name)) {
			break
		}

		if len(land) != 0 && !strings.Contains(strings.ToLower(land), strings.ToLower(hotel.Address.Land)) {
			break
		}
		if len(tags) != 0 && !containsTag(tags, hotel.Tags) {
			break
		}
		result[i] = hotel
		/*
			resultT := make([]*model.Travel, len(hotel.Travels))
			for j := 0; i < len(hotel.Travels); i++ {
				if from.Unix() >= hotel.Travels[i].From.Unix() && to.Unix() <= hotel.Travels[i].From.Unix() {
					resultT[] = hotel.Travels[j]
				}
			}
		*/
	}
	return result, error
}

func (service TravelServiceImpl) GetHotel(id uint) (*model.Hotel, error) {
	return service.hotels.FindByID(id)
}

func (service TravelServiceImpl) GetTravel(id uint) (*model.Travel, error) {
	return service.travels.FindByID(id)
}

func containsTag(s []uint, e []*model.Tag) bool {
	for i := 0; i < len(s); i++ {
		for j := 0; i < len(e); i++ {
			if s[i] == e[j].Id {
				return true
			}
		}
	}
	return false
}

func (service TravelServiceImpl) UpdateHotel(hotelUpdate *model.Hotel) (*model.Hotel, error) {
	return service.hotels.Update(hotelUpdate)
}

func (service TravelServiceImpl) UpdateTravel(offerUpdate *model.Travel, hotelRef uint) (*model.Travel, error) {
	return service.travels.Update(offerUpdate, hotelRef)

}

func (service TravelServiceImpl) RemoveHotel(id uint) error {
	return service.hotels.Delete(id)
}
func (service TravelServiceImpl) RemoveTravel(id uint) error {
	return service.travels.Delete(id)
}
