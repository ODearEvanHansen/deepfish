package api

import (
	"os"
	"path/filepath"
	"testing"
)

func setupTestConfig(t *testing.T) string {
	dir := t.TempDir()
	configPath := filepath.Join(dir, "phishing.json")
	configContent := `{
		"prompts": {
			"zh": {
				"system": "Test system prompt",
				"user": ""
			}
		}
	}`
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to create test config: %v", err)
	}
	
	// Verify file was created
	if _, err := os.Stat(configPath); err != nil {
		t.Fatalf("Test config file not found at %s: %v", configPath, err)
	}
	t.Logf("Created test config at: %s", configPath)
	return configPath
}

func TestMain(m *testing.M) {
	// Setup test environment
	os.Setenv("DEEPSEEK_API_KEY", "test_key")
	os.Setenv("PHISHING_PROMPTS_PATH", "testdata/phishing.json")
	
	// Run tests
	code := m.Run()
	os.Exit(code)
}

func TestDeepSeekClient_GenerateChineseEmail(t *testing.T) {
	// Skip test if API key is not set
	apiKey := os.Getenv("DEEPSEEK_API_KEY")
	if apiKey == "" {
		t.Skip("DEEPSEEK_API_KEY environment variable not set, skipping test")
	}

	// Setup test config
	configPath := setupTestConfig(t)
	os.Setenv("DEEPSEEK_API_KEY", apiKey)
	os.Setenv("PHISHING_PROMPTS_PATH", configPath)
	
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