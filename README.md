# MSN Map API

## Architecture

```
yet-another-itsm/
├── cmd/server/           # Application entry point
├── internal/
│   ├── config/           # Configuration management
│   ├── controller/       # HTTP controllers (separated by functionality)
│   ├── database/         # Database connection setup
│   ├── middleware/       # Custom middleware
│   ├── repository/       # Data access layer (sqlc generated)
│   └── service/          # Business logic layer
├── migrations/           # Database migration files
├── sql/
│   ├── queries/          # SQL queries for sqlc
│   └── schema/           # Database schema files
├── docker-compose.yml    # Development environment
├── Dockerfile            # Production container
├── Makefile              # Development commands
└── sqlc.yaml             # sqlc configuration
```

## Conventions
- Follow Zalando API guidelines: https://opensource.zalando.com/restful-api-guidelines/

## Quick Start

### Prerequisites

- Go 1.24+
- Docker & Docker Compose
- Make (optional, for convenience commands)

### 1. Clone and Setup

```bash
git clone <repository-url>
cd yet-another-itsm

# Install dependencies
go mod download

# Generate sqlc code
make sqlc-generate
# or
go run github.com/sqlc-dev/sqlc/cmd/sqlc generate
```

### 2. Start Development Environment

```bash
# Start PostgreSQL with Docker
make docker-up

# Wait a moment for database to initialize, then run migrations
make migrate-up

# Start the development server
make dev
```

The API will be available at `http://localhost:8080`

### 3. Test the API

```bash
# Hello World endpoint
curl http://localhost:8080/api/v1/hello

# Health check
curl http://localhost:8080/health
```

## Development

### Available Commands

```bash
make help              # Show all available commands
make dev               # Run in development mode
make build             # Build the application
make test              # Run tests
make fmt               # Format code
make lint              # Run linter
```

### Database Migration Workflow

This project uses [Goose](https://github.com/pressly/goose) for database migrations and [SQLC](https://sqlc.dev/) for type-safe code generation from SQL. Here's the complete workflow for database schema changes:

#### 1. Database Migration Commands

```bash
make migrate-up        # Apply migrations
make migrate-down      # Rollback migrations
make migrate-status    # Check migration status
make migrate-reset     # Reset all migrations (⚠️ destructive)
make migrate-create NAME=add_users_table  # Create new migration
make sqlc-generate     # Regenerate sqlc code
```

#### 2. Creating Schema Changes

**Step 1: Create a new migration**
```bash
# Create a new migration file
make migrate-create NAME=add_posts_table

# This creates: sql/migrations/YYYYMMDDHHMMSS_add_posts_table.sql
```

**Step 2: Edit the migration file**
```sql
-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_posts_user_id ON posts(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_posts_user_id;
DROP TABLE IF EXISTS posts;
-- +goose StatementEnd
```

**Step 3: Apply the migration**
```bash
make migrate-up
```

#### 3. Adding Database Queries

**Step 1: Create SQL queries in [`sql/queries/`](sql/queries/)**
```sql
-- sql/queries/posts.sql

-- name: GetPost :one
SELECT id, title, content, user_id, created_at, updated_at
FROM posts
WHERE id = $1 LIMIT 1;

-- name: CreatePost :one
INSERT INTO posts (title, content, user_id)
VALUES ($1, $2, $3)
RETURNING id, title, content, user_id, created_at, updated_at;

-- name: ListPostsByUser :many
SELECT id, title, content, user_id, created_at, updated_at
FROM posts
WHERE user_id = $1
ORDER BY created_at DESC;
```

**Step 2: Generate Go code**
```bash
make sqlc-generate
```

This generates type-safe Go functions in [`internal/repository/`](internal/repository/):
- `posts.sql.go` - Generated functions like `GetPost()`, `CreatePost()`, etc.
- `models.go` - Go structs for database tables
- `querier.go` - Interface definitions

#### 4. Migration Best Practices

**Migration Guidelines:**
- Always include both `-- +goose Up` and `-- +goose Down` sections
- Use `IF NOT EXISTS` for CREATE statements to avoid conflicts
- Use `IF EXISTS` for DROP statements in down migrations
- Create indexes for foreign keys and frequently queried columns
- Test migrations on a copy of production data before applying

**Query Guidelines:**
- Use descriptive query names that indicate the operation
- Specify return types: `:one`, `:many`, `:exec`
- Always use parameterized queries (`$1`, `$2`, etc.)
- Group related queries in the same `.sql` file

#### 5. Complete Development Workflow

```bash
# 1. Create new migration
make migrate-create NAME=add_posts_table

# 2. Edit the migration file in sql/migrations/
# 3. Apply migration
make migrate-up

# 4. Add queries in sql/queries/posts.sql
# 5. Generate code
make sqlc-generate

# 6. Check migration status
make migrate-status

# 7. Test your changes
make dev
```

#### 6. Troubleshooting

**Migration fails:**
```bash
# Check current status
make migrate-status

# Rollback if needed
make migrate-down

# Fix the migration file and try again
make migrate-up
```

**Code generation issues:**
- Ensure SQL syntax is correct
- Check [`sqlc.yaml`](sqlc.yaml) configuration
- Verify query comments match SQLC format

**Database connection issues:**
- Ensure PostgreSQL is running: `make docker-up`
- Check database credentials in environment variables
- Wait a moment after starting Docker before running migrations

### Docker Operations

```bash
make docker-up         # Start containers
make docker-down       # Stop containers
make docker-logs       # View logs
make docker-rebuild    # Rebuild containers
```

## Configuration

The application uses environment variables for configuration:

### Server Configuration
- `PORT`: Server port (default: 8080)
- `READ_TIMEOUT`: HTTP read timeout (default: 10s)
- `WRITE_TIMEOUT`: HTTP write timeout (default: 10s)
- `IDLE_TIMEOUT`: HTTP idle timeout (default: 120s)

### Database Configuration
- `DB_HOST`: PostgreSQL host (default: localhost)
- `DB_PORT`: PostgreSQL port (default: 5432)
- `DB_USER`: Database user (default: postgres)
- `DB_PASSWORD`: Database password (default: postgres)
- `DB_NAME`: Database name (default: msn_map_api)
- `DB_SSL_MODE`: SSL mode (default: disable)
- `DB_MAX_CONNS`: Maximum connections (default: 25)
- `DB_MIN_CONNS`: Minimum connections (default: 5)

### Logging Configuration
- `LOG_LEVEL`: Log level - debug, info, warn, error (default: info)
- `LOG_FORMAT`: Log format - json, console (default: json)

## Database Schema

The application includes a sample `users` table:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);
```

## Production Deployment

### Using Docker

```bash
# Build production image
docker build -t yet-another-itsm .

# Run with environment variables
docker run -p 8080:8080 \
  -e DB_HOST=your-postgres-host \
  -e DB_PASSWORD=your-password \
  yet-another-itsm
```

### Using Docker Compose

```bash
# Production deployment
docker-compose up -d
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests: `make test`
5. Format code: `make fmt`
6. Submit a pull request

## License

This project is licensed under the MIT License.