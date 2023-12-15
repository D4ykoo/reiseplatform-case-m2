package dbGorm

import (
	"fmt"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter/dbGorm/entities"
	model "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

// getDB returns a database connection from a configured database pool.
// TODO: May need to be in the interface
func getDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("SSL_TLS"),
		os.Getenv("TIMEZONE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	// migrate the db model
	errMigrate := db.AutoMigrate(&entities.UserEntity{})
	if errMigrate != nil {
		return nil, errMigrate
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()

	if err != nil {
		log.Panic(err)
	}
	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(100)

	sqlDB.SetConnMaxLifetime(time.Hour)
	return db, nil
}

func Save(user model.User) error {
	db, err := getDB()

	dbUser := entities.UserEntity{
		Model:     gorm.Model{},
		Username:  user.Username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Password:  user.Password,
		Salt:      os.Getenv("SALT"),
	}

	if err != nil {
		log.Panic("Error connecting to the database:", err)
	}

	result := db.Create(&dbUser)
	return result.Error

}

// Update returns id, nil or 0 and error
func Update(updateID uint, user model.User) error {
	db, err := getDB()

	if err != nil {
		log.Panic("Error connecting to the database:", err)
	}

	errUser, _ := FindById(updateID)

	if errUser != nil {
		return err
	}

	user.Password = utils.HashPassword(user.Password, []byte(os.Getenv("SALT")))

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
	result := db.Save(&dbUser)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func Delete(id int) error {
	db, err := getDB()
	user := entities.UserEntity{Model: gorm.Model{ID: uint(id)}}

	if err != nil {
		log.Panic("Error connecting to the database:", err)
	}
	result := db.Delete(&user)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected != 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func FindByUsername(username string) (*entities.UserEntity, error) {
	db, err := getDB()
	var user entities.UserEntity

	if err != nil {
		log.Panic("Error connecting to the database:", err)
		return nil, err
	}

	result := db.First(&user, "username = ?", username)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected < 1 {
		return nil, result.Error
	}

	return &user, nil
}

func FindById(id uint) (*entities.UserEntity, error) {
	db, err := getDB()
	var user entities.UserEntity
	if err != nil {
		log.Panic("Error connecting to the database:", err)
	}
	result := db.First(&user, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected < 1 {
		return nil, result.Error
	}

	return &user, nil
}

func ListAll() (*[]entities.UserEntity, error) {
	db, err := getDB()
	var user []entities.UserEntity

	if err != nil {
		log.Panic("Error connecting to the database:", err)
		return nil, err
	}

	result := db.Find(&user)

	if result.Error != nil || result.RowsAffected < 1 {
		return nil, result.Error
	}

	return &user, nil
}
