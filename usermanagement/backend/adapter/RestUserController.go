package adapter

import (
	domain "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUserRequest(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// TODO: Event push
		return
	}
	id, err := CreateUser(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// TODO: Event push
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "id": id})
}

func UpdateUserRequest(c *gin.Context) {
	return
}

func DeleteUserRequest(c *gin.Context) {
	return
}

func GetUserRequest(c *gin.Context) {
	return
}

func ListUserRequest(c *gin.Context) {
	return
}
