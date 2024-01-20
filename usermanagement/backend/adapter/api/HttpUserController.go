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
	"net/url"
	"os"
	"strconv"
	"strings"
)

type HttpUserController struct {
	userHttpService domain.IUserHttpService
	userEvents      outbound.IUserEvents
	emailNotify     outbound.IEmailNotification
	auth            inbound.IUserAuthentication
}

func InitHttpUserController(userHttpService domain.IUserHttpService, userEvents outbound.IUserEvents, emailNotify outbound.IEmailNotification, auth inbound.IUserAuthentication) HttpUserController {
	return HttpUserController{
		userHttpService: userHttpService,
		userEvents:      userEvents,
		emailNotify:     emailNotify,
		auth:            auth,
	}
}

func (controller HttpUserController) extractDomain(inputUrl string) (string, error) {
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

func (controller HttpUserController) RegisterRequest(c *gin.Context) {
	var user dto.CreateUserRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		controller.userEvents.SendEvent(model.EventRegister, err.Error())
		return
	}

	userId, err := controller.userHttpService.RegisterUser(user.ToUser())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		controller.userEvents.SendEvent(model.EventRegister, err.Error())
		return
	}

	controller.userEvents.SendEvent(model.EventRegister, user.Username)

	isProd := false
	isProd, _ = strconv.ParseBool(os.Getenv("PRODUCTION"))

	jwt, jwtErr := controller.auth.CreateJWT(user.Username, userId, os.Getenv("JWT_SECRET"), false)

	if jwtErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": jwtErr.Error()})
		controller.userEvents.SendEvent(model.EventRegister, string(rune(http.StatusInternalServerError))+jwtErr.Error())
		return
	}
	extractDomain, errDomain := controller.extractDomain(os.Getenv("DOMAIN"))

	if errDomain != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errDomain.Error()})
		controller.userEvents.SendEvent(model.EventRegister, string(rune(http.StatusInternalServerError))+errDomain.Error())
		return
	}

	c.SetCookie("authTravel", jwt, 3600*24, "/", extractDomain, isProd, true)
	c.SetSameSite(http.SameSiteDefaultMode)

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (controller HttpUserController) LoginRequest(c *gin.Context) {
	var user dto.LoginRequest
	// check if credentials are valid
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, loginErr := controller.userHttpService.LoginUser(user.Username, user.Password)

	if loginErr != nil {
		if errors.Is(loginErr, errors.New("falsePassword")) {
			controller.userEvents.SendEvent(model.EventLogin, string(rune(http.StatusUnauthorized))+loginErr.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": loginErr.Error()})
			return
		} else {
			controller.userEvents.SendEvent(model.EventLogin, string(rune(http.StatusNotFound))+loginErr.Error())
			c.JSON(http.StatusNotFound, gin.H{"error": loginErr.Error()})
			return
		}
	}

	// if valid create jwt
	jwt, jwtErr := controller.auth.CreateJWT(user.Username, userId, os.Getenv("JWT_SECRET"), false)

	if jwtErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": jwtErr.Error()})
		controller.userEvents.SendEvent(model.EventLogin, string(rune(http.StatusInternalServerError))+jwtErr.Error())
		return
	}

	isProd := false
	isProd, _ = strconv.ParseBool(os.Getenv("PRODUCTION"))

	domainExtracted, errDomain := controller.extractDomain(os.Getenv("DOMAIN"))

	if errDomain != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errDomain.Error()})
		controller.userEvents.SendEvent(model.EventRegister, string(rune(http.StatusInternalServerError))+errDomain.Error())
		return
	}

	c.SetCookie("authTravel", jwt, 3600*24, "/", domainExtracted, isProd, true)
	c.SetSameSite(http.SameSiteDefaultMode)

	c.JSON(http.StatusOK, gin.H{"status": "success", "jwt": jwt})
}

func (controller HttpUserController) ResetPasswordRequest(c *gin.Context) {
	// search user with pw
	var user dto.ResetPasswordRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate reset urls, currently only mocked
	controller.emailNotify.SendEmail(outbound.EmailContent{
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
	userId, errUpdate := controller.userHttpService.ResetPassword(user.Username, user.NewPassword)

	if errUpdate != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errUpdate.Error()})
		controller.userEvents.SendEvent(model.EventPasswordReset, string(rune(http.StatusInternalServerError))+errUpdate.Error())
		return
	}

	controller.userEvents.SendEvent(model.EventPasswordReset, "user "+user.Username+"password reset")

	// if valid create jwt
	jwt, jwtErr := controller.auth.CreateJWT(user.Username, userId, os.Getenv("JWT_SECRET"), false)

	if jwtErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": jwtErr.Error()})
		controller.userEvents.SendEvent(model.EventPasswordReset, string(rune(http.StatusInternalServerError))+jwtErr.Error())
		return
	}

	isProd := false
	isProd, _ = strconv.ParseBool(os.Getenv("PRODUCTION"))

	domainExtracted, err := controller.extractDomain(os.Getenv("DOMAIN"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		controller.userEvents.SendEvent(model.EventRegister, string(rune(http.StatusInternalServerError))+err.Error())
		return
	}

	c.SetCookie("authTravel", jwt, 3600*24, "/", domainExtracted, isProd, true)
	c.SetSameSite(http.SameSiteDefaultMode)

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (controller HttpUserController) LogoutRequest(c *gin.Context) {
	isProd := false
	isProd, _ = strconv.ParseBool(os.Getenv("PRODUCTION"))

	extractDomain, err := controller.extractDomain(os.Getenv("DOMAIN"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		controller.userEvents.SendEvent(model.EventRegister, string(rune(http.StatusInternalServerError))+err.Error())
		return
	}

	c.SetCookie("authTravel", "", -1, "/", extractDomain, isProd, true)

	controller.userEvents.SendEvent(model.EventLogout, "cookie deleted")

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
