package config

import "os"

type Config struct {
	AppEnv    string
	Port      string
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
	JWTSecret string
	RedisHost string
	RedisPort string
}

func Load() *Config {
	return &Config{
		AppEnv:    getEnv("APP_ENV", "development"),
		Port:      getEnv("PORT", "8080"),
		DBHost:    getEnv("DB_HOST", "localhost"),
		DBPort:    getEnv("DB_PORT", "5432"),
		DBUser:    getEnv("DB_USER", "postgres"),
		DBPass:    getEnv("DB_PASSWORD", "postgres"),
		DBName:    getEnv("DB_NAME", "ecommerce"),
		JWTSecret: getEnv("JWT_SECRET", "secret"),
		RedisHost: getEnv("REDIS_HOST", "localhost"),
		RedisPort: getEnv("REDIS_PORT", "6379"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
