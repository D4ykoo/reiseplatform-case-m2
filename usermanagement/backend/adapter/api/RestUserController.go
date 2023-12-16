package api

import (
	"errors"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter/api/dto"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter/kafka"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/application"
	model "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

// Checks if the cookie is present and validates the containing JWT.
//
// Returned error can be nil.
func checkCookie(c *gin.Context) error {
	var err error

	cookie, cookieErr := c.Cookie("authTravel")

	if cookieErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": cookieErr.Error()})
		return cookieErr
	}

	_, valErr, _ := adapter.ValidateJWT(cookie, os.Getenv("JWT_SECRET"))

	if valErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": valErr.Error()})
		return valErr
	}

	return err
}

func CreateUserRequest(c *gin.Context) {

	if cErr := checkCookie(c); cErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": cErr.Error()})
		kafka.SendEvent(model.EventUserCreate, cErr.Error())
		return
	}

	var user dto.CreateUserRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		kafka.SendEvent(model.EventUserCreate, err.Error())
		return
	}

	saveUser := model.User{
		//Id:        "",
		Username:  user.Username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Password:  user.Password,
	}

	err := application.CreateUser(saveUser)

	if err != nil {
		kafka.SendEvent(model.EventUserCreate, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	kafka.SendEvent(model.EventUserCreate, "")
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func UpdateUserRequest(c *gin.Context) {
	if cErr := checkCookie(c); cErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": cErr.Error()})
		kafka.SendEvent(model.EventUserUpdate, cErr.Error())
		return
	}

	var user dto.UpdateUserRequest

	id, errID := strconv.Atoi(c.Param("id"))

	if errID != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errID.Error()})
		kafka.SendEvent(model.EventUserUpdate, errID.Error())
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		kafka.SendEvent(model.EventUserUpdate, err.Error())
		return
	}

	updateUser := model.User{
		Id:        uint(id),
		Username:  user.Username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Password:  user.NewPassword,
	}

	//errDBU, dbUser := dbGorm.FindById(uint(id))
	err := application.ChangeUser(uint(id), updateUser, user.OldPassword)

	if err != nil {
		if !errors.Is(err, errors.New("falsePassword")) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			kafka.SendEvent(model.EventUserUpdate, err.Error())
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			kafka.SendEvent(model.EventUserUpdate, "invalid password")
			return
		}
	}

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		kafka.SendEvent(model.EventUserUpdate, err.Error())
		return
	}

	kafka.SendEvent(model.EventUserUpdate, user.Username)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func DeleteUserRequest(c *gin.Context) {
	if cErr := checkCookie(c); cErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": cErr.Error()})
		kafka.SendEvent(model.EventUserDelete, cErr.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		kafka.SendEvent(model.EventUserDelete, err.Error())
		return
	}

	err = application.DeleteUser(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		kafka.SendEvent(model.EventUserDelete, err.Error())
		return
	}

	kafka.SendEvent(model.EventUserDelete, "")
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func GetUserRequest(c *gin.Context) {
	if cErr := checkCookie(c); cErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": cErr.Error()})
		kafka.SendEvent(model.EventUserGet, cErr.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		kafka.SendEvent(model.EventUserGet, err.Error())
		return
	}

	dbUser, errGet := application.FindUser(id)

	if errGet != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errGet.Error()})
		kafka.SendEvent(model.EventUserGet, errGet.Error())
		return
	}

	user := dto.UserResponse{
		ID:        dbUser.Id,
		Username:  dbUser.Username,
		Firstname: dbUser.Firstname,
		Lastname:  dbUser.Lastname,
		Email:     dbUser.Email,
	}

	kafka.SendEvent(model.EventUserGet, "")
	c.JSON(http.StatusOK, user)
}

func ListUserRequest(c *gin.Context) {

	if cErr := checkCookie(c); cErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": cErr.Error()})
		kafka.SendEvent(model.EventUserGet, cErr.Error())
		return
	}

	dbUsers, err := application.ListAllUser()

	var users []dto.UserResponse

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		kafka.SendEvent(model.EventUserGet, err.Error())
		return
	}

	for _, dbUser := range *dbUsers {
		user := dto.UserResponse{
			ID:        dbUser.Id,
			Username:  dbUser.Username,
			Firstname: dbUser.Firstname,
			Lastname:  dbUser.Lastname,
			Email:     dbUser.Email,
		}
		users = append(users, user)
	}

	kafka.SendEvent(model.EventUserGet, "list all")
	c.JSON(http.StatusOK, users)
}
