# Mini E-Commerce API Specification

## Overview

This is a comprehensive API specification for a mini e-commerce platform built with Go Fiber. The API supports authentication, user management, product catalog, shopping cart, orders, and payment processing.

## Base URL

```
http://localhost:3000/api/v1
```

## Authentication

The API uses JWT (JSON Web Tokens) for authentication. Include the token in the Authorization header:

```
Authorization: Bearer <jwt_token>
```

## User Roles

- **customer**: Regular customer with full shopping capabilities
- **admin**: Administrative access to manage products, categories, and orders

---

## 1. Authentication Endpoints

### POST /auth/register

Register a new user account.

**Request Body:**

```json
{
  "email": "user@example.com",
  "password": "password123",
  "name": "John Doe",
  "phone": "+1234567890",
  "role": "customer"
}
```

**Response (201):**

```json
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "user": {
      "id": 1,
      "email": "user@example.com",
      "name": "John Doe",
      "phone": "+1234567890",
      "role": "customer",
      "created_at": "2025-09-01T10:00:00Z"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### POST /auth/login

Authenticate user and get access token.

**Request Body:**

```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response (200):**

```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "user": {
      "id": 1,
      "email": "user@example.com",
      "name": "John Doe",
      "role": "customer"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### POST /auth/logout

Logout user (invalidate token).

**Headers:** `Authorization: Bearer <token>`

**Response (200):**

```json
{
  "success": true,
  "message": "Logout successful"
}
```

---

## 2. User Management Endpoints

### GET /users/profile

Get current user profile.

**Headers:** `Authorization: Bearer <token>`

**Response (200):**

```json
{
  "success": true,
  "data": {
    "id": 1,
    "email": "user@example.com",
    "name": "John Doe",
    "phone": "+1234567890",
    "role": "customer",
    "addresses": [
      {
        "id": 1,
        "label": "Home",
        "recipient_name": "John Doe",
        "phone": "+1234567890",
        "address_line_1": "123 Main Street",
        "address_line_2": "Apt 4B",
        "city": "New York",
        "state": "NY",
        "postal_code": "10001",
        "country": "USA",
        "is_default": true
      }
    ],
    "created_at": "2025-09-01T10:00:00Z",
    "updated_at": "2025-09-01T10:00:00Z"
  }
}
```

### PUT /users/profile

Update user profile.

**Headers:** `Authorization: Bearer <token>`

**Request Body:**

```json
{
  "name": "John Smith",
  "phone": "+1234567891"
}
```

**Response (200):**

```json
{
  "success": true,
  "message": "Profile updated successfully",
  "data": {
    "id": 1,
    "email": "user@example.com",
    "name": "John Smith",
    "phone": "+1234567891",
    "role": "customer",
    "updated_at": "2025-09-01T11:00:00Z"
  }
}
```

### POST /users/addresses

Add new address.

**Headers:** `Authorization: Bearer <token>`

**Request Body:**

```json
{
  "label": "Work",
  "recipient_name": "John Doe",
  "phone": "+1234567890",
  "address_line_1": "456 Business Ave",
  "address_line_2": "Suite 200",
  "city": "New York",
  "state": "NY",
  "postal_code": "10002",
  "country": "USA",
  "is_default": false
}
```

**Response (201):**

```json
{
  "success": true,
  "message": "Address added successfully",
  "data": {
    "id": 2,
    "label": "Work",
    "recipient_name": "John Doe",
    "phone": "+1234567890",
    "address_line_1": "456 Business Ave",
    "address_line_2": "Suite 200",
    "city": "New York",
    "state": "NY",
    "postal_code": "10002",
    "country": "USA",
    "is_default": false,
    "created_at": "2025-09-01T10:30:00Z"
  }
}
```

### GET /users/addresses

Get all user addresses.

**Headers:** `Authorization: Bearer <token>`

**Response (200):**

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "label": "Home",
      "recipient_name": "John Doe",
      "phone": "+1234567890",
      "address_line_1": "123 Main Street",
      "address_line_2": "Apt 4B",
      "city": "New York",
      "state": "NY",
      "postal_code": "10001",
      "country": "USA",
      "is_default": true,
      "created_at": "2025-09-01T10:00:00Z"
    }
  ]
}
```

### PUT /users/addresses/:id

Update address.

**Headers:** `Authorization: Bearer <token>`

**Request Body:**

