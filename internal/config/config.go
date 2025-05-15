package config

import (
	"encoding/json"
	"os"
	"sync"
)

// Config holds the application configuration
type Config struct {
	DeepSeekAPIKey string
	DeepSeekBaseURL string
	DeepSeekModel  string
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
			DeepSeekBaseURL: getEnv("DEEPSEEK_BASE_URL", "https://api.deepseek.com/v1"),
			DeepSeekModel:  getEnv("DEEPSEEK_MODEL", "deepseek-chat"),
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

// PhishingPrompts holds the structured phishing email prompts
type PhishingPrompts struct {
	Prompts map[string]struct {
		System string `json:"system"`
		User   string `json:"user"`
	} `json:"prompts"`
}

// LoadPhishingPrompts loads phishing prompts from a JSON file
func LoadPhishingPrompts(path string) (*PhishingPrompts, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var prompts PhishingPrompts
	if err := json.Unmarshal(data, &prompts); err != nil {
		return nil, err
	}

	return &prompts, nil
}