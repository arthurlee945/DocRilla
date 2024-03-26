package config

import (
	"os"
)

type Config struct {
	DatabaseUrl string
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
