# Project Structure

## Overview

This document describes the complete project structure for the E-Commerce API application.

## Directory Structure

```
e-commerce-API/
├── cmd/
│   └── api/
│       └── main.go                     # Application entry point
├── config/
│   ├── config.go                       # Configuration struct
│   ├── database.go                     # Database configuration
│   └── env.go                          # Environment variables
├── internal/
│   ├── domain/
│   │   ├── entities/
│   │   │   ├── user.go                 # User entity
│   │   │   ├── category.go             # Category entity
│   │   │   ├── product.go              # Product entity
│   │   │   ├── cart.go                 # Cart entity
│   │   │   ├── order.go                # Order entity
│   │   │   ├── payment.go              # Payment entity
│   │   │   └── address.go              # Address entity
│   │   └── repositories/
│   │       ├── user_repository.go      # User repository interface
│   │       ├── category_repository.go  # Category repository interface
│   │       ├── product_repository.go   # Product repository interface
│   │       ├── cart_repository.go      # Cart repository interface
│   │       ├── order_repository.go     # Order repository interface
│   │       ├── payment_repository.go   # Payment repository interface
│   │       └── address_repository.go   # Address repository interface
│   ├── usecases/
│   │   ├── auth_usecase.go             # Authentication use cases
│   │   ├── user_usecase.go             # User management use cases
│   │   ├── category_usecase.go         # Category management use cases
│   │   ├── product_usecase.go          # Product management use cases
│   │   ├── cart_usecase.go             # Cart management use cases
│   │   ├── order_usecase.go            # Order management use cases
│   │   └── payment_usecase.go          # Payment processing use cases
│   ├── interfaces/
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── auth_handler.go     # Authentication endpoints
│   │       │   ├── user_handler.go     # User management endpoints
│   │       │   ├── category_handler.go # Category management endpoints
│   │       │   ├── product_handler.go  # Product management endpoints
│   │       │   ├── cart_handler.go     # Cart management endpoints
│   │       │   ├── order_handler.go    # Order management endpoints
│   │       │   └── payment_handler.go  # Payment processing endpoints
│   │       ├── middleware/
│   │       │   ├── auth.go             # JWT authentication middleware
│   │       │   ├── cors.go             # CORS middleware
│   │       │   ├── logger.go           # Logging middleware
│   │       │   ├── rate_limit.go       # Rate limiting middleware
│   │       │   └── admin.go            # Admin role middleware
│   │       ├── routes/
│   │       │   ├── auth_routes.go      # Authentication routes
│   │       │   ├── user_routes.go      # User routes
│   │       │   ├── category_routes.go  # Category routes
│   │       │   ├── product_routes.go   # Product routes
│   │       │   ├── cart_routes.go      # Cart routes
│   │       │   ├── order_routes.go     # Order routes
│   │       │   ├── payment_routes.go   # Payment routes
│   │       │   └── routes.go           # Main route setup
│   │       └── dto/
│   │           ├── auth_dto.go         # Authentication DTOs
│   │           ├── user_dto.go         # User DTOs
│   │           ├── category_dto.go     # Category DTOs
│   │           ├── product_dto.go      # Product DTOs
│   │           ├── cart_dto.go         # Cart DTOs
│   │           ├── order_dto.go        # Order DTOs
│   │           ├── payment_dto.go      # Payment DTOs
│   │           └── response.go         # Common response DTOs
│   └── infrastructure/
│       ├── database/
│       │   ├── postgres.go             # PostgreSQL connection
│       │   ├── migrations.go           # Database migrations
│       │   └── repositories/
│       │       ├── user_repository_impl.go     # User repository implementation
│       │       ├── category_repository_impl.go # Category repository implementation
│       │       ├── product_repository_impl.go  # Product repository implementation
│       │       ├── cart_repository_impl.go     # Cart repository implementation
│       │       ├── order_repository_impl.go    # Order repository implementation
│       │       ├── payment_repository_impl.go  # Payment repository implementation
│       │       └── address_repository_impl.go  # Address repository implementation
│       ├── auth/
│       │   ├── jwt.go                  # JWT token generation/validation
│       │   └── password.go             # Password hashing utilities
│       └── payment/
│           ├── credit_card.go          # Credit card payment processor
│           └── cash.go                 # Cash payment processor
├── pkg/
│   ├── utils/
│   │   ├── response.go                 # HTTP response utilities
│   │   ├── pagination.go               # Pagination utilities
│   │   └── uuid.go                     # UUID utilities
│   ├── validator/
│   │   ├── validator.go                # Custom validation rules
│   │   └── rules.go                    # Validation rule definitions
│   └── logger/
│       └── logger.go                   # Structured logging
├── migrations/
│   ├── 001_create_users_table.sql      # Initial users table
│   ├── 002_create_categories_table.sql # Categories table
│   ├── 003_create_products_table.sql   # Products table
│   ├── 004_create_addresses_table.sql  # User addresses table
│   ├── 005_create_carts_table.sql      # Shopping carts table
│   ├── 006_create_orders_table.sql     # Orders table
│   ├── 007_create_payments_table.sql   # Payments table
│   └── 008_create_indexes.sql          # Performance indexes
├── docs/
│   ├── api/                            # API documentation
│   ├── deployment/                     # Deployment guides
│   └── development/                    # Development setup
├── scripts/
│   ├── build.sh                        # Build script
│   ├── test.sh                         # Test script
│   └── deploy.sh                       # Deployment script
├── tests/
│   ├── integration/                    # Integration tests
│   ├── unit/                           # Unit tests
│   └── fixtures/                       # Test data fixtures
├── init-scripts/
│   └── init.sql                        # Database initialization script
├── .env.example                        # Environment variables example
├── .env                                # Environment variables (gitignored)
├── .gitignore                          # Git ignore file
├── .air.toml                           # Air configuration for hot reloading
├── docker-compose.yml                  # Docker development setup
├── Dockerfile                          # Docker production image
├── Makefile                            # Build automation
├── go.mod                              # Go modules
├── go.sum                              # Go checksums
└── README.md                           # Project documentation
```

