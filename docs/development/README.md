# Development Guide

## Overview

This guide explains how to set up and develop the E-Commerce API application.

## Prerequisites

- Go 1.21+
- Docker and Docker Compose
- PostgreSQL 13+
- Redis 6+
- Git
- IDE (Visual Studio Code recommended)

## Project Structure

```
e-commerce-API/
├── cmd/
│   └── api/              # Application entry point
├── config/               # Configuration files
├── internal/             # Internal application code
│   ├── domain/           # Domain entities and repositories
│   ├── usecases/         # Business logic
│   ├── interfaces/       # HTTP handlers, middleware, routes
│   └── infrastructure/   # Database, external services
├── pkg/                  # Shared packages
├── migrations/           # Database migrations
├── docs/                 # Documentation
├── scripts/              # Utility scripts
├── .env.example          # Environment variables example
├── .env                  # Environment variables (gitignored)
├── go.mod                # Go modules
├── go.sum                # Go checksums
├── Dockerfile            # Docker configuration
├── docker-compose.yml    # Docker Compose configuration
├── Makefile              # Build automation
└── README.md             # Project documentation
```

## Setup Development Environment

### 1. Clone the Repository

```bash
git clone <repository-url>
cd e-commerce-API
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Set Up Environment Variables

```bash
cp .env.example .env
# Edit .env with your local configuration
```

### 4. Start Development Services

```bash
# Start database and other services
docker-compose up -d db redis

# Wait for services to start
# Check with: docker-compose ps
```

### 5. Run Database Migrations

```bash
# Run migrations
go run scripts/migrate.go up
```

### 6. Start the Application

```bash
# Option 1: Run directly
go run cmd/api/main.go

# Option 2: Build and run
go build -o ecommerce-api cmd/api/main.go
./ecommerce-api

# Option 3: Use Makefile
make run

# Option 4: Use hot reloading (requires air)
make install-air
make run-dev
```

## Development Workflow

### 1. Code Structure

Follow the Clean Architecture principles:

1. **Domain Layer**: Business entities and repository interfaces
2. **Use Cases Layer**: Business logic implementation
3. **Interface Adapters Layer**: HTTP handlers, middleware, DTOs
4. **Infrastructure Layer**: Database implementations, external services

### 2. Adding New Features

1. Define domain entities in `internal/domain/entities/`
2. Create repository interfaces in `internal/domain/repositories/`
3. Implement use cases in `internal/usecases/`
4. Create HTTP handlers in `internal/interfaces/http/handlers/`
5. Add routes in `internal/interfaces/http/routes/`
6. Implement repository in `internal/infrastructure/database/repositories/`

### 3. Database Changes

1. Create new migration file:
   ```bash
   # Create migration file manually in migrations/ directory
   # Or use a migration tool
   ```

2. Add migration to `migrations/` directory

3. Run migration:
   ```bash
   go run scripts/migrate.go up
   ```

### 4. Testing

Run tests with:
```bash
# Run all tests
make test

# Run tests with coverage
make test-cover

# Run specific test
go test -v ./internal/usecases/...
```

## Development Tools

### 1. Hot Reloading

Install and use air for hot reloading:
```bash
make install-air
make run-dev
```

### 2. Linting

Use golangci-lint for code quality:
```bash
make install-lint
make lint
```

### 3. Code Formatting

Format code with:
```bash
make fmt
```

### 4. Dependency Management

```bash
# Add dependency
go get <package>

# Update dependencies
make deps-update

# Tidy modules
go mod tidy
```

## Database Development

### 1. Access Database

```bash
# Using Docker
docker exec -it ecommerce-db psql -U ecommerce_user -d ecommerce_db

# Using psql directly
psql -h localhost -p 5432 -U ecommerce_user -d ecommerce_db
```

### 2. Database Migrations

```bash
# Create migration
# Manually create file in migrations/ directory

# Run migrations
make migrate-up

# Rollback migrations
make migrate-down
```

### 3. Database Schema

Refer to `DATABASE_SCHEMA.md` for detailed schema information.

## API Development

### 1. Adding New Endpoints

1. Create DTOs in `internal/interfaces/http/dto/`
2. Create handler in `internal/interfaces/http/handlers/`
3. Add route in `internal/interfaces/http/routes/`
4. Register route in `internal/interfaces/http/routes/routes.go`

### 2. Testing Endpoints

Use curl or Postman to test endpoints:
```bash
# Example
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123","name":"Test User"}'
```

## Debugging

### 1. Logging

Check logs for debugging information:
```bash
# Application logs
tail -f logs/app.log

# Docker logs
docker logs -f ecommerce-api
```

### 2. Debugging with Delve

```bash
# Install delve
go install github.com/go-delve/delve/cmd/dlv@latest

# Debug application
dlv debug cmd/api/main.go
```

## Code Quality

### 1. Testing

Write unit tests for:
- Use cases
- Handlers
- Repositories
- Utility functions

Run tests with coverage:
```bash
make test-cover
```

### 2. Code Review

Before committing:
1. Run `make fmt` to format code
2. Run `make vet` to check for issues
3. Run `make lint` to check code quality
4. Run `make test` to ensure all tests pass

### 3. Git Workflow

1. Create feature branch:
   ```bash
   git checkout -b feature/new-feature
   ```

2. Commit changes with clear messages:
   ```bash
   git commit -m "Add new feature: description"
   ```

3. Push and create pull request

## Common Development Tasks

### 1. Add New Entity

1. Create domain entity in `internal/domain/entities/`
2. Create repository interface in `internal/domain/repositories/`
3. Create GORM model in `internal/infrastructure/database/models/`
4. Create repository implementation in `internal/infrastructure/database/repositories/`
5. Create use case in `internal/usecases/`
6. Create DTOs in `internal/interfaces/http/dto/`
7. Create handler in `internal/interfaces/http/handlers/`
8. Add routes in `internal/interfaces/http/routes/`

### 2. Add New API Endpoint

1. Create DTOs in `internal/interfaces/http/dto/`
2. Create handler in `internal/interfaces/http/handlers/`
3. Add route in `internal/interfaces/http/routes/`
4. Register route in `internal/interfaces/http/routes/routes.go`

### 3. Update Database Schema

1. Create migration file in `migrations/`
2. Update GORM models in `internal/infrastructure/database/models/`
3. Update domain entities if needed
4. Run migration with `make migrate-up`

## Troubleshooting

### 1. Common Issues

1. **Module not found errors**
   - Run `go mod tidy`
   - Check go.mod file

2. **Database connection failed**
   - Check .env configuration
   - Verify database is running
   - Check Docker Compose services

3. **Port already in use**
   - Change SERVER_PORT in .env
   - Kill process using the port

### 2. Useful Commands

```bash
# Check running services
docker-compose ps

# View logs
docker-compose logs -f

# Stop services
docker-compose down

# Start services
docker-compose up -d

# Rebuild services
docker-compose up -d --build
```

## Best Practices

1. Follow Clean Architecture principles
2. Write unit tests for all business logic
3. Use meaningful commit messages
4. Document complex code
5. Keep functions small and focused
6. Use context for request-scoped values
7. Handle errors appropriately
8. Use logging for debugging and monitoring
9. Validate all inputs
10. Follow Go idioms and conventions