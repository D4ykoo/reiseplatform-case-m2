package adapter

import (
	"fmt"
	model "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

// getDB returns a database connection from a configured database pool
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

func CreateUser(user model.User) error {
	db, err := getDB()

	dbUser := model.DBUser{
		Model: gorm.Model{},
		User:  user,
		Salt:  os.Getenv("SALT"),
	}

	if err != nil {
		log.Panic("Error connecting to the database:", err)
	}

	result := db.Create(&dbUser)
	return result.Error

}

// UpdateUser returns id, nil or 0 and error
func UpdateUser(user model.User) error {
	db, err := getDB()

	if err != nil {
		log.Panic("Error connecting to the database:", err)
	}

	dbUser := model.DBUser{
		Model: gorm.Model{},
		User:  user,
		Salt:  os.Getenv("SALT"),
	}

	userByUname, err := getUserByUsername(user.Username)
	if err != nil {
		return err
	}

	resIsPresent := db.First(&dbUser, "id = ?", userByUname.ID)

	if resIsPresent.Error != nil {
		return resIsPresent.Error
	}

	// updates user when id is set, otherwise save -> check for id above
	result := db.Save(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteUser(id int) error {
	db, err := getDB()
	var user model.DBUser

	if err != nil {
		log.Panic("Error connecting to the database:", err)
	}
	result := db.Delete(&user, "id = ", id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func getUserByUsername(username string) (*model.DBUser, error) {
	db, err := getDB()
	var user model.DBUser

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

func GetUser(id int64) (error, *model.DBUser) {
	db, err := getDB()
	var user model.DBUser
	if err != nil {
		log.Panic("Error connecting to the database:", err)
	}
	result := db.First(&user, "id = ?", id)

	if result.Error != nil {
		return result.Error, nil
	}

	if result.RowsAffected < 1 {
		return result.Error, nil
	}

	return nil, &user
}

func ListUser() (*[]model.DBUser, error) {
	db, err := getDB()
	var user []model.DBUser

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
