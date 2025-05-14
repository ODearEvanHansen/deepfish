#!/bin/bash

# Make sure the script exits on any error
set -e

# Build the application
cd "$(dirname "$0")/.."
go build -o deepfish ./cmd/deepfish

# Load API key from .env file if it exists
if [ -f ../.env ]; then
    source ../.env
elif [ -f ./.env ]; then
    source ./.env
fi

# Check if API key is set
if [ -z "$DEEPSEEK_API_KEY" ]; then
    echo "Error: DEEPSEEK_API_KEY is not set."
    echo "Please create a .env file in the project root with the following content:"
    echo "DEEPSEEK_API_KEY=your-api-key"
    echo "You can get an API key from https://platform.deepseek.com/"
    exit 1
fi

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