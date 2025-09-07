# Database Migrations

## Overview

This document explains how to work with database migrations in the E-Commerce API project.

## Migration Files

Migration files are located in the `migrations/` directory and follow this naming convention:
```
{version}_{description}.sql
```

Example:
```
001_create_users_table.sql
002_create_categories_table.sql
```

## Running Migrations

### Using Go Script

```bash
# Run all pending migrations
go run scripts/migrate.go up

# Rollback last migration
go run scripts/migrate.go down

# Rollback all migrations
go run scripts/migrate.go down all

# Check migration status
go run scripts/migrate.go status
```

### Using Makefile

```bash
# Run all pending migrations
make migrate-up

# Rollback last migration
make migrate-down
```

### Using Docker

```bash
# Run migrations in Docker container
docker exec -it ecommerce-api ./migrate up
```

## Creating New Migrations

1. Create a new file in the `migrations/` directory
2. Follow the naming convention: `{next_version}_{description}.sql`
3. Add UP and DOWN sections

Example migration file:
```sql
-- +migrate Up
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- +migrate Down
DROP TABLE products;
```

## Migration Best Practices

1. **Always include Down migrations** for rollback capability
2. **Test migrations** on a copy of production data
3. **Backup database** before running migrations in production
4. **Use transactions** for complex migrations
5. **Add indexes** for frequently queried columns
6. **Use appropriate data types** for storage efficiency
7. **Add constraints** to maintain data integrity

## Common Migration Tasks

### Adding a New Column

```sql
-- +migrate Up
ALTER TABLE users ADD COLUMN phone VARCHAR(20);

-- +migrate Down
ALTER TABLE users DROP COLUMN phone;
```

### Adding an Index

```sql
-- +migrate Up
CREATE INDEX idx_products_category_id ON products(category_id);

-- +migrate Down
DROP INDEX idx_products_category_id;
```

### Adding a Constraint

```sql
-- +migrate Up
ALTER TABLE products ADD CONSTRAINT chk_price CHECK (price >= 0);

-- +migrate Down
ALTER TABLE products DROP CONSTRAINT chk_price;
```

## Migration Status

Check the status of migrations:
```bash
go run scripts/migrate.go status
```

This will show:
- Which migrations have been applied
- Which migrations are pending
- Any migration errors

## Rollback Strategy

Always have a rollback plan:
1. Test down migrations in development
2. Backup production database before migration
3. Have a plan to restore from backup if needed
4. Monitor application after migration

## Troubleshooting

### Common Issues

1. **Migration already applied**
   - Check migration status
   - Manually mark as applied if needed

2. **Migration failed**
   - Check error message
   - Fix issue and retry
   - Rollback if necessary

3. **Missing migration files**
   - Verify file names and locations
   - Check file permissions

### Manual Migration Management

If you need to manually manage migrations:
1. Check the `schema_migrations` table
2. Add/remove entries as needed
3. Ensure consistency with actual schema

## Production Considerations

1. **Run migrations during maintenance windows**
2. **Test migrations on staging environment first**
3. **Have rollback plan ready**
4. **Monitor application after migration**
5. **Keep migration scripts in version control**