package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ODearEvanHansen/deepfish/internal/api"
	"github.com/ODearEvanHansen/deepfish/internal/templates"
)

func main() {
	// Define command-line flags
	apiKey := flag.String("api-key", "", "DeepSeek API key")
	outputFile := flag.String("output", "", "Output file path (optional)")
	format := flag.String("format", "text", "Output format: text or html")
	prompt := flag.String("prompt", "", "Prompt for generating the email")
	flag.Parse()

	// Check if API key is provided via flag or environment variable
	if *apiKey != "" {
		os.Setenv("DEEPSEEK_API_KEY", *apiKey)
	} else if os.Getenv("DEEPSEEK_API_KEY") == "" {
		fmt.Println("Error: DeepSeek API key is required. Provide it via -api-key flag or DEEPSEEK_API_KEY environment variable.")
		os.Exit(1)
	}

	// Check if prompt is provided
	if *prompt == "" {
		// If no prompt is provided via flag, check if there are remaining arguments
		args := flag.Args()
		if len(args) > 0 {
			*prompt = strings.Join(args, " ")
		} else {
			fmt.Println("Error: Prompt is required. Provide it via -prompt flag or as arguments.")
			os.Exit(1)
		}
	}

	// Create DeepSeek client
	client := api.NewDeepSeekClient()

	// Generate email content
	fmt.Println("Generating email content...")
	content, err := client.GenerateChineseEmail(*prompt)
	if err != nil {
		fmt.Printf("Error generating email: %v\n", err)
		os.Exit(1)
	}

	// Extract subject from content
	lines := strings.Split(content, "\n")
	subject := "Generated Email" // Default subject if none found
	
	// Check for subject line with various patterns
	subjectPrefixes := []string{
		"subject:", "主题:", "标题:", "邮件主题:", "email subject:", 
		"subject：", "主题：", "标题：", "邮件主题：", "email subject：", // Chinese colon
		"re:", "fw:", "fwd:", "回复:", "转发:",  // Common email prefixes
		"【", "[", // Common subject delimiters in Chinese emails
	}
	
	// First pass: Look for explicit subject markers
	subjectFound := false
	for i, line := range lines {
		if i > 10 { // Check more lines (10 instead of 5)
			break
		}
		
		lineLower := strings.ToLower(strings.TrimSpace(line))
		if lineLower == "" {
			continue // Skip empty lines
		}
		
		for _, prefix := range subjectPrefixes {
			if strings.HasPrefix(lineLower, strings.ToLower(prefix)) {
				// Extract the subject text after the prefix
				prefixLen := len(prefix)
				if len(line) >= prefixLen {
					actualPrefix := line[:prefixLen] // Get the actual case of the prefix
					subject = strings.TrimSpace(strings.TrimPrefix(line, actualPrefix))
					
					// Remove the subject line from the content
					contentLines := strings.Split(content, "\n")
					content = strings.Join(append(contentLines[:i], contentLines[i+1:]...), "\n")
					subjectFound = true
					break
				}
			}
		}
		
		// Check for patterns like "Subject: text" anywhere in the line
		if !subjectFound {
			for _, prefix := range subjectPrefixes {
				prefixPattern := strings.ToLower(prefix)
				if idx := strings.Index(lineLower, prefixPattern); idx >= 0 {
					// Extract text after the prefix
					startIdx := idx + len(prefixPattern)
					if startIdx < len(line) {
						subject = strings.TrimSpace(line[startIdx:])
						// Remove the subject line from the content
						contentLines := strings.Split(content, "\n")
						content = strings.Join(append(contentLines[:i], contentLines[i+1:]...), "\n")
						subjectFound = true
						break
					}
				}
			}
		}
		
		if subjectFound {
			break
		}
	}
	
	// Second pass: If no explicit subject found, use heuristics
	if !subjectFound {
		// If the first non-empty line is short, use it as subject
		for i, line := range lines {
			trimmedLine := strings.TrimSpace(line)
			if trimmedLine != "" && len(trimmedLine) < 100 && !strings.Contains(trimmedLine, ":") {
				subject = trimmedLine
				// Remove the subject line from the content
				contentLines := strings.Split(content, "\n")
				content = strings.Join(append(contentLines[:i], contentLines[i+1:]...), "\n")
				break
			}
		}
	}
	
	// Clean up the subject
	subject = strings.TrimSpace(subject)
	// Remove common decorators like [], (), ""
	for _, decorator := range []string{"[", "]", "(", ")", "\"", "【", "】", "《", "》"} {
		subject = strings.ReplaceAll(subject, decorator, "")
	}
	subject = strings.TrimSpace(subject)

	// Create email template
	emailTemplate := &templates.EmailTemplate{
		Subject: subject,
		Content: content,
	}

	// Render email based on format
	var output string
	if *format == "html" {
		output, err = emailTemplate.RenderHTML()
		if err != nil {
			fmt.Printf("Error rendering HTML: %v\n", err)
			os.Exit(1)
		}
	} else {
		output = emailTemplate.RenderText()
	}

	// Write output to file or stdout
	if *outputFile != "" {
		err := os.WriteFile(*outputFile, []byte(output), 0644)
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Email written to %s\n", *outputFile)
	} else {
		fmt.Println("\n--- Generated Email ---")
		fmt.Println(output)
	}
}