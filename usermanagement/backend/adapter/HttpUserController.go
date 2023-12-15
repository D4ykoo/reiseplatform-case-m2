package adapter

import (
	model "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/ports"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

func RegisterRequest(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		SendEvent(brokerUrls, topic, ports.UserCreate, err.Error())
		return
	}
	user.Password = utils.HashPassword(user.Password, []byte(os.Getenv("SALT")))

	err := CreateUser(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		SendEvent(brokerUrls, topic, ports.Register, err.Error())
		return
	}

	SendEvent(brokerUrls, topic, ports.Register, user.Username)

	isProd := false
	isProd, _ = strconv.ParseBool(os.Getenv("PRODUCTION"))

	jwt, jwtErr := CreateJWT(user.Username, os.Getenv("JWT_SECRET"), false)

	if jwtErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		SendEvent(brokerUrls, topic, ports.Login, string(rune(http.StatusInternalServerError))+err.Error())
		return
	}

	c.SetCookie("authTravel", jwt, 3600*24, "/", "localhost", isProd, true)
	c.SetSameSite(http.SameSiteDefaultMode)

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func LoginRequest(c *gin.Context) {
	println("sasdfadsfsd")
	var user model.LoginUser
	// check if credentials are valid
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbUser, err := getUserByUsername(user.Username)

	brokerUrls := []string{os.Getenv("BROKERS")}
	topic := os.Getenv("TOPIC")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		SendEvent(brokerUrls, topic, ports.Login, string(rune(http.StatusNotFound))+err.Error())
		return
	}

	salt := []byte(os.Getenv("SALT"))

	isOk := utils.ComparePasswords(dbUser.Password, user.Password, salt)

	if !isOk {
		SendEvent(brokerUrls, topic, ports.Login, string(rune(http.StatusUnauthorized))+err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// if valid create jwt
	jwt, jwtErr := CreateJWT(user.Username, os.Getenv("JWT_SECRET"), false)

	if jwtErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		SendEvent(brokerUrls, topic, ports.Login, string(rune(http.StatusInternalServerError))+err.Error())
		return
	}

	isProd := false
	isProd, _ = strconv.ParseBool(os.Getenv("PRODUCTION"))

	c.SetCookie("authTravel", jwt, 3600*24, "/", "localhost", isProd, true)
	c.SetSameSite(http.SameSiteDefaultMode)

	c.JSON(http.StatusOK, gin.H{"status": "success", "jwt": jwt})
}

func ResetPasswordRequest(c *gin.Context) {
	brokerUrls := []string{os.Getenv("BROKERS")}
	topic := os.Getenv("TOPIC")

	// search user with pw
	var user model.ResetUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbUser, err := getUserByUsername(user.Username)
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
		To:     user.Email,
		Body:   "your-reset-link",
	})

	// TODO: Needed when generating actual reset urls
	//isOk := utils.ComparePasswords(dbUser.Password, user.Password, salt)
	//
	//if !isOk {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	//	SendEvent(brokerUrls, topic, ports.Login, string(rune(http.StatusUnauthorized))+err.Error())
	//	return
	//}

	var updatedUser model.User

	updatedUser.Password = user.NewPassword
	updatedUser.Username = dbUser.Username
	updatedUser.Email = dbUser.Email
	updatedUser.Firstname = dbUser.Firstname
	updatedUser.Lastname = dbUser.Lastname

	errUpdate := UpdateUser(dbUser.ID, updatedUser)

	if errUpdate != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		SendEvent(brokerUrls, topic, ports.Login, string(rune(http.StatusInternalServerError))+err.Error())
		return
	}

	SendEvent(brokerUrls, topic, ports.Login, "user "+updatedUser.Username+"password reset")

	// if valid create jwt
	jwt, jwtErr := CreateJWT(user.Username, os.Getenv("JWT_SECRET"), false)

	if jwtErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		SendEvent(brokerUrls, topic, ports.Login, string(rune(http.StatusInternalServerError))+err.Error())
		return
	}

	isProd := false
	isProd, _ = strconv.ParseBool(os.Getenv("PRODUCTION"))

	c.SetCookie("authTravel", jwt, 3600*24, "/", "localhost", isProd, true)
	c.SetSameSite(http.SameSiteDefaultMode)

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func LogoutRequest(c *gin.Context) {
	isProd := false
	isProd, _ = strconv.ParseBool(os.Getenv("PRODUCTION"))

	brokerUrls := []string{os.Getenv("BROKERS")}
	topic := os.Getenv("TOPIC")

	c.SetCookie("authTravel", "", -1, "/", "localhost", isProd, true)

	SendEvent(brokerUrls, topic, ports.Logout, "cookie deleted")

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
