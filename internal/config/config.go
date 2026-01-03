package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppName  string
	AppEnv   string
	HTTPPort string
	DB       DBConfig
	Auth     AuthConfig
}

type DBConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	SSLMode  string
}
type AuthConfig struct {
	JWTSecret   string
	JWTExpHours int
}
func Load() (*Config, error) {
	cfg := &Config{
		AppName:  os.Getenv("APP_NAME"),
		AppEnv:   os.Getenv("APP_ENV"),
		HTTPPort: os.Getenv("HTTP_PORT"),
		DB: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		},
		Auth: AuthConfig{
			JWTSecret:   os.Getenv("JWT_SECRET"),
			JWTExpHours: 24,
		},
	}

	if cfg.HTTPPort == "" {
		return nil, fmt.Errorf("HTTP_PORT not set")
	}

	return cfg, nil
}
