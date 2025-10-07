.PHONY: build run clean test docker-build docker-run help

# Binary name
BINARY_NAME=transfer-system
OS=$(shell go env GOOS)

# Build the application
build:
	@echo "Building application..."
	GO111MODULE=on GOOS=$(OS) go build -o bin/$(BINARY_NAME) main.go
	@echo "Build complete: bin/$(BINARY_NAME)"

# Run the application locally
run-local:
	@echo "Running application..."
	go run main.go

# Clean build artifacts
clean-artifacts:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	@echo "Clean complete"

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run with Docker Compose
docker-run:
	@echo "Starting application with Docker Compose..."
	docker-compose up --build -d

# Stop Docker Compose
docker-stop:
	@echo "Stopping Docker Compose..."
	docker-compose down

# Initialize database
db-init:
	@echo "Initializing database..."
	docker-compose exec postgres psql -U postgres -d transfer_system -f docker-entrypoint-initdb.d/init.sql
	@echo "Database initialized"

# Help
help:
	@echo "Available targets:"
	@echo "  build             - Build the application binary"
	@echo "  run-local         - Run the application locally"
	@echo "  clean-artifacts   - Remove build artifacts"
	@echo "  test              - Run tests"
	@echo "  docker-run        - Run application with Docker Compose"
	@echo "  docker-stop       - Stop Docker Compose"
	@echo "  db-init           - Initialize database schema"
	@echo "  help              - Show this help message"
