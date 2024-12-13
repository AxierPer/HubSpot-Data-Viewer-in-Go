package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error al cargar el archivo .env", err)
	}

	return os.Getenv("TOKEN_HUBSPOT")
}
