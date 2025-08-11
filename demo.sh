#!/bin/bash

# Motor Finance Complaint Generator Demo Script
# This script demonstrates the capabilities of the generator

set -e

echo "Motor Finance Complaint Generator Demo"
echo "======================================"
echo ""

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "ERROR: Go is not installed. Please install Go 1.21+ to continue."
    exit 1
fi

echo "Go version: $(go version)"
echo ""

# Install dependencies
echo "Installing dependencies..."
go mod tidy
echo "Dependencies installed"
echo ""

# Build the application
echo "Building application..."
make build
echo "Build complete"
echo ""

# Run demo scenarios
echo "Running Demo Scenarios"
echo "====================="
echo ""

echo "Scenario 1: Basic complaint with passwords included"
echo "--------------------------------------------------"
echo "Generating 2 cases with passwords in main email..."
printf "y\nn\n2\n" | go run cmd/generator/main.go
echo ""

echo "Scenario 2: Separate password emails"
echo "------------------------------------"
echo "Generating 2 cases with separate password emails..."
printf "n\nn\n2\n" | go run cmd/generator/main.go
echo ""

echo "Scenario 3: Full workflow with chaser emails"
echo "--------------------------------------------"
echo "Generating 3 cases with chaser emails (7 day delay)..."
printf "y\ny\n7\n3\n" | go run cmd/generator/main.go
echo ""

echo "Scenario 4: Performance test"
echo "----------------------------"
echo "Generating 20 cases to demonstrate performance..."
printf "y\nn\n20\n" | go run cmd/generator/main.go
echo ""

# Show generated files
echo "Generated Files:"
echo "==============="
ls -la *.eml | head -10
echo ""
echo "Total files generated: $(ls *.eml | wc -l)"
echo ""

# File analysis
echo "File Analysis:"
echo "============="
complaint_files=$(ls case_*_complaint_*.eml | wc -l)
password_files=$(ls case_*_passwords_*.eml 2>/dev/null | wc -l || echo "0")
chaser_files=$(ls case_*_chaser_*.eml 2>/dev/null | wc -l || echo "0")

echo "Complaint emails: $complaint_files"
echo "Password emails: $password_files"
echo "Chaser emails: $chaser_files"
echo ""

# Cleanup option
echo "Demo complete! Would you like to clean up generated files? (y/n)"
read -r cleanup_choice
if [[ $cleanup_choice =~ ^[Yy]$ ]]; then
    rm -f *.eml
    echo "Cleanup complete"
else
    echo "Files preserved for inspection"
fi

echo ""
echo "Demo completed successfully!"
echo "Check README.md for more usage examples"
echo "Star the repository if you found it useful!"
