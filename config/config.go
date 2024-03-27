package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl string
}

func Initialize(filename string) error {
	return godotenv.Load(filename)
}

func New() *Config {
	return &Config{
		DatabaseUrl: getEnv("DATABASE_URL", ""),
	}
}

func getEnv(key, defaultVal string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return defaultVal
}
