package api

import (
	"os"
	"testing"
)

func TestDeepSeekClient_GenerateChineseEmail(t *testing.T) {
	// Skip test if API key is not set
	apiKey := os.Getenv("DEEPSEEK_API_KEY")
	if apiKey == "" {
		t.Skip("DEEPSEEK_API_KEY environment variable not set, skipping test")
	}

	// Set the API key in the environment for the config package
	os.Setenv("DEEPSEEK_API_KEY", apiKey)
	
	// Create a new client
	client := NewDeepSeekClient()

	// Test generating content
	prompt := "Generate a short test message in Chinese (less than 50 characters)"
	content, err := client.GenerateChineseEmail(prompt)
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