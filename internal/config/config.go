package config

import (
	"os"
	"sync"
)

// Config holds the application configuration
type Config struct {
	DeepSeekAPIKey string
	DeepSeekBaseURL string
}

var (
	instance *Config
	once     sync.Once
)

// GetConfig returns the singleton instance of Config
func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{
			DeepSeekAPIKey: getEnv("DEEPSEEK_API_KEY", ""),
			DeepSeekBaseURL: getEnv("DEEPSEEK_BASE_URL", "https://api.deepseek.com"),
		}
	})
	return instance
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}