# Testing Guide

## Overview

This document explains how to test the E-Commerce API application.

## Test Structure

```
e-commerce-API/
├── tests/
│   ├── integration/      # Integration tests
│   ├── unit/            # Unit tests
│   └── fixtures/        # Test data fixtures
├── coverage.out         # Coverage report
└── coverage.html        # Coverage report (HTML)
```

## Testing Frameworks

- **Unit Testing**: Built-in `testing` package
- **Assertion Library**: `github.com/stretchr/testify`
- **Mocking**: `github.com/golang/mock`
- **Integration Testing**: `testing` package with real database

## Running Tests

### Unit Tests

Run unit tests:
```bash
# Run all unit tests
make test

# Run specific package tests
go test -v ./internal/usecases/...

# Run tests with coverage
make test-cover

# Run tests with race detection
go test -race ./...
```

### Integration Tests

Run integration tests:
```bash
# Run all integration tests
go test -v ./tests/integration/...

# Run specific integration test
go test -v ./tests/integration/products_test.go
```

### All Tests

Run all tests:
```bash
# Run all tests
go test -v ./...

# Run all tests with coverage
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

## Writing Tests

### Unit Tests

Unit tests should:
1. Test one function/method at a time
2. Use mocks for dependencies
3. Cover all code paths
4. Be fast and isolated

Example unit test:
```go
func TestUserUseCase_CreateUser(t *testing.T) {
    // Setup
    mockRepo := new(mocks.UserRepository)
    uc := NewUserUseCase(mockRepo)
    
    // Mock expectations
    mockRepo.On("Create", mock.Anything, mock.Anything).Return(nil)
    
    // Test
    user := &entities.User{
        Email: "test@example.com",
        Name:  "Test User",
    }
    
    err := uc.CreateUser(context.Background(), user)
    
    // Assertions
    assert.NoError(t, err)
    mockRepo.AssertExpectations(t)
}
```

### Integration Tests

Integration tests should:
1. Test the interaction between components
2. Use real database connections
3. Test with actual data
4. Clean up after themselves

Example integration test:
```go
func TestProductRepository_Create(t *testing.T) {
    // Setup
    db := setupTestDB()
    repo := NewProductRepository(db)
    
    // Test
    product := &entities.Product{
        Name:  "Test Product",
        Price: 99.99,
    }
    
    err := repo.Create(context.Background(), product)
    
    // Assertions
    assert.NoError(t, err)
    assert.NotZero(t, product.ID)
    
    // Cleanup
    repo.Delete(context.Background(), product.ID)
}
```

## Test Coverage

### Checking Coverage

Check test coverage:
```bash
# Generate coverage report
go test -coverprofile=coverage.out ./...

# View coverage in terminal
go tool cover -func=coverage.out

# View coverage in browser
go tool cover -html=coverage.out
```

### Coverage Requirements

Minimum coverage requirements:
- Overall: 80%
- Critical business logic: 90%
- API handlers: 85%
- Repositories: 90%

## Mocking

### Generating Mocks

Generate mocks using mockgen:
```bash
# Generate mocks for all repositories
make mock

# Generate mock for specific interface
mockgen -source=internal/domain/repositories/user_repository.go -destination=internal/domain/repositories/mocks/user_repository_mock.go
```

### Using Mocks

Use mocks in tests:
```go
// Create mock
mockRepo := new(mocks.UserRepository)

// Set expectations
mockRepo.On("FindByID", mock.Anything, 1).Return(&entities.User{ID: 1, Name: "John"}, nil)

// Use in use case
uc := NewUserUseCase(mockRepo)
user, err := uc.GetUser(context.Background(), 1)

// Assert expectations
mockRepo.AssertExpectations(t)
```

## Test Data

### Fixtures

Use fixtures for test data:
```go
// Load fixture
func loadFixture(filename string) []byte {
    data, err := ioutil.ReadFile("tests/fixtures/" + filename)
    if err != nil {
        panic(err)
    }
    return data
}

// Use in test
userData := loadFixture("user.json")
```

### Factories

Create test data factories:
```go
func createUserFactory() *entities.User {
    return &entities.User{
        Email: "test@example.com",
        Name:  "Test User",
        Role:  "customer",
    }
}
```

## Continuous Integration

### GitHub Actions

Example GitHub Actions workflow:
```yaml
name: Test

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    services:
      postgres:
        image: postgres:13
        env:
          POSTGRES_PASSWORD: postgres
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
        ports:
          - 5432:5432
    steps:
      - uses: actions/checkout@v2
      - name: Set up Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.21
      - name: Install dependencies
        run: go mod download
      - name: Run tests
        run: make test-cover
        env:
          DB_HOST: localhost
          DB_PORT: 5432
          DB_USER: postgres
          DB_PASSWORD: postgres
          DB_NAME: postgres
