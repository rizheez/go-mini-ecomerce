# Variables
BINARY_NAME=ecommerce-api
VERSION ?= dev
BUILD_TIME := $(shell date -u +%Y%m%d.%H%M%S)
GIT_COMMIT := $(shell git rev-parse HEAD)
GIT_DIRTY := $(shell git diff --shortstat 2> /dev/null | tail -n1)

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GORUN=$(GOCMD) run

# Directories
CMD_DIR=./cmd/api
INTERNAL_DIR=./internal
PKG_DIR=./pkg

# Build the application
build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(CMD_DIR)/main.go

# Build with version information
build-version:
	$(GOBUILD) -ldflags="-X 'main.Version=$(VERSION)' -X 'main.BuildTime=$(BUILD_TIME)' -X 'main.GitCommit=$(GIT_COMMIT)'" -o $(BINARY_NAME) -v $(CMD_DIR)/main.go

# Run the application
run:
	$(GORUN) $(CMD_DIR)/main.go

# Run the application with air for hot reloading
run-dev:
	air

# Run tests
test:
	$(GOTEST) -v ./...

# Run tests with coverage
test-cover:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

# Clean build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f coverage.out
	rm -f coverage.html

# Install dependencies
deps:
	$(GOMOD) download

# Update dependencies
deps-update:
	$(GOMOD) tidy

# Format code
fmt:
	$(GOCMD) fmt ./...

# Vet code
vet:
	$(GOCMD) vet ./...

# Install air for hot reloading (development)
install-air:
	$(GOGET) github.com/cosmtrek/air@latest

# Run golangci-lint
lint:
	golangci-lint run

# Install golangci-lint
install-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2

# Generate mocks
mock:
	mockgen -source=$(INTERNAL_DIR)/domain/repositories/user_repository.go -destination=$(INTERNAL_DIR)/domain/repositories/mocks/user_repository_mock.go
	mockgen -source=$(INTERNAL_DIR)/domain/repositories/category_repository.go -destination=$(INTERNAL_DIR)/domain/repositories/mocks/category_repository_mock.go
	mockgen -source=$(INTERNAL_DIR)/domain/repositories/product_repository.go -destination=$(INTERNAL_DIR)/domain/repositories/mocks/product_repository_mock.go

# Docker build
docker-build:
	docker build -t ecommerce-api:$(VERSION) .

# Docker run
docker-run:
	docker run --rm -p 8080:8080 --env-file .env ecommerce-api:$(VERSION)

# Docker compose up
docker-up:
	docker-compose up -d

# Docker compose down
docker-down:
	docker-compose down

# Docker compose logs
docker-logs:
	docker-compose logs -f

# Migrate database
migrate-up:
	$(GORUN) ./scripts/migrate.go up

migrate-down:
	$(GORUN) ./scripts/migrate.go down

# Generate documentation
docs:
	swag init -g $(CMD_DIR)/main.go -o ./docs/swagger

# Install swag for documentation
install-swag:
	$(GOGET) github.com/swaggo/swag/cmd/swag@latest

# Help
help:
	@echo "Available commands:"
	@echo "  build          - Build the application"
	@echo "  run            - Run the application"
	@echo "  run-dev        - Run the application with hot reloading"
	@echo "  test           - Run tests"
	@echo "  clean          - Clean build files"
	@echo "  deps           - Install dependencies"
	@echo "  fmt            - Format code"
	@echo "  vet            - Vet code"
	@echo "  lint           - Run golangci-lint"
	@echo "  docker-build   - Build Docker image"
	@echo "  docker-up      - Start Docker containers"
	@echo "  docker-down    - Stop Docker containers"
	@echo "  migrate-up     - Run database migrations"
	@echo "  docs           - Generate documentation"
	@echo ""
	@echo "For development, use 'make run-dev' to start with hot reloading"

.PHONY: build run run-dev test clean deps fmt vet lint docker-build docker-up docker-down migrate-up migrate-down docs help