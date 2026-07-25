package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Name string
	Env  string
	Port int
}

type DBConfig struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

type JWTConfig struct {
	Secret      string
	ExpireHours int
}

type Config struct {
	App AppConfig
	DB  DBConfig
	JWT JWTConfig
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("load .env file: %w", err)
	}

	cfg := &Config{
		App: AppConfig{
			Name: getEnv("APP_NAME", "GoKit"),
			Env:  getEnv("APP_ENV", "development"),
			Port: getEnvInt("APP_PORT", 8080),
		},
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvInt("DB_PORT", 3306),
			Database: getEnv("DB_DATABASE", "gokit"),
			Username: getEnv("DB_USERNAME", "root"),
			Password: getEnv("DB_PASSWORD", ""),
		},
		JWT: JWTConfig{
			Secret:      getEnv("JWT_SECRET", "changeme"),
			ExpireHours: getEnvInt("JWT_EXPIRE_HOURS", 24),
		},
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if val, ok := os.LookupEnv(key); ok {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return fallback
}