```json
{
  "label": "Home Address",
  "address_line_1": "123 Main Street Updated",
  "is_default": true
}
```

**Response (200):**

```json
{
  "success": true,
  "message": "Address updated successfully",
  "data": {
    "id": 1,
    "label": "Home Address",
    "address_line_1": "123 Main Street Updated",
    "is_default": true,
    "updated_at": "2025-09-01T11:00:00Z"
  }
}
```

### DELETE /users/addresses/:id

Delete address.

**Headers:** `Authorization: Bearer <token>`

**Response (200):**

```json
{
  "success": true,
  "message": "Address deleted successfully"
}
```

---

## 3. Category Management Endpoints

### GET /categories

Get all categories.

**Query Parameters:**

- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)

**Response (200):**

```json
{
  "success": true,
  "data": {
    "categories": [
      {
        "id": 1,
        "name": "Electronics",
        "description": "Electronic devices and accessories",
        "image_url": "https://example.com/electronics.jpg",
        "is_active": true,
        "product_count": 25,
        "created_at": "2025-09-01T10:00:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 3,
      "total_items": 25,
      "per_page": 10
    }
  }
}
```

### GET /categories/:id

Get category by ID.

**Response (200):**

```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "Electronics",
    "description": "Electronic devices and accessories",
    "image_url": "https://example.com/electronics.jpg",
    "is_active": true,
    "product_count": 25,
    "created_at": "2025-09-01T10:00:00Z"
  }
}
```

### POST /categories

Create new category (Admin only).

**Headers:** `Authorization: Bearer <admin_token>`

**Request Body:**

```json
{
  "name": "Fashion",
  "description": "Clothing and accessories",
  "image_url": "https://example.com/fashion.jpg"
}
```

**Response (201):**

```json
{
  "success": true,
  "message": "Category created successfully",
  "data": {
    "id": 2,
    "name": "Fashion",
    "description": "Clothing and accessories",
    "image_url": "https://example.com/fashion.jpg",
    "is_active": true,
    "product_count": 0,
    "created_at": "2025-09-01T10:30:00Z"
  }
}
```

### PUT /categories/:id

Update category (Admin only).

**Headers:** `Authorization: Bearer <admin_token>`

**Request Body:**

```json
{
  "name": "Fashion & Apparel",
  "description": "Clothing, shoes, and accessories",
  "is_active": true
}
```

**Response (200):**

```json
{
  "success": true,
  "message": "Category updated successfully",
  "data": {
    "id": 2,
    "name": "Fashion & Apparel",
    "description": "Clothing, shoes, and accessories",
    "is_active": true,
    "updated_at": "2025-09-01T11:00:00Z"
  }
}
```

### DELETE /categories/:id

Delete category (Admin only).

**Headers:** `Authorization: Bearer <admin_token>`

**Response (200):**

```json
{
  "success": true,
  "message": "Category deleted successfully"
}
```

---

## 4. Product Management Endpoints

### GET /products

Get all products with filtering and search.

**Query Parameters:**

- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 12)
- `category_id` (optional): Filter by category
- `search` (optional): Search in name and description
- `min_price` (optional): Minimum price filter
- `max_price` (optional): Maximum price filter
- `sort_by` (optional): Sort by field (name, price, created_at)
- `sort_order` (optional): Sort order (asc, desc)

**Response (200):**

```json
{
  "success": true,
  "data": {
    "products": [
      {
        "id": 1,
        "name": "iPhone 15 Pro",
        "description": "Latest iPhone with advanced features",
        "price": 999.99,
        "stock_quantity": 50,
        "category": {
          "id": 1,
          "name": "Electronics"
        },
        "images": [
          {
            "id": 1,
            "url": "https://example.com/iphone-1.jpg",
            "alt_text": "iPhone 15 Pro front view",
            "is_primary": true
          }
        ],
        "is_active": true,
        "rating": 4.8,
        "review_count": 124,
        "created_at": "2025-09-01T10:00:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 8,
      "total_items": 95,
      "per_page": 12
    }
  }
}
```

### GET /products/:id

Get product details by ID.

**Response (200):**

