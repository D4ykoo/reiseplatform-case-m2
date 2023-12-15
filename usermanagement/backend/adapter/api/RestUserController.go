package api

import (
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter/api/dto"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter/dbGorm"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter/kafka"
	model "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/utils"
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

	user.Password = utils.HashPassword(user.Password, []byte(os.Getenv("SALT")))

	saveUser := model.User{
		//Id:        "",
		Username:  user.Username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Password:  user.Password,
	}

	err := dbGorm.Save(saveUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		kafka.SendEvent(model.EventUserCreate, err.Error())
		return
	}

	kafka.SendEvent(model.EventUserCreate, err.Error())
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

	errDBU, dbUser := dbGorm.FindById(uint(id))
	if errDBU != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errDBU.Error()})
		kafka.SendEvent(model.EventUserUpdate, errDBU.Error())
		return
	}
	isOk := utils.ComparePasswords(dbUser.Password, user.OldPassword, []byte(os.Getenv("SALT")))

	if !isOk {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errDBU.Error()})
		kafka.SendEvent(model.EventUserUpdate, "invalid password")
		return
	}

	// TODO: is that allowed in hexagonal architecture?
	newUser := model.User{
		Username:  user.Username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Password:  user.NewPassword,
	}

	err := dbGorm.Update(uint(id), newUser)

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

	err = dbGorm.Delete(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		kafka.SendEvent(model.EventUserDelete, err.Error())
		return
	}

	kafka.SendEvent(model.EventUserDelete, err.Error())
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

	errGet, dbUser := dbGorm.FindById(uint(id))

	if errGet != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errGet.Error()})
		kafka.SendEvent(model.EventUserGet, errGet.Error())
		return
	}

	user := dto.UserResponse{
		ID:        dbUser.ID,
		Username:  dbUser.Username,
		Firstname: dbUser.Firstname,
		Lastname:  dbUser.Lastname,
		Email:     dbUser.Email,
	}

	c.JSON(http.StatusOK, user)
}

func ListUserRequest(c *gin.Context) {

	if cErr := checkCookie(c); cErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": cErr.Error()})
		kafka.SendEvent(model.EventUserGet, cErr.Error())
		return
	}

	dbUsers, err := dbGorm.ListAll()
	var users []dto.UserResponse

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		kafka.SendEvent(model.EventUserGet, err.Error())
		return
	}

	for _, dbUser := range *dbUsers {
		user := dto.UserResponse{
			ID:        dbUser.ID,
			Username:  dbUser.Username,
			Firstname: dbUser.Firstname,
			Lastname:  dbUser.Lastname,
			Email:     dbUser.Email,
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}
