package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mig3177/travelmanagement/adapter/api/dto"
	"github.com/mig3177/travelmanagement/domain"
	"github.com/mig3177/travelmanagement/domain/model"
)

type RestHotelController struct {
	service domain.HotelService
}

func (controller RestHotelController) CreateHotelRequest(c *gin.Context) {

	var hotel dto.CreateHotelRequest

	if err := c.ShouldBindJSON(&hotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pics, address := toAddressPictures(&hotel)
	controller.service.CreateHotel(hotel.HotelName, address, uuid.New(), hotel.Description, pics)
}

func (controller RestHotelController) GetHotelRequest(c *gin.Context) {

	stringId := c.Param("id")
	id, err := uuid.Parse(stringId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hotel, err := controller.service.FindHotelByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hotelResponse := dto.HotelRespnse{
		Id:          hotel.ID.String(),
		HotelName:   hotel.Name,
		Street:      hotel.Address.Street,
		State:       hotel.Address.State,
		Land:        hotel.Address.Land,
		VendorID:    hotel.Vendor.ID.String(),
		VendorName:  hotel.Vendor.Username,
		Description: hotel.Description,
		Pictures:    toPictures(hotel),
	}

	c.JSON(http.StatusOK, hotelResponse)

}

func toAddressPictures(hotel *dto.CreateHotelRequest) ([]*model.Picture, model.Address) {
	pictures := make([]*model.Picture, len(hotel.Pictures))
	for i, pic := range hotel.Pictures {
		pictures[i] = &model.Picture{ID: uuid.New(), Payload: pic.Payload, Description: pic.Description}
	}
	return pictures, model.Address{Street: hotel.Street, State: hotel.State, Land: hotel.Land}
}

func toPictures(hotel *model.Hotel) []dto.PictureResponse {
	pictures := make([]dto.PictureResponse, len(hotel.Pictures))
	for i, pic := range hotel.Pictures {
		pictures[i] = dto.PictureResponse{Id: pic.ID.String(), Payload: pic.Payload, Description: pic.Description}
	}
	return pictures
}
