## SimpleBank — Production Banking API

Go (Golang) backend for a simple banking system with users, accounts, and inter-account transfers. Built with Gin, PostgreSQL, sqlc, and Docker. Migrations run automatically on container start.

### Stack
- Go + Gin
- PostgreSQL
- sqlc (type-safe queries)
- Viper (config)
- Docker/Compose

---

## Quick start (Docker)

From the `simplebank` directory:

```bash
docker compose up -d postgres api
docker compose ps
```

Expected: `postgres` healthy, `api` up and listening on `0.0.0.0:8082`.

Create a user:

```bash
curl -i -X POST http://localhost:8082/users \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"secret123","full_name":"Test User","email":"test@example.com"}'
```

Login:

```bash
curl -i -X POST http://localhost:8082/user/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"secret123"}'
```

Stop services:

```bash
docker compose down -v
```

---

## Run locally without Docker

Prereqs: Go 1.24+, Postgres running on your machine.

1) Configure `simplebank/app.env`:

```env
DB_SOURCE=postgresql://root:secret@127.0.0.1:5432/simple_bank?sslmode=disable
SERVER_ADDRESS=0.0.0.0:8082
TOKEN_SYMMETRIC_KEY=12345678901234567890123456789012
ACCESS_TOKEN_DURATION=15m
```

2) Run the app from the `simplebank` directory so Viper can load `app.env`:

```bash
cd simplebank
go run .
```

Run migrations (optional when not using the container’s auto-migrate):

```bash
docker run --rm -v "$(pwd)/db/migration:/migrations" --network host migrate/migrate:4 \
  -path=/migrations -database "postgresql://root:secret@127.0.0.1:5432/simple_bank?sslmode=disable" up
```

---

## API Endpoints (selection)

- POST `/users` — create user
- POST `/user/login` — login, returns access token
- POST `/account` — create account
- GET `/account/:id` — get account by id
- GET `/accounts` — list accounts (with pagination)
- DELETE `/account/:id` — delete account
- POST `/transfer` — transfer money between accounts

---

## Configuration

Config is loaded by Viper from `app.env` (when running from `simplebank`) and environment variables.

Important keys:
- `DB_SOURCE` — Postgres DSN
- `SERVER_ADDRESS` — `host:port` to bind
- `ACCESS_TOKEN_DURATION` — e.g. `15m`

Compose overrides `DB_SOURCE` to use the Docker network host (`postgres`).

---

## Troubleshooting

- Connection refused to Postgres on macOS:
  - Use IPv4 in DSN: `127.0.0.1` instead of `localhost` (avoids resolving to `::1`).

- `ACCESS_TOKEN_DURATION expected a map or struct` in logs:
  - Ensure `AccessTokenDuration` type is `time.Duration` in code and env value like `15m`.

- Compose says “No such container” on create:
  - Restart Docker Desktop and run `docker compose up -d` again.

- API not reachable:
  - Check logs: `docker compose logs api --since=1m`
  - Verify port mapped: `docker compose ps`

---

## Development

Useful commands from `simplebank/`:

```bash
docker compose build --no-cache api
docker compose up -d postgres api
docker compose logs api --since=1m
docker compose down -v
```
