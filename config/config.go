package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseConfig DatabaseConfig
	ServerConfig   ServerConfig
	JWTConfig      JWTConfig
	Environment    string
}

type ServerConfig struct {
	PORT string
	HOST string
}
type DatabaseConfig struct {
	Name     string
	User     string
	Password string
	Port     string
	Host     string
	SSLMode  string
}
type JWTConfig struct {
	Secret            string
	AccessExpiration  int
	RefreshExpiration int
}

func LoadConfig() *Config {
	// Try loading .env from current dir and a few parent dirs so running
	// from subfolders (for example `cmd/api`) still picks up the project's .env
	var loadedFrom string
	candidates := []string{".env", "../.env", "../../.env", "../../../.env"}
	for _, p := range candidates {
		if err := godotenv.Load(p); err == nil {
			loadedFrom = p
			break
		}
	}
	if loadedFrom != "" {
		log.Println("loaded .env from", loadedFrom)
	} else {
		log.Println("no .env file found (looked in .env and parent directories)")
	}

	return &Config{
		ServerConfig: ServerConfig{
			PORT: getEnv("PORT", "8000"),
			HOST: getEnv("HOST", "0.0.0.0"),
		},
		DatabaseConfig: DatabaseConfig{
			Name:     getEnv("DB_NAME", "store"),
			Password: getEnv("DB_PASSWORD", ""),
			Port:     getEnv("DB_PORT", "5432"),
			Host:     getEnv("DB_HOST", "localhost"),
			User:     getEnv("DB_USER", "postgres"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		JWTConfig: JWTConfig{
			Secret:            getEnv("SECRET_KEY", ""),
			AccessExpiration:  getEnvAsInt("JWT_ACCESS_EXPIRATION", 60),
			RefreshExpiration: getEnvAsInt("JWT_REFRESH_EXPIRATION", 7),
		},
		Environment: getEnv("ENVIRONMENT", "development"),
	}
}
func getEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
func getEnvAsInt(key string, defaultValue int) int {
	valueSTR := getEnv(key, "")
	if value, err := strconv.Atoi(valueSTR); err == nil {
		return value
	}
	return defaultValue
}
