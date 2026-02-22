# Car Wash Backend

A RESTful API backend for a car wash appointment booking system, built with Go.

## Features

- User registration and authentication (JWT-based)
- Appointment scheduling and management
- Appointment status tracking (pending, confirmed, completed, cancelled)

## Tech Stack

- **Language**: Go
- **Router**: [Chi v5](https://github.com/go-chi/chi)
- **Database**: PostgreSQL (via [pgx v5](https://github.com/jackc/pgx))
- **Auth**: JWT ([golang-jwt](https://github.com/golang-jwt/jwt))
- **Query generation**: [sqlc](https://sqlc.dev/)
- **Config**: [godotenv](https://github.com/joho/godotenv)

## Project Structure

```
cmd/api/          - Entry point
internal/
  handler/        - HTTP request handlers
  service/        - Business logic
  repository/     - Data access layer
  domain/         - Domain models and DTOs
  token/          - JWT token management
  database/       - DB connection setup
  db/             - sqlc-generated query code
db/
  migrations/     - Database schema migrations
  queries/        - SQL query definitions
pkg/httputil/     - HTTP response utilities
```

## API Endpoints

| Method | Path             | Description          | Auth required |
|--------|------------------|----------------------|---------------|
| POST   | `/api/login`     | Authenticate user    | No            |
| GET    | `/api/users`     | List all users       | Yes           |
| POST   | `/api/users`     | Create a new user    | No            |

## Getting Started

### Prerequisites

- Go 1.22+
- PostgreSQL

### Setup

1. Clone the repository:
   ```bash
   git clone <repo-url>
   cd car-wash-backend
   ```

2. Create a `.env` file:
   ```env
   DATABASE_URL=postgres://user:password@localhost:5432/carwash?sslmode=disable
   JWT_SECRET=your-secret-key
   ```

3. Run database migrations:
   ```bash
   migrate -source file://db/migrations -database $DATABASE_URL up
   ```

4. Start the server:
   ```bash
   go run cmd/api/main.go
   ```

The server runs on port `3000`.
