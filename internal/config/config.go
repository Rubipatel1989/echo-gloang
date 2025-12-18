package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	// Server
	Port string
	Env  string

	// Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBCharset  string

	// JWT
	JWTSecret            string
	JWTExpiration        time.Duration
	JWTRefreshExpiration time.Duration

	// File Upload
	UploadDir     string
	MaxUploadSize int64

	// CORS
	CORSAllowedOrigins []string
}

var AppConfig *Config

func LoadConfig() error {
	// Load .env file if it exists
	_ = godotenv.Load()

	AppConfig = &Config{
		Port: getEnv("PORT", "8080"),
		Env:  getEnv("ENV", "development"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "new_password"),
		DBName:     getEnv("DB_NAME", "basketball_db"),
		DBCharset:  getEnv("DB_CHARSET", "utf8mb4"),

		JWTSecret:            getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
		JWTExpiration:        parseDuration(getEnv("JWT_EXPIRATION", "15m")),
		JWTRefreshExpiration: parseDuration(getEnv("JWT_REFRESH_EXPIRATION", "168h")),

		UploadDir:     getEnv("UPLOAD_DIR", "./uploads"),
		MaxUploadSize: parseInt64(getEnv("MAX_UPLOAD_SIZE", "10485760")), // 10MB

		CORSAllowedOrigins: getEnvSlice("CORS_ALLOWED_ORIGINS", []string{"*"}),
	}

	return nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvSlice(key string, defaultValue []string) []string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	// Simple split by comma
	result := []string{}
	for _, v := range splitString(value, ",") {
		if v != "" {
			result = append(result, v)
		}
	}
	return result
}

func splitString(s, sep string) []string {
	result := []string{}
	current := ""
	for _, char := range s {
		if string(char) == sep {
			if current != "" {
				result = append(result, current)
				current = ""
			}
		} else {
			current += string(char)
		}
	}
	if current != "" {
		result = append(result, current)
	}
	return result
}

func parseInt64(s string) int64 {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 10485760 // default 10MB
	}
	return val
}

func parseDuration(s string) time.Duration {
	duration, err := time.ParseDuration(s)
	if err != nil {
		return 15 * time.Minute // default 15 minutes
	}
	return duration
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
		c.DBCharset,
	)
}
