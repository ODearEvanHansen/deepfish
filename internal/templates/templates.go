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
	
	// Basic Markdown formatting
	// Bold text
	htmlContent = strings.ReplaceAll(htmlContent, "**", "<strong>")
	
	// Italic text
	htmlContent = strings.ReplaceAll(htmlContent, "*", "<em>")
	
	// Links - simple regex-based replacement
	// This is a simplified approach and might not handle all Markdown link formats
	for {
		startIdx := strings.Index(htmlContent, "[")
		if startIdx == -1 {
			break
		}
		
		endParenIdx := strings.Index(htmlContent[startIdx:], ")")
		if endParenIdx == -1 {
			break
		}
		endParenIdx += startIdx
		
		linkText := ""
		url := ""
		
		// Extract link text and URL
		linkSegment := htmlContent[startIdx:endParenIdx+1]
		bracketCloseIdx := strings.Index(linkSegment, "]")
		if bracketCloseIdx != -1 {
			linkText = linkSegment[1:bracketCloseIdx]
			parenOpenIdx := strings.Index(linkSegment, "(")
			if parenOpenIdx != -1 && parenOpenIdx > bracketCloseIdx {
				url = linkSegment[parenOpenIdx+1:len(linkSegment)-1]
			}
		}
		
		if linkText != "" && url != "" {
			// Replace the Markdown link with HTML link
			htmlLink := "<a href=\"" + url + "\">" + linkText + "</a>"
			htmlContent = htmlContent[:startIdx] + htmlLink + htmlContent[endParenIdx+1:]
		} else {
			// If we can't parse the link properly, just move past this occurrence
			htmlContent = htmlContent[:startIdx] + "[" + htmlContent[startIdx+1:]
		}
	}
	
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