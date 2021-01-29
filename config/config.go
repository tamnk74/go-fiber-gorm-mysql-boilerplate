package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func getEnv(key string, defaultVal string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultVal
}

var PORT = getEnv("PORT", "8000")

var JWT_SECRET = getEnv("JWT_SECRET", "golang")
var JWT_EXP, _ = strconv.Atoi(getEnv("JWT_EXP", "3600"))

var DB_USERNAME = getEnv("DB_USERNAME", "root")
var DB_PASSWORD = getEnv("DB_PASSWORD", "")
var DB_HOST = getEnv("DB_HOST", "127.0.0.1")
var DB_PORT = getEnv("DB_PORT", "3306")
var DB_NAME = getEnv("DB_NAME", "test")
