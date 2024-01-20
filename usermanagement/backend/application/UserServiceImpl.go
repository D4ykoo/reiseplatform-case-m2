package application

import (
	"errors"
	model "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/ports/outbound"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/utils"
	"os"
	"reflect"
)

type UserServiceImpl struct {
	users outbound.IUserRepository
}

func InitUserService(userRepo outbound.IUserRepository) UserServiceImpl {
	return UserServiceImpl{userRepo}
}

func (service UserServiceImpl) CreateUser(user model.User) error {
	user.Password = utils.HashPassword(user.Password, []byte(os.Getenv("SALT")))

	err := service.users.Save(user)

	if err != nil {
		return err
	}
	return nil
}

func (service UserServiceImpl) ChangeUser(id uint, user model.User) error {
	// check if user with id exists
	dbUser, errDBU := service.users.FindById(id)

	if errDBU != nil {
		return errDBU
	}

	// check if password is the ok
	//isOk := utils.ComparePasswords(dbUser.Password, oldPassword, []byte(os.Getenv("SALT")))
	//if !isOk {
	//	return errors.New("falsePassword")
	//}

	if user.Password == "" {
		user.Password = dbUser.Password
	}

	// update user
	err := service.users.Update(id, user)
	if err != nil {
		return err
	}

	return nil
}

func (service UserServiceImpl) DeleteUser(id int) error {
	err := service.users.Delete(uint(id))

	if err != nil {
		return err
	}

	return nil
}

func (service UserServiceImpl) FindUser(identifier interface{}) (*model.User, error) {
	switch identifier.(type) {
	case int:
		id := reflect.ValueOf(identifier).Int()
		user, err := service.users.FindById(uint(id))

		if err != nil {
			return nil, err
		}

		return user, nil

	case string:
		username := reflect.ValueOf(identifier).String()
		user, err := service.users.FindByUsername(username)

		if err != nil {
			return nil, err
		}

		return user, nil

	default:
		return nil, errors.New("identifier is not a int or string")
	}
}

func (service UserServiceImpl) ListAllUser() (*[]model.User, error) {
	return service.users.ListAll()
}
