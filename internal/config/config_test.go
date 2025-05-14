package config

import (
	"os"
	"sync"
	"testing"
)

func TestGetConfig(t *testing.T) {
	// Save original environment variables
	origAPIKey := os.Getenv("DEEPSEEK_API_KEY")
	origBaseURL := os.Getenv("DEEPSEEK_BASE_URL")
	origModel := os.Getenv("DEEPSEEK_MODEL")

	// Restore environment variables after test
	defer func() {
		os.Setenv("DEEPSEEK_API_KEY", origAPIKey)
		os.Setenv("DEEPSEEK_BASE_URL", origBaseURL)
		os.Setenv("DEEPSEEK_MODEL", origModel)
	}()

	// Reset the singleton instance for testing
	instance = nil
	once = sync.Once{}

	// Test with all environment variables set
	os.Setenv("DEEPSEEK_API_KEY", "test-api-key")
	os.Setenv("DEEPSEEK_BASE_URL", "https://test-api.example.com")
	os.Setenv("DEEPSEEK_MODEL", "test-model")

	cfg := GetConfig()
	if cfg.DeepSeekAPIKey != "test-api-key" {
		t.Errorf("Expected API key to be 'test-api-key', got '%s'", cfg.DeepSeekAPIKey)
	}
	if cfg.DeepSeekBaseURL != "https://test-api.example.com" {
		t.Errorf("Expected base URL to be 'https://test-api.example.com', got '%s'", cfg.DeepSeekBaseURL)
	}
	if cfg.DeepSeekModel != "test-model" {
		t.Errorf("Expected model to be 'test-model', got '%s'", cfg.DeepSeekModel)
	}

	// Reset the singleton instance for testing
	instance = nil
	once = sync.Once{}

	// Test with only API key set (should use defaults for other values)
	os.Unsetenv("DEEPSEEK_BASE_URL")
	os.Unsetenv("DEEPSEEK_MODEL")
	os.Setenv("DEEPSEEK_API_KEY", "test-api-key")

	cfg = GetConfig()
	if cfg.DeepSeekAPIKey != "test-api-key" {
		t.Errorf("Expected API key to be 'test-api-key', got '%s'", cfg.DeepSeekAPIKey)
	}
	if cfg.DeepSeekBaseURL != "https://api.deepseek.com/v1" {
		t.Errorf("Expected base URL to be '%s', got '%s'", "https://api.deepseek.com/v1", cfg.DeepSeekBaseURL)
	}
	if cfg.DeepSeekModel != "deepseek-chat" {
		t.Errorf("Expected model to be '%s', got '%s'", "deepseek-chat", cfg.DeepSeekModel)
	}
}