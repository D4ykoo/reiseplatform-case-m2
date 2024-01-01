package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadFile() {
	path := ".env"
	err := godotenv.Load(path)

	if err != nil {
		log.Panic("Error loading .env file at path: " + path)
	}
}
