package outbound

import (
	model "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
)

type IUserRepository interface {
	Save(user model.User) error
	Update(updateID uint, user model.User) error
	Delete(id uint) error

	FindById(id uint) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
	ListAll() (*[]model.User, error)
}
