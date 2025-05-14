#!/bin/bash

# Exit on any error
set -e

# Build the application
cd "$(dirname "$0")/.."
go build -v ./cmd/deepfish

# Check if DEEPSEEK_API_KEY is set
if [ -z "$DEEPSEEK_API_KEY" ]; then
    echo "Error: DEEPSEEK_API_KEY environment variable is not set"
    echo "Please set it before running this test"
    exit 1
fi

# Create a test directory
TEST_DIR="$(pwd)/test_output"
mkdir -p "$TEST_DIR"

# Test 1: Generate a simple email
echo "Test 1: Generate a simple email"
./deepfish -prompt "Create a short phishing email in Chinese" -output "$TEST_DIR/test1.txt"

# Check if output file exists and is not empty
if [ ! -s "$TEST_DIR/test1.txt" ]; then
    echo "Error: Output file is empty or does not exist"
    exit 1
fi

# Test 2: Generate an HTML email
echo "Test 2: Generate an HTML email"
./deepfish -prompt "Create a short phishing email in Chinese" -format html -output "$TEST_DIR/test2.html"

# Check if output file exists and is not empty
if [ ! -s "$TEST_DIR/test2.html" ]; then
    echo "Error: Output file is empty or does not exist"
    exit 1
fi

# Test 3: Check if output contains Chinese characters
echo "Test 3: Check if output contains Chinese characters"
if ! grep -P "[\x{4e00}-\x{9fff}]" "$TEST_DIR/test1.txt"; then
    echo "Error: Output does not contain Chinese characters"
    exit 1
fi

echo "All integration tests passed!"