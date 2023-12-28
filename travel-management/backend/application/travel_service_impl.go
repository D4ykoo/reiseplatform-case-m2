package application

import (
	"errors"
	"strings"
	"time"

	"github.com/mig3177/travelmanagement/domain/model"
	"github.com/mig3177/travelmanagement/ports/outbound"
)

type TravelServiceImpl struct {
	hotels  outbound.HotelRepository
	travels outbound.TravelRepository
	tags    outbound.TagRepository
}

func New(hrepo outbound.HotelRepository, trepo outbound.TravelRepository, tagrepo outbound.TagRepository) TravelServiceImpl {
	return TravelServiceImpl{hotels: hrepo, travels: trepo, tags: tagrepo}
}

func (service TravelServiceImpl) NewHotel(name string, address model.Address, vendor model.Vendor, description string, pics []*model.Picture, tags []*model.Tag) (*model.Hotel, error) {
	// TODO check user is valid
	// TODO check hotel already exist
	hotel := &model.Hotel{Address: address, Name: name, Description: description, Vendor: vendor, Pictures: pics, Tags: tags}
	return service.hotels.Create(hotel)
}

func (service TravelServiceImpl) NewTravel(hotelRef uint, vendor uint, from time.Time, to time.Time, price float32, description string) (*model.Travel, error) {
	offer := &model.Travel{Vendor: model.Vendor{Id: vendor}, From: from, To: to, Price: price, Description: description}
	return service.travels.Create(offer, hotelRef)
}

func (service TravelServiceImpl) FindHotelByTravel(hotelId, travelId uint) (*model.Hotel, error) {
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

func (service TravelServiceImpl) FindHotelTravel(name string, land string, from *time.Time, to *time.Time, tags []uint) ([]*model.Hotel, error) {
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

func (service TravelServiceImpl) GetHotel(id uint) (*model.Hotel, error) {
	return service.hotels.FindByID(id)
}

func (service TravelServiceImpl) GetTravel(id uint) (*model.Travel, error) {
	return service.travels.FindByID(id)
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

func (service TravelServiceImpl) UpdateHotel(hotelUpdate *model.Hotel) (*model.Hotel, error) {
	return service.hotels.Update(hotelUpdate)
}

func (service TravelServiceImpl) UpdateTravel(offerUpdate *model.Travel, id uint, hotelRef uint) (*model.Travel, error) {
	return service.travels.Update(offerUpdate, id, hotelRef)

}

func (service TravelServiceImpl) RemoveHotel(id uint) error {
	return service.hotels.Delete(id)
}
func (service TravelServiceImpl) RemoveTravel(id uint) error {
	return service.travels.Delete(id)
}
func (service TravelServiceImpl) ListHotelTravel() ([]*model.Hotel, error) {
	return service.hotels.ListAll()
}

func (service TravelServiceImpl) GetTag(id uint) (*model.Tag, error) {
	return service.tags.FindByID(id)
}

func (service TravelServiceImpl) ListTags() ([]*model.Tag, error) {
	return service.tags.ListAll()
}

func (service TravelServiceImpl) NewTag(name string) (*model.Tag, error) {
	return service.tags.Create(&model.Tag{Name: name})
}

func (service TravelServiceImpl) RemoveTag(id uint) error {
	return service.tags.Delete(id)
}

func (service TravelServiceImpl) UpdateTags(tag *model.Tag) (*model.Tag, error) {
	return service.tags.Update(tag)
}
