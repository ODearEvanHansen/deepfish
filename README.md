# DeepFish

> **⚠️ IMPORTANT DISCLAIMER**: This tool is provided for **educational and security testing purposes only**. The generated phishing emails should only be used in controlled environments with proper authorization. Unauthorized use of phishing emails may violate laws and regulations. Users are responsible for ensuring compliance with all applicable laws and regulations.

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

## Disclaimer

This tool is provided for educational and testing purposes only. The generated phishing emails should only be used in controlled environments with proper authorization. Unauthorized use of phishing emails may violate laws and regulations. Users are responsible for ensuring compliance with all applicable laws and regulations.

## Configuration Options

You can configure the following options via environment variables:

- `DEEPSEEK_API_KEY`: Your DeepSeek API key (required)
- `DEEPSEEK_BASE_URL`: DeepSeek API base URL (default: https://api.deepseek.com/v1)
- `DEEPSEEK_MODEL`: DeepSeek model to use (default: deepseek-chat)

## Testing

DeepFish includes unit tests and integration tests to ensure functionality.

### Running Unit Tests

```bash
# Run all unit tests
go test -v ./internal/...

# Run tests for a specific package
go test -v ./internal/api
go test -v ./internal/templates
go test -v ./internal/config
```

### Running Integration Tests

```bash
# Make sure your API key is set
export DEEPSEEK_API_KEY="your-api-key"

# Run the integration test script
./test/integration_test.sh
```

### GitHub Actions

This repository includes GitHub Actions workflows that automatically run tests on push and pull requests. The workflow uses the `DEEPSEEK_API_KEY` secret configured in the repository settings.

To set up the GitHub Actions workflow:

1. Go to your repository settings on GitHub
2. Navigate to "Secrets and variables" > "Actions"
3. Click "New repository secret"
4. Add a secret with the name `DEEPSEEK_API_KEY` and your DeepSeek API key as the value
5. Save the secret

The GitHub Actions workflow will run unit tests and integration tests on every push to the main branch and on pull requests to the main branch.

## License

This project is licensed under the Apache License 2.0 - see the LICENSE file for details.