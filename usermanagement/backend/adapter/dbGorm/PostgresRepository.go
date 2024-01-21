package dbGorm

import (
	"fmt"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter/dbGorm/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type PostgresRepository struct {
	Connection *gorm.DB
}

func getDB() PostgresRepository {
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
		log.Panic(errMigrate.Error())
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, errDb := db.DB()

	if errDb != nil {
		log.Panic(errDb)
	}
	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(100)

	sqlDB.SetConnMaxLifetime(time.Hour)

	return PostgresRepository{Connection: db}
}
