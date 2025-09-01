# UUID Migration Guide for Mini E-Commerce API

## Overview

This document outlines the changes made to update the e-commerce API from using integer IDs to UUIDs for better security, scalability, and distributed system compatibility.

## Database Schema Changes

### 1. Enable UUID Extension

```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pg_trgm";
CREATE EXTENSION IF NOT EXISTS "btree_gin";
```

### 2. Updated Tables with UUIDs

All primary keys have been changed from `SERIAL` to `UUID DEFAULT uuid_generate_v4()`:

#### Core Entity Tables:

- **users**: `id UUID PRIMARY KEY DEFAULT uuid_generate_v4()`
- **categories**: `id UUID PRIMARY KEY DEFAULT uuid_generate_v4()`
- **products**: `id UUID PRIMARY KEY DEFAULT uuid_generate_v4()`
- **user_addresses**: `id UUID PRIMARY KEY DEFAULT uuid_generate_v4()`
- **product_images**: `id UUID PRIMARY KEY DEFAULT uuid_generate_v4()`
- **carts**: `id UUID PRIMARY KEY DEFAULT uuid_generate_v4()`
- **cart_items**: `id UUID PRIMARY KEY DEFAULT uuid_generate_v4()`
- **order_items**: `id UUID PRIMARY KEY DEFAULT uuid_generate_v4()`
- **order_status_history**: `id UUID PRIMARY KEY DEFAULT uuid_generate_v4()`
- **product_reviews**: `id UUID PRIMARY KEY DEFAULT uuid_generate_v4()`

#### Special Cases:

- **orders**: Uses both UUID primary key and human-readable order number
  - `id UUID PRIMARY KEY DEFAULT uuid_generate_v4()`
  - `order_number VARCHAR(50) UNIQUE NOT NULL` (e.g., "ORD-20250901001")
- **payments**: Uses both UUID primary key and human-readable payment number
  - `id UUID PRIMARY KEY DEFAULT uuid_generate_v4()`
  - `payment_number VARCHAR(50) UNIQUE NOT NULL` (e.g., "PAY-20250901001")

### 3. Enhanced Indexing Strategy

#### Performance Indexes Added:

```sql
-- User performance
CREATE INDEX idx_users_phone ON users(phone) WHERE phone IS NOT NULL;
CREATE INDEX idx_users_created_at ON users(created_at);

-- Address lookups
CREATE INDEX idx_user_addresses_city ON user_addresses(city);
CREATE INDEX idx_user_addresses_state ON user_addresses(state);
CREATE INDEX idx_user_addresses_postal_code ON user_addresses(postal_code);
CREATE INDEX idx_user_addresses_country ON user_addresses(country);
CREATE INDEX idx_user_addresses_location ON user_addresses(city, state, country);

-- Category performance
CREATE INDEX idx_categories_created_at ON categories(created_at);
CREATE INDEX idx_categories_active_sort ON categories(is_active, sort_order) WHERE is_active = TRUE;

-- Product search and filtering
CREATE INDEX idx_products_sku ON products(sku) WHERE sku IS NOT NULL;
CREATE INDEX idx_products_category_active ON products(category_id, is_active) WHERE is_active = TRUE;
CREATE INDEX idx_products_price_range ON products(price, is_active) WHERE is_active = TRUE;
CREATE INDEX idx_products_created_at ON products(created_at);
CREATE INDEX idx_products_stock_active ON products(stock_quantity, is_active) WHERE is_active = TRUE;
CREATE INDEX idx_products_name_trgm ON products USING gin(name gin_trgm_ops);
CREATE INDEX idx_products_full_text ON products USING gin((to_tsvector('english', name || ' ' || COALESCE(description, ''))));

-- Product images
CREATE INDEX idx_product_images_sort ON product_images(product_id, sort_order);
CREATE INDEX idx_product_images_created_at ON product_images(created_at);

-- Cart performance
CREATE INDEX idx_carts_updated_at ON carts(updated_at);
CREATE INDEX idx_cart_items_added_at ON cart_items(added_at);
CREATE INDEX idx_cart_items_updated_at ON cart_items(updated_at);

-- Order performance
CREATE INDEX idx_orders_user_status ON orders(user_id, status);
CREATE INDEX idx_orders_user_created ON orders(user_id, created_at);
CREATE INDEX idx_orders_tracking_number ON orders(tracking_number) WHERE tracking_number IS NOT NULL;
CREATE INDEX idx_orders_shipped_at ON orders(shipped_at) WHERE shipped_at IS NOT NULL;
CREATE INDEX idx_orders_delivered_at ON orders(delivered_at) WHERE delivered_at IS NOT NULL;

-- Order items
CREATE INDEX idx_order_items_created_at ON order_items(created_at);
CREATE INDEX idx_order_items_total_price ON order_items(total_price);

-- Payment tracking
CREATE INDEX idx_payments_created_at ON payments(created_at);
CREATE INDEX idx_payments_amount ON payments(amount);
CREATE INDEX idx_payments_payment_method ON payments(payment_method);

-- Composite indexes for common queries
CREATE INDEX idx_orders_user_status_created ON orders(user_id, status, created_at);
CREATE INDEX idx_products_category_price_active ON products(category_id, price, is_active) WHERE is_active = TRUE;
CREATE INDEX idx_orders_date_total ON orders(DATE(created_at), total_amount);
CREATE INDEX idx_order_items_product_date ON order_items(product_id, created_at);
CREATE INDEX idx_payments_order_status ON payments(order_id, status);
```

