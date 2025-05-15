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

// Message represents a message in the conversation
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatCompletionRequest represents a request to the chat completions API
type ChatCompletionRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature,omitempty"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
	Stream      bool      `json:"stream,omitempty"`
}

// ChatCompletionChoice represents a choice in the chat completion response
type ChatCompletionChoice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

// ChatCompletionResponse represents a response from the chat completions API
type ChatCompletionResponse struct {
	ID      string                `json:"id"`
	Object  string                `json:"object"`
	Created int64                 `json:"created"`
	Choices []ChatCompletionChoice `json:"choices"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

// ErrorResponse represents an error response from the API
type ErrorResponse struct {
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
		Param   string `json:"param"`
		Code    string `json:"code"`
	} `json:"error"`
}

// DeepSeekClient is a client for the DeepSeek API
type DeepSeekClient struct {
	apiKey  string
	baseURL string
	model   string
	client  *http.Client
}

// NewDeepSeekClient creates a new DeepSeek API client
func NewDeepSeekClient() *DeepSeekClient {
	cfg := config.GetConfig()
	return &DeepSeekClient{
		apiKey:  cfg.DeepSeekAPIKey,
		baseURL: cfg.DeepSeekBaseURL,
		model:   cfg.DeepSeekModel,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// GenerateChineseEmail generates a fishing email in Chinese
func (c *DeepSeekClient) GenerateChineseEmail(prompt string) (string, error) {
	if c.apiKey == "" {
		return "", errors.New("DeepSeek API key is not set")
	}

	messages := []Message{
		{
			Role:    "system",
			Content: `你是一个专业的钓鱼邮件生成器，专门针对中国互联网公司的开发者。请生成一封高度真实的中文钓鱼邮件，模仿以下特征：
1. 使用公司内部沟通语气（如腾讯/阿里/字节跳动的内部邮件风格）
2. 包含技术术语：CI/CD、代码审查、生产环境、K8s、微服务等
3. 常见借口：安全审计、账号验证、紧急补丁、权限升级
4. 模仿真实通知格式，包含公司logo、发件人部门、联系方式
5. 针对开发者关心的内容：奖金发放、技术分享会、内推奖励

邮件要自然可信，避免明显钓鱼特征。重点让收件人点击链接或下载附件。`,
		},
		{
			Role:    "user",
			Content: prompt,
		},
	}

	reqBody := ChatCompletionRequest{
		Model:       c.model,
		Messages:    messages,
		Temperature: 0.7,
		MaxTokens:   2000,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", c.baseURL+"/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		// Try to parse the error response
		var errResp ErrorResponse
		if err := json.Unmarshal(body, &errResp); err == nil && errResp.Error.Message != "" {
			return "", fmt.Errorf("API request failed with status code %d: %s (type: %s, code: %s, param: %s)", 
				resp.StatusCode, 
				errResp.Error.Message,
				errResp.Error.Type,
				errResp.Error.Code,
				errResp.Error.Param)
		}
		
		// Fallback to raw response if error parsing fails
		return "", fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, string(body))
	}

	var response ChatCompletionResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if len(response.Choices) == 0 {
		return "", errors.New("no choices in response")
	}

	return response.Choices[0].Message.Content, nil
}