package api

import (
	"errors"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter/api/dto"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain"
	model "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/ports/inbound"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/ports/outbound"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

type RestUserController struct {
	userService domain.IUserService
	userEvents  outbound.IUserEvents
	auth        inbound.IUserAuthentication
}

func Init(userService domain.IUserService, userEvents outbound.IUserEvents, auth inbound.IUserAuthentication) RestUserController {
	return RestUserController{userService: userService, userEvents: userEvents, auth: auth}
}

// Checks if the cookie is present and validates the containing JWT.
//
// Returned error can be nil.
func (userController RestUserController) checkCookie(c *gin.Context) error {
	var err error

	cookie, cookieErr := c.Cookie("authTravel")

	if cookieErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": cookieErr.Error()})
		return cookieErr
	}

	_, valErr, _ := userController.auth.ValidateJWT(cookie, os.Getenv("JWT_SECRET"))

	if valErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": valErr.Error()})
		return valErr
	}

	return err
}

func (userController RestUserController) CreateUserRequest(c *gin.Context) {

	if cErr := userController.checkCookie(c); cErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": cErr.Error()})
		userController.userEvents.SendEvent(model.EventUserCreate, cErr.Error())
		return
	}

	var user dto.CreateUserRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		userController.userEvents.SendEvent(model.EventUserCreate, err.Error())
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

	err := userController.userService.CreateUser(saveUser)

	if err != nil {
		userController.userEvents.SendEvent(model.EventUserCreate, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userController.userEvents.SendEvent(model.EventUserCreate, "")
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (userController RestUserController) UpdateUserRequest(c *gin.Context) {
	if cErr := userController.checkCookie(c); cErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": cErr.Error()})
		userController.userEvents.SendEvent(model.EventUserUpdate, cErr.Error())
		return
	}

	var user dto.UpdateUserRequest

	id, errID := strconv.Atoi(c.Param("id"))

	if errID != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errID.Error()})
		userController.userEvents.SendEvent(model.EventUserUpdate, errID.Error())
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		userController.userEvents.SendEvent(model.EventUserUpdate, err.Error())
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
	err := userController.userService.ChangeUser(uint(id), updateUser)

	if err != nil {
		if !errors.Is(err, errors.New("falsePassword")) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			userController.userEvents.SendEvent(model.EventUserUpdate, err.Error())
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			userController.userEvents.SendEvent(model.EventUserUpdate, "invalid password")
			return
		}
	}

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		userController.userEvents.SendEvent(model.EventUserUpdate, err.Error())
		return
	}

	userController.userEvents.SendEvent(model.EventUserUpdate, user.Username)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (userController RestUserController) DeleteUserRequest(c *gin.Context) {
	if cErr := userController.checkCookie(c); cErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": cErr.Error()})
		userController.userEvents.SendEvent(model.EventUserDelete, cErr.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		userController.userEvents.SendEvent(model.EventUserDelete, err.Error())
		return
	}

	err = userController.userService.DeleteUser(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		userController.userEvents.SendEvent(model.EventUserDelete, err.Error())
		return
	}

	userController.userEvents.SendEvent(model.EventUserDelete, "")
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (userController RestUserController) GetUserRequest(c *gin.Context) {
	if cErr := userController.checkCookie(c); cErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": cErr.Error()})
		userController.userEvents.SendEvent(model.EventUserGet, cErr.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		userController.userEvents.SendEvent(model.EventUserGet, err.Error())
		return
	}

	dbUser, errGet := userController.userService.FindUser(id)

	if errGet != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errGet.Error()})
		userController.userEvents.SendEvent(model.EventUserGet, errGet.Error())
		return
	}

	user := dto.UserResponse{
		ID:        dbUser.Id,
		Username:  dbUser.Username,
		Firstname: dbUser.Firstname,
		Lastname:  dbUser.Lastname,
		Email:     dbUser.Email,
	}

	userController.userEvents.SendEvent(model.EventUserGet, "")
	c.JSON(http.StatusOK, user)
}

func (userController RestUserController) ListUserRequest(c *gin.Context) {

	if cErr := userController.checkCookie(c); cErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": cErr.Error()})
		userController.userEvents.SendEvent(model.EventUserGet, cErr.Error())
		return
	}

	dbUsers, err := userController.userService.ListAllUser()

	var users []dto.UserResponse

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		userController.userEvents.SendEvent(model.EventUserGet, err.Error())
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

	userController.userEvents.SendEvent(model.EventUserGet, "list all")
	c.JSON(http.StatusOK, users)
}
