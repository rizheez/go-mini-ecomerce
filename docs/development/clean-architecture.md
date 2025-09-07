# Clean Architecture Implementation

## Overview

This document explains how Clean Architecture principles are implemented in the E-Commerce API project.

## Clean Architecture Layers

The project follows the Clean Architecture pattern with four main layers:

```
┌────────────────────────────────────┐
│              Domain                │
│  (Entities, Repository Interfaces) │
├────────────────────────────────────┤
│             Use Cases              │
│        (Business Logic)            │
├────────────────────────────────────┤
│        Interface Adapters          │
│  (Handlers, DTOs, Presenters)      │
├────────────────────────────────────┤
│           Infrastructure           │
│  (Database, External Services)     │
└────────────────────────────────────┘
```

## Layer Details

### 1. Domain Layer

**Location**: `internal/domain/`

**Responsibilities**:
- Business entities (Entities)
- Business rules interfaces (Repository Interfaces)
- No dependencies on frameworks or external libraries

**Structure**:
```
internal/domain/
├── entities/
│   ├── user.go
│   ├── product.go
│   └── order.go
└── repositories/
    ├── user_repository.go
    ├── product_repository.go
    └── order_repository.go
```

### 2. Use Cases Layer

**Location**: `internal/usecases/`

**Responsibilities**:
- Application business rules
- Orchestrates the flow of data to and from entities
- Depends only on domain layer

**Structure**:
```
internal/usecases/
├── user_usecase.go
├── product_usecase.go
└── order_usecase.go
```

### 3. Interface Adapters Layer

**Location**: `internal/interfaces/`

**Responsibilities**:
- Adapts data from use cases to external formats
- HTTP handlers, middleware, routes
- Data transfer objects (DTOs)

**Structure**:
```
internal/interfaces/
├── http/
│   ├── handlers/
│   │   ├── user_handler.go
│   │   ├── product_handler.go
│   │   └── order_handler.go
│   ├── middleware/
│   │   ├── auth.go
│   │   ├── cors.go
│   │   └── logger.go
│   ├── routes/
│   │   ├── user_routes.go
│   │   ├── product_routes.go
│   │   └── routes.go
│   └── dto/
│       ├── user_dto.go
│       ├── product_dto.go
│       └── response.go
```

### 4. Infrastructure Layer

**Location**: `internal/infrastructure/`

**Responsibilities**:
- Implementation of repository interfaces
- Database connections and operations
- External service integrations
- Framework-specific code

**Structure**:
```
internal/infrastructure/
├── database/
│   ├── postgres.go
│   ├── migrations.go
│   └── repositories/
│       ├── user_repository_impl.go
│       ├── product_repository_impl.go
│       └── order_repository_impl.go
├── auth/
│   ├── jwt.go
│   └── password.go
└── payment/
    ├── credit_card.go
    └── cash.go
```

## Dependency Rule

Dependencies point inward only:

```
cmd/api/main.go → interfaces → usecases → domain
pkg/ → (can be used by any layer)
infrastructure → domain (implements interfaces)
```

### Implementation Details

1. **Domain Layer** has no imports from other layers
2. **Use Cases Layer** only imports from domain layer
3. **Interface Adapters Layer** imports from domain and use cases layers
4. **Infrastructure Layer** imports from domain layer and implements repository interfaces

## Key Benefits

### 1. Separation of Concerns

Each layer has a single responsibility:
- Domain: Business entities and rules
- Use Cases: Application business logic
- Interface Adapters: Data transformation and presentation
- Infrastructure: Technical implementation details

### 2. Testability

- Easy to mock interfaces for unit testing
- Business logic isolated from infrastructure
- Framework-independent core business logic

### 3. Flexibility

- Easy to swap implementations (e.g., database, payment providers)
- Framework-independent core business logic
- Technology-agnostic domain layer

### 4. Maintainability

- Clear dependency direction (inward)
- Easier to understand and modify
- Changes in one layer have minimal impact on others

## Implementation Examples

### Repository Interface (Domain Layer)

```go
// internal/domain/repositories/user_repository.go
package repositories

import (
    "context"
    "github.com/your-project/internal/domain/entities"
)

type UserRepository interface {
    Create(ctx context.Context, user *entities.User) error
    FindByID(ctx context.Context, id int) (*entities.User, error)
    FindByEmail(ctx context.Context, email string) (*entities.User, error)
    Update(ctx context.Context, user *entities.User) error
    Delete(ctx context.Context, id int) error
}
```

### Repository Implementation (Infrastructure Layer)

```go
// internal/infrastructure/database/repositories/user_repository_impl.go
package repositories

import (
    "context"
    "github.com/your-project/internal/domain/entities"
    "github.com/your-project/internal/domain/repositories"
    "github.com/your-project/internal/infrastructure/database/models"
    "gorm.io/gorm"
)

type userRepositoryImpl struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
    return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) Create(ctx context.Context, user *entities.User) error {
    userModel := &models.User{
        Email:        user.Email,
        PasswordHash: user.PasswordHash,
        Name:         user.Name,
        Phone:        user.Phone,
        Role:         user.Role,
    }
    
    if err := r.db.WithContext(ctx).Create(userModel).Error; err != nil {
        return err
    }
    
    user.ID = userModel.ID
    return nil
}
```

