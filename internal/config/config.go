package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDriver  string
	DBName    string
	JWTSecret string
	Port      string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables.")
	}

	return &Config{
		DBDriver:  getEnv("DB_DRIVER", "sqlite"),
		DBName:    getEnv("DB_NAME", "expense_tracker.db"),
		JWTSecret: getEnv("JWT_SECRET", "default_secret"),
		Port:      getEnv("PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
