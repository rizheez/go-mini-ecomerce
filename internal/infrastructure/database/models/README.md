# Entity Models Structure

## Overview

This project follows Clean Architecture principles by separating domain entities from data models:

1. **Domain Entities** (`internal/domain/entities/`) - Pure business entities without framework dependencies
2. **Data Models** (`internal/infrastructure/database/models/`) - GORM-specific models with database tags

## Domain Entities vs Data Models

### Domain Entities
- Located in `internal/domain/entities/`
- Contain only business logic and properties
- No framework-specific tags or dependencies
- Used throughout the business logic layer

### Data Models
- Located in `internal/infrastructure/database/models/`
- Contain GORM tags for database mapping
- Handle database-specific concerns
- Used only in the infrastructure layer

## Entity Relationships

The following entities are defined in the system:

1. **User** - Represents a system user
2. **UserAddress** - Shipping addresses for users
3. **Category** - Product categories
4. **Product** - Products in the catalog
5. **ProductImage** - Images associated with products
6. **Cart** - Shopping carts for users
7. **CartItem** - Items within shopping carts
8. **Order** - Customer orders
9. **OrderItem** - Items within orders
10. **OrderStatusHistory** - History of order status changes
11. **Payment** - Payment transactions
12. **ProductReview** - Product reviews (optional feature)

## Implementation Notes

1. **JSONB Fields**: Both domain entities and data models handle JSONB fields differently:
   - Domain entities use `map[string]interface{}`
   - Data models use a shared `JSONB` type defined in `types.go` with Scan/Value methods for GORM

2. **Time Fields**: Time fields are handled in the data models with GORM tags but omitted in domain entities for purity.

3. **Primary Keys**: Defined in data models with GORM tags but kept simple in domain entities.

4. **Shared Types**: Custom types like `JSONB` are defined in `types.go` to avoid redeclaration errors.

This separation ensures that business logic remains framework-agnostic and testable.