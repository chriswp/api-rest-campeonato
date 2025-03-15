package config

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
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
	
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "api",
		Key:         []byte(Envs.JWTSecret),
		Timeout:     time.Duration(Envs.JWTExpiresIn) * time.Second,
		MaxRefresh:  time.Hour,
		IdentityKey: "user_id",
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var login struct {
				Username string `json:"username"`
				Password string `json:"password"`
			}
			if err := c.ShouldBindJSON(&login); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if login.Username == "admin" && login.Password == "admin" {
				return map[string]interface{}{"user_id": 1, "role": "admin"}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(map[string]interface{}); ok && v["role"] == "admin" {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{"error": message})
		},
		TokenLookup:   "header: Authorization, query: token",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("Erro ao criar middleware JWT:", err)
	}

	Envs.TokenAuth = authMiddleware

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
