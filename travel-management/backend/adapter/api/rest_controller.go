package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mig3177/travelmanagement/adapter/api/dto"
	"github.com/mig3177/travelmanagement/domain"
	"github.com/mig3177/travelmanagement/domain/model"
)

type RestController struct {
	service domain.TravelService
}

func New(service domain.TravelService) RestController {
	return RestController{service: service}
}

func (ctr RestController) CreateHotelRequest(c *gin.Context) {

	var hotel dto.CreateHotelRequest

	if err := c.ShouldBindJSON(&hotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pics := dto.ToPictureModel(hotel.Pictures)

	hotelRes, err := ctr.service.NewHotel(hotel.HotelName, model.Address{Street: hotel.Street, State: hotel.State, Land: hotel.HotelName}, model.Vendor{Id: hotel.VendorId, Username: hotel.VendorName}, hotel.Description, pics)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.ToHotelResoponse(hotelRes))

}

func (ctr RestController) FindHotels(c *gin.Context) {

	query := c.Request.URL.Query()

	var hotels []*model.Hotel

	name := query.Get("name")
	land := query.Get("land")
	from := query.Get("from")
	to := query.Get("to")
	tags := query.Get("tags")

	if len(query) == 0 {
		var err error
		hotels, err = ctr.service.ListHotelTravel()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

	}
	// TODO
	hotels, err := ctr.service.FindHotelTravel(name, land, "", "", []uint{})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hotelResponse := make([]*dto.HotelResponse, len(hotels))

	for index, hotel := range hotels {
		hotelResponse[index] = dto.ToHotelResoponse(hotel)
	}

	c.JSON(http.StatusOK, hotelResponse)
}

func (ctr RestController) GetHotelById(c *gin.Context) {

	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hotel, err := ctr.service.GetHotel(uint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hotelResponse := dto.ToHotelResoponse(hotel)

	c.JSON(http.StatusOK, hotelResponse)
}

func (ctr RestController) DeleteHotelRequest(c *gin.Context) {
	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = ctr.service.RemoveHotel(uint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})

}

func (ctr RestController) UpdateHotel(c *gin.Context) {

	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var hotel dto.UpdateHotelRequest

	if err := c.ShouldBindJSON(&hotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if uint(id) != hotel.Id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id mismatch"})
	}

	updatetedHotel, err := ctr.service.UpdateHotel(dto.ToHotelModel(&hotel))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hotelResponse := dto.ToHotelResoponse(updatetedHotel)
	c.JSON(http.StatusOK, &hotelResponse)

}
