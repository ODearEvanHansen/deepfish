# DeepFish

DeepFish is an AI-powered fishing email generator written in Go. It uses DeepSeek's API to generate convincing phishing emails in Chinese.

## Features

- Generate realistic phishing emails in Chinese
- Output in plain text or HTML format
- Save output to a file or display in the console
- Secure API key handling via environment variables

## Installation

```bash
# Clone the repository
git clone https://github.com/ODearEvanHansen/deepfish.git
cd deepfish

# Build the application
go build -o deepfish ./cmd/deepfish
```

## Usage

```bash
# Set your DeepSeek API key as an environment variable
export DEEPSEEK_API_KEY="your-api-key"

# Generate a phishing email
./deepfish -prompt "Create a phishing email pretending to be from a bank"

# Generate HTML output and save to a file
./deepfish -format html -output email.html -prompt "Create a phishing email pretending to be from a tech company"

# Provide API key via command line (not recommended for production)
./deepfish -api-key "your-api-key" -prompt "Create a phishing email pretending to be from a government agency"
```

## Security Note

Never store your API key in the source code. Always use environment variables or a secure configuration management system.

## License

This project is licensed under the Apache License 2.0 - see the LICENSE file for details.