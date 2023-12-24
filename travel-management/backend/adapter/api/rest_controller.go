package api

import (
	"net/http"
	"strconv"
	"strings"
	"time"

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
	fromStr := query.Get("from")
	toStr := query.Get("to")
	tagsStr := query.Get("tags")
	tagsArray := strings.Split(tagsStr, ",")

	// Retun all result without querry parameter
	if len(query) == 0 {
		var err error
		hotels, err = ctr.service.ListHotelTravel()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			var hotelRes []*dto.HotelResponse
			for _, hotel := range hotels {
				hotelRes = append(hotelRes, dto.ToHotelResoponse(hotel))
			}
			c.JSON(http.StatusOK, &hotelRes)
			return
		}
	}

	var from *time.Time
	var to *time.Time

	if len(fromStr) > 0 && len(toStr) > 0 {
		fromTmp, errFrom := time.Parse("2006-01-02", fromStr)

		if errFrom != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errFrom.Error()})
			return
		}

		toTmp, errTo := time.Parse("2006-01-02", toStr)

		if errTo != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errTo.Error()})
			return
		}

		from = &fromTmp
		to = &toTmp

	}

	var tags []uint

	for _, tag := range tagsArray {
		i, err2 := strconv.Atoi(tag)
		if err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
			return
		}
		tags = append(tags, uint(i))
	}

	// TODO
	hotels, err3 := ctr.service.FindHotelTravel(name, land, from, to, tags)

	if err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err3.Error()})
		return
	}

	var hotelResponse []*dto.HotelResponse

	for _, hotel := range hotels {
		hotelResponse = append(hotelResponse, dto.ToHotelResoponse(hotel))
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

// travel

func (ctr RestController) CreateTravelRequest(c *gin.Context) {

	stringId := c.Param("id")
	hotelId, err := strconv.Atoi(stringId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var travel dto.CreateTravelRequest

	if err := c.ShouldBindJSON(&travel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	from, err := time.Parse("2006-01-02", travel.From)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	to, err := time.Parse("2006-01-02", travel.To)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	travelRes, err := ctr.service.NewTravel(uint(hotelId), uint(travel.VendorId), from, to, travel.Price, travel.Description)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.ToTravelResoponse(travelRes))

}

func (ctr RestController) GetTravelById(c *gin.Context) {

	hotelStrId := c.Param("id")
	hotelId, err := strconv.Atoi(hotelStrId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	travelStrId := c.Param("tid")
	travelId, err := strconv.Atoi(travelStrId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hotel, err := ctr.service.FindHotelByTravel(uint(hotelId), uint(travelId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	travelResponse := dto.ToHotelResoponse(hotel)

	c.JSON(http.StatusOK, travelResponse)
}

// tag

func (ctr RestController) CreateTagRequest(c *gin.Context) {

	var tag dto.TagRequest

	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tagRes, err := ctr.service.NewTag(tag.Name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.ToTagResponse(tagRes))

}

func (ctr RestController) GetTagById(c *gin.Context) {

	tagStrId := c.Param("id")
	tagId, err := strconv.Atoi(tagStrId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tag, err := ctr.service.GetTag(uint(tagId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToTagResponse(tag))
}

func (ctr RestController) DeleteTagRequest(c *gin.Context) {
	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = ctr.service.RemoveTag(uint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})

}

func (ctr RestController) ListAllTags(c *gin.Context) {

	tags, err := ctr.service.ListTags()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var tagsRes []*dto.TagResponse
	for _, tag := range tags {
		tagsRes = append(tagsRes, dto.ToTagResponse(tag))
	}

	c.JSON(http.StatusOK, tagsRes)

}

func (ctr RestController) UpdateTag(c *gin.Context) {

	tagStrId := c.Param("id")
	tagId, err := strconv.Atoi(tagStrId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var tag dto.TagRequest

	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if uint(tagId) != tag.Id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id mismatch"})
	}

	tagUpdate, err := ctr.service.UpdateTags(&model.Tag{Id: tag.Id, Name: tag.Name})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToTagResponse(tagUpdate))
}