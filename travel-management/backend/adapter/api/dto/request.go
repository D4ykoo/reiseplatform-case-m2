package dto

import (
	"github.com/mig3177/travelmanagement/domain/model"
)

type CreateHotelRequest struct {
	HotelName   string           `json:"hotelname"`
	Street      string           `json:"street"`
	State       string           `json:"state"`
	Land        string           `json:"land"`
	VendorId    uint             `json:"vendorid"`
	VendorName  string           `json:"vendorname"`
	Description string           `json:"description"`
	Pictures    []PictureRequest `json:"pictures"`
	Tags        []TagRequest     `json:"tagids"`
}

type PictureRequest struct {
	Id          uint   `json:"id"`
	Description string `json:"description"`
	Payload     string `json:"payload"`
}

type UpdateHotelRequest struct {
	Id          uint             `json:"id"`
	HotelName   string           `json:"hotelname"`
	Street      string           `json:"street"`
	State       string           `json:"state"`
	Land        string           `json:"land"`
	VendorId    uint             `json:"vendorid"`
	VendorName  string           `json:"vendorname"`
	Description string           `json:"description"`
	Pictures    []PictureRequest `json:"pictures"`
	Tags        []TagRequest     `json:"tagids"`
}

type TagRequest struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type CreateTravelRequest struct {
	VendorId    uint    `json:"vendorid"`
	VendorName  string  `json:"vendorname"`
	From        string  `json:"from"`
	To          string  `json:"to"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}

type DeleteTravelRequest struct {
}

func ToPictureModel(pics []PictureRequest) []*model.Picture {
	pictures := make([]*model.Picture, len(pics))
	for i, pic := range pics {
		pictures[i] = &model.Picture{Payload: pic.Payload, Description: pic.Description}
	}
	return pictures
}

func ToHotelModel(hotel *UpdateHotelRequest) *model.Hotel {
	pics := ToPictureModel(hotel.Pictures)
	return &model.Hotel{Id: hotel.Id, Name: hotel.HotelName, Address: model.Address{Street: hotel.Street, State: hotel.State, Land: hotel.Land},
		Description: hotel.Description, Vendor: model.Vendor{Id: hotel.VendorId, Username: hotel.VendorName}, Pictures: pics}
}

func ToTagsModel(tags []TagRequest) []*model.Tag {
	tagsRes := make([]*model.Tag, len(tags))
	for i, tag := range tags {
		tagsRes[i] = &model.Tag{Id: tag.Id, Name: tag.Name}
	}
	return tagsRes
}
