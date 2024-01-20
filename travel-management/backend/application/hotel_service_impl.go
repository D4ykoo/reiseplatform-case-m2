package application

import (
	"errors"
	"strings"
	"time"

	"github.com/mig3177/travelmanagement/domain/model"
	"github.com/mig3177/travelmanagement/ports/outbound"
)

type HotelServiceImpl struct {
	hotels outbound.HotelRepository
	events outbound.TravelEvents
}

func NewHotelService(hrepo outbound.HotelRepository, events outbound.TravelEvents) HotelServiceImpl {
	return HotelServiceImpl{hotels: hrepo, events: events}
}

func (service HotelServiceImpl) NewHotel(name string, address model.Address, vendor model.Vendor, description string, pics []*model.Picture, tags []*model.Tag) (*model.Hotel, error) {

	hotel := &model.Hotel{Address: address, Name: name, Description: description, Vendor: vendor, Pictures: pics, Tags: tags}

	hotelres, err := service.hotels.Create(hotel)
	service.events.HotelAdded(hotelres)
	return hotelres, err
}

func (service HotelServiceImpl) FindHotelByTravel(hotelId, travelId uint) (*model.Hotel, error) {
	hotel, err := service.hotels.FindByID(hotelId)

	if err != nil {
		return &model.Hotel{}, err
	}
	var travels []*model.Travel

	for _, t := range hotel.Travels {
		if t.Id == travelId {
			travels = append(travels, t)
			break
		}
	}

	if len(travels) != 1 {
		return &model.Hotel{}, errors.New("Travel id not found")
	}

	hotel.Travels = travels
	return hotel, err
}

func (service HotelServiceImpl) FindHotelTravel(name string, land string, from *time.Time, to *time.Time, tags []uint) ([]*model.Hotel, error) {
	hotels, error := service.hotels.ListAll()

	var result []*model.Hotel
	for _, hotel := range hotels {
		if len(name) != 0 && !strings.Contains(strings.ToLower(hotel.Name), strings.ToLower(name)) {
			continue
		}

		if len(land) != 0 && !strings.Contains(strings.ToLower(hotel.Address.Land), strings.ToLower(land)) {
			continue
		}
		if len(tags) != 0 && !containsTag(tags, hotel.Tags) {
			continue
		}

		var resultT = make([]*model.Travel, 0)
		if from != nil && to != nil {
			for j := 0; j < len(hotel.Travels); j++ {
				if hotel.Travels[j].From.Unix() >= from.Unix() && hotel.Travels[j].To.Unix() <= to.Unix() {
					resultT = append(resultT, hotel.Travels[j])
				}
			}
			hotel.Travels = resultT
		}
		result = append(result, hotel)
	}
	return result, error
}

func (service HotelServiceImpl) GetHotel(id uint) (*model.Hotel, error) {
	return service.hotels.FindByID(id)
}

func containsTag(s []uint, e []*model.Tag) bool {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(e); j++ {
			if s[i] == e[j].Id {
				return true
			}
		}
	}
	return false
}

func (service HotelServiceImpl) UpdateHotel(hotelUpdate *model.Hotel) (*model.Hotel, error) {

	hotelres, err := service.hotels.Update(hotelUpdate)
	service.events.HotelUpdated(hotelres)
	return hotelres, err
}

func (service HotelServiceImpl) RemoveHotel(id uint) error {
	err := service.hotels.Delete(id)
	service.events.HotelRemoved(&model.Hotel{Id: id})
	return err
}

func (service HotelServiceImpl) ListHotelTravel() ([]*model.Hotel, error) {
	return service.hotels.ListAll()
}
