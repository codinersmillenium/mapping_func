package config

import (
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	AppPort            string
	DBDriver           string
	DSN                string
	CityPath           string
	CorsAllowedOrigins string
	GinMode            string
	TrustedProxies     []string
	MigrationsPath     string
}

// Load config dari .env
func Load() *Config {
	_ = godotenv.Load() // load .env jika ada

	trustedProxies := strings.Split(os.Getenv("TRUSTED_PROXIES"), ",")

	return &Config{
		AppPort:            getEnv("APP_PORT", "8080"),
		DBDriver:           getEnv("DB_DRIVER", "mysql"),
		DSN:                getEnv("DB_DSN", ""),
		CityPath:           getEnv("CITY_PATH", "./cities.json"),
		CorsAllowedOrigins: getEnv("CORS_ALLOWED_ORIGINS", "*"),
		GinMode:            getEnv("GIN_MODE", "release"),
		TrustedProxies:     trustedProxies,
		MigrationsPath:     getEnv("MIGRATIONS_PATH", "./migrations"),
	}
}

func getEnv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
