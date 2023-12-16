package application

import (
	"errors"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter/dbGorm"
	model "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/utils"
	"os"
	"reflect"
)

func CreateUser(user model.User) error {
	user.Password = utils.HashPassword(user.Password, []byte(os.Getenv("SALT")))

	err := dbGorm.Save(user)

	if err != nil {
		return err
	}
	return nil
}

func ChangeUser(id uint, user model.User, oldPassword string) error {
	// check if user with id exists
	dbUser, errDBU := dbGorm.FindById(id)

	if errDBU != nil {
		return errDBU
	}

	// check if password is the ok
	isOk := utils.ComparePasswords(dbUser.Password, oldPassword, []byte(os.Getenv("SALT")))
	if !isOk {
		return errors.New("falsePassword")
	}

	// update user
	err := dbGorm.Update(id, user)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(id int) error {
	err := dbGorm.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func FindUser(identifier interface{}) (*model.User, error) {
	switch identifier.(type) {
	case int:
		id := reflect.ValueOf(identifier).Int()
		user, err := dbGorm.FindById(uint(id))

		if err != nil {
			return nil, err
		}

		return user, nil

	case string:
		username := reflect.ValueOf(identifier).String()
		user, err := dbGorm.FindByUsername(username)

		if err != nil {
			return nil, err
		}

		return user, nil

	default:
		return nil, errors.New("identifier is not a int or string")
	}
}

func ListAllUser() (*[]model.User, error) {
	return dbGorm.ListAll()
}

func RegisterUser(user model.User) error {
	user.Password = utils.HashPassword(user.Password, []byte(os.Getenv("SALT")))

	err := dbGorm.Save(user)

	if err != nil {
		return err
	}
	return nil
}

func LoginUser(username string, password string) error {
	dbUser, err := dbGorm.FindByUsername(username)

	if err != nil {
		return err
	}
	salt := []byte(os.Getenv("SALT"))

	isOk := utils.ComparePasswords(dbUser.Password, password, salt)
	if !isOk {
		return errors.New("falsePassword")
	}

	return nil
}

func ResetPassword(username string, newPassword string) error {
	dbUser, err := dbGorm.FindByUsername(username)

	if err != nil {
		return err
	}

	var updatedUser model.User

	updatedUser.Password = newPassword
	updatedUser.Username = dbUser.Username
	updatedUser.Email = dbUser.Email
	updatedUser.Firstname = dbUser.Firstname
	updatedUser.Lastname = dbUser.Lastname

	errUpdate := dbGorm.Update(dbUser.Id, updatedUser)

	if errUpdate != nil {
		return errUpdate
	}

	return nil
}
