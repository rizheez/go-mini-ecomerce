# Deployment Guide

## Overview

This guide explains how to deploy the E-Commerce API application to different environments.

## Prerequisites

- Docker and Docker Compose installed
- PostgreSQL database (if not using Docker)
- Redis server (if not using Docker)
- Go 1.21+ (for building from source)

## Deployment Options

### 1. Docker Deployment (Recommended)

#### Using Docker Compose (Development/Testing)

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd e-commerce-API
   ```

2. Update the `.env` file with your configuration:
   ```bash
   cp .env.example .env
   # Edit .env with your values
   ```

3. Start the services:
   ```bash
   docker-compose up -d
   ```

4. Access the services:
   - API: http://localhost:8080
   - Database Admin (Adminer): http://localhost:8081
   - Redis Admin (Redis Commander): http://localhost:8082

#### Using Docker Only (Production)

1. Build the Docker image:
   ```bash
   docker build -t ecommerce-api:latest .
   ```

2. Run the container:
   ```bash
   docker run -d \
     --name ecommerce-api \
     -p 8080:8080 \
     --env-file .env \
     ecommerce-api:latest
   ```

### 2. Manual Deployment

#### Building from Source

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd e-commerce-API
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Build the application:
   ```bash
   go build -o ecommerce-api cmd/api/main.go
   ```

4. Update the `.env` file with your configuration:
   ```bash
   cp .env.example .env
   # Edit .env with your values
   ```

5. Run the application:
   ```bash
   ./ecommerce-api
   ```

#### Using Makefile

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd e-commerce-API
   ```

2. Update the `.env` file with your configuration:
   ```bash
   cp .env.example .env
   # Edit .env with your values
   ```

3. Build and run the application:
   ```bash
   make build
   make run
   ```

## Environment Configuration

### Required Environment Variables

| Variable | Description | Example |
|----------|-------------|---------|
| DB_HOST | Database host | localhost |
| DB_PORT | Database port | 5432 |
| DB_USER | Database user | ecommerce_user |
| DB_PASSWORD | Database password | ecommerce_password |
| DB_NAME | Database name | ecommerce_db |
| JWT_SECRET | JWT secret key | my_secret_key |
| SERVER_PORT | Server port | 8080 |

### Optional Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| DB_SSL_MODE | Database SSL mode | disable |
| LOG_LEVEL | Logging level | info |
| DEBUG | Debug mode | false |

## Database Migration

Run database migrations after deployment:

```bash
# Using Docker
docker exec -it ecommerce-api ./migrate up

# Using binary
./migrate up
```

## Health Checks

The application provides health check endpoints:

- `/health` - Basic health check
- `/health/db` - Database connectivity check
- `/health/redis` - Redis connectivity check

Example:
```bash
curl http://localhost:8080/health
```

## Monitoring and Logging

### Logs

Application logs are written to:
- Console (stdout/stderr)
- File (if LOG_FILE is configured)

### Metrics

The application exposes Prometheus metrics at:
- `/metrics`

## Scaling

### Horizontal Scaling

To scale the application horizontally:

1. Deploy multiple instances behind a load balancer
2. Use shared database and Redis instances
3. Ensure consistent environment configuration across instances

### Vertical Scaling

To scale vertically:
1. Increase container resources (CPU, memory)
2. Optimize database queries
3. Add database indexes

## Security Considerations

1. Use HTTPS in production
2. Store secrets securely (HashiCorp Vault, AWS Secrets Manager)
3. Regularly update dependencies
4. Implement proper authentication and authorization
5. Use rate limiting to prevent abuse
6. Sanitize all user inputs

## Backup and Recovery

### Database Backup

```bash
# Backup
pg_dump -h $DB_HOST -p $DB_PORT -U $DB_USER $DB_NAME > backup.sql

# Restore
psql -h $DB_HOST -p $DB_PORT -U $DB_USER $DB_NAME < backup.sql
```

### Configuration Backup

Always backup your `.env` file and store it securely.

## Troubleshooting

### Common Issues

1. **Database connection failed**
   - Check database credentials in `.env`
   - Verify database is running and accessible
   - Check firewall settings

2. **Port already in use**
   - Change SERVER_PORT in `.env`
   - Kill process using the port

3. **Permission denied**
   - Check file permissions
   - Run with appropriate user privileges

### Logs

Check logs for detailed error information:

```bash
# Docker logs
docker logs ecommerce-api

# File logs (if configured)
tail -f logs/app.log
```

## Updating the Application

1. Pull the latest code:
   ```bash
   git pull origin main
   ```

2. Update dependencies:
   ```bash
   go mod tidy
   ```

3. Rebuild the application:
   ```bash
   make build
   ```

4. Restart the service:
   ```bash
   # For Docker
   docker-compose down
   docker-compose up -d
   
   # For manual deployment
   ./ecommerce-api
   ```

## Rollback

To rollback to a previous version:

1. Stop current service
2. Deploy previous version
3. Restore database backup if needed
4. Start service