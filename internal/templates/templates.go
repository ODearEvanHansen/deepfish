package templates

import (
	"bytes"
	"html/template"
	"strings"
)

// EmailTemplate represents a template for generating emails
type EmailTemplate struct {
	Subject string
	Content string
}

// HTMLContent returns the content as template.HTML for safe rendering in HTML templates
// This should only be used with content from trusted sources
func (e *EmailTemplate) HTMLContent() template.HTML {
	// Replace newlines with <br> tags for proper HTML rendering
	htmlContent := strings.ReplaceAll(e.Content, "\n", "<br>")
	return template.HTML(htmlContent)
}

// RenderHTML renders the email template as HTML
func (e *EmailTemplate) RenderHTML() (string, error) {
	const htmlTemplate = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>{{.Subject}}</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            margin: 0;
            padding: 20px;
            color: #333;
        }
        .container {
            max-width: 600px;
            margin: 0 auto;
            border: 1px solid #ddd;
            border-radius: 5px;
            padding: 20px;
        }
        .header {
            border-bottom: 1px solid #eee;
            padding-bottom: 10px;
            margin-bottom: 20px;
        }
        .footer {
            margin-top: 30px;
            padding-top: 10px;
            border-top: 1px solid #eee;
            font-size: 12px;
            color: #777;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h2>{{.Subject}}</h2>
        </div>
        <div class="content">
            {{.HTMLContent}}
        </div>
        <div class="footer">
            <p>此邮件由AI自动生成，仅用于演示目的。</p>
        </div>
    </div>
</body>
</html>
`

	tmpl, err := template.New("email").Parse(htmlTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, e); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// RenderText renders the email template as plain text
func (e *EmailTemplate) RenderText() string {
	return e.Content
}