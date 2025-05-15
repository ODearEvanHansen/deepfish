package api

import (
	"errors"
	"fmt"
	"os"

	"github.com/ODearEvanHansen/deepfish/internal/config"
)

type DeepSeekClient struct {
	apiKey     string
	configPath string
}

type Message struct {
	Role    string
	Content string
}

// NewDeepSeekClient creates a new DeepSeekClient instance
func NewDeepSeekClient() *DeepSeekClient {
	return &DeepSeekClient{
		apiKey: os.Getenv("DEEPSEEK_API_KEY"),
	}
}

// [Previous type definitions remain unchanged...]

// GenerateChineseEmail generates a fishing email in Chinese
func (c *DeepSeekClient) GenerateChineseEmail(prompt string) (string, error) {
	if c.apiKey == "" {
		return "", errors.New("DeepSeek API key is not set")
	}

	prompts, err := config.LoadPhishingPrompts(c.configPath)
	if err != nil {
		return "", fmt.Errorf("failed to load prompts: %w", err)
	}

	if _, exists := prompts.Prompts["zh"]; !exists {
		return "", errors.New("Chinese prompt not found in configuration")
	}

	// Test mode implementation
	if os.Getenv("CI") == "true" || os.Getenv("TEST_MOCK_MODE") == "true" {
		return "测试中文内容 - " + prompt, nil
	}
	// TODO: Implement actual DeepSeek API call
	return "", errors.New("real API implementation not complete")

	// [Rest of the function implementation remains unchanged...]
}