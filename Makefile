.PHONY: generate-mocks test test-coverage clean install-tools build run dev

default: help

# Development tools installation
install-tools:
	@echo "Installing development tools..."
	@go install github.com/vektra/mockery/v2@latest
	@echo "Tools installed successfully"

# Mock generation
generate-mocks:
	@echo "Generating mocks..."
	@mockery --all
	@echo "Mocks generated successfully"

# Testing
test: generate-mocks
	@go test -v ./internal/repositories/... ./internal/services/...

test-coverage: generate-mocks
	@go test -v -coverprofile=coverage.out ./internal/repositories/... ./internal/services/...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Build and run
build:
	@echo "Building application..."
	@go build -o bin/server cmd/server/main.go
	@echo "Build completed: bin/server"

run: build
	@echo "Starting server..."
	@./bin/server

dev:
	@echo "Starting development server with air..."
	@air

# Cleanup
clean:
	@echo "Cleaning up..."
	@rm -rf mocks/
	@rm -f coverage.out coverage.html
	@rm -rf bin/
	@echo "Cleanup completed"

swagger:
	@echo "Generating swagger documentation..."
	@swag init -g cmd/server/main.go --parseDependency --parseInternal
	@echo "Swagger documentation generated successfully"

# Help
help:
	@echo "Available commands:"
	@echo "  install-tools    - Install development tools (mockery)"
	@echo "  generate-mocks   - Generate mocks for interfaces"
	@echo "  test            - Run all tests"
	@echo "  test-coverage   - Run tests with coverage report"
	@echo "  build           - Build the application"
	@echo "  run             - Build and run the application"
	@echo "  dev             - Start development server with air"
	@echo "  clean           - Clean up generated files"
	@echo "  swagger         - Generate swagger documentation"
	@echo "  help            - Show this help message"