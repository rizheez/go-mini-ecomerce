# Clean Architecture Folder Structure for Go Fiber E-Commerce API

## Project Structure

```
d:\mkmk\go\e-commerce-API\
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
├── .env.example                        # Environment variables example
├── .gitignore                          # Git ignore file
├── docker-compose.yml                  # Docker development setup
├── Dockerfile                          # Docker production image
├── Makefile                            # Build automation
└── README.md                           # Project documentation
```

## Clean Architecture Layers

### 1. **Domain Layer** (`internal/domain/`)

- **Entities**: Core business entities with business rules
- **Repositories**: Interfaces for data access (repository pattern)
- **No dependencies on external frameworks**

### 2. **Use Cases Layer** (`internal/usecases/`)

- **Business logic and application rules**
- **Orchestrates the flow of data to and from entities**
- **Depends only on domain layer**

### 3. **Interface Adapters Layer** (`internal/interfaces/`)

- **HTTP handlers, routes, middleware**
- **DTOs for request/response transformation**
- **Converts data between use cases and external format**

### 4. **Infrastructure Layer** (`internal/infrastructure/`)

- **Database implementations**
- **External service integrations**
- **Framework-specific code**

### 5. **External Layer** (`pkg/`, `cmd/`)

- **Application entry point**
- **Shared utilities and packages**
- **Configuration management**

## Key Benefits

### 1. **Separation of Concerns**

- Each layer has a single responsibility
- Clear boundaries between business logic and technical details

### 2. **Testability**

- Easy to mock interfaces for unit testing
- Business logic isolated from infrastructure

### 3. **Flexibility**

- Easy to swap implementations (e.g., database, payment providers)
- Framework-independent core business logic

### 4. **Maintainability**

- Clear dependency direction (inward)
- Easier to understand and modify

## Dependency Rule

```
cmd → interfaces → usecases → domain
pkg → (can be used by any layer)
infrastructure → domain (implements interfaces)
```

**Dependencies point inward only** - outer layers can depend on inner layers, but not vice versa.

## Example Implementation Flow

1. **HTTP Request** → Handler (`interfaces/http/handlers`)
2. **Handler** → Use Case (`usecases`)
3. **Use Case** → Repository Interface (`domain/repositories`)
4. **Repository Interface** → Repository Implementation (`infrastructure/database/repositories`)
5. **Database** → Returns data through the same chain
6. **Response** → Transformed by DTOs and returned to client

This structure ensures your e-commerce API is scalable, maintainable, and follows clean architecture principles!
