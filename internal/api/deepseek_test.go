package api

import (
	"os"
	"testing"

	"github.com/ODearEvanHansen/deepfish/internal/config"
)

func TestDeepSeekClient_GenerateContent(t *testing.T) {
	// Skip test if API key is not set
	apiKey := os.Getenv("DEEPSEEK_API_KEY")
	if apiKey == "" {
		t.Skip("DEEPSEEK_API_KEY environment variable not set, skipping test")
	}

	// Create a new client
	cfg := &config.Config{
		DeepSeekAPIKey: apiKey,
		DeepSeekModel:  "deepseek-chat", // Use default model for testing
	}
	client := NewDeepSeekClient(cfg)

	// Test generating content
	prompt := "Generate a short test message in Chinese (less than 50 characters)"
	content, err := client.GenerateContent(prompt)
	if err != nil {
		t.Fatalf("Failed to generate content: %v", err)
	}

	// Verify content is not empty
	if content == "" {
		t.Error("Generated content is empty")
	}

	// Verify content contains Chinese characters
	hasChineseChar := false
	for _, r := range content {
		if r >= '\u4e00' && r <= '\u9fff' {
			hasChineseChar = true
			break
		}
	}
	if !hasChineseChar {
		t.Error("Generated content does not contain Chinese characters")
	}
}