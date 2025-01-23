package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Redis struct {
	Host     string
	Password string
	DB       int
}

type Config struct {
	AppEnv   string
	Port     string
	
	DB_HOST string
	DB_USER string
	DB_PASSWORD string
	DB_NAME string
	DB_PORT string
	SSL_MODE string
	TIMEZONE string

	BCRYPT_COST string

	Redis Redis
}

var (
	configInstance *Config
	once           sync.Once
)

// LoadConfig initializes and loads the configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		AppEnv:   getEnv("APP_ENV", "development"),
		Port:     getEnv("PORT", "8080"),

		DB_HOST: getEnv("DB_HOST", "localhost"),
		DB_USER: getEnv("DB_USER", "postgres"),
		DB_PASSWORD: getEnv("DB_PASSWORD", "postgres"),
		DB_NAME: getEnv("DB_NAME", "postgres5"),
		DB_PORT: getEnv("DB_PORT", "5432"),
		SSL_MODE: getEnv("SSL_MODE", "disable"),
		TIMEZONE: getEnv("TIMEZONE", "Asia/Shanghai"),

		BCRYPT_COST: getEnv("BCRYPT_COST", "5"),

		Redis: Redis{
			Host:     getEnv("REDIS_HOST", "localhost:6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       0,
		},
	}, nil
}

// GetConfig provides access to the singleton Config instance
func GetConfig() *Config {
	once.Do(func() {
		configInstance, _ = LoadConfig()
	})
	return configInstance
}

// getEnv fetches the value of an environment variable or returns a default value
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
