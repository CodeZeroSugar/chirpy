# Chirpy API

Chirpy is a simple social media backend API built in Go. It supports user registration, authentication (with JWT tokens), posting/deleting chirps, and basic admin metrics.

## Features

- User registration and login (with password hashing)
- JWT-based authentication (access tokens and refresh tokens)
- Create, read, and delete chirps
- Update user email/password
- Webhook for upgrading users (e.g., via Polka payment service)
- Health checks, readiness probe, and metrics endpoint
- Profanity filtering on chirps
- Simple file-based database (JSON) – no external DB required
- Serves a static `index.html` frontend page

## Tech Stack

- **Language**: Go 
- **Database**: Postgres

## Prerequisites

- Go 1.25.5+
- PostgreSQL 18.1+
## Running Locally

1. Clone the repository:
   ```bash
   git clone https://github.com/CodeZeroSugar/chirpy.git
   cd chirpy
   ```

2. Run the server:
   ```bash
   go run .
   ```
   Or build and run the binary:
   ```bash
   go build -o chirpy
   ./chirpy
   ```

The server runs on **port 8080** by default.

Environment variables:
- `DB_URL="<database connection string>"`
- `PLATFORM="dev"`
- `SECRET` – Secret key for signing JWT tokens **(required for auth endpoints)**
- `POLKA_KEY` – API key for webhook authentication (upgrade endpoint)

## API Reference

Base URL: `http://localhost:8080`

### Health Check
- **GET** `/api/healthz`  
  Response: `"OK"`  
  (Used for readiness/liveness probes)

### Metrics
- **GET** `/admin/metrics`  
  Serves an HTML page with server hit counter and file statistics.

### Reset (Development only)
- **POST** `/api/reset`  
  Clears all chirps and users (resets database files).

### Chirps

| Method | Endpoint                  | Description                                      | Auth Required      |
|--------|---------------------------|--------------------------------------------------|--------------------|
| GET    | `/api/chirps`             | Get all chirps (sorted by ID asc; supports `author_id` query param and `sort=desc`) | No                 |
| GET    | `/api/chirps/{chirpID}`   | Get a single chirp by ID                         | No                 |
| POST   | `/api/chirps`             | Create a new chirp                               | Yes (Bearer token) |
| DELETE | `/api/chirps/{chirpID}`   | Delete a chirp (only by its author)              | Yes (Bearer token) |

**Create Chirp Request Body**:
```json
{
  "body": "Your chirp text here (max 140 characters)"
}
```

Chirps are automatically filtered for profanity.

### Users

| Method | Endpoint         | Description                          | Auth Required      |
|--------|------------------|--------------------------------------|--------------------|
| POST   | `/api/users`     | Create a new user                    | No                 |
| POST   | `/api/login`     | Login (returns JWT + refresh token)  | No                 |
| PUT    | `/api/users`     | Update email/password                | Yes (Bearer token) |

**Create/Login User Request Body**:
```json
{
  "email": "user@example.com",
  "password": "strongpassword123"
}
```

### Authentication Refresh/Revoke

| Method | Endpoint         | Description                                 | Auth Required                |
|--------|------------------|---------------------------------------------|------------------------------|
| POST   | `/api/refresh`   | Get new access token using refresh token    | Yes (Bearer refresh token)   |
| POST   | `/api/revoke`    | Invalidate refresh token                    | Yes (Bearer refresh token)   |

### Webhook (User Upgrade)

- **POST** `/api/polka/webhooks`  
  Handles payment events from Polka to upgrade a user to "red" (premium) status.  
  Requires header: `Authorization: ApiKey {POLKA_KEY}`

## Example Requests

```bash
# Create a user
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"email": "test@example.com", "password": "password123"}'

# Login
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"email": "test@example.com", "password": "password123"}'

# Create a chirp (replace YOUR_JWT_TOKEN)
curl -X POST http://localhost:8080/api/chirps \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"body": "Hello world!"}'
```

## Error Responses

All errors return JSON in this format:
```json
{
  "error": "Error message here"
}
```

Common status codes:
- 400 – Bad request / validation error
- 401 – Unauthorized
- 403 – Forbidden
- 404 – Not found
- 500 – Internal server error

---
