package adapter

import (
	domain "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUserRequest(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// TODO: Event push
		return
	}
	err := CreateUser(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// TODO: Event push
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func UpdateUserRequest(c *gin.Context) {
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// TODO: Event push
		return
	}
	err := UpdateUser(user)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		// TODO: Event push
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func DeleteUserRequest(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// TODO: Event push
		return
	}

	err = DeleteUser(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		// TODO: Event push
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func GetUserRequest(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// TODO: Event push
		return
	}

	err, user := GetUser(int64(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		// TODO: Event push
		return
	}

	c.JSON(http.StatusOK, user)
}

func ListUserRequest(c *gin.Context) {
	users, err := ListUser()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		// TODO: Event push
		return
	}

	c.JSON(http.StatusOK, users)
}
