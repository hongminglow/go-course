# Beginner Go Workout API

This repository captures a hands-on Golang course project that builds a small workout tracking API. It walks through beginner-friendly patterns—structuring a Go application, wiring HTTP handlers and middleware, integrating Postgres, and layering authentication—while keeping the code approachable for learners.

## Highlights

- **CRUD workouts** with nested workout entries stored in PostgreSQL.
- **User onboarding** via `POST /users` with passwords hashed using `bcrypt`.
- **Token-based authentication** that issues short-lived bearer tokens and enforces access with middleware.
- **Request middleware** for transparent authentication and route protection (`Authenticate` and `RequireUser`).
- **Database migrations** managed with [Goose](https://github.com/pressly/goose) to keep schemas versioned.
- **Integration tests** for the workout store using a dedicated test database.
- **Go basics refresher** inside `main.go` showcasing core language features (functions, slices, structs, pointers, receivers).

## Project layout

```
internal/
  api/         // HTTP handlers for users, tokens, and workouts
  app/         // Dependency wiring for stores, handlers, middleware
  middleware/  // Authentication helpers
  routes/      // Chi router setup and route registration
  store/       // Postgres stores, migrations helpers, tests
  tokens/      // Token generation utilities
  utils/       // JSON helpers and request parsing utilities
migrations/    // Goose SQL migrations
main.go        // Server bootstrap + Go refresher snippets
```

## Prerequisites

- Go **1.25** or newer (`go.mod` targets Go 1.25.1).
- Docker Desktop (for the local Postgres instances defined in `docker-compose.yml`).
- The Goose CLI (`go install github.com/pressly/goose/v3/cmd/goose@latest`) if you want to run migrations outside of tests.

## Quick start

1. **Start Postgres**
   ```powershell
   docker compose up -d db
   ```
2. **Install Go dependencies**
   ```powershell
   go mod tidy
   ```
3. **Run database migrations** (option A – Goose CLI):
   ```powershell
   goose -dir migrations postgres "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up
   ```
   Option B – reuse the helper from tests:
   ```powershell
   go test ./internal/store -run TestCreateWorkout -count=0
   ```
   (the helper calls `store.Migrate` which applies every migration).
4. **Start the API server**
   ```powershell
   go run . -port 8080
   ```

The server logs to stdout and exposes a readiness probe at `GET /health`.

## API overview

| Method | Path                     | Description                                     | Auth |
|--------|--------------------------|-------------------------------------------------|------|
| GET    | `/health`                | Basic health probe                              | No   |
| POST   | `/users`                 | Register a new user                             | No   |
| POST   | `/tokens/authentication` | Issue a bearer token for an existing user       | No   |
| GET    | `/workouts/{id}`         | Fetch a workout (with entries)                  | Yes  |
| POST   | `/workouts`              | Create a workout with nested entries            | Yes  |
| PUT    | `/workouts/{id}`         | Update workout metadata and entries             | Yes  |
| DELETE | `/workouts/{id}`         | Delete a workout you own                        | Yes  |

Authenticated routes expect `Authorization: Bearer <token>` headers. Tokens are generated via `/tokens/authentication` and are hashed + stored server-side with SHA-256 for validation.

## Authentication pipeline

1. Users register through `/users`; passwords are hashed with `bcrypt` via the `password.Set` helper.
2. Clients exchange credentials for a token at `/tokens/authentication`. Tokens are random 32-byte values encoded in Base32 and persisted with hashes.
3. Incoming requests pass through the `UserMiddleware.Authenticate` middleware, which resolves the bearer token, loads the user, and attaches it to the request context.
4. Protected routes wrap handlers with `RequireUser` to reject anonymous callers.

## Testing

A dedicated test database (`test_db` service in `docker-compose.yml`, mapped to port 5433) keeps integration tests isolated. To run the workout store tests:

```powershell
# start the dedicated test database once
docker compose up -d test_db

# execute the integration tests
go test ./internal/store -run TestCreateWorkout
```

The test helper truncates tables between runs and reapplies migrations, ensuring reproducible results.

## Next steps & experimentation ideas

- Add more workout queries (pagination, filters, per-user listing).
- Extend token scopes (e.g., password reset) using the existing token store.
- Swap in environment-driven database configuration instead of literals in `store.Open`.
- Build a frontend or CLI that consumes the workout endpoints.

Happy hacking and enjoy the Go learning journey! 🎯
