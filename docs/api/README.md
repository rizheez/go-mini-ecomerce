# API Documentation

## Overview

This document provides documentation for the E-Commerce API endpoints.

## Authentication

Most endpoints require authentication using JWT tokens. To authenticate, include the token in the Authorization header:

```
Authorization: Bearer <token>
```

## Base URL

```
http://localhost:8080/api/v1
```

## Endpoints

### Authentication

#### POST /auth/register
Register a new user

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "password123",
  "name": "John Doe",
  "phone": "+1234567890"
}
```

**Response:**
```json
{
  "token": "jwt_token_here",
  "user": {
    "id": 1,
    "email": "user@example.com",
    "name": "John Doe",
    "phone": "+1234567890",
    "role": "customer"
  }
}
```

#### POST /auth/login
Login to the system

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "token": "jwt_token_here",
  "user": {
    "id": 1,
    "email": "user@example.com",
    "name": "John Doe",
    "phone": "+1234567890",
    "role": "customer"
  }
}
```

#### GET /auth/profile
Get current user profile

**Response:**
```json
{
  "id": 1,
  "email": "user@example.com",
  "name": "John Doe",
  "phone": "+1234567890",
  "role": "customer"
}
```

### Categories

#### GET /categories
Get all categories

**Response:**
```json
[
  {
    "id": 1,
    "name": "Electronics",
    "description": "Electronic devices and accessories",
    "image_url": "https://example.com/electronics.jpg",
    "is_active": true,
    "sort_order": 1
  }
]
```

#### GET /categories/:id
Get a specific category

**Response:**
```json
{
  "id": 1,
  "name": "Electronics",
  "description": "Electronic devices and accessories",
  "image_url": "https://example.com/electronics.jpg",
  "is_active": true,
  "sort_order": 1
}
```

### Products

#### GET /products
Get all products with pagination

**Query Parameters:**
- page (optional): Page number (default: 1)
- limit (optional): Items per page (default: 10)
- category_id (optional): Filter by category
- search (optional): Search by name or description

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "name": "Smartphone",
      "description": "Latest model smartphone",
      "price": 699.99,
      "stock_quantity": 50,
      "category_id": 1,
      "sku": "SP-001",
      "specifications": {
        "brand": "TechBrand",
        "model": "X1"
      },
      "is_active": true,
      "weight": 0.2,
      "dimensions": {
        "length": 15,
        "width": 7,
        "height": 0.8
      }
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 100,
    "pages": 10
  }
}
```

#### GET /products/:id
Get a specific product

**Response:**
```json
{
  "id": 1,
  "name": "Smartphone",
  "description": "Latest model smartphone",
  "price": 699.99,
  "stock_quantity": 50,
  "category_id": 1,
  "sku": "SP-001",
  "specifications": {
    "brand": "TechBrand",
    "model": "X1"
  },
  "is_active": true,
  "weight": 0.2,
  "dimensions": {
    "length": 15,
    "width": 7,
    "height": 0.8
  }
}
```

### Cart

#### GET /cart
Get current user's cart

**Response:**
```json
{
  "id": 1,
  "user_id": 1,
  "items": [
    {
      "id": 1,
      "product_id": 1,
      "product_name": "Smartphone",
      "quantity": 2,
      "unit_price": 699.99,
      "total_price": 1399.98
    }
  ],
  "total_items": 2,
  "total_amount": 1399.98
}
```

#### POST /cart/items
Add item to cart

**Request Body:**
```json
{
  "product_id": 1,
  "quantity": 2
}
```

**Response:**
```json
{
  "message": "Item added to cart successfully"
}
```

#### PUT /cart/items/:id
Update cart item quantity

**Request Body:**
```json
{
  "quantity": 3
}
```

**Response:**
```json
{
  "message": "Cart item updated successfully"
}
```

#### DELETE /cart/items/:id
Remove item from cart

**Response:**
```json
{
  "message": "Item removed from cart successfully"
}
```

### Orders

#### POST /orders
Create a new order

**Request Body:**
```json
{
  "shipping_address_id": 1,
  "payment_method": "credit_card"
}
```

**Response:**
```json
{
  "id": "ORD-202301010001",
  "user_id": 1,
  "status": "pending",
  "subtotal": 1399.98,
  "shipping_cost": 10.00,
  "tax_amount": 112.00,
  "total_amount": 1521.98,
  "payment_method": "credit_card",
  "payment_status": "pending",
  "items": [
    {
      "id": 1,
      "product_id": 1,
      "product_name": "Smartphone",
      "quantity": 2,
      "unit_price": 699.99,
      "total_price": 1399.98
    }
  ]
}
```

#### GET /orders
Get user's orders

**Response:**
```json
[
  {
    "id": "ORD-202301010001",
    "user_id": 1,
    "status": "pending",
    "total_amount": 1521.98,
    "payment_method": "credit_card",
    "payment_status": "pending",
    "created_at": "2023-01-01T10:00:00Z"
  }
]
```

#### GET /orders/:id
Get a specific order

**Response:**
```json
{
  "id": "ORD-202301010001",
  "user_id": 1,
  "status": "pending",
  "subtotal": 1399.98,
  "shipping_cost": 10.00,
  "tax_amount": 112.00,
  "total_amount": 1521.98,
  "payment_method": "credit_card",
  "payment_status": "pending",
  "shipping_address": {
    "recipient_name": "John Doe",
    "address_line_1": "123 Main St",
    "city": "New York",
    "state": "NY",
    "postal_code": "10001",
    "country": "USA"
  },
  "items": [
    {
      "id": 1,
      "product_id": 1,
      "product_name": "Smartphone",
      "quantity": 2,
      "unit_price": 699.99,
      "total_price": 1399.98
    }
  ],
  "created_at": "2023-01-01T10:00:00Z"
}
```

### Payments

#### POST /payments
Process a payment

**Request Body:**
```json
{
  "order_id": "ORD-202301010001",
  "payment_method": "credit_card",
  "payment_details": {
    "card_number": "4111111111111111",
    "expiry_month": 12,
    "expiry_year": 2025,
    "cvv": "123"
  }
}
```

**Response:**
```json
{
  "id": "PAY-202301010001",
  "order_id": "ORD-202301010001",
  "amount": 1521.98,
  "payment_method": "credit_card",
  "status": "completed",
  "transaction_id": "txn_1234567890"
}
```

## Error Responses

All error responses follow this format:

```json
{
  "error": "Error message",
  "code": "ERROR_CODE"
}
```

Common error codes:
- VALIDATION_ERROR: Request validation failed
- UNAUTHORIZED: Authentication required
- FORBIDDEN: Access denied
- NOT_FOUND: Resource not found
- INTERNAL_ERROR: Server error