### Use Case (Use Cases Layer)

```go
// internal/usecases/user_usecase.go
package usecases

import (
    "context"
    "github.com/your-project/internal/domain/entities"
    "github.com/your-project/internal/domain/repositories"
)

type UserUseCase struct {
    userRepo repositories.UserRepository
}

func NewUserUseCase(userRepo repositories.UserRepository) *UserUseCase {
    return &UserUseCase{userRepo: userRepo}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, user *entities.User) error {
    // Business logic here
    if user.Email == "" {
        return errors.New("email is required")
    }
    
    // Hash password
    hashedPassword, err := hashPassword(user.PasswordHash)
    if err != nil {
        return err
    }
    user.PasswordHash = hashedPassword
    
    // Save to repository
    return uc.userRepo.Create(ctx, user)
}
```

### HTTP Handler (Interface Adapters Layer)

```go
// internal/interfaces/http/handlers/user_handler.go
package handlers

import (
    "net/http"
    "github.com/your-project/internal/interfaces/http/dto"
    "github.com/your-project/internal/usecases"
)

type UserHandler struct {
    userUseCase *usecases.UserUseCase
}

func NewUserHandler(userUseCase *usecases.UserUseCase) *UserHandler {
    return &UserHandler{userUseCase: userUseCase}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
    var req dto.CreateUserRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
            Error: "Invalid request body",
        })
    }
    
    // Convert DTO to entity
    user := &entities.User{
        Email:    req.Email,
        Name:     req.Name,
        PasswordHash: req.Password,
    }
    
    // Call use case
    if err := h.userUseCase.CreateUser(c.Context(), user); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
            Error: err.Error(),
        })
    }
    
    // Convert entity to response DTO
    resp := dto.UserResponse{
        ID:    user.ID,
        Email: user.Email,
        Name:  user.Name,
    }
    
    return c.Status(http.StatusCreated).JSON(dto.SuccessResponse{
        Data: resp,
    })
}
```

## Data Flow Example

1. **HTTP Request** → Handler (`interfaces/http/handlers`)
2. **Handler** → Use Case (`usecases`)
3. **Use Case** → Repository Interface (`domain/repositories`)
4. **Repository Interface** → Repository Implementation (`infrastructure/database/repositories`)
5. **Database** → Returns data through the same chain
6. **Response** → Transformed by DTOs and returned to client

## Package Structure

### Domain Entities

All business entities are defined in `internal/domain/entities/` with no external dependencies:

```go
// internal/domain/entities/user.go
package entities

type User struct {
    ID           int
    Email        string
    PasswordHash string
    Name         string
    Phone        string
    Role         string
    EmailVerified bool
    IsActive     bool
}
```

### Repository Interfaces

Repository interfaces define the contract for data access:

```go
// internal/domain/repositories/user_repository.go
package repositories

type UserRepository interface {
    Create(ctx context.Context, user *entities.User) error
    FindByID(ctx context.Context, id int) (*entities.User, error)
    // ... other methods
}
```

### Use Cases

Use cases contain application business logic:

```go
// internal/usecases/user_usecase.go
package usecases

type UserUseCase struct {
    userRepo repositories.UserRepository
}

func (uc *UserUseCase) CreateUser(ctx context.Context, user *entities.User) error {
    // Business logic implementation
}
```

### HTTP Handlers

HTTP handlers convert HTTP requests to use case calls and return HTTP responses:

```go
// internal/interfaces/http/handlers/user_handler.go
package handlers

type UserHandler struct {
    userUseCase *usecases.UserUseCase
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
    // HTTP request handling
}
```

## Benefits of This Structure

1. **Framework Independence**: Business logic is not tied to any specific framework
2. **Testability**: Each layer can be tested independently
3. **Independence**: Layers can be developed and modified independently
4. **Isolation**: Business rules are isolated from technical details
5. **Reusability**: Business logic can be reused across different interfaces

## Migration from Traditional MVC

Traditional MVC structures often mix business logic with presentation logic. Clean Architecture separates these concerns:

### MVC Approach
```
controllers/
├── user_controller.go  (contains business logic)
├── product_controller.go
└── order_controller.go
models/
├── user.go
├── product.go
└── order.go
```

### Clean Architecture Approach
```
domain/
├── entities/
└── repositories/
usecases/
├── user_usecase.go
├── product_usecase.go
└── order_usecase.go
interfaces/
├── http/
│   └── handlers/
│       ├── user_handler.go
│       ├── product_handler.go
│       └── order_handler.go
infrastructure/
└── database/
    └── repositories/
```

## Best Practices

1. **Keep domain entities pure** - No framework dependencies
2. **Define clear interfaces** - Repository interfaces should be in domain layer
3. **Implement interfaces in infrastructure** - Concrete implementations in infrastructure layer
4. **Use dependency injection** - Wire up dependencies in main function
5. **Keep use cases focused** - Each use case should have a single responsibility
6. **Separate data transfer objects** - Use DTOs for external data representation
7. **Follow dependency rule** - Dependencies should point inward only

This Clean Architecture implementation ensures that the E-Commerce API is scalable, maintainable, and follows industry best practices for software architecture.