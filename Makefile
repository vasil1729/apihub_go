.PHONY: help run build test test-unit test-integration coverage lint swagger docker-up docker-down clean

# Variables
APP_NAME=apihub_go
MAIN_PATH=./cmd/server
BINARY_PATH=./bin/$(APP_NAME)
COVERAGE_FILE=coverage.out

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

run: ## Run the application
	@echo "Starting $(APP_NAME)..."
	@go run $(MAIN_PATH)/main.go

build: ## Build the application
	@echo "Building $(APP_NAME)..."
	@mkdir -p bin
	@go build -o $(BINARY_PATH) $(MAIN_PATH)/main.go
	@echo "Binary created at $(BINARY_PATH)"

test: ## Run all tests
	@echo "Running all tests..."
	@go test -v -race ./...

test-unit: ## Run unit tests only
	@echo "Running unit tests..."
	@go test -v -race ./tests/unit/...

test-integration: ## Run integration tests only
	@echo "Running integration tests..."
	@go test -v -race ./tests/integration/...

coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@go test -v -race -coverprofile=$(COVERAGE_FILE) -covermode=atomic ./...
	@go tool cover -html=$(COVERAGE_FILE) -o coverage.html
	@echo "Coverage report generated: coverage.html"

lint: ## Run linter
	@echo "Running linter..."
	@golangci-lint run --timeout=5m

swagger: ## Generate Swagger documentation
	@echo "Generating Swagger docs..."
	@swag init -g $(MAIN_PATH)/main.go -o ./api/swagger
	@echo "Swagger docs generated in ./api/swagger"

docker-up: ## Start Docker containers
	@echo "Starting Docker containers..."
	@docker-compose up -d

docker-down: ## Stop Docker containers
	@echo "Stopping Docker containers..."
	@docker-compose down

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	@docker build -t $(APP_NAME):latest .

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -rf $(COVERAGE_FILE) coverage.html
	@rm -rf api/swagger/docs.go api/swagger/swagger.json api/swagger/swagger.yaml
	@echo "Clean complete"

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy

install-tools: ## Install development tools
	@echo "Installing development tools..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/swaggo/swag/cmd/swag@latest
	@echo "Tools installed"

.DEFAULT_GOAL := help
