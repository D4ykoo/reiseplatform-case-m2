package adapter

import (
	model "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginRequest(c *gin.Context) {
	var user model.LoginUser
	// check if credentials are valid
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbUser, err := getUserByUsername(user.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// TODO: env salt
	isOk := utils.ComparePasswords(dbUser.Password, user.Password, []byte(""))

	if !isOk {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// if valid create jwt
	jwt, jwtErr := CreateJWT(user.Username, "", false)

	if jwtErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// TODO: set cookie or session storage
	// return jwt and 200
	c.JSON(http.StatusOK, gin.H{"status": "success", "jwt": jwt})
}

func ResetPasswordRequest(c *gin.Context) {
	// search user with pw
	// set the new val
	// regnerate jwt
	// return jwt and 200
	return
}
