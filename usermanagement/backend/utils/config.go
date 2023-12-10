package utils

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadFile() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Panic("Error loading .env file")
	}
}
