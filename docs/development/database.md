# Database Structure

## Overview

This document describes the database structure for the E-Commerce API application.

## Entity Relationship Diagram

```
Users ||--o{ Orders : creates
Users ||--o{ UserAddresses : has
Users ||--|| Carts : owns
Categories ||--o{ Products : contains
Products ||--o{ CartItems : in
Products ||--o{ OrderItems : in
Products ||--o{ ProductImages : has
Orders ||--o{ OrderItems : contains
Orders ||--|| Payments : has
Carts ||--o{ CartItems : contains
Orders ||--|| UserAddresses : ships_to
```

## Tables

### 1. users

Primary table for user management with comprehensive customer information.

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    role VARCHAR(20) DEFAULT 'customer' CHECK (role IN ('customer', 'admin')),
    email_verified BOOLEAN DEFAULT FALSE,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
```

**Indexes:**
- `idx_users_email` on `email`
- `idx_users_role` on `role`
- `idx_users_active` on `is_active`

### 2. user_addresses

Customer shipping addresses with comprehensive location details.

```sql
CREATE TABLE user_addresses (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    label VARCHAR(50) NOT NULL, -- 'Home', 'Work', 'Other'
    recipient_name VARCHAR(255) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    address_line_1 VARCHAR(255) NOT NULL,
    address_line_2 VARCHAR(255),
    city VARCHAR(100) NOT NULL,
    state VARCHAR(100) NOT NULL,
    postal_code VARCHAR(20) NOT NULL,
    country VARCHAR(100) NOT NULL DEFAULT 'USA',
    is_default BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
```

**Indexes:**
- `idx_user_addresses_user_id` on `user_id`
- `idx_user_addresses_default` on `user_id, is_default`

**Constraints:**
- Unique constraint on `user_id` where `is_default = TRUE`

### 3. categories

Product categorization system.

```sql
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    image_url VARCHAR(500),
    is_active BOOLEAN DEFAULT TRUE,
    sort_order INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
```

**Indexes:**
- `idx_categories_active` on `is_active`
- `idx_categories_sort_order` on `sort_order`

### 4. products

Main product catalog with pricing and inventory.

```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL CHECK (price >= 0),
    stock_quantity INTEGER NOT NULL DEFAULT 0 CHECK (stock_quantity >= 0),
    category_id INTEGER NOT NULL REFERENCES categories(id) ON DELETE RESTRICT,
    sku VARCHAR(100) UNIQUE,
    specifications JSONB, -- Store flexible product specifications
    is_active BOOLEAN DEFAULT TRUE,
    weight DECIMAL(8,2), -- in kg
    dimensions JSONB, -- {"length": 10, "width": 5, "height": 2} in cm
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
```

**Indexes:**
- `idx_products_category_id` on `category_id`
- `idx_products_active` on `is_active`
- `idx_products_price` on `price`
- `idx_products_stock` on `stock_quantity`

### 5. product_images

Product image management with multiple images per product.

```sql
CREATE TABLE product_images (
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    url VARCHAR(500) NOT NULL,
    alt_text VARCHAR(255),
    is_primary BOOLEAN DEFAULT FALSE,
    sort_order INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
```

**Indexes:**
- `idx_product_images_product_id` on `product_id`
- `idx_product_images_primary` on `product_id, is_primary`

**Constraints:**
- Unique constraint on `product_id` where `is_primary = TRUE`

### 6. carts

Shopping cart management (one per user).

```sql
CREATE TABLE carts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
```

**Indexes:**
- `idx_carts_user_id` on `user_id` (unique)

### 7. cart_items

Individual items in shopping carts.

```sql
CREATE TABLE cart_items (
    id SERIAL PRIMARY KEY,
    cart_id INTEGER NOT NULL REFERENCES carts(id) ON DELETE CASCADE,
    product_id INTEGER NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    unit_price DECIMAL(10,2) NOT NULL, -- Price at time of adding to cart
    added_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
```

**Indexes:**
- `idx_cart_items_cart_id` on `cart_id`
- `idx_cart_items_product_id` on `product_id`

**Constraints:**
- Unique constraint on `cart_id, product_id` (one product per cart)

### 8. orders

Order management with comprehensive tracking.

```sql
CREATE TABLE orders (
    id VARCHAR(50) PRIMARY KEY, -- Format: ORD-YYYYMMDDNNNNN
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    status VARCHAR(20) NOT NULL DEFAULT 'pending'
        CHECK (status IN ('pending', 'confirmed', 'processing', 'shipped', 'delivered', 'cancelled', 'returned')),
    subtotal DECIMAL(10,2) NOT NULL CHECK (subtotal >= 0),
    shipping_cost DECIMAL(10,2) NOT NULL DEFAULT 0 CHECK (shipping_cost >= 0),
    tax_amount DECIMAL(10,2) NOT NULL DEFAULT 0 CHECK (tax_amount >= 0),
    total_amount DECIMAL(10,2) NOT NULL CHECK (total_amount >= 0),
    payment_method VARCHAR(20) NOT NULL CHECK (payment_method IN ('credit_card', 'cash')),
    payment_status VARCHAR(20) DEFAULT 'pending'
        CHECK (payment_status IN ('pending', 'completed', 'failed', 'refunded', 'cancelled')),
    shipping_address JSONB NOT NULL, -- Store complete address snapshot
    tracking_number VARCHAR(100),
    notes TEXT,
    cancelled_at TIMESTAMP WITH TIME ZONE,
    shipped_at TIMESTAMP WITH TIME ZONE,
    delivered_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
```

**Indexes:**
- `idx_orders_user_id` on `user_id`
- `idx_orders_status` on `status`
- `idx_orders_payment_status` on `payment_status`
- `idx_orders_created_at` on `created_at`
- `idx_orders_total_amount` on `total_amount`

### 9. order_items

Individual items within orders.

```sql
CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id VARCHAR(50) NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    product_id INTEGER NOT NULL REFERENCES products(id) ON DELETE RESTRICT,
    product_name VARCHAR(255) NOT NULL, -- Snapshot at time of order
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    unit_price DECIMAL(10,2) NOT NULL CHECK (unit_price >= 0),
    total_price DECIMAL(10,2) NOT NULL CHECK (total_price >= 0),
    product_snapshot JSONB, -- Store product details at time of order
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
```

**Indexes:**
- `idx_order_items_order_id` on `order_id`
- `idx_order_items_product_id` on `product_id`

### 10. order_status_history

Track order status changes for audit trail.

```sql
CREATE TABLE order_status_history (
    id SERIAL PRIMARY KEY,
    order_id VARCHAR(50) NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    from_status VARCHAR(20),
    to_status VARCHAR(20) NOT NULL,
    note TEXT,
    changed_by INTEGER REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
```

**Indexes:**
- `idx_order_status_history_order_id` on `order_id`
- `idx_order_status_history_created_at` on `created_at`

### 11. payments

Payment transaction records.

```sql
CREATE TABLE payments (
    id VARCHAR(50) PRIMARY KEY, -- Format: PAY-YYYYMMDDNNNNN
    order_id VARCHAR(50) NOT NULL REFERENCES orders(id) ON DELETE RESTRICT,
    amount DECIMAL(10,2) NOT NULL CHECK (amount > 0),
    payment_method VARCHAR(20) NOT NULL CHECK (payment_method IN ('credit_card', 'cash')),
    status VARCHAR(20) NOT NULL DEFAULT 'pending'
        CHECK (status IN ('pending', 'completed', 'failed', 'refunded', 'cancelled')),
    transaction_id VARCHAR(100), -- External payment processor transaction ID
    payment_details JSONB, -- Store payment method specific details (masked)
    gateway_response JSONB, -- Store payment gateway response
    processed_at TIMESTAMP WITH TIME ZONE,
    failed_at TIMESTAMP WITH TIME ZONE,
    failure_reason TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
```

**Indexes:**
- `idx_payments_order_id` on `order_id`
- `idx_payments_status` on `status`
- `idx_payments_processed_at` on `processed_at`
- `idx_payments_transaction_id` on `transaction_id`

### 12. product_reviews (Optional - for future enhancement)

```sql
CREATE TABLE product_reviews (
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    order_item_id INTEGER REFERENCES order_items(id) ON DELETE SET NULL,
    rating INTEGER NOT NULL CHECK (rating >= 1 AND rating <= 5),
    title VARCHAR(255),
    review_text TEXT,
    is_verified_purchase BOOLEAN DEFAULT FALSE,
    is_published BOOLEAN DEFAULT TRUE,
    helpful_count INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
```

**Indexes:**
- `idx_product_reviews_product_id` on `product_id`
- `idx_product_reviews_user_id` on `user_id`
- `idx_product_reviews_rating` on `rating`

**Constraints:**
- Unique constraint on `product_id, user_id` (one review per user per product)

## Views

### Product catalog view with aggregated data

```sql
CREATE VIEW product_catalog AS
SELECT
    p.id,
    p.name,
    p.description,
    p.price,
    p.stock_quantity,
    p.sku,
    p.specifications,
    p.is_active,
    c.name AS category_name,
    c.id AS category_id,
    pi.url AS primary_image_url,
    COALESCE(AVG(pr.rating), 0) AS average_rating,
    COUNT(pr.id) AS review_count
FROM products p
JOIN categories c ON p.category_id = c.id
LEFT JOIN product_images pi ON p.id = pi.product_id AND pi.is_primary = TRUE
LEFT JOIN product_reviews pr ON p.id = pr.product_id AND pr.is_published = TRUE
GROUP BY p.id, c.id, pi.url;
```

### Order summary view

```sql
CREATE VIEW order_summary AS
SELECT
    o.id,
    o.user_id,
    u.name AS customer_name,
    u.email AS customer_email,
    o.status,
    o.payment_status,
    o.total_amount,
    o.payment_method,
    COUNT(oi.id) AS item_count,
    o.created_at,
    o.updated_at
FROM orders o
JOIN users u ON o.user_id = u.id
LEFT JOIN order_items oi ON o.id = oi.order_id
GROUP BY o.id, u.id;
```

## Triggers

### Update timestamps

```sql
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Apply to all tables with updated_at column
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_user_addresses_updated_at BEFORE UPDATE ON user_addresses FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_categories_updated_at BEFORE UPDATE ON categories FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_products_updated_at BEFORE UPDATE ON products FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_carts_updated_at BEFORE UPDATE ON carts FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_cart_items_updated_at BEFORE UPDATE ON cart_items FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_orders_updated_at BEFORE UPDATE ON orders FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_payments_updated_at BEFORE UPDATE ON payments FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
```

### Order status history trigger

```sql
CREATE OR REPLACE FUNCTION record_order_status_change()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.status IS DISTINCT FROM NEW.status THEN
        INSERT INTO order_status_history (order_id, from_status, to_status, note)
        VALUES (NEW.id, OLD.status, NEW.status, 'Status changed');
    END IF;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER order_status_change_trigger
AFTER UPDATE ON orders
FOR EACH ROW
EXECUTE FUNCTION record_order_status_change();
```

## Sample Data

### Initial Categories

```sql
INSERT INTO categories (name, description, image_url, sort_order) VALUES
('Electronics', 'Electronic devices and accessories', 'https://example.com/electronics.jpg', 1),
('Fashion', 'Clothing, shoes, and accessories', 'https://example.com/fashion.jpg', 2),
('Home & Garden', 'Home improvement and garden supplies', 'https://example.com/home.jpg', 3),
('Books', 'Books and educational materials', 'https://example.com/books.jpg', 4),
('Sports', 'Sports and outdoor equipment', 'https://example.com/sports.jpg', 5);
```

### Admin User

```sql
INSERT INTO users (email, password_hash, name, phone, role) VALUES
('admin@ecommerce.com', '$2a$12$hash_here', 'System Administrator', '+1234567890', 'admin');
```

## Database Configuration

### Connection String Format

```
postgresql://username:password@localhost:5432/ecommerce_db?sslmode=disable
```

### Recommended Settings

- `max_connections = 100`
- `shared_buffers = 256MB`
- `effective_cache_size = 1GB`
- `work_mem = 4MB`
- `maintenance_work_mem = 64MB`

## Security Considerations

1. **Password Storage**: Always hash passwords using bcrypt with salt
2. **Data Encryption**: Encrypt sensitive payment details
3. **Access Control**: Implement row-level security for multi-tenant scenarios
4. **Audit Trail**: Maintain comprehensive logs for financial transactions
5. **Backup Strategy**: Regular automated backups with point-in-time recovery
6. **Connection Security**: Use SSL/TLS for database connections in production

## Performance Optimization

1. **Indexing**: Comprehensive indexes on frequently queried columns
2. **Partitioning**: Consider partitioning large tables (orders, payments) by date
3. **Archiving**: Implement data archiving strategy for old orders
4. **Query Optimization**: Use EXPLAIN ANALYZE to optimize slow queries
5. **Connection Pooling**: Implement connection pooling for better performance

## Migration Strategy

1. **Version Control**: Use migration files for schema changes
2. **Rollback Plan**: Always have rollback scripts for production deployments
3. **Testing**: Test all migrations on staging environment first
4. **Monitoring**: Monitor performance impact of schema changes