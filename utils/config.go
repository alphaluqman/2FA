package utils

import (
	"github.com/joho/godotenv"
	"log"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
		return
	}
}
