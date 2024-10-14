package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	PublicHost              string
	Port                    int
	CookiesAuthSecret       string
	CookiesAuthAgeInSeconds int
	CookiesAuthIsSecure     bool
	CookiesAuthIsHttpOnly   bool
	AzureADClientID         string
	AzureADClientSecret     string
	AzureADTenantID         string
	DSN                     string
}

const (
	twoDaysInSeconds = 60 * 60 * 24 * 2
)

var Config AppConfig

func init() {
	envPath := ".env"
	// Load the .env file
	err := godotenv.Load(envPath)
	if err != nil {
		fmt.Println("Error loading .env file")
	} else {
		fmt.Println("Loaded .env file")
	}

	Config = initConfig()
}

func initConfig() AppConfig {
	return AppConfig{
		PublicHost:              getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                    getEnvAsInt("PORT", 8080),
		CookiesAuthSecret:       getEnv("COOKIES_AUTH_SECRET", "some-secret-key"),
		CookiesAuthAgeInSeconds: getEnvAsInt("COOKIES_AUTH_AGE_IN_SECONDS", twoDaysInSeconds),
		CookiesAuthIsSecure:     getEnvAsBool("COOKIES_AUTH_IS_SECURE", false),
		CookiesAuthIsHttpOnly:   getEnvAsBool("COOKIES_AUTH_IS_HTTP_ONLY", false),
		AzureADClientID:         getEnvOrError("AZURE_AD_CLIENT_ID"),
		AzureADClientSecret:     getEnvOrError("AZURE_AD_CLIENT_SECRET"),
		AzureADTenantID:         getEnvOrError("AZURE_AD_TENANT_ID"),
		DSN:                     getEnvOrError("DATABASE_URL"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvOrError(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	// should server panic or not?
	// I thing it should since the environment variable is required to run the application
	panic(fmt.Sprintf("Environment variable %s is not set", key))
}

func getEnvAsInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}

func getEnvAsBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		b, err := strconv.ParseBool(value)
		if err != nil {
			return fallback
		}

		return b
	}

	return fallback
}