## API Response Format Changes

### Before (Integer IDs):

```json
{
  "id": 1,
  "name": "John Doe",
  "email": "user@example.com"
}
```

### After (UUID IDs):

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440001",
  "name": "John Doe",
  "email": "user@example.com"
}
```

## Key Benefits of UUID Implementation

### 1. Security

- Prevents ID enumeration attacks
- No predictable sequential IDs
- Better privacy protection

### 2. Scalability

- No ID conflicts in distributed systems
- Better for database sharding
- Supports multi-region deployments

### 3. Integration

- Better for microservices architecture
- External system integration friendly
- API versioning compatibility

### 4. Performance Considerations

- Added comprehensive indexing strategy
- Optimized for common query patterns
- Full-text search capabilities
- Location-based address lookups

## Migration Strategy

### 1. Development Environment

```sql
-- Create new tables with UUID
-- Migrate existing data (if any)
-- Update application code
-- Test thoroughly
```

### 2. Production Deployment

```sql
-- Backup existing database
-- Run migration scripts during maintenance window
-- Update application
-- Verify data integrity
```

### 3. Rollback Plan

```sql
-- Keep backup of pre-migration state
-- Documented rollback procedures
-- Health checks and monitoring
```

## Implementation Notes

### 1. Order and Payment Numbers

- Keep human-readable identifiers for customer-facing operations
- Use UUIDs internally for system operations
- Maintain uniqueness constraints on both

### 2. Foreign Key Relationships

- All foreign keys updated to reference UUID columns
- Cascade deletion maintained where appropriate
- Referential integrity preserved

### 3. Application Code Changes

- Update all model definitions
- Modify API serializers/deserializers
- Update database query patterns
- Add UUID validation

### 4. Testing Requirements

- Test all CRUD operations
- Verify foreign key constraints
- Performance testing with indexes
- Load testing with realistic data volumes

## Performance Optimization

### 1. Query Patterns

- Use composite indexes for multi-column searches
- Implement full-text search for products
- Optimize pagination with cursor-based pagination
- Cache frequently accessed data

### 2. Database Configuration

```sql
-- Recommended PostgreSQL settings
shared_buffers = 256MB
effective_cache_size = 1GB
work_mem = 4MB
maintenance_work_mem = 64MB
random_page_cost = 1.1  -- For SSD storage
```

### 3. Monitoring

- Track query performance
- Monitor index usage
- Watch for slow queries
- Set up alerts for performance degradation

This migration provides a solid foundation for a scalable, secure e-commerce platform with proper indexing for optimal performance.
