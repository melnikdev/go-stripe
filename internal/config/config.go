package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Server *Server
	DB     *Database
	Stripe *Stripe
}

type Stripe struct {
	SecretKey string
}

type Server struct {
	Port int
}

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
	Timezone string
}

func NewConfig() *Config {
	return &Config{
		Server: &Server{
			Port: getEnvAsInt("SERVER_PORT", 8080),
		},
		DB: &Database{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvAsInt("DB_PORT", 5432),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "stripe"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
			Timezone: getEnv("DB_TIMEZONE", "Europe/Berlin"),
		},
		Stripe: &Stripe{
			SecretKey: getEnv("STRIPE_API_KEY", ""),
		},
	}
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		c.DB.Host,
		c.DB.User,
		c.DB.Password,
		c.DB.DBName,
		c.DB.Port,
		c.DB.SSLMode,
		c.DB.Timezone,
	)
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}
