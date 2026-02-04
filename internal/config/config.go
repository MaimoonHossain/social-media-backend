package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	JWT      JWTConfig
	Upload   UploadConfig
	RateLimit RateLimitConfig
	Redis    RedisConfig
	AWS      AWSConfig
	Email    EmailConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type ServerConfig struct {
	Port        string
	Environment string
	FrontendURL string
}

type JWTConfig struct {
	Secret string
	Expiry time.Duration
}

type UploadConfig struct {
	MaxSize           int64
	Path              string
	AllowedExtensions []string
}

type RateLimitConfig struct {
	Requests int
	Duration time.Duration
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type AWSConfig struct {
	Region    string
	Bucket    string
	AccessKey string
	SecretKey string
}

type EmailConfig struct {
	SMTPHost string
	SMTPPort string
	Username string
	Password string
	From     string
}

var AppConfig *Config

func Load() (*Config, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	jwtExpiry, err := time.ParseDuration(getEnv("JWT_EXPIRY", "24h"))
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_EXPIRY: %w", err)
	}

	rateLimitDuration, err := time.ParseDuration(getEnv("RATE_LIMIT_DURATION", "1m"))
	if err != nil {
		return nil, fmt.Errorf("invalid RATE_LIMIT_DURATION: %w", err)
	}

	maxUploadSize, err := strconv.ParseInt(getEnv("MAX_UPLOAD_SIZE", "10485760"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid MAX_UPLOAD_SIZE: %w", err)
	}

	rateLimitRequests, err := strconv.Atoi(getEnv("RATE_LIMIT_REQUESTS", "100"))
	if err != nil {
		return nil, fmt.Errorf("invalid RATE_LIMIT_REQUESTS: %w", err)
	}

	redisDB, err := strconv.Atoi(getEnv("REDIS_DB", "0"))
	if err != nil {
		return nil, fmt.Errorf("invalid REDIS_DB: %w", err)
	}

	config := &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "social_media_db"),
		},
		Server: ServerConfig{
			Port:        getEnv("PORT", "8080"),
			Environment: getEnv("APP_ENV", "development"),
			FrontendURL: getEnv("FRONTEND_URL", "http://localhost:3000"),
		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", "change-this-secret-key"),
			Expiry: jwtExpiry,
		},
		Upload: UploadConfig{
			MaxSize:           maxUploadSize,
			Path:              getEnv("UPLOAD_PATH", "./storage/uploads"),
			AllowedExtensions: parseExtensions(getEnv("ALLOWED_EXTENSIONS", "jpg,jpeg,png,gif,mp4,mov")),
		},
		RateLimit: RateLimitConfig{
			Requests: rateLimitRequests,
			Duration: rateLimitDuration,
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       redisDB,
		},
		AWS: AWSConfig{
			Region:    getEnv("AWS_REGION", "us-east-1"),
			Bucket:    getEnv("AWS_BUCKET", ""),
			AccessKey: getEnv("AWS_ACCESS_KEY", ""),
			SecretKey: getEnv("AWS_SECRET_KEY", ""),
		},
		Email: EmailConfig{
			SMTPHost: getEnv("SMTP_HOST", "smtp.gmail.com"),
			SMTPPort: getEnv("SMTP_PORT", "587"),
			Username: getEnv("SMTP_USERNAME", ""),
			Password: getEnv("SMTP_PASSWORD", ""),
			From:     getEnv("FROM_EMAIL", "noreply@socialmedia.com"),
		},
	}

	AppConfig = config
	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func parseExtensions(extensions string) []string {
	var result []string
	for i := 0; i < len(extensions); {
		end := i
		for end < len(extensions) && extensions[end] != ',' {
			end++
		}
		if i < end {
			result = append(result, extensions[i:end])
		}
		i = end + 1
	}
	return result
}