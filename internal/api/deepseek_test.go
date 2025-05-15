package api

import (
	"os"
	"path/filepath"
	"strings"
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
	// Configure test environment
	os.Setenv("DEEPSEEK_API_KEY", "test_key_ci")
	os.Setenv("PHISHING_PROMPTS_PATH", "testdata/phishing.json")
	
	// Start mock server in all environments
	srv := startMockAPIServer()
	defer srv.Close()
	os.Setenv("DEEPSEEK_BASE_URL", srv.URL)

	// Run tests
	code := m.Run()
	os.Exit(code)
}

func TestDeepSeekClient_GenerateChineseEmail(t *testing.T) {
	t.Log("=== Starting GenerateChineseEmail test ===")
	defer t.Log("=== GenerateChineseEmail test completed ===")

	// CI Environment detection
	if os.Getenv("CI") == "true" {
		t.Log("Running in CI environment - enabling CI-specific configurations")
		os.Setenv("TEST_IN_CI", "true")
		t.Logf("CI Environment Variables: DEEPSEEK_BASE_URL=%s", os.Getenv("DEEPSEEK_BASE_URL"))
	}

	// Skip test if running in CI without mock mode
	if os.Getenv("CI") == "true" && os.Getenv("TEST_MOCK_MODE") != "true" {
		t.Skip("Skipping live API test in CI environment")
	}

	// Setup and validate test environment
	configPath := setupTestConfig(t)
	t.Logf("Using config file: %s", configPath)
	if _, err := os.Stat(configPath); err != nil {
		t.Fatalf("Config file not accessible in test environment: %v", err)
	}
	t.Logf("Using config file at: %s", configPath)

	// Set required environment variables
	apiKey := os.Getenv("DEEPSEEK_API_KEY")
	if apiKey == "" {
		apiKey = "test_key_ci"
	}
	os.Setenv("DEEPSEEK_API_KEY", apiKey)
	os.Setenv("PHISHING_PROMPTS_PATH", configPath)
	
	// Create client with test config path
	client := NewDeepSeekClient()
	client.configPath = configPath

	// Test generating content (mock mode)
	prompt := "Generate a short test message in Chinese (less than 50 characters)"
	content, err := client.GenerateChineseEmail(prompt)
	if err != nil {
		t.Fatalf("Failed to generate content: %v", err)
	}

	// Verify mock response pattern
	expectedPrefix := "测试中文内容 - "
	if !strings.HasPrefix(content, expectedPrefix) {
		t.Errorf("Generated content does not match mock pattern")
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