package adapter

import (
	model "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/ports"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func LoginRequest(c *gin.Context) {
	var user model.LoginUser
	// check if credentials are valid
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbUser, err := getUserByUsername(user.Username)

	brokerUrls := []string{os.Getenv("BROKERS")}

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		SendEvent(brokerUrls, topic, ports.Login, string(rune(http.StatusNotFound))+err.Error())
		return
	}

	salt := []byte(os.Getenv("SALT"))

	isOk := utils.ComparePasswords(dbUser.Password, user.Password, salt)

	if !isOk {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		SendEvent(brokerUrls, topic, ports.Login, string(rune(http.StatusUnauthorized))+err.Error())
		return
	}

	// if valid create jwt
	jwt, jwtErr := CreateJWT(user.Username, "", false)

	if jwtErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		SendEvent(brokerUrls, topic, ports.Login, string(rune(http.StatusInternalServerError))+err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "jwt": jwt})
}

func ResetPasswordRequest(c *gin.Context) {
	brokerUrls := []string{os.Getenv("BROKERS")}
	salt := []byte(os.Getenv("SALT"))

	// search user with pw
	var user model.ResetUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbUser, err := getUserByUsername(user.OldLoginUser.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		SendEvent(brokerUrls, topic, ports.Login, string(rune(http.StatusNotFound))+err.Error())
		return
	}

	// email
	SendEmail(ports.EmailContent{
		Header: "Password Reset Travel-Management",
		Title:  "Your password reset",
		From:   "mail@travel",
		To:     dbUser.Email,
		Body:   "your-reset-link",
	})

	isOk := utils.ComparePasswords(dbUser.Password, user.OldLoginUser.Username, salt)

	if !isOk {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		SendEvent(brokerUrls, topic, ports.Login, string(rune(http.StatusUnauthorized))+err.Error())
		return
	}

	var updatedUser model.User

	updatedUser.Password = user.NewPassword
	updatedUser.Username = dbUser.Username
	updatedUser.Email = dbUser.Email
	updatedUser.Firstname = dbUser.Firstname
	updatedUser.Lastname = dbUser.Lastname

	errUpdate := UpdateUser(updatedUser)

	if errUpdate != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		SendEvent(brokerUrls, topic, ports.Login, string(rune(http.StatusInternalServerError))+err.Error())
		return
	}

	SendEvent(brokerUrls, topic, ports.Login, "user "+updatedUser.Username+"password reset")
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
