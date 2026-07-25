# GoIgniter

![Go Version](https://img.shields.io/badge/Go-1.24-blue)
![License](https://img.shields.io/badge/License-MIT-green)

**CI4-flavored Go starter kit** — Fiber + GORM + JWT.

GoIgniter is a Go starter kit / boilerplate that feels familiar to developers coming from the CodeIgniter 4 ecosystem. It's not a new framework — it's glue code that connects the best Go libraries with a CI4-inspired structure.

---

## Quick Start

```bash
# 1. Clone
git clone https://github.com/iskandar221201/goigniter.git
cd goigniter

# 2. Copy env
cp .env.example .env
# Edit .env with your database credentials

# 3. Run
go run main.go

# Or with hot-reload (dev)
air
```

---

## Folder Structure

```
goigniter/
├── app/
│   ├── controllers/    # HTTP handlers (embed BaseController)
│   ├── models/         # GORM models + query methods
│   ├── services/       # Business logic layer
│   ├── validations/    # Request validation structs
│   ├── helpers/        # Pure utility functions
│   └── middleware/     # Fiber middleware (JWT auth, RBAC)
├── config/             # Typed config (load dari .env)
├── database/
│   └── migrations/     # SQL migration files (golang-migrate)
├── routes/             # Semua route definition
├── system/             # Core: DB, BaseController, Response, JWT
├── public/             # Static assets
├── storage/            # Logs, uploads, temp
├── .air.toml           # Hot reload config
├── .env.example
└── main.go
```

---

## Conventions

### Controller
- One file per resource: `user_controller.go`, `auth_controller.go`
- Embed `system.BaseController` for access to Respond, BodyParse, CurrentUser
- Method names: `Index`, `Show`, `Create`, `Update`, `Delete`

### Service
- Business logic **must not** be in controllers or models
- Accept dependencies via constructor (DB, config)
- Return `(result, error)` — idiomatic Go

### Model
- One file per entity
- Embed `gorm.Model` for ID, CreatedAt, UpdatedAt, DeletedAt
- Query methods as struct methods, not static

### Validation
- One struct per request with `validate:` tags from go-playground/validator
- Parse errors via `system.ParseValidationErrors` to standard response format

---

## API Endpoints

### Public

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/auth/register` | Register a new user |
| POST | `/api/v1/auth/login` | Login, returns JWT token |

### Protected (Bearer Token)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/users` | List all users |
| POST | `/api/v1/users` | Create a new user |
| GET | `/api/v1/users/:id` | Get user by ID |
| PUT | `/api/v1/users/:id` | Update user |
| DELETE | `/api/v1/users/:id` | Soft delete user |

### Response Format

Success:
```json
{
  "status": true,
  "message": "OK",
  "data": {}
}
```

Error:
```json
{
  "status": false,
  "message": "Validation failed",
  "errors": {
    "email": "email is required"
  }
}
```

---

## Tech Stack

| Layer | Library |
|-------|---------|
| HTTP Framework | [Fiber v2](https://github.com/gofiber/fiber) |
| ORM | [GORM](https://gorm.io) |
| Database Driver | [MySQL](https://github.com/go-sql-driver/mysql) |
| Validation | [go-playground/validator](https://github.com/go-playground/validator) |
| JWT | [golang-jwt/jwt v5](https://github.com/golang-jwt/jwt) |
| Config / Env | [godotenv](https://github.com/joho/godotenv) |
| Migration | [golang-migrate](https://github.com/golang-migrate/migrate) |
| Hot Reload | [air](https://github.com/air-verse/air) |

---

## License

MIT
