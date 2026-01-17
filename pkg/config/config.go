package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Config holds all configuration for the application
type Config struct {
	// Server
	Port   string
	Env    string
	HostURL string
	DataPath string

	// Database
	MongoDBURI string
	DBName     string

	// JWT
	AccessTokenSecret  string
	AccessTokenExpiry  string
	RefreshTokenSecret string
	RefreshTokenExpiry string

	// CORS
	CORSOrigin []string

	// Rate Limiting
	RateLimitWindowMS      int
	RateLimitMaxRequests   int

	// Redis
	RedisURI string

	// Email
	SMTPHost     string
	SMTPPort     int
	SMTPUser     string
	SMTPPassword string
	SMTPFrom     string

	// File Upload
	MaxFileSize int64
	UploadDir   string

	// Payment
	RazorpayKeyID     string
	RazorpayKeySecret string
	StripeSecretKey   string
	StripePublishableKey string
	StripeWebhookSecret string
	PayPalClientID    string
	PayPalClientSecret string
	PayPalMode        string

	// OAuth
	GoogleClientID     string
	GoogleClientSecret string
	GoogleCallbackURL  string
	GitHubClientID     string
	GitHubClientSecret string
	GitHubCallbackURL  string

	// Logging
	LogLevel  string
	LogFormat string
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	cfg := &Config{
		Port:    getEnv("PORT", "8080"),
		Env:     getEnv("NODE_ENV", "development"),
		HostURL: getEnv("FREEAPI_HOST_URL", "http://localhost:8080"),
		DataPath: getEnv("DATA_PATH", "./data"),

		MongoDBURI: getEnv("MONGODB_URI", "mongodb://localhost:27017/apihub_go"),
		DBName:     getEnv("DB_NAME", "apihub_go"),

		AccessTokenSecret:  getEnv("ACCESS_TOKEN_SECRET", ""),
		AccessTokenExpiry:  getEnv("ACCESS_TOKEN_EXPIRY", "1d"),
		RefreshTokenSecret: getEnv("REFRESH_TOKEN_SECRET", ""),
		RefreshTokenExpiry: getEnv("REFRESH_TOKEN_EXPIRY", "10d"),

		CORSOrigin: strings.Split(getEnv("CORS_ORIGIN", "*"), ","),

		RateLimitWindowMS:    getEnvAsInt("RATE_LIMIT_WINDOW_MS", 900000),
		RateLimitMaxRequests: getEnvAsInt("RATE_LIMIT_MAX_REQUESTS", 5000),

		RedisURI: getEnv("REDIS_URI", ""),

		SMTPHost:     getEnv("SMTP_HOST", ""),
		SMTPPort:     getEnvAsInt("SMTP_PORT", 587),
		SMTPUser:     getEnv("SMTP_USER", ""),
		SMTPPassword: getEnv("SMTP_PASSWORD", ""),
		SMTPFrom:     getEnv("SMTP_FROM", "noreply@apihub.com"),

		MaxFileSize: getEnvAsInt64("MAX_FILE_SIZE", 5242880),
		UploadDir:   getEnv("UPLOAD_DIR", "./public/images"),

		RazorpayKeyID:     getEnv("RAZORPAY_KEY_ID", ""),
		RazorpayKeySecret: getEnv("RAZORPAY_KEY_SECRET", ""),
		StripeSecretKey:   getEnv("STRIPE_SECRET_KEY", ""),
		StripePublishableKey: getEnv("STRIPE_PUBLISHABLE_KEY", ""),
		StripeWebhookSecret: getEnv("STRIPE_WEBHOOK_SECRET", ""),
		PayPalClientID:    getEnv("PAYPAL_CLIENT_ID", ""),
		PayPalClientSecret: getEnv("PAYPAL_CLIENT_SECRET", ""),
		PayPalMode:        getEnv("PAYPAL_MODE", "sandbox"),

		GoogleClientID:     getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
		GoogleCallbackURL:  getEnv("GOOGLE_CALLBACK_URL", ""),
		GitHubClientID:     getEnv("GITHUB_CLIENT_ID", ""),
		GitHubClientSecret: getEnv("GITHUB_CLIENT_SECRET", ""),
		GitHubCallbackURL:  getEnv("GITHUB_CALLBACK_URL", ""),

		LogLevel:  getEnv("LOG_LEVEL", "debug"),
		LogFormat: getEnv("LOG_FORMAT", "json"),
	}

	// Validate required fields
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

// Validate checks if required configuration fields are set
func (c *Config) Validate() error {
	if c.AccessTokenSecret == "" {
		return fmt.Errorf("ACCESS_TOKEN_SECRET is required")
	}
	if c.RefreshTokenSecret == "" {
		return fmt.Errorf("REFRESH_TOKEN_SECRET is required")
	}
	if c.MongoDBURI == "" {
		return fmt.Errorf("MONGODB_URI is required")
	}
	return nil
}

// IsDevelopment returns true if running in development mode
func (c *Config) IsDevelopment() bool {
	return c.Env == "development"
}

// IsProduction returns true if running in production mode
func (c *Config) IsProduction() bool {
	return c.Env == "production"
}

// Helper functions
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsInt64(key string, defaultValue int64) int64 {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
		return value
	}
	return defaultValue
}
