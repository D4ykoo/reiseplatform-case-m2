package adapter

import (
	domain "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/ports"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

var brokerUrls = []string{os.Getenv("BROKERS")}
var topic = os.Getenv("TOPIC")

func CreateUserRequest(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		SendEvent(brokerUrls, topic, ports.UserCreate, err.Error())
		return
	}
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
	var user domain.User
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

	err := UpdateUser(uint(id), user)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		SendEvent(brokerUrls, topic, ports.UserUpdate, err.Error())
		return
	}

	SendEvent(brokerUrls, topic, ports.UserUpdate, user.Username)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func DeleteUserRequest(c *gin.Context) {
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
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		SendEvent(brokerUrls, topic, ports.UserGet, err.Error())
		return
	}

	errGet, user := GetUser(uint(id))

	if errGet != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errGet.Error()})
		SendEvent(brokerUrls, topic, ports.UserGet, errGet.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func ListUserRequest(c *gin.Context) {
	users, err := ListUser()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		SendEvent(brokerUrls, topic, ports.UserGet, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}
