#!/bin/bash

# Make sure the script exits on any error
set -e

# Build the application
cd "$(dirname "$0")/.."
go build -o deepfish ./cmd/deepfish

# Set the API key (replace with your actual key or use environment variable)
export DEEPSEEK_API_KEY="your-api-key"

# Example 1: Generate a simple phishing email
echo "Example 1: Simple phishing email"
./deepfish -prompt "Create a phishing email pretending to be from a bank asking for account verification"

# Example 2: Generate HTML output and save to a file
echo -e "\nExample 2: HTML phishing email saved to file"
./deepfish -format html -output example_email.html -prompt "Create a phishing email pretending to be from a tech company about a security breach"
echo "Email saved to example_email.html"

# Example 3: Use a more specific prompt
echo -e "\nExample 3: Specific phishing scenario"
./deepfish -prompt "Create a phishing email pretending to be from China Mobile about an urgent account issue that requires immediate action"