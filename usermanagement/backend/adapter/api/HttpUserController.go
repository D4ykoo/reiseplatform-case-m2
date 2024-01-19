package api

import (
	"errors"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter/api/dto"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter/kafka"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/application"
	model "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/ports/outbound"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func extractDomain(inputUrl string) (string, error) {
	parsedUrl, err := url.Parse(inputUrl)
	if err != nil {
		return "", err
	}

	host := parsedUrl.Hostname()

	dotIndex := strings.Index(host, ".")
	if dotIndex == -1 {
		host = host[dotIndex+1:]
	}

	return host, nil
}

func RegisterRequest(c *gin.Context) {
	var user dto.CreateUserRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		kafka.SendEvent(model.EventRegister, err.Error())
		return
	}

	userId, err := application.RegisterUser(user.ToUser())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		kafka.SendEvent(model.EventRegister, err.Error())
		return
	}

	kafka.SendEvent(model.EventRegister, user.Username)

	isProd := false
	isProd, _ = strconv.ParseBool(os.Getenv("PRODUCTION"))

	jwt, jwtErr := adapter.CreateJWT(user.Username, userId, os.Getenv("JWT_SECRET"), false)

	if jwtErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		kafka.SendEvent(model.EventRegister, string(rune(http.StatusInternalServerError))+err.Error())
		return
	}
	domain, err := extractDomain(os.Getenv("DOMAIN"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		kafka.SendEvent(model.EventRegister, string(rune(http.StatusInternalServerError))+err.Error())
		return
	}

	c.SetCookie("authTravel", jwt, 3600*24, "/", domain, isProd, true)
	c.SetSameSite(http.SameSiteDefaultMode)

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func LoginRequest(c *gin.Context) {
	var user dto.LoginRequest
	// check if credentials are valid
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, loginErr := application.LoginUser(user.Username, user.Password)

	if loginErr != nil {
		if errors.Is(loginErr, errors.New("falsePassword")) {
			kafka.SendEvent(model.EventLogin, string(rune(http.StatusUnauthorized))+loginErr.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": loginErr.Error()})
			return
		} else {
			kafka.SendEvent(model.EventLogin, string(rune(http.StatusNotFound))+loginErr.Error())
			c.JSON(http.StatusNotFound, gin.H{"error": loginErr.Error()})
			return
		}
	}

	// if valid create jwt
	jwt, jwtErr := adapter.CreateJWT(user.Username, userId, os.Getenv("JWT_SECRET"), false)

	if jwtErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": jwtErr.Error()})
		kafka.SendEvent(model.EventLogin, string(rune(http.StatusInternalServerError))+jwtErr.Error())
		return
	}

	isProd := false
	isProd, _ = strconv.ParseBool(os.Getenv("PRODUCTION"))

	domain, err := extractDomain(os.Getenv("DOMAIN"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		kafka.SendEvent(model.EventRegister, string(rune(http.StatusInternalServerError))+err.Error())
		return
	}

	c.SetCookie("authTravel", jwt, 3600*24, "/", domain, isProd, true)
	c.SetSameSite(http.SameSiteDefaultMode)

	c.JSON(http.StatusOK, gin.H{"status": "success", "jwt": jwt})
}

func ResetPasswordRequest(c *gin.Context) {
	// search user with pw
	var user dto.ResetPasswordRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate reset urls, currently only mocked
	adapter.SendEmail(outbound.EmailContent{
		Header: "Password Reset Travel-Management",
		Title:  "Your password reset",
		From:   "mail@travel",
		To:     user.Email,
		Body:   "your-reset-link",
	})

	// Needed when generating actual reset urls
	//isOk := utils.ComparePasswords(dbUser.Password, user.Password, salt)
	//
	//if !isOk {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	//	SendEvent(brokerUrls, topic, ports.Login, string(rune(http.StatusUnauthorized))+err.Error())
	//	return
	//}

	// when password ok after clicking on link
	userId, errUpdate := application.ResetPassword(user.Username, user.NewPassword)

	if errUpdate != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errUpdate.Error()})
		kafka.SendEvent(model.EventPasswordReset, string(rune(http.StatusInternalServerError))+errUpdate.Error())
		return
	}

	kafka.SendEvent(model.EventPasswordReset, "user "+user.Username+"password reset")

	// if valid create jwt
	jwt, jwtErr := adapter.CreateJWT(user.Username, userId, os.Getenv("JWT_SECRET"), false)

	if jwtErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": jwtErr.Error()})
		kafka.SendEvent(model.EventPasswordReset, string(rune(http.StatusInternalServerError))+jwtErr.Error())
		return
	}

	isProd := false
	isProd, _ = strconv.ParseBool(os.Getenv("PRODUCTION"))

	domain, err := extractDomain(os.Getenv("DOMAIN"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		kafka.SendEvent(model.EventRegister, string(rune(http.StatusInternalServerError))+err.Error())
		return
	}

	c.SetCookie("authTravel", jwt, 3600*24, "/", domain, isProd, true)
	c.SetSameSite(http.SameSiteDefaultMode)

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func LogoutRequest(c *gin.Context) {
	isProd := false
	isProd, _ = strconv.ParseBool(os.Getenv("PRODUCTION"))

	domain, err := extractDomain(os.Getenv("DOMAIN"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		kafka.SendEvent(model.EventRegister, string(rune(http.StatusInternalServerError))+err.Error())
		return
	}

	c.SetCookie("authTravel", "", -1, "/", domain, isProd, true)

	kafka.SendEvent(model.EventLogout, "cookie deleted")

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