```

## Performance Testing

### Benchmark Tests

Write benchmark tests:
```go
func BenchmarkUserUseCase_GetUser(b *testing.B) {
    // Setup
    mockRepo := new(mocks.UserRepository)
    uc := NewUserUseCase(mockRepo)
    
    mockRepo.On("FindByID", mock.Anything, 1).Return(&entities.User{ID: 1, Name: "John"}, nil)
    
    // Benchmark
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        uc.GetUser(context.Background(), 1)
    }
}
```

Run benchmark tests:
```bash
# Run benchmarks
go test -bench=. ./...

# Run specific benchmark
go test -bench=BenchmarkUserUseCase_GetUser ./...
```

## Test Best Practices

### 1. Test Organization

- Group related tests in test suites
- Use descriptive test names
- Keep tests focused and small
- Separate unit and integration tests

### 2. Test Data

- Use factories for test data creation
- Clean up test data after tests
- Use fixtures for complex data structures
- Avoid hardcoding test values

### 3. Mocking

- Mock only external dependencies
- Verify mock expectations
- Use realistic mock behavior
- Don't mock value objects

### 4. Assertions

- Use specific assertions
- Provide meaningful assertion messages
- Assert only what's necessary
- Use table-driven tests for multiple scenarios

### 5. Test Speed

- Keep tests fast
- Use parallel testing when possible
- Avoid unnecessary setup/teardown
- Use in-memory databases for integration tests

## Debugging Tests

### Verbose Output

Run tests with verbose output:
```bash
go test -v ./...
```

### Debugging with Delve

Debug tests with Delve:
```bash
# Install delve
go install github.com/go-delve/delve/cmd/dlv@latest

# Debug test
dlv test ./internal/usecases/ -- -test.run TestUserUseCase_CreateUser
```

### Logging in Tests

Add logging to tests:
```go
func TestSomething(t *testing.T) {
    t.Log("Starting test")
    // ... test code ...
    t.Logf("Result: %v", result)
}
```

## Common Test Patterns

### Table-Driven Tests

Use table-driven tests for multiple scenarios:
```go
func TestValidateEmail(t *testing.T) {
    tests := []struct {
        email   string
        isValid bool
    }{
        {"valid@example.com", true},
        {"invalid.email", false},
        {"", false},
    }
    
    for _, tt := range tests {
        t.Run(tt.email, func(t *testing.T) {
            result := validateEmail(tt.email)
            assert.Equal(t, tt.isValid, result)
        })
    }
}
```

### Subtests

Use subtests for related test cases:
```go
func TestUserService(t *testing.T) {
    t.Run("CreateUser", func(t *testing.T) {
        // Test create user
    })
    
    t.Run("GetUser", func(t *testing.T) {
        // Test get user
    })
}
```

## Test Utilities

### Helper Functions

Create helper functions for common test operations:
```go
func setupTestDB() *gorm.DB {
    // Setup test database
}

func teardownTestDB(db *gorm.DB) {
    // Teardown test database
}

func createTestUser(db *gorm.DB) *entities.User {
    // Create test user
}
```

### Test Configuration

Use test configuration:
```go
func TestMain(m *testing.M) {
    // Setup
    os.Setenv("DB_HOST", "localhost")
    os.Setenv("DB_PORT", "5432")
    
    // Run tests
    code := m.Run()
    
    // Teardown
    os.Exit(code)
}
```

## Troubleshooting

### Common Issues

1. **Test database connection failed**
   - Check database credentials
   - Verify database is running
   - Check firewall settings

2. **Mock expectations not met**
   - Check method signatures
   - Verify call counts
   - Ensure proper mock setup

3. **Test data not cleaned up**
   - Add cleanup functions
   - Use transactional tests
   - Reset database between tests

4. **Tests running slow**
   - Use in-memory databases
   - Parallelize tests
   - Optimize test setup

### Debugging Tips

1. Use `t.Log` for debugging output
2. Run specific tests to isolate issues
3. Use Delve for interactive debugging
4. Check test coverage to identify untested code