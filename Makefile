.PHONY: all test lint clean build

# Default target
all: lint test build

# Build the application
build:
	@echo "Building application..."
	go build -v ./...

# Run all tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run linting checks
lint:
	@echo "Running linter..."
	go vet ./...
	@if command -v golangci-lint > /dev/null; then \
		echo "Running golangci-lint..."; \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not found, skipping. Consider installing it for more thorough linting."; \
	fi

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	go clean
	rm -f $(shell find . -name '*.test')

