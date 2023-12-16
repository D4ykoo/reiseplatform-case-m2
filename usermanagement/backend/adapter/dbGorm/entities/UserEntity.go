package entities

import (
	model "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
	"gorm.io/gorm"
)

type UserEntity struct {
	gorm.Model
	Username  string
	Firstname string
	Lastname  string
	Email     string
	Password  string
	Salt      string
}

func (ue UserEntity) ToUser() model.User {
	return model.User{
		Id:        ue.ID,
		Username:  ue.Username,
		Firstname: ue.Firstname,
		Lastname:  ue.Lastname,
		Email:     ue.Email,
		Password:  ue.Password,
	}
}
