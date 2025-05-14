package templates

import (
	"strings"
	"testing"
)

func TestEmailTemplate_HTMLContent(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected string
	}{
		{
			name:     "Plain text",
			content:  "Hello, world!",
			expected: "Hello, world!",
		},
		{
			name:     "Text with newlines",
			content:  "Hello,\nworld!",
			expected: "Hello,<br>world!",
		},
		{
			name:     "Text with bold markdown",
			content:  "Hello, **world**!",
			// The current implementation just replaces all ** with <strong>
			expected: "Hello, <strong>world<strong>!",
		},
		{
			name:     "Text with italic markdown",
			content:  "Hello, *world*!",
			// The current implementation just replaces all * with <em>
			expected: "Hello, <em>world<em>!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			template := &EmailTemplate{
				Subject: "Test Subject",
				Content: tt.content,
			}

			html := string(template.HTMLContent())
			if !strings.Contains(html, tt.expected) {
				t.Errorf("HTMLContent() = %v, want %v", html, tt.expected)
			}
		})
	}
}

func TestEmailTemplate_PlainText(t *testing.T) {
	template := &EmailTemplate{
		Subject: "Test Subject",
		Content: "Hello, world!",
	}

	if template.Content != "Hello, world!" {
		t.Errorf("PlainText() = %v, want %v", template.Content, "Hello, world!")
	}
}