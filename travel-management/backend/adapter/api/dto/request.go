package dto

import "github.com/mig3177/travelmanagement/domain/model"

type CreateHotelRequest struct {
	HotelName   string           `json:"hotelname"`
	Street      string           `json:"street"`
	State       string           `json:"state"`
	Land        string           `json:"land"`
	Vendor      uint             `json:"vendor"`
	Description string           `json:"description"`
	Pictures    []PictureRequest `json:"pictures"`
}

type PictureRequest struct {
	Description string `json:"description"`
	Payload     string `json:"payload"`
}

type UpdateHotelRequest struct {
}

type CreateTravelRequest struct {
}

type DeleteTravelRequest struct {
}

func toPictures(hotel *model.Hotel) []dto.PictureResponse {
	pictures := make([]dto.PictureResponse, len(hotel.Pictures))
	for i, pic := range hotel.Pictures {
		pictures[i] = dto.PictureResponse{Id: pic.ID.String(), Payload: pic.Payload, Description: pic.Description}
	}
	return pictures
}
