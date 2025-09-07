package config

import (
	"mini-ecommerce/pkg/utils"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server ServerConfig
	DB     DBConfig
	JWT    JWTConfig
	Rate   RateLimitConfig
	CORS   CORSConfig
	App    AppEnv
}

type ServerConfig struct {
	Port string
	Host string
}

type JWTConfig struct {
	SecretKey string
	Expire    time.Duration
}

type RateLimitConfig struct {
	RequestPerMinute int
	Minutes          int
}

type CORSConfig struct {
	AllowedOrigins []string
	AllowedMethods []string
	AllowedHeaders []string
}

type DBConfig struct {
	Host    string
	Port    string
	User    string
	Pass    string
	DBName  string
	SSLMode string
}

type AppEnv struct {
	Environment string
	Debug       bool
}

func LoadConfig() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	RateLimitRequestsPerMinute, err := utils.GetEnvAsInt("RATE_LIMIT_REQUESTS_PER_MINUTE", 60)
	if err != nil {
		return nil, err
	}
	RateLimitMinutes, err := utils.GetEnvAsInt("RATE_LIMIT_MINUTES", 1)
	if err != nil {
		return nil, err
	}
	JWTExpirationHours, err := utils.GetEnvAsInt("JWT_EXPIRATION_HOURS", 24)
	if err != nil {
		return nil, err
	}
	AppDebug, err := utils.GetEnvAsBool("APP_DEBUG", true)
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "3000"),
			Host: getEnv("SERVER_HOST", "localhost"),
		},
		DB: DBConfig{
			Host:    getEnv("DB_HOST", "localhost"),
			Port:    getEnv("DB_PORT", "5432"),
			User:    getEnv("DB_USER", "postgres"),
			Pass:    getEnv("DB_PASS", "password"),
			DBName:  getEnv("DB_NAME", "mini_ecommerce_db"),
			SSLMode: getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			SecretKey: getEnv("JWT_SECRET_KEY", "your_secret_key"),
			Expire:    time.Duration(JWTExpirationHours) * time.Hour,
		},
		Rate: RateLimitConfig{
			RequestPerMinute: RateLimitRequestsPerMinute,
			Minutes:          RateLimitMinutes,
		},
		CORS: CORSConfig{
			AllowedOrigins: utils.GetEnvAsSlice("CORS_ALLOWED_ORIGINS", []string{"*"}, ","),
			AllowedMethods: utils.GetEnvAsSlice("CORS_ALLOWED_METHODS", []string{"GET", "POST", "PUT", "DELETE", "PATCH"}, ","),
			AllowedHeaders: utils.GetEnvAsSlice("CORS_ALLOWED_HEADERS", []string{"Origin", "Content-Type", "Authorization"}, ","),
		},
		App: AppEnv{
			Environment: getEnv("APP_ENV", "development"),
			Debug:       AppDebug,
		},
	}
	if cfg.App.Environment == "production" {
		cfg.App.Debug = false
	}

	return cfg, nil
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
