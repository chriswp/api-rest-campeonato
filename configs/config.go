package config

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	DBDriver         string `mapstructure:"DB_DRIVER"`
	DBHost           string `mapstructure:"DB_HOST"`
	DBPort           string `mapstructure:"DB_PORT"`
	DBUser           string `mapstructure:"DB_USERNAME"`
	DBPass           string `mapstructure:"DB_PASSWORD"`
	DBDatabase       string `mapstructure:"DB_DATABASE"`
	WebServerPort    string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret        string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn     int64  `mapstructure:"JWT_EXPIRES_IN"`
	FootballAPIURL   string `mapstructure:"FOOTBALL_API_URL"`
	FootballAPIToken string `mapstructure:"FOOTBALL_API_TOKEN"`
	TokenAuth        *jwt.GinJWTMiddleware
}

var Envs Config

func LoadConfig() {
	err := godotenv.Load("configs/.env")
	if err != nil {
		panic("Erro ao carregar arquivo .env")
	}

	Envs = Config{
		DBDriver:         getEnv("DB_DRIVER", "postgres"),
		WebServerPort:    getEnv("WEB_SERVER_PORT", "http://localhost"),
		DBPort:           getEnv("DB_PORT", "8080"),
		DBUser:           getEnv("DB_USERNAME", "root"),
		DBPass:           getEnv("DB_PASSWORD", "mypassword"),
		DBHost:           getEnv("DB_HOST", "localhost"),
		DBDatabase:       getEnv("DB_DATABASE", "ecom"),
		JWTSecret:        getEnv("JWT_SECRET", "not-so-secret-now-is-it?"),
		JWTExpiresIn:     getEnvAsInt("JWT_EXPIRES_IN", 3600*24*7),
		FootballAPIURL:   getEnv("FOOTBALL_API_URL", "https://api.football-data.org/v2"),
		FootballAPIToken: getEnv("FOOTBALL_API_TOKEN", "c6c7c7b8d2f64b4f8d7b4b6f6c7b4b8d"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
