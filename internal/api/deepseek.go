package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ODearEvanHansen/deepfish/internal/config"
)

// [Previous type definitions remain unchanged...]

// GenerateChineseEmail generates a fishing email in Chinese
func (c *DeepSeekClient) GenerateChineseEmail(prompt string) (string, error) {
	if c.apiKey == "" {
		return "", errors.New("DeepSeek API key is not set")
	}

	prompts, err := config.LoadPhishingPrompts("config/prompts/phishing.json")
	if err != nil {
		return "", fmt.Errorf("failed to load prompts: %w", err)
	}

	zhPrompt, exists := prompts.Prompts["zh"]
	if !exists {
		return "", errors.New("Chinese prompt not found in configuration")
	}

	messages := []Message{
		{
			Role:    "system",
			Content: zhPrompt.System,
		},
		{
			Role:    "user",
			Content: prompt,
		},
	}

	// [Rest of the function implementation remains unchanged...]
}