```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "iPhone 15 Pro",
    "description": "Latest iPhone with pro features including titanium design, Action Button, and 48MP camera system.",
    "price": 999.99,
    "stock_quantity": 50,
    "category": {
      "id": 1,
      "name": "Electronics",
      "description": "Electronic devices and accessories"
    },
    "images": [
      {
        "id": 1,
        "url": "https://example.com/iphone-1.jpg",
        "alt_text": "iPhone 15 Pro front view",
        "is_primary": true
      },
      {
        "id": 2,
        "url": "https://example.com/iphone-2.jpg",
        "alt_text": "iPhone 15 Pro back view",
        "is_primary": false
      }
    ],
    "specifications": {
      "storage": "128GB",
      "color": "Natural Titanium",
      "display": "6.1-inch Super Retina XDR",
      "chip": "A17 Pro"
    },
    "is_active": true,
    "rating": 4.8,
    "review_count": 124,
    "created_at": "2025-09-01T10:00:00Z",
    "updated_at": "2025-09-01T10:00:00Z"
  }
}
```

### POST /products

Create new product (Admin only).

**Headers:** `Authorization: Bearer <admin_token>`

**Request Body:**

```json
{
  "name": "Samsung Galaxy S24",
  "description": "Premium Android smartphone with AI features",
  "price": 799.99,
  "stock_quantity": 30,
  "category_id": 1,
  "images": [
    {
      "url": "https://example.com/galaxy-1.jpg",
      "alt_text": "Galaxy S24 front view",
      "is_primary": true
    }
  ],
  "specifications": {
    "storage": "256GB",
    "color": "Phantom Black",
    "display": "6.2-inch Dynamic AMOLED 2X",
    "chip": "Snapdragon 8 Gen 3"
  }
}
```

**Response (201):**

```json
{
  "success": true,
  "message": "Product created successfully",
  "data": {
    "id": 2,
    "name": "Samsung Galaxy S24",
    "description": "Premium Android smartphone with AI features",
    "price": 799.99,
    "stock_quantity": 30,
    "category_id": 1,
    "is_active": true,
    "rating": 0,
    "review_count": 0,
    "created_at": "2025-09-01T10:30:00Z"
  }
}
```

### PUT /products/:id

Update product (Admin only).

**Headers:** `Authorization: Bearer <admin_token>`

**Request Body:**

```json
{
  "name": "Samsung Galaxy S24 Ultra",
  "price": 899.99,
  "stock_quantity": 25,
  "is_active": true
}
```

**Response (200):**

```json
{
  "success": true,
  "message": "Product updated successfully",
  "data": {
    "id": 2,
    "name": "Samsung Galaxy S24 Ultra",
    "price": 899.99,
    "stock_quantity": 25,
    "is_active": true,
    "updated_at": "2025-09-01T11:00:00Z"
  }
}
```

### DELETE /products/:id

Delete product (Admin only).

**Headers:** `Authorization: Bearer <admin_token>`

**Response (200):**

```json
{
  "success": true,
  "message": "Product deleted successfully"
}
```

---

## 5. Shopping Cart Endpoints

### GET /cart

Get current user's cart.

**Headers:** `Authorization: Bearer <token>`

**Response (200):**

```json
{
  "success": true,
  "data": {
    "id": 1,
    "items": [
      {
        "id": 1,
        "product": {
          "id": 1,
          "name": "iPhone 15 Pro",
          "price": 999.99,
          "image_url": "https://example.com/iphone-1.jpg",
          "stock_quantity": 50
        },
        "quantity": 2,
        "unit_price": 999.99,
        "total_price": 1999.98,
        "added_at": "2025-09-01T10:00:00Z"
      }
    ],
    "total_items": 2,
    "total_amount": 1999.98,
    "updated_at": "2025-09-01T10:00:00Z"
  }
}
```

### POST /cart/items

Add item to cart.

**Headers:** `Authorization: Bearer <token>`

**Request Body:**

```json
{
  "product_id": 1,
  "quantity": 2
}
```

**Response (201):**

```json
{
  "success": true,
  "message": "Item added to cart successfully",
  "data": {
    "id": 1,
    "product": {
      "id": 1,
      "name": "iPhone 15 Pro",
      "price": 999.99
    },
    "quantity": 2,
    "unit_price": 999.99,
    "total_price": 1999.98,
    "added_at": "2025-09-01T10:00:00Z"
  }
}
```

### PUT /cart/items/:id

Update cart item quantity.

**Headers:** `Authorization: Bearer <token>`

**Request Body:**

```json
{
  "quantity": 3
}
```

**Response (200):**

```json
{
  "success": true,
  "message": "Cart item updated successfully",
  "data": {
    "id": 1,
    "quantity": 3,
    "total_price": 2999.97,
    "updated_at": "2025-09-01T10:30:00Z"
  }
}
```

