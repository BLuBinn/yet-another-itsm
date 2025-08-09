.PHONY: help build run dev test clean docker-up docker-down migrate-up migrate-down migrate-create sqlc-generate fmt lint

# Default target
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-20s %s\n", $$1, $$2}'

# Development
dev: ## Run the application in development mode
	@echo "Starting development server..."
	@LOG_LEVEL=debug LOG_FORMAT=console go run github.com/air-verse/air@latest

build: ## Build the application
	@echo "Building application..."
	@go build -o bin/server cmd/server/main.go

run: build ## Build and run the application
	@echo "Running application..."
	@./bin/server

# Testing
test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html

# Code quality
fmt: ## Format Go code
	@echo "Formatting code..."
	@go fmt ./...

# Database operations
migrate-up: ## Run database migrations up
	@echo "Running migrations up..."
	@go run github.com/pressly/goose/v3/cmd/goose@latest -dir sql/migrations postgres "host=localhost port=5432 user=postgres password=postgres dbname=msn_map_api sslmode=disable" up

migrate-down: ## Run database migrations down
	@echo "Running migrations down..."
	@go run github.com/pressly/goose/v3/cmd/goose@latest -dir sql/migrations postgres "host=localhost port=5432 user=postgres password=postgres dbname=msn_map_api sslmode=disable" down

migrate-reset: ## Reset database migrations
	@echo "Resetting database..."
	@go run github.com/pressly/goose/v3/cmd/goose@latest -dir sql/migrations postgres "host=localhost port=5432 user=postgres password=postgres dbname=msn_map_api sslmode=disable" reset

migrate-status: ## Check migration status
	@echo "Checking migration status..."
	@go run github.com/pressly/goose/v3/cmd/goose@latest -dir sql/migrations postgres "host=localhost port=5432 user=postgres password=postgres dbname=msn_map_api sslmode=disable" status

migrate-create: ## Create a new migration file (usage: make migrate-create NAME=migration_name)
	@if [ -z "$(NAME)" ]; then \
		echo "Usage: make migrate-create NAME=migration_name"; \
		exit 1; \
	fi
	@echo "Creating migration: $(NAME)"
	@go run github.com/pressly/goose/v3/cmd/goose@latest -dir sql/migrations create $(NAME) sql

# SQLC
sqlc-generate: ## Generate sqlc code
	@echo "Generating sqlc code..."
	@go run github.com/sqlc-dev/sqlc/cmd/sqlc@latest generate

# Docker operations
docker-up: ## Start Docker containers
	@echo "Starting Docker containers..."
	@docker-compose up -d

docker-down: ## Stop Docker containers
	@echo "Stopping Docker containers..."
	@docker-compose down

docker-logs: ## Show Docker logs
	@echo "Showing Docker logs..."
	@docker-compose logs -f

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	@docker build -t yet-another-itsm .

docker-rebuild: ## Rebuild and restart Docker containers
	@echo "Rebuilding Docker containers..."
	@docker-compose down
	@docker-compose build --no-cache
	@docker-compose up -d

# Database setup
db-setup: docker-up ## Setup database with Docker and run migrations
	@echo "Waiting for database to be ready..."
	@sleep 5
	@make migrate-up

# Development workflow
setup: ## Initial project setup
	@echo "Setting up development environment..."
	@go mod download
	@make sqlc-generate
	@make db-setup
	@echo "Setup complete! Run 'make dev' to start the server"

# Cleanup
clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html
	@docker system prune -f

# Health check
health: ## Check API health
	@echo "Checking API health..."
	@curl -f http://localhost:8080/health || echo "API is not running"