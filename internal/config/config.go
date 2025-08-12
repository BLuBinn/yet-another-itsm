package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"yet-another-itsm/internal/constants"

	"github.com/MicahParks/keyfunc"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"golang.org/x/oauth2"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Logger   LoggerConfig
	OAuth    OAuthConfig
}

type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	MaxConns int32
	MinConns int32
}

type LoggerConfig struct {
	Level  string
	Format string // "json" or "console"
}

type OAuthConfig struct {
	EntraConfig  *oauth2.Config
	JWKSEntra    *keyfunc.JWKS
	ClientID     string
	ClientSecret string
	TenantID     string
	RedirectURI  string
	GraphScope   string
}

func Load() (*Config, error) {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Warn().Err(err).Msg(constants.ErrNoEnvFileFoundMsg)
	}

	config := &Config{
		Server: ServerConfig{
			Port:         getEnv("PORT", "8080"),
			ReadTimeout:  getDurationEnv("READ_TIMEOUT", 10*time.Second),
			WriteTimeout: getDurationEnv("WRITE_TIMEOUT", 10*time.Second),
			IdleTimeout:  getDurationEnv("IDLE_TIMEOUT", 120*time.Second),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "msn_map_api"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
			MaxConns: int32(getIntEnv("DB_MAX_CONNS", 25)),
			MinConns: int32(getIntEnv("DB_MIN_CONNS", 5)),
		},
		Logger: LoggerConfig{
			Level:  getEnv("LOG_LEVEL", "info"),
			Format: getEnv("LOG_FORMAT", "json"),
		},
		OAuth: OAuthConfig{
			ClientID:     getEnv("ENTRA_CLIENT_ID", ""),
			ClientSecret: getEnv("ENTRA_CLIENT_SECRET", ""),
			TenantID:     getEnv("ENTRA_TENANT_ID", ""),
			RedirectURI:  getEnv("REDIRECT_URI", "http://localhost:8080/callback"),
			GraphScope:   getEnv("APPLICATION_GRAPH_API_SCOPE", "https://graph.microsoft.com/.default"),
		},
	}

	// Configure zerolog
	configureLogger(config.Logger)

	if err := initOAuth(&config.OAuth); err != nil {
		return nil, fmt.Errorf(constants.ErrFailedToInitializeOAuthMsg, err)
	}

	return config, nil
}

func configureLogger(cfg LoggerConfig) {
	// Set log level
	level, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		log.Warn().Str("level", cfg.Level).Msg(constants.ErrInvalidLogLevelMsg)
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)

	// Set log format
	if cfg.Format == "console" {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		})
	}

	// Add caller information in development
	if cfg.Level == "debug" {
		log.Logger = log.Logger.With().Caller().Logger()
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getIntEnv(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}
