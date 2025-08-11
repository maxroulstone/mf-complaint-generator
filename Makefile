.PHONY: run build clean test install benchmark

# Default target
all: build

# Run the application
run:
	go run cmd/generator/main.go

# Build the application
build:
	go build -o bin/mf-complaint-generator cmd/generator/main.go

# Install dependencies
install:
	go mod tidy
	go mod download

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f case_*.eml
	rm -f complaint_*.eml
	rm -f email_*.eml

# Run tests
test:
	go test ./...

# Build for multiple platforms
build-all:
	mkdir -p bin
	GOOS=darwin GOARCH=amd64 go build -o bin/mf-complaint-generator-darwin-amd64 cmd/generator/main.go
	GOOS=linux GOARCH=amd64 go build -o bin/mf-complaint-generator-linux-amd64 cmd/generator/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/mf-complaint-generator-windows-amd64.exe cmd/generator/main.go

# Quick demo run with default settings
demo:
	printf "y\ny\n14\n3\n" | go run cmd/generator/main.go

# Demo with separate passwords
demo-separate:
	printf "n\nn\n3\n" | go run cmd/generator/main.go

# Performance benchmark - generate 50 cases
benchmark:
	@echo "Benchmarking with 50 cases..."
	printf "y\nn\n50\n" | go run cmd/generator/main.go

# Large scale test - 100 cases
stress-test:
	@echo "Stress testing with 100 cases..."
	printf "y\ny\n7\n100\n" | go run cmd/generator/main.go

# Show help
help:
	@echo "Available targets:"
	@echo "  run          - Run the application interactively"
	@echo "  build        - Build the application"
	@echo "  install      - Install dependencies"
	@echo "  clean        - Clean build artifacts and generated files"
	@echo "  test         - Run tests"
	@echo "  build-all    - Build for multiple platforms"
	@echo "  demo         - Run with default demo settings (3 cases)"
	@echo "  demo-separate- Demo with separate password emails"
	@echo "  benchmark    - Performance test with 50 cases"
	@echo "  stress-test  - Large scale test with 100 cases"
	@echo "  help         - Show this help"
