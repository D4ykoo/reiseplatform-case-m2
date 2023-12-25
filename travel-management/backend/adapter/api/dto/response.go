package dto

import (
	"time"

	"github.com/mig3177/travelmanagement/domain/model"
)

type TravelResponse struct {
	Id          uint    `json:"id"`
	VendorID    uint    `json:"vendorid"`
	VendorName  string  `json:"vendorname"`
	From        string  `json:"from"`
	To          string  `json:"to"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"createdat"`
	UpdatedAt   string  `json:"updatedat"`
}

type HotelResponse struct {
	Id          uint               `json:"id"`
	HotelName   string             `json:"hotelname"`
	Street      string             `json:"street"`
	State       string             `json:"state"`
	Land        string             `json:"land"`
	VendorID    string             `json:"vendorid"`
	VendorName  string             `json:"vendorname"`
	Description string             `json:"description"`
	Pictures    []*PictureResponse `json:"pictures"`
	Tags        []*TagResponse     `json:"tags"`
	Travels     []*TravelResponse  `json:"travels"`
}

type PictureResponse struct {
	Id          uint   `json:"id"`
	Description string `json:"description"`
	Payload     string `json:"payload"`
}

type TagResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func ToPictureResopnse(pics []*model.Picture) []*PictureResponse {
	pictures := make([]*PictureResponse, len(pics))
	for i, pic := range pics {
		pictures[i] = &PictureResponse{Id: pic.Id, Payload: pic.Payload, Description: pic.Description}
	}
	return pictures
}

func ToHotelResoponse(hotel *model.Hotel) *HotelResponse {
	pics := ToPictureResopnse(hotel.Pictures)
	travels := make([]*TravelResponse, len(hotel.Travels))
	for i, travel := range hotel.Travels {
		travels[i] = ToTravelResoponse(travel)
	}

	tags := make([]*TagResponse, len(hotel.Tags))
	for i, tag := range hotel.Tags {
		tags[i] = ToTagResponse(tag)
	}

	return &HotelResponse{Id: hotel.Id, HotelName: hotel.Name, Street: hotel.Address.Street, State: hotel.Address.State, Land: hotel.Address.Land,
		VendorID: hotel.Vendor.Username, VendorName: hotel.Vendor.Username, Description: hotel.Description, Pictures: pics, Travels: travels, Tags: tags}
}

func ToTravelResoponse(travel *model.Travel) *TravelResponse {
	return &TravelResponse{Id: travel.Id, VendorID: travel.Vendor.Id, VendorName: travel.Vendor.Username,
		From: travel.From.Format(time.RFC3339), To: travel.To.Format(time.RFC3339), Price: travel.Price,
		Description: travel.Description, CreatedAt: travel.CreatedAt.Format(time.RFC3339), UpdatedAt: travel.UpdatedAt.Format(time.RFC3339)}
}

func ToTagResponse(tag *model.Tag) *TagResponse {
	return &TagResponse{Id: tag.Id, Name: tag.Name}
}
