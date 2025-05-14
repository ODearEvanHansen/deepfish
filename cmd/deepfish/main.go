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

	// Extract subject from content (assuming first line is the subject)
	lines := strings.Split(content, "\n")
	subject := lines[0]
	if strings.HasPrefix(strings.ToLower(subject), "subject:") {
		subject = strings.TrimSpace(strings.TrimPrefix(subject, "Subject:"))
	} else if strings.HasPrefix(strings.ToLower(subject), "主题:") {
		subject = strings.TrimSpace(strings.TrimPrefix(subject, "主题:"))
	}

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