### DELETE /cart/items/:id

Remove item from cart.

**Headers:** `Authorization: Bearer <token>`

**Response (200):**

```json
{
  "success": true,
  "message": "Item removed from cart successfully"
}
```

### DELETE /cart

Clear entire cart.

**Headers:** `Authorization: Bearer <token>`

**Response (200):**

```json
{
  "success": true,
  "message": "Cart cleared successfully"
}
```

---

## 6. Order Management Endpoints

### POST /orders

Create new order from cart.

**Headers:** `Authorization: Bearer <token>`

**Request Body:**

```json
{
  "shipping_address_id": 1,
  "payment_method": "credit_card",
  "notes": "Please handle with care"
}
```

**Response (201):**

```json
{
  "success": true,
  "message": "Order created successfully",
  "data": {
    "id": "ORD-2025090100001",
    "status": "pending",
    "items": [
      {
        "id": 1,
        "product": {
          "id": 1,
          "name": "iPhone 15 Pro"
        },
        "quantity": 2,
        "unit_price": 999.99,
        "total_price": 1999.98
      }
    ],
    "shipping_address": {
      "recipient_name": "John Doe",
      "address_line_1": "123 Main Street",
      "city": "New York",
      "state": "NY",
      "postal_code": "10001"
    },
    "subtotal": 1999.98,
    "shipping_cost": 15.0,
    "tax_amount": 160.0,
    "total_amount": 2174.98,
    "payment_method": "credit_card",
    "notes": "Please handle with care",
    "created_at": "2025-09-01T10:00:00Z"
  }
}
```

### GET /orders

Get user's orders.

**Headers:** `Authorization: Bearer <token>`

**Query Parameters:**

- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)
- `status` (optional): Filter by status

**Response (200):**

```json
{
  "success": true,
  "data": {
    "orders": [
      {
        "id": "ORD-2025090100001",
        "status": "pending",
        "total_amount": 2174.98,
        "payment_method": "credit_card",
        "payment_status": "pending",
        "item_count": 2,
        "created_at": "2025-09-01T10:00:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 2,
      "total_items": 15,
      "per_page": 10
    }
  }
}
```

### GET /orders/:id

Get order details.

**Headers:** `Authorization: Bearer <token>`

**Response (200):**

```json
{
  "success": true,
  "data": {
    "id": "ORD-2025090100001",
    "status": "confirmed",
    "items": [
      {
        "id": 1,
        "product": {
          "id": 1,
          "name": "iPhone 15 Pro",
          "image_url": "https://example.com/iphone-1.jpg"
        },
        "quantity": 2,
        "unit_price": 999.99,
        "total_price": 1999.98
      }
    ],
    "shipping_address": {
      "recipient_name": "John Doe",
      "phone": "+1234567890",
      "address_line_1": "123 Main Street",
      "address_line_2": "Apt 4B",
      "city": "New York",
      "state": "NY",
      "postal_code": "10001",
      "country": "USA"
    },
    "subtotal": 1999.98,
    "shipping_cost": 15.0,
    "tax_amount": 160.0,
    "total_amount": 2174.98,
    "payment_method": "credit_card",
    "payment_status": "paid",
    "tracking_number": "TRK123456789",
    "notes": "Please handle with care",
    "status_history": [
      {
        "status": "pending",
        "timestamp": "2025-09-01T10:00:00Z",
        "note": "Order created"
      },
      {
        "status": "confirmed",
        "timestamp": "2025-09-01T10:30:00Z",
        "note": "Payment confirmed"
      }
    ],
    "created_at": "2025-09-01T10:00:00Z",
    "updated_at": "2025-09-01T10:30:00Z"
  }
}
```

### PUT /orders/:id/cancel

Cancel order (if status allows).

**Headers:** `Authorization: Bearer <token>`

**Response (200):**

```json
{
  "success": true,
  "message": "Order cancelled successfully",
  "data": {
    "id": "ORD-2025090100001",
    "status": "cancelled",
    "cancelled_at": "2025-09-01T11:00:00Z"
  }
}
```

### GET /admin/orders

Get all orders (Admin only).

**Headers:** `Authorization: Bearer <admin_token>`

**Query Parameters:**

- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)
- `status` (optional): Filter by status
- `user_id` (optional): Filter by user

**Response (200):**

```json
{
  "success": true,
  "data": {
    "orders": [
      {
        "id": "ORD-2025090100001",
        "user": {
          "id": 1,
          "name": "John Doe",
          "email": "user@example.com"
        },
        "status": "pending",
        "total_amount": 2174.98,
        "payment_status": "pending",
        "created_at": "2025-09-01T10:00:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 5,
      "total_items": 47,
      "per_page": 10
    }
  }
}
```

### PUT /admin/orders/:id/status

Update order status (Admin only).

**Headers:** `Authorization: Bearer <admin_token>`

**Request Body:**

```json
{
  "status": "shipped",
  "tracking_number": "TRK123456789",
  "note": "Order shipped via FedEx"
}
```

**Response (200):**

```json
{
  "success": true,
  "message": "Order status updated successfully",
  "data": {
    "id": "ORD-2025090100001",
    "status": "shipped",
    "tracking_number": "TRK123456789",
    "updated_at": "2025-09-01T11:00:00Z"
  }
}
```

---

## 7. Payment Endpoints

### POST /payments/process

Process payment for an order.

**Headers:** `Authorization: Bearer <token>`

**Request Body (Credit Card):**

```json
{
  "order_id": "ORD-2025090100001",
  "payment_method": "credit_card",
  "payment_details": {
    "card_number": "4111111111111111",
    "expiry_month": "12",
    "expiry_year": "2025",
    "cvv": "123",
    "cardholder_name": "John Doe"
  }
}
```

**Request Body (Cash):**

```json
{
  "order_id": "ORD-2025090100001",
  "payment_method": "cash",
  "payment_details": {
    "note": "Cash on delivery"
  }
}
```

**Response (200):**

```json
{
  "success": true,
  "message": "Payment processed successfully",
  "data": {
    "payment_id": "PAY-2025090100001",
    "order_id": "ORD-2025090100001",
    "amount": 2174.98,
    "payment_method": "credit_card",
    "status": "completed",
    "transaction_id": "TXN123456789",
    "processed_at": "2025-09-01T10:30:00Z"
  }
}
```

### GET /payments/:payment_id

Get payment details.

**Headers:** `Authorization: Bearer <token>`

**Response (200):**

```json
{
  "success": true,
  "data": {
    "id": "PAY-2025090100001",
    "order_id": "ORD-2025090100001",
    "amount": 2174.98,
    "payment_method": "credit_card",
    "status": "completed",
    "transaction_id": "TXN123456789",
    "payment_details": {
      "last_four": "1111",
      "card_type": "Visa"
    },
    "processed_at": "2025-09-01T10:30:00Z",
    "created_at": "2025-09-01T10:30:00Z"
  }
}
```

### GET /payments/order/:order_id

Get payments for a specific order.

**Headers:** `Authorization: Bearer <token>`

**Response (200):**

```json
{
  "success": true,
  "data": [
    {
      "id": "PAY-2025090100001",
      "amount": 2174.98,
      "payment_method": "credit_card",
      "status": "completed",
      "processed_at": "2025-09-01T10:30:00Z"
    }
  ]
}
```

---

## Error Responses

### Common Error Format

```json
{
  "success": false,
  "message": "Error description",
  "errors": [
    {
      "field": "email",
      "message": "Email is required"
    }
  ]
}
```

### HTTP Status Codes

- `200` - Success
- `201` - Created
- `400` - Bad Request
- `401` - Unauthorized
- `403` - Forbidden
- `404` - Not Found
- `422` - Validation Error
- `500` - Internal Server Error

---

## Order Status Flow

1. **pending** - Order created, awaiting payment
2. **confirmed** - Payment received, order confirmed
3. **processing** - Order being prepared
4. **shipped** - Order shipped to customer
5. **delivered** - Order delivered successfully
6. **cancelled** - Order cancelled
7. **returned** - Order returned by customer

## Payment Status

- **pending** - Payment initiated
- **completed** - Payment successful
- **failed** - Payment failed
- **refunded** - Payment refunded
- **cancelled** - Payment cancelled

---

## Rate Limiting

- Authentication endpoints: 5 requests per minute
- General API endpoints: 100 requests per minute per user
- Admin endpoints: 200 requests per minute

## Data Validation Rules

- Email must be valid format
- Password minimum 8 characters
- Phone numbers should include country code
- Prices must be positive numbers
- Stock quantities must be non-negative integers
- Order amounts calculated server-side for security
