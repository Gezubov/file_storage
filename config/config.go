package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type JWTConfig struct {
	Secret     string
	Expiration int
}

var Cfg Config

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := ServerConfig{
		Port: os.Getenv("APP_PORT"),
	}
	database := DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}

	jwt := JWTConfig{
		Secret:     os.Getenv("SECRET_KEY"),
		Expiration: getEnvAsInt("JWT_EXPIRATION", 3600),
	}

	Cfg = Config{
		Server:   server,
		Database: database,
		JWT:      jwt,
	}
}

func GetConfig() *Config {
	return &Cfg
}

func getEnvAsInt(key string, defaultVal int) int {
	valStr := os.Getenv(key)
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}
	return defaultVal
}
