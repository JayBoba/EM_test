package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	LogLevel   string
}

func loadEnvFile(path string) error {
	if _, err := os.Stat(path); err == nil {
		return godotenv.Load(path)
	}
	return nil
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func Load() (*Config, error) {
	_ = loadEnvFile(".env")

	cfg := &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "secret"),
		DBName:     getEnv("DB_NAME", "subscriptions"),
		LogLevel:   getEnv("LOG_LEVEL", "info"),
	}
	return cfg, nil
}
