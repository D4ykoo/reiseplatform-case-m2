package utils

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadFile() {
	path := ".env"
	err := godotenv.Load(path)

	if err != nil {
		log.Panic("Error loading .env file at path: " + path)
	}
}