## Key Directories

### cmd/

Contains the main application entry points. Each subdirectory represents a different executable.

**Structure**:
```
cmd/
└── api/
    └── main.go  # Main entry point for the API server
```

### config/

Contains configuration-related code including environment variable parsing and configuration structs.

**Structure**:
```
config/
├── config.go   # Main configuration struct
├── database.go # Database configuration
└── env.go      # Environment variable parsing
```

### internal/

Contains all the internal application code that should not be imported by external projects.

**Structure**:
```
internal/
├── domain/        # Domain layer (entities and repository interfaces)
├── usecases/      # Use cases layer (business logic)
├── interfaces/    # Interface adapters layer (HTTP handlers, middleware, etc.)
└── infrastructure/ # Infrastructure layer (database, external services)
```

### pkg/

Contains shared packages that could potentially be used by other projects.

**Structure**:
```
pkg/
├── utils/      # Utility functions
├── validator/  # Validation utilities
└── logger/     # Logging utilities
```

### migrations/

Contains database migration files.

**Structure**:
```
migrations/
├── 001_*.sql  # First migration
├── 002_*.sql  # Second migration
└── ...        # Additional migrations
```

### docs/

Contains all project documentation.

**Structure**:
```
docs/
├── api/         # API documentation
├── deployment/  # Deployment guides
└── development/ # Development documentation
```

### scripts/

Contains utility scripts for building, testing, and deploying the application.

**Structure**:
```
scripts/
├── build.sh   # Build script
├── test.sh    # Test script
└── deploy.sh  # Deployment script
```

### tests/

Contains all test files.

**Structure**:
```
tests/
├── integration/ # Integration tests
├── unit/        # Unit tests
└── fixtures/    # Test data fixtures
```

## File Naming Conventions

### Go Files

1. **Package files**: `package_name.go`
2. **Implementation files**: `package_name_impl.go`
3. **Interface files**: `package_name_interface.go` or simply `package_name.go` in interface packages
4. **Test files**: `file_name_test.go`

