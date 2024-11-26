package config

import (
	"os"
	"strconv"
	"sync"

	"github.com/flastors/songius/pkg/client/postgresql"
	"github.com/flastors/songius/pkg/utils/logging"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		logger := logging.GetLogger()
		logger.Fatal("Error loading .env file")
	}
}

type Config struct {
	Debug       bool
	Http        HttpServerConfig
	Storage     postgresql.StorageConfig
	ExternalAPI ExternalAPIConfig
}

type HttpServerConfig struct {
	Host string
	Port string
}

type ExternalAPIConfig struct {
	Url string
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("Application configuration")
		instance = &Config{
			Debug: getEnvAsBool("DEBUG", true),
			Http: HttpServerConfig{
				Host: getEnv("HTTP_HOST", "localhost"),
				Port: getEnv("HTTP_PORT", "8080"),
			},
			Storage: postgresql.StorageConfig{
				Host:     getEnv("POSTGRES_HOST", "localhost"),
				Port:     getEnv("POSTGRES_PORT", "5432"),
				Username: getEnv("POSTGRES_USER", "postgres"),
				Password: getEnv("POSTGRES_PASSWORD", "postgres"),
				Database: getEnv("POSTGRES_DB", "postgres"),
			},
			ExternalAPI: ExternalAPIConfig{
				Url: getEnv("EXTERNAL_API_URL", "www.google.com"),
			},
		}
	})
	return instance

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
	if value, err := strconv.ParseBool(valStr); err == nil {
		return value
	}

	return defaultVal
}
