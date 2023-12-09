package adapter

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mig3177/hotelmanagement/adapter/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ManagementDB struct {
	db *sql.DB
}

func NewDB() ManagementDB {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Berlin",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Failed to connect to database", err.Error())
	}
	// Create TABLES
	err = connection.AutoMigrate(&entities.HotelEntity{})
	if err != nil {
		log.Panic("Failed to create Tables", err.Error())
	}
	// Connection Pool
	sqlDB, err := connection.DB()

	if err != nil {
		log.Panic("Failed to create Pool", err.Error())
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	return ManagementDB{db: sqlDB}
}