### Configuration Files

1. **Environment files**: `.env`, `.env.example`
2. **Configuration files**: `*.toml`, `*.yaml`, `*.json`

### Documentation Files

1. **README files**: `README.md`
2. **Documentation files**: `*.md`

### Script Files

1. **Shell scripts**: `*.sh`
2. **Batch files**: `*.bat` (Windows)

## Package Organization

### Domain Packages

```
internal/domain/
├── entities/      # Business entities
└── repositories/  # Repository interfaces
```

### Use Case Packages

```
internal/usecases/
├── auth_usecase.go     # Authentication use cases
├── user_usecase.go     # User management use cases
└── product_usecase.go  # Product management use cases
```

### Interface Packages

```
internal/interfaces/
└── http/
    ├── handlers/   # HTTP handlers
    ├── middleware/ # Middleware functions
    ├── routes/     # Route definitions
    └── dto/        # Data transfer objects
```

### Infrastructure Packages

```
internal/infrastructure/
├── database/   # Database-related code
├── auth/       # Authentication infrastructure
└── payment/    # Payment processing infrastructure
```

## Dependency Management

### Go Modules

The project uses Go modules for dependency management:

```bash
# Initialize module
go mod init github.com/your-project/e-commerce-API

# Add dependency
go get github.com/some/package

# Update dependencies
go mod tidy

# Vendor dependencies (optional)
go mod vendor
```

### External Dependencies

Major external dependencies include:
- **Web Framework**: Fiber (or Gin)
- **Database**: GORM with PostgreSQL
- **Authentication**: JWT
- **Validation**: Validator
- **Logging**: Zap or Logrus
- **Testing**: Testify, Mock

## Environment Configuration

### Environment Files

- `.env.example` - Template for environment variables
- `.env` - Actual environment variables (gitignored)

### Required Variables

```
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=ecommerce_user
DB_PASSWORD=ecommerce_password
DB_NAME=ecommerce_db

# JWT
JWT_SECRET=your_jwt_secret
JWT_EXPIRATION_HOURS=24

# Server
SERVER_PORT=8080
```

## Build and Deployment

### Local Development

```bash
# Install dependencies
go mod download

# Run application
go run cmd/api/main.go

# Or using Makefile
make run
```

### Building for Production

```bash
# Build binary
go build -o ecommerce-api cmd/api/main.go

# Or using Makefile
make build
```

### Docker Deployment

```bash
# Build Docker image
docker build -t ecommerce-api .

# Run with Docker
docker run -p 8080:8080 --env-file .env ecommerce-api
```

### Docker Compose

```bash
# Start all services
docker-compose up -d

# Stop all services
docker-compose down
```

## Testing

### Unit Tests

```bash
# Run unit tests
go test ./internal/...

# Run with coverage
go test -coverprofile=coverage.out ./internal/...
```

### Integration Tests

```bash
# Run integration tests
go test ./tests/integration/...
```

## Documentation

### API Documentation

Located in `docs/api/` directory.

### Development Documentation

Located in `docs/development/` directory.

### Deployment Documentation

Located in `docs/deployment/` directory.

## Version Control

### Git Ignore

The `.gitignore` file excludes:
- Build artifacts
- Environment files (except .env.example)
- Logs
- IDE files
- OS-specific files

### Branching Strategy

- `main` - Production code
- `develop` - Development code
- `feature/*` - Feature branches
- `hotfix/*` - Hotfix branches
- `release/*` - Release branches

## Continuous Integration

### GitHub Actions

Workflows are defined in `.github/workflows/`:
- Test workflow
- Build workflow
- Deploy workflow

## Security Considerations

### Sensitive Files

- `.env` is gitignored
- SSL certificates are not included in repository
- API keys and secrets are not hardcoded

### Security Best Practices

- Input validation and sanitization
- Authentication and authorization
- Secure password storage
- Rate limiting
- CORS configuration
- HTTPS in production

This project structure follows industry best practices for Go applications and Clean Architecture principles, ensuring maintainability, scalability, and testability.