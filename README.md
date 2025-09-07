# E-Commerce API

A robust, scalable e-commerce API built with Go following Clean Architecture principles.

## Table of Contents

- [Features](#features)
- [Architecture](#architecture)
- [Tech Stack](#tech-stack)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
  - [Installation](#installation)
  - [Configuration](#configuration)
  - [Running the Application](#running-the-application)
- [API Documentation](#api-documentation)
- [Project Structure](#project-structure)
- [Database Schema](#database-schema)
- [Testing](#testing)
- [Deployment](#deployment)
- [Contributing](#contributing)
- [License](#license)

## Features

- User authentication and authorization (JWT)
- Product catalog management
- Shopping cart functionality
- Order processing
- Payment integration
- Category management
- User address management
- Admin dashboard
- Rate limiting
- Logging and monitoring
- Docker support
- Comprehensive API documentation

## Architecture

This project follows Clean Architecture principles with a clear separation of concerns:

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

For detailed architecture documentation, see [Clean Architecture](docs/development/clean-architecture.md).

## Tech Stack

- **Language**: Go 1.21+
- **Web Framework**: Fiber
- **Database**: PostgreSQL
- **ORM**: GORM
- **Authentication**: JWT
- **Validation**: Validator
- **Logging**: Zap
- **Testing**: Testify, Mock
- **Containerization**: Docker, Docker Compose
- **Documentation**: Swagger/OpenAPI
- **Build Tool**: Make

## Prerequisites

- Go 1.21+
- Docker and Docker Compose
- PostgreSQL 13+ (if running without Docker)
- Redis 6+ (if running without Docker)

## Getting Started

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/e-commerce-API.git
   cd e-commerce-API
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

### Configuration

1. Copy the example environment file:
   ```bash
   cp .env.example .env
   ```

2. Update the `.env` file with your configuration:
   ```bash
   # Database
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=ecommerce_user
   DB_PASSWORD=ecommerce_password
   DB_NAME=ecommerce_db
   
   # JWT
   JWT_SECRET=your_jwt_secret_key
   JWT_EXPIRATION_HOURS=24
   
   # Server
   SERVER_PORT=8080
   ```

### Running the Application

#### Option 1: Using Docker (Recommended)

```bash
# Start all services
docker-compose up -d

# Access the API at http://localhost:8080
```

#### Option 2: Local Development

```bash
# Start database and Redis
docker-compose up -d db redis

# Run the application
go run cmd/api/main.go
```

#### Option 3: Using Makefile

```bash
# Install air for hot reloading (first time only)
make install-air

# Run with hot reloading
make run-dev
```

## API Documentation

API documentation is available in multiple formats:

- [API Structure](docs/api/structure.md)
- [Endpoints Documentation](docs/api/README.md)
- Swagger/OpenAPI documentation (available at `/swagger` when running)

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
└── ...
```

For detailed project structure, see [Project Structure](docs/development/project-structure.md).

## Database Schema

The database schema includes tables for users, products, categories, orders, payments, and more.

For detailed schema documentation, see [Database Schema](docs/development/database.md).

## Testing

Run tests with the following commands:

```bash
# Run all tests
make test

# Run tests with coverage
make test-cover

# Run specific package tests
go test -v ./internal/usecases/...
```

For detailed testing documentation, see [Testing Guide](docs/development/testing.md).

## Deployment

### Docker Deployment

```bash
# Build the Docker image
docker build -t ecommerce-api .

# Run the container
docker run -d -p 8080:8080 --env-file .env ecommerce-api
```

### Docker Compose Deployment

```bash
# Start all services
docker-compose up -d

# Scale the application
docker-compose up -d --scale app=3
```

For detailed deployment documentation, see [Deployment Guide](docs/deployment/README.md).

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a pull request

For development guidelines, see [Development Guide](docs/development/README.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Additional Documentation

- [Clean Architecture Implementation](docs/development/clean-architecture.md)
- [Database Migrations](docs/development/migrations.md)
- [API Structure](docs/api/structure.md)
- [Deployment Guide](docs/deployment/README.md)
- [Development Guide](docs/development/README.md)
