package dbGorm

import (
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter/dbGorm/entities"
	model "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/utils"
	"gorm.io/gorm"
	"os"
)

type UserRepositoryImpl struct {
	db PostgresRepository
}

func Init() UserRepositoryImpl {
	db := getDB()
	return UserRepositoryImpl{db: db}
}

func (repo UserRepositoryImpl) Save(user model.User) error {
	dbUser := entities.UserEntity{
		Model:     gorm.Model{},
		Username:  user.Username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Password:  user.Password,
		Salt:      os.Getenv("SALT"),
	}

	result := repo.db.Connection.Create(&dbUser)
	return result.Error

}

// Update returns id, nil or 0 and error
func (repo UserRepositoryImpl) Update(updateID uint, user model.User) error {
	userById, errUser := repo.FindById(updateID)

	if errUser != nil {
		return errUser
	}

	if userById.Password != user.Password {
		user.Password = utils.HashPassword(user.Password, []byte(os.Getenv("SALT")))
	}

	dbUser := entities.UserEntity{
		Model:     gorm.Model{ID: updateID},
		Username:  user.Username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Password:  user.Password,
		Salt:      os.Getenv("SALT"),
	}

	// updates user when id is set, otherwise save -> check for id above
	result := repo.db.Connection.Save(&dbUser)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo UserRepositoryImpl) Delete(id uint) error {
	user := entities.UserEntity{Model: gorm.Model{ID: uint(id)}}

	result := repo.db.Connection.Delete(&user)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected != 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (repo UserRepositoryImpl) FindByUsername(username string) (*model.User, error) {
	var user entities.UserEntity
	var retUser model.User

	result := repo.db.Connection.First(&user, "username = ?", username)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected < 1 {
		return nil, result.Error
	}

	retUser = user.ToUser()

	return &retUser, nil
}

func (repo UserRepositoryImpl) FindById(id uint) (*model.User, error) {
	var user entities.UserEntity
	var retUser model.User

	result := repo.db.Connection.First(&user, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected < 1 {
		return nil, result.Error
	}

	retUser = user.ToUser()

	return &retUser, nil
}

func (repo UserRepositoryImpl) ListAll() (*[]model.User, error) {
	var user []entities.UserEntity
	var retUser []model.User

	result := repo.db.Connection.Find(&user)

	if result.Error != nil || result.RowsAffected < 1 {
		return nil, result.Error
	}

	for _, elem := range user {
		retUser = append(retUser, elem.ToUser())
	}

	return &retUser, nil
}
