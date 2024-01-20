package application

import (
	"errors"
	model "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/ports/outbound"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/utils"
	"os"
)

type HttpUserServiceImpl struct {
	users outbound.IUserRepository
}

func InitHttpService(userRepo outbound.IUserRepository) HttpUserServiceImpl {
	return HttpUserServiceImpl{userRepo}
}
func (service HttpUserServiceImpl) RegisterUser(user model.User) (*uint, error) {
	user.Password = utils.HashPassword(user.Password, []byte(os.Getenv("SALT")))

	err := service.users.Save(user)

	if err != nil {
		return nil, err
	}

	dbUser, findErr := service.users.FindByUsername(user.Username)

	if findErr != nil {
		return nil, err
	}

	return &dbUser.Id, nil
}

func (service HttpUserServiceImpl) LoginUser(username string, password string) (*uint, error) {
	dbUser, err := service.users.FindByUsername(username)

	if err != nil {
		return nil, err
	}
	salt := []byte(os.Getenv("SALT"))

	isOk := utils.ComparePasswords(dbUser.Password, password, salt)
	if !isOk {
		return nil, errors.New("falsePassword")
	}

	return &dbUser.Id, nil
}

func (service HttpUserServiceImpl) ResetPassword(username string, newPassword string) (*uint, error) {
	dbUser, err := service.users.FindByUsername(username)

	if err != nil {
		return nil, err
	}

	var updatedUser model.User

	updatedUser.Password = newPassword
	updatedUser.Username = dbUser.Username
	updatedUser.Email = dbUser.Email
	updatedUser.Firstname = dbUser.Firstname
	updatedUser.Lastname = dbUser.Lastname

	errUpdate := service.users.Update(dbUser.Id, updatedUser)

	if errUpdate != nil {
		return nil, errUpdate
	}

	return &dbUser.Id, nil
}
