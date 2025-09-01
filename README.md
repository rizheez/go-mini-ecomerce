# 🛒 Mini E-Commerce API

A comprehensive e-commerce REST API built with Go Fiber, implementing Clean Architecture principles. This API provides complete e-commerce functionality including user management, product catalog, shopping cart, order processing, and payment handling.

## 🚀 Features

- **👤 User Management**

  - User registration and authentication
  - JWT-based authorization
  - Role-based access control (Customer/Admin)
  - User profile management
  - Multiple shipping addresses

- **📦 Product Management**

  - Product CRUD operations
  - Category management
  - Product search and filtering
  - Image management
  - Stock tracking

- **🛒 Shopping Cart**

  - Add/remove items
  - Update quantities
  - Persistent cart storage

- **📋 Order Management**

  - Order creation and tracking
  - Order status updates
  - Order history

- **💳 Payment Processing**
  - Credit card payments (simulation)
  - Cash on delivery
  - Payment tracking

## 🏗️ Architecture

This project follows **Clean Architecture** principles with clear separation of concerns:

```
├── cmd/api/                 # Application entry point
├── internal/
│   ├── domain/             # Business entities and interfaces
│   ├── usecases/           # Business logic
│   ├── interfaces/http/    # HTTP handlers, middleware, routes
│   └── infrastructure/     # Database, external services
├── pkg/                    # Shared utilities
├── config/                 # Configuration management
└── migrations/             # Database migrations
```

## 🛠️ Tech Stack

- **Framework:** [Fiber v2](https://docs.gofiber.io/) - Fast HTTP framework
- **Database:** PostgreSQL with [GORM](https://gorm.io/) ORM
- **Authentication:** JWT tokens
- **Validation:** [go-playground/validator](https://github.com/go-playground/validator)
- **Environment:** [godotenv](https://github.com/joho/godotenv)
- **Testing:** [Testify](https://github.com/stretchr/testify)

## 📋 Prerequisites

- Go 1.21 or higher
- PostgreSQL 12 or higher
- Git

## 📖 API Documentation

### Base URL

```
http://localhost:3000/api/v1
```

### Authentication

Include JWT token in the Authorization header:

```
Authorization: Bearer <your_jwt_token>
```

### Key Endpoints

#### 🔐 Authentication

- `POST /auth/register` - Register new user
- `POST /auth/login` - User login
- `POST /auth/logout` - User logout

#### 👤 User Management

- `GET /users/profile` - Get user profile
- `PUT /users/profile` - Update profile
- `GET /users/addresses` - Get user addresses
- `POST /users/addresses` - Add new address

#### 📦 Products

- `GET /products` - Get all products (with search & filters)
- `GET /products/:id` - Get product details
- `POST /admin/products` - Create product (Admin only)
- `PUT /admin/products/:id` - Update product (Admin only)

#### 🛒 Shopping Cart

- `GET /cart` - Get user's cart
- `POST /cart/items` - Add item to cart
- `PUT /cart/items/:id` - Update cart item
- `DELETE /cart/items/:id` - Remove item from cart

#### 📋 Orders

- `GET /orders` - Get user orders
- `POST /orders` - Create new order
- `GET /orders/:id` - Get order details

#### 💳 Payments

- `POST /payments/process` - Process payment

## 🛡️ Security Features

- **JWT Authentication** with secure token generation
- **Password Hashing** using bcrypt
- **Rate Limiting** to prevent abuse
- **CORS Protection** for web browser security
- **Input Validation** on all endpoints
- **SQL Injection Prevention** through GORM ORM
- **UUID Primary Keys** to prevent enumeration attacks

## 📊 Database Schema

The database uses UUID primary keys and includes the following main entities:

- **users** - User accounts and authentication
- **user_addresses** - Customer shipping addresses
- **categories** - Product categories
- **products** - Product catalog
- **carts** - Shopping carts
- **orders** - Customer orders
- **payments** - Payment transactions

## 🔧 Development Tools

### Live Reload

```bash
# Install air
go install github.com/cosmtrek/air@latest

# Run with live reload
air
```

### Database Migrations

```bash
# Create migration
migrate create -ext sql -dir migrations -seq create_users_table

# Run migrations
migrate -path migrations -database "postgres://user:pass@localhost/dbname?sslmode=disable" up
```

### Code Quality

```bash
# Format code
go fmt ./...

# Lint code
golangci-lint run

# Security check
gosec ./...
```

## 📁 Project Structure

```
mini-ecommerce/
├── cmd/api/                    # Application entrypoint
├── config/                     # Configuration management
├── internal/
│   ├── domain/
│   │   ├── entities/          # Business entities
│   │   └── repositories/      # Repository interfaces
│   ├── usecases/              # Business logic
│   ├── interfaces/http/
│   │   ├── handlers/          # HTTP handlers
│   │   ├── middleware/        # HTTP middleware
│   │   ├── routes/            # Route definitions
│   │   └── dto/               # Data transfer objects
│   └── infrastructure/
│       ├── database/          # Database implementations
│       ├── auth/              # Authentication services
│       └── payment/           # Payment services
├── pkg/                       # Shared utilities
├── migrations/                # Database migrations
├── docs/                      # Documentation
├── scripts/                   # Build and deployment scripts
├── .env.example               # Environment variables template
├── .air.toml                  # Live reload configuration
├── docker-compose.yml         # Docker development setup
├── Dockerfile                 # Production Docker image
└── README.md                  # This file
```

## 🙏 Acknowledgments

- [Fiber](https://docs.gofiber.io/) - Amazing Go web framework
- [GORM](https://gorm.io/) - Fantastic Go ORM
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) - Architecture principles
- Go community for excellent tools and libraries
