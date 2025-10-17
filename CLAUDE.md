# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go Fiber REST API following Clean Architecture principles with MySQL and Redis integration. The application provides a user management system with full CRUD operations.

## Essential Commands

### Development
```bash
# Start dependencies (MySQL + Redis)
docker-compose up mysql redis -d

# Run application locally
go run cmd/server/main.go

# Build application
go build -o bin/main cmd/server/main.go

# Download dependencies
go mod tidy
```

### Docker Development
```bash
# Full stack (app + dependencies)
docker-compose up --build

# Background mode
docker-compose up -d --build

# Stop all services
docker-compose down

# Stop with volume cleanup
docker-compose down -v
```

### Using Makefile
```bash
# Install development tools (mockery)
make install-tools

# Generate mocks for testing
make generate-mocks

# Run tests
make test

# Run tests with coverage report
make test-coverage

# Build the application
make build

# Build and run
make run

# Development mode with hot reload (requires air)
make dev

# Generate Swagger documentation
make swagger

# Clean up generated files
make clean
```

## Architecture Overview

### Clean Architecture Implementation
The codebase follows strict Clean Architecture with dependency inversion:

**Domain Layer** (`internal/domain/`):
- `entities/` - Core business entities (User) with GORM annotations
- `interfaces/` - Repository contracts (UserRepository interface)

**Data Layer** (`internal/repositories/`):
- Concrete implementations of repository interfaces using GORM
- Handles database operations with proper error handling

**Service Layer** (`internal/services/`):
- Business logic and validation rules
- Orchestrates repository calls and implements business constraints

**Handler Layer** (`internal/handlers/`):
- HTTP request/response handling with Fiber
- Input validation using go-playground/validator
- Request/response DTOs with validation tags

### Dependency Injection Flow
```go
// Repository layer
userRepo := repositories.NewUserRepository(database.GetDB())

// Service layer (business logic)
userService := services.NewUserService(userRepo)

// Handler layer (HTTP)
userHandler := handlers.NewUserHandler(userService)
```

### Database Architecture
- **MySQL**: Primary database using GORM with connection pooling
- **Redis**: Caching layer with go-redis client
- **Global Singletons**: `database.GetDB()` and `database.GetRedis()` for access
- **Graceful Shutdown**: Both connections properly closed on app termination

### Configuration System
Environment-based configuration in `internal/config/config.go`:
- Database, Redis, and Server configurations
- Environment variable override with sensible defaults
- Type-safe configuration loading

## API Structure

### Base URL Pattern
- Health check: `GET /health`
- API endpoints: `/api/v1/*`
- User endpoints: `/api/v1/users/*`

### Middleware Stack
- CORS (permissive for development)
- Request logging
- Panic recovery
- Custom error handler with consistent JSON responses

### Response Format
Standardized responses via `pkg/utils/response.go`:
```json
{
  "success": true/false,
  "message": "descriptive message",
  "data": {...} // optional
}
```

## Key Dependencies

- **GoFiber v2**: High-performance HTTP framework
- **GORM**: ORM with MySQL driver
- **go-redis/v9**: Redis client
- **go-playground/validator/v10**: Request validation
- **Viper**: Configuration management

## Environment Variables

Required for database and Redis connections:
```bash
DB_HOST=localhost
DB_PORT=3306
DB_USER=user_1
DB_PASSWORD=password_1
DB_NAME=db_1

REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

SERVER_PORT=3000
```

## Development Notes

### Adding New Entities
1. Create entity in `internal/domain/entities/`
2. Define repository interface in `internal/domain/interfaces/`
3. Implement repository in `internal/repositories/`
4. Create service in `internal/services/`
5. Add handlers in `internal/handlers/`
6. Register routes in `cmd/server/main.go`

### Database Operations
- Use GORM methods through repository pattern
- Soft deletes are enabled on User entity
- Connection pooling configured (10 idle, 100 max connections)
- No auto-migration - database schema managed externally

### Error Handling
- Repository layer returns GORM errors
- Service layer adds business logic validation
- Handler layer converts to appropriate HTTP status codes
- Use `utils.ErrorResponse()` and `utils.SuccessResponse()` for consistency

### Testing
Currently no test files exist. When adding tests:
- Unit tests for services (business logic)
- Integration tests for repositories (database layer)
- HTTP tests for handlers (API layer)
- Use dependency injection for easy mocking
