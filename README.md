# Basketball Management System

A comprehensive basketball management application similar to Cricbuzz, built with Golang (Echo framework) backend and Flutter frontend.

## 🏀 Features

- **Multi-tenant System**: Support for multiple colleges/organizations
- **Team Management**: Create and manage teams with players
- **Match Management**: Schedule matches, track live scores
- **Tournament System**: Organize tournaments and leagues
- **Statistics**: Comprehensive team and player statistics
- **Live Scoring**: Real-time match updates
- **Admin Panel**: GoAdmin-based admin interface
- **Role-based Access**: Super Admin, Organization Admin, Team Members, Public Users

## 🏗️ Architecture

See [ARCHITECTURE.md](./ARCHITECTURE.md) for detailed architecture documentation.

### Tech Stack

**Backend:**
- Echo Framework (Golang)
- PostgreSQL (Database)
- GORM (ORM)
- Redis (Caching & Real-time)
- JWT (Authentication)
- GoAdmin (Admin Panel)

**Frontend:**
- Flutter (Mobile App)

## 📁 Project Structure

```
basketball-app/
├── cmd/server/          # Application entry point
├── internal/
│   ├── config/         # Configuration
│   ├── database/       # DB connection & migrations
│   ├── models/         # Data models
│   ├── handlers/       # HTTP handlers
│   ├── services/       # Business logic
│   ├── repositories/   # Data access layer
│   ├── middleware/     # Middleware (auth, CORS, etc.)
│   ├── utils/          # Utilities
│   └── admin/          # GoAdmin setup
├── pkg/errors/         # Custom errors
├── migrations/         # SQL migrations
└── uploads/            # File uploads
```

## 🚀 Quick Start

### Prerequisites

- Go 1.24+
- PostgreSQL 15+
- Redis 7+
- Docker & Docker Compose (optional)

### Setup

1. **Clone and setup:**
   ```bash
   cd echo-gloang
   go mod download
   ```

2. **Start services with Docker:**
   ```bash
   docker-compose up -d
   ```

3. **Configure environment:**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. **Run migrations:**
   ```bash
   # TODO: Add migration command
   ```

5. **Start server:**
   ```bash
   go run cmd/server/main.go
   ```

## 📚 API Documentation

API endpoints will be documented with Swagger (to be added).

Base URL: `http://localhost:8080/api/v1`

### Main Endpoints

- `/auth/*` - Authentication
- `/organizations/*` - Organization management
- `/teams/*` - Team management
- `/players/*` - Player management
- `/matches/*` - Match management
- `/tournaments/*` - Tournament management
- `/statistics/*` - Statistics
- `/admin` - Admin panel

## 👥 User Roles

1. **Super Admin**: Full system access
2. **Organization Admin**: Manage their organization's teams and matches
3. **Team Members**: View team matches and statistics
4. **Public Users**: View matches, scores, and standings

## 🔐 Security

- JWT-based authentication
- Role-based access control (RBAC)
- Password hashing with bcrypt
- Input validation
- CORS configuration
- Rate limiting (to be implemented)

## 📝 Development Status

🚧 **In Development** - Architecture and project structure defined. Implementation in progress.

## 📄 License

See [LICENSE](./LICENSE) file.

