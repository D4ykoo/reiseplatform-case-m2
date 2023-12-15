package dbGorm

import "github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter/dbGorm/entities"

type IGORMUserCRUDRepository interface {
	Save(user entities.UserEntity) error
	Update(user entities.UserEntity) error
	Delete(id uint) error

	FindById(id uint) (*entities.UserEntity, error)
	FindByUsername(username string) (*entities.UserEntity, error)
	ListAll() (*[]entities.UserEntity, error)
}
