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

func NewRestHotelController(service domain.HotelService) RestHotelController {
	return RestHotelController{service: service}
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

func (controller RestHotelController) GetHotelsByNameRequest(c *gin.Context) {

	query := c.Request.URL.Query()

	if len(query) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "filter is not set"})
	}

	name := query.Get("name")

	hotels, err := controller.service.FindHotelByName(name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hotelResponse := make([]dto.HotelRespnse, len(hotels))

	for index, hotel := range hotels {
		hotelResponse[index] = dto.HotelRespnse{
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
	}
	c.JSON(http.StatusOK, &hotelResponse)
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

	c.JSON(http.StatusOK, &hotelResponse)
}

func (controller RestHotelController) DeleteHotelRequest(c *gin.Context) {
	stringId := c.Param("id")
	id, err := uuid.Parse(stringId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = controller.service.DeleteHotel(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})

}

func (controller RestHotelController) UpdateHotelRequest(c *gin.Context) {

	stringId := c.Param("id")
	id, err := uuid.Parse(stringId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var hotel dto.CreateHotelRequest

	if err := c.ShouldBindJSON(&hotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	stringVendorId := c.Param(hotel.Vendor)
	vendorId, err := uuid.Parse(stringVendorId)

	if err := c.ShouldBindJSON(&hotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pics, address := toAddressPictures(&hotel)

	hotelUpdate := model.Hotel{
		ID:          id,
		Name:        hotel.HotelName,
		Address:     address,
		Description: hotel.Description,
		Vendor:      model.Vendor{ID: vendorId},
		Pictures:    pics,
	}

	updatetedHotel, err := controller.service.UpdateHotel(&hotelUpdate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hotelResponse := dto.HotelRespnse{
		Id:          updatetedHotel.ID.String(),
		HotelName:   updatetedHotel.Name,
		Street:      updatetedHotel.Address.Street,
		State:       updatetedHotel.Address.State,
		Land:        updatetedHotel.Address.Land,
		VendorID:    updatetedHotel.Vendor.ID.String(),
		VendorName:  updatetedHotel.Vendor.Username,
		Description: updatetedHotel.Description,

		Pictures: toPictures(updatetedHotel),
	}

	c.JSON(http.StatusOK, &hotelResponse)

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
