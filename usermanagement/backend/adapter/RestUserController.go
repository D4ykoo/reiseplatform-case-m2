package adapter

import (
	domain "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/ports"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
)

var brokerUrls = []string{os.Getenv("BROKERS")}
var topic = os.Getenv("TOPIC")

func checkCookie(c *gin.Context) error {
	var err error
	cookie, cookieErr := c.Cookie("authTravel")

	if cookieErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": cookieErr.Error()})
	}

	_, valErr, _ := ValidateJWT(cookie, os.Getenv("JWT_SECRET"))

	if valErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": cookieErr.Error()})
	}
	return err
}

func CreateUserRequest(c *gin.Context) {

	if cErr := checkCookie(c); cErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": cErr.Error()})
		SendEvent(brokerUrls, topic, ports.UserCreate, cErr.Error())
		return
	}

	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		SendEvent(brokerUrls, topic, ports.UserCreate, err.Error())
		return
	}
	user.Password = utils.HashPassword(user.Password, []byte(os.Getenv("SALT")))
	log.Print(user.Password)
	err := CreateUser(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		SendEvent(brokerUrls, topic, ports.UserCreate, err.Error())
		return
	}

	SendEvent(brokerUrls, topic, ports.UserCreate, user.Username)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func UpdateUserRequest(c *gin.Context) {
	if cErr := checkCookie(c); cErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": cErr.Error()})
		SendEvent(brokerUrls, topic, ports.UserUpdate, cErr.Error())
		return
	}

	var user domain.UpdateUser

	id, errID := strconv.Atoi(c.Param("id"))

	if errID != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errID.Error()})
		SendEvent(brokerUrls, topic, ports.UserUpdate, errID.Error())
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		SendEvent(brokerUrls, topic, ports.UserUpdate, err.Error())
		return
	}

	errDBU, dbUser := GetUser(uint(id))
	if errDBU != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errDBU.Error()})
		SendEvent(brokerUrls, topic, ports.UserUpdate, errDBU.Error())
		return
	}
	isOk := utils.ComparePasswords(dbUser.Password, user.OldPassword, []byte(os.Getenv("SALT")))

	if !isOk {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errDBU.Error()})
		SendEvent(brokerUrls, topic, ports.UserUpdate, "invalid password")
		return
	}
	newUser := domain.User{
		Username:  user.Username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Password:  user.NewPassword,
	}

	err := UpdateUser(uint(id), newUser)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		SendEvent(brokerUrls, topic, ports.UserUpdate, err.Error())
		return
	}

	SendEvent(brokerUrls, topic, ports.UserUpdate, user.Username)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func DeleteUserRequest(c *gin.Context) {
	if cErr := checkCookie(c); cErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": cErr.Error()})
		SendEvent(brokerUrls, topic, ports.UserDelete, cErr.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		SendEvent(brokerUrls, topic, ports.UserDelete, err.Error())
		return
	}

	err = DeleteUser(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		SendEvent(brokerUrls, topic, ports.UserDelete, err.Error())
		return
	}

	SendEvent(brokerUrls, topic, ports.UserDelete, strconv.Itoa(id))
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func GetUserRequest(c *gin.Context) {
	if cErr := checkCookie(c); cErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": cErr.Error()})
		SendEvent(brokerUrls, topic, ports.UserGet, cErr.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		SendEvent(brokerUrls, topic, ports.UserGet, err.Error())
		return
	}

	errGet, dbUser := GetUser(uint(id))

	if errGet != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errGet.Error()})
		SendEvent(brokerUrls, topic, ports.UserGet, errGet.Error())
		return
	}

	user := domain.ResponseUser{
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
		SendEvent(brokerUrls, topic, ports.UserGet, cErr.Error())
		return
	}

	dbUsers, err := ListUser()
	var users []domain.ResponseUser

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		SendEvent(brokerUrls, topic, ports.UserGet, err.Error())
		return
	}

	for _, dbUser := range *dbUsers {
		user := domain.ResponseUser{
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
