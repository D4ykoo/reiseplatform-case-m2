package api

import (
	"fmt"
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
	travelService domain.TravelService
	tagService    domain.TagService
	hotelService  domain.HotelService
}

func New(travelService domain.TravelService, tagService domain.TagService, hotelService domain.HotelService) RestController {

	return RestController{travelService: travelService, tagService: tagService, hotelService: hotelService}
}

func (ctr RestController) CheckLoginStatus(c *gin.Context) {

	claims, err := ValidateLoginStatus(c)

	if err != nil {
		return
	}

	id, username := getUserData(claims)

	c.JSON(http.StatusOK, dto.UserResponse{Id: id, Name: username})

}

func (ctr RestController) CreateHotelRequest(c *gin.Context) {
	// Function can only be executed with a valid login status
	if _, err := ValidateLoginStatus(c); err != nil {
		return
	}

	var hotel dto.CreateHotelRequest

	if err := c.ShouldBindJSON(&hotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pics := dto.ToPictureModel(hotel.Pictures)
	tags := dto.ToTagsModel(hotel.Tags)

	hotelRes, err := ctr.hotelService.NewHotel(hotel.HotelName, model.Address{Street: hotel.Street, State: hotel.State, Land: hotel.HotelName},
		model.Vendor{Id: hotel.VendorId, Username: hotel.VendorName}, hotel.Description, pics, tags)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.ToHotelResoponse(hotelRes))

}

func (ctr RestController) FindHotels(c *gin.Context) {

	name := c.Query("name")
	land := c.Query("land")
	fromStr := c.Query("from")
	toStr := c.Query("to")
	tagsStr := c.Query("tags")
	tagsArray := strings.Split(tagsStr, ",")
	var hotels []*model.Hotel

	// Retun all result without querry parameter

	if len(name) == 0 && len(land) == 0 && len(fromStr) == 0 && len(toStr) == 0 && len(tagsStr) == 0 {
		var err error
		hotels, err = ctr.hotelService.ListHotelTravel()
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
		fromTmp, errFrom := time.Parse(time.RFC3339, fromStr)

		if errFrom != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errFrom.Error()})
			return
		}

		toTmp, errTo := time.Parse(time.RFC3339, toStr)

		if errTo != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errTo.Error()})
			return
		}

		from = &fromTmp
		to = &toTmp

	}

	var tags []uint

	for _, tag := range tagsArray {
		if len(tag) == 0 {
			break
		}
		i, err2 := strconv.Atoi(tag)
		if err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
			return
		}
		tags = append(tags, uint(i))
	}

	hotels, err3 := ctr.hotelService.FindHotelTravel(name, land, from, to, tags)

	if err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err3.Error()})
		return
	}

	var hotelResponse []*dto.HotelResponse

	for _, hotel := range hotels {
		hotelResponse = append(hotelResponse, dto.ToHotelResoponse(hotel))
	}
	fmt.Println(hotelResponse)

	c.JSON(http.StatusOK, hotelResponse)
}

func (ctr RestController) GetHotelById(c *gin.Context) {

	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hotel, err := ctr.hotelService.GetHotel(uint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hotelResponse := dto.ToHotelResoponse(hotel)

	c.JSON(http.StatusOK, hotelResponse)
}

func (ctr RestController) DeleteHotelRequest(c *gin.Context) {

	// Function can only be executed with a valid login status
	if _, err := ValidateLoginStatus(c); err != nil {
		return
	}

	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = ctr.hotelService.RemoveHotel(uint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})

}

func (ctr RestController) UpdateHotel(c *gin.Context) {

	// Function can only be executed with a valid login status
	if _, err := ValidateLoginStatus(c); err != nil {
		return
	}

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

	fmt.Println("Rst input", hotel.Tags)
	updatetedHotel, err := ctr.hotelService.UpdateHotel(dto.ToHotelModel(&hotel))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hotelResponse := dto.ToHotelResoponse(updatetedHotel)
	c.JSON(http.StatusOK, &hotelResponse)
}

// travel

func (ctr RestController) CreateTravelRequest(c *gin.Context) {

	// Function can only be executed with a valid login status
	if _, err := ValidateLoginStatus(c); err != nil {
		return
	}

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

	from, err := time.Parse(time.RFC3339, travel.From)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	to, err := time.Parse(time.RFC3339, travel.To)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	travelRes, err := ctr.travelService.NewTravel(uint(hotelId), uint(travel.VendorId), from, to, travel.Price, travel.Description)

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

	hotel, err := ctr.hotelService.FindHotelByTravel(uint(hotelId), uint(travelId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	travelResponse := dto.ToHotelResoponse(hotel)

	c.JSON(http.StatusOK, travelResponse)
}

func (ctr RestController) DeleteTravel(c *gin.Context) {

	// Function can only be executed with a valid login status
	if _, err := ValidateLoginStatus(c); err != nil {
		return
	}

	travelStrId := c.Param("tid")
	travelId, err := strconv.Atoi(travelStrId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ctr.travelService.RemoveTravel(uint(travelId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "sucess")
}

func (ctr RestController) UpdateTravel(c *gin.Context) {

	// Function can only be executed with a valid login status
	if _, err := ValidateLoginStatus(c); err != nil {
		return
	}

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

	var travel dto.UpdateTravelRequest

	if err = c.ShouldBindJSON(&travel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ctr.travelService.UpdateTravel(dto.ToTravelModel(travel), uint(travelId), uint(hotelId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToTravelResoponse(res))
}

// tag

func (ctr RestController) CreateTagRequest(c *gin.Context) {

	// Function can only be executed with a valid login status
	if _, err := ValidateLoginStatus(c); err != nil {
		return
	}

	var tag dto.TagRequest

	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tagRes, err := ctr.tagService.NewTag(tag.Name)

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

	tag, err := ctr.tagService.GetTag(uint(tagId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToTagResponse(tag))
}

func (ctr RestController) DeleteTagRequest(c *gin.Context) {
	// Function can only be executed with a valid login status
	if _, err := ValidateLoginStatus(c); err != nil {
		return
	}

	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = ctr.tagService.RemoveTag(uint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})

}

func (ctr RestController) ListAllTags(c *gin.Context) {

	tags, err := ctr.tagService.ListTags()

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

	// Function can only be executed with a valid login status
	if _, err := ValidateLoginStatus(c); err != nil {
		return
	}

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

	tagUpdate, err := ctr.tagService.UpdateTags(&model.Tag{Id: tag.Id, Name: tag.Name})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToTagResponse(tagUpdate))
}
