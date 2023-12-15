package application

import (
	"errors"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter/dbGorm"
	model "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/utils"
	"os"
)

func CreateUser(user model.User) error {
	user.Password = utils.HashPassword(user.Password, []byte(os.Getenv("SALT")))

	err := dbGorm.Save(user)

	if err != nil {
		return err
	}
	return nil
}

func DeleteUser() {

}

func FindUser() {

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

func ListAllUser() {}

func RegisterUser() {}

func LoginUser() {

}

func ResetPassword() {

}
