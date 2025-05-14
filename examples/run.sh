#!/bin/bash

# Make sure the script exits on any error
set -e

# Build the application
cd "$(dirname "$0")/.."
go build -o deepfish ./cmd/deepfish

# IMPORTANT: Replace "your-api-key" with your actual DeepSeek API key
# You can get an API key from https://platform.deepseek.com/
# DO NOT commit this script with your real API key
export DEEPSEEK_API_KEY="your-api-key"

# Alternatively, you can set the API key in your environment before running this script
# export DEEPSEEK_API_KEY="your-api-key"

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