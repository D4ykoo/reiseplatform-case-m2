package adapter

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type pgRepository struct {
	Connection *gorm.DB
}

func initPGConnection(idle int, max int) pgRepository {
	/*
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Berlin",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
	*/

	dsn := "host=127.0.0.1 user=postgres password=postgres dbname=hotel port=5432 sslmode=disable TimeZone=Europe/Berlin"

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Failed to connect to database", err.Error())
	}
	// Create TABLES
	//err = connection.AutoMigrate(new(T))
	//if err != nil {
	//	log.Panic("Failed to create Tables", err.Error())
	//}
	// Connection Pool
	sqlDB, err := connection.DB()

	if err != nil {
		log.Panic("Failed to create Pool", err.Error())
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(idle)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(max)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	return pgRepository{Connection: connection}
}

func (repo pgRepository) createTable(entites ...interface{}) bool {
	err := repo.Connection.AutoMigrate(entites...)
	if err != nil {
		log.Panic("Failed to create Tables", err.Error())
		return false
	}
	return true
}

/*

func (db pgRepository[T]) create(t T) {
	db.connection.Create(&t)
}

func (db pgRepository[T]) update(t T) {
	db.connection.Save(&t)
}

func (db pgRepository[T]) get(uuid string) T {
	var t T
	db.connection.First(&t)
	fmt.Println(t, uuid)
	return t
}

func (db pgRepository[T]) getBy(statemant string, parameter ...string) []entities.HotelEntity {
	var hotels []entities.HotelEntity
	db.connection.Where(statemant, parameter).Find(&hotels)
	return hotels
}

func (db pgRepository[T]) delete(t T) bool {
	res := db.connection.Delete(t)
	if res.Error != nil {
		return false
	}100
	return true
}
*/
