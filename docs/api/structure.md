# API Structure

## Overview

This document describes the API structure and endpoint organization for the E-Commerce API application.

## Base URL

```
http://localhost:8080/api/v1
```

## API Endpoints

### Authentication

#### POST /auth/register
Register a new user

#### POST /auth/login
Login to the system

#### GET /auth/profile
Get current user profile

#### PUT /auth/profile
Update current user profile

#### POST /auth/password/change
Change password

#### POST /auth/password/forgot
Forgot password

#### POST /auth/password/reset
Reset password

### Users (Admin Only)

#### GET /users
Get all users

#### GET /users/:id
Get a specific user

#### PUT /users/:id
Update a user

#### DELETE /users/:id
Delete a user

### Categories

#### GET /categories
Get all categories

#### GET /categories/:id
Get a specific category

#### POST /categories (Admin Only)
Create a new category

#### PUT /categories/:id (Admin Only)
Update a category

#### DELETE /categories/:id (Admin Only)
Delete a category

### Products

#### GET /products
Get all products with pagination and filtering

#### GET /products/:id
Get a specific product

#### POST /products (Admin Only)
Create a new product

#### PUT /products/:id (Admin Only)
Update a product

#### DELETE /products/:id (Admin Only)
Delete a product

### Product Images

#### POST /products/:id/images (Admin Only)
Upload product images

#### DELETE /products/:id/images/:imageId (Admin Only)
Delete a product image

### Cart

#### GET /cart
Get current user's cart

#### POST /cart/items
Add item to cart

#### PUT /cart/items/:id
Update cart item quantity

#### DELETE /cart/items/:id
Remove item from cart

#### DELETE /cart
Clear cart

### Addresses

#### GET /addresses
Get current user's addresses

#### GET /addresses/:id
Get a specific address

#### POST /addresses
Create a new address

#### PUT /addresses/:id
Update an address

#### DELETE /addresses/:id
Delete an address

#### PUT /addresses/:id/default
Set address as default

### Orders

#### POST /orders
Create a new order

#### GET /orders
Get user's orders

#### GET /orders/:id
Get a specific order

#### PUT /orders/:id/cancel
Cancel an order

### Order Items

#### GET /orders/:id/items
Get order items

### Payments

#### POST /payments
Process a payment

#### GET /payments
Get user's payments

#### GET /payments/:id
Get a specific payment

### Admin Orders

#### GET /admin/orders (Admin Only)
Get all orders

#### GET /admin/orders/:id (Admin Only)
Get a specific order

#### PUT /admin/orders/:id/status (Admin Only)
Update order status

### Admin Payments

#### GET /admin/payments (Admin Only)
Get all payments

#### GET /admin/payments/:id (Admin Only)
Get a specific payment

#### PUT /admin/payments/:id/status (Admin Only)
Update payment status

## HTTP Status Codes

| Code | Description |
|------|-------------|
| 200 | OK |
| 201 | Created |
| 204 | No Content |
| 400 | Bad Request |
| 401 | Unauthorized |
| 403 | Forbidden |
| 404 | Not Found |
| 409 | Conflict |
| 422 | Unprocessable Entity |
| 500 | Internal Server Error |

## Request/Response Format

### Request Format

All requests should use JSON format:
```
Content-Type: application/json
```

### Response Format

All responses are in JSON format:
```json
{
  "data": {},
  "meta": {},
  "error": {}
}
```

## Authentication

Most endpoints require authentication using JWT tokens. To authenticate, include the token in the Authorization header:

```
Authorization: Bearer <token>
```

## Rate Limiting

The API implements rate limiting to prevent abuse:
- 100 requests per minute per IP address
- 1000 requests per hour per user

## Pagination

List endpoints support pagination with the following query parameters:
- `page`: Page number (default: 1)
- `limit`: Items per page (default: 10, max: 100)

## Filtering and Sorting

List endpoints support filtering and sorting:
- Filtering: `?filter[field]=value`
- Sorting: `?sort=field` or `?sort=-field` (descending)

## Error Handling

All errors follow this format:
```json
{
  "error": {
    "code": "ERROR_CODE",
    "message": "Human readable error message",
    "details": {}
  }
}
```

## Versioning

The API is versioned using URL path versioning:
```
/api/v1/endpoint
```

## CORS

The API supports CORS with the following configuration:
- Allowed origins: Configurable via environment variables
- Allowed methods: GET, POST, PUT, DELETE, OPTIONS
- Allowed headers: Content-Type, Authorization, X-Requested-With, Accept, Origin

## Security

### HTTPS

In production, the API should be served over HTTPS.

### Input Validation

All input is validated using struct tags and custom validation rules.

### Sanitization

All user input is sanitized to prevent XSS attacks.

### Rate Limiting

Rate limiting is implemented to prevent abuse.

### Authentication

JWT tokens are used for authentication with proper expiration.

### Authorization

Role-based access control (RBAC) is implemented for admin endpoints.

## Monitoring

### Health Checks

- `/health` - Basic health check
- `/health/db` - Database connectivity check
- `/health/redis` - Redis connectivity check

### Metrics

Prometheus metrics are exposed at `/metrics`.

## Documentation

API documentation is available in OpenAPI (Swagger) format.

## Testing

API endpoints can be tested using the provided Postman collection or curl commands.