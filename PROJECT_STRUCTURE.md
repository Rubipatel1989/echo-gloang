# Project Structure Guide

This document provides a detailed breakdown of the project structure and what each component does.

## Directory Structure

```
basketball-app/
│
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point, initializes Echo server
│
├── internal/                        # Private application code
│   │
│   ├── config/
│   │   └── config.go               # Loads and manages environment variables
│   │
│   ├── database/
│   │   ├── connection.go           # PostgreSQL connection setup
│   │   └── migrations/             # Database migration files
│   │
│   ├── models/                     # GORM models (database entities)
│   │   ├── user.go                 # User model with roles
│   │   ├── organization.go         # Organization/College model
│   │   ├── team.go                 # Team model
│   │   ├── player.go               # Player model
│   │   ├── match.go                # Match model
│   │   ├── match_event.go          # Live match events
│   │   ├── tournament.go           # Tournament/League model
│   │   ├── player_statistics.go    # Player match statistics
│   │   └── standing.go             # Tournament standings
│   │
│   ├── handlers/                   # HTTP request handlers (Echo handlers)
│   │   ├── auth_handler.go         # Authentication endpoints
│   │   ├── organization_handler.go # Organization CRUD
│   │   ├── team_handler.go         # Team CRUD and queries
│   │   ├── player_handler.go       # Player CRUD and queries
│   │   ├── match_handler.go        # Match CRUD and live scoring
│   │   ├── tournament_handler.go   # Tournament management
│   │   └── statistics_handler.go   # Statistics endpoints
│   │
│   ├── services/                   # Business logic layer
│   │   ├── auth_service.go         # Authentication logic
│   │   ├── organization_service.go # Organization business logic
│   │   ├── team_service.go         # Team business logic
│   │   ├── player_service.go       # Player business logic
│   │   ├── match_service.go        # Match scheduling and scoring
│   │   ├── tournament_service.go   # Tournament management logic
│   │   └── statistics_service.go   # Statistics calculations
│   │
│   ├── repositories/               # Data access layer (database operations)
│   │   ├── user_repository.go      # User database operations
│   │   ├── organization_repository.go
│   │   ├── team_repository.go
│   │   ├── player_repository.go
│   │   ├── match_repository.go
│   │   ├── tournament_repository.go
│   │   └── statistics_repository.go
│   │
│   ├── middleware/                 # Echo middleware
│   │   ├── auth.go                 # JWT authentication middleware
│   │   ├── cors.go                 # CORS configuration
│   │   ├── logger.go               # Request logging
│   │   ├── role_check.go           # Role-based access control
│   │   └── error_handler.go        # Global error handling
│   │
│   ├── utils/                      # Utility functions
│   │   ├── jwt.go                  # JWT token generation/validation
│   │   ├── password.go             # Password hashing/verification
│   │   ├── validator.go            # Input validation helpers
│   │   ├── response.go             # Standardized API responses
│   │   └── file_upload.go          # File upload handling
│   │
│   └── admin/                      # GoAdmin configuration
│       └── admin.go                # Admin panel setup and routes
│
├── pkg/                            # Public packages (can be imported by other projects)
│   └── errors/
│       └── errors.go               # Custom error types
│
├── migrations/                     # SQL migration files (if using raw SQL)
│   └── *.sql
│
├── uploads/                        # Uploaded files (logos, photos)
│   ├── teams/
│   ├── players/
│   └── organizations/
│
├── .env                            # Environment variables (not in git)
├── .env.example                    # Example environment file
├── .gitignore
├── docker-compose.yml              # Docker services (PostgreSQL, Redis)
├── go.mod                          # Go module file
├── go.sum                          # Go dependencies checksum
├── ARCHITECTURE.md                 # Detailed architecture documentation
├── PROJECT_STRUCTURE.md            # This file
└── README.md                       # Project overview

```

## Layer Responsibilities

### 1. Handlers Layer (`internal/handlers/`)
- **Purpose**: Handle HTTP requests and responses
- **Responsibilities**:
  - Parse request data (JSON, query params, path params)
  - Validate input
  - Call appropriate service methods
  - Format and return responses
  - Handle HTTP status codes

**Example Flow:**
```
Request → Handler → Service → Repository → Database
Response ← Handler ← Service ← Repository ← Database
```

### 2. Services Layer (`internal/services/`)
- **Purpose**: Business logic and orchestration
- **Responsibilities**:
  - Implement business rules
  - Coordinate between multiple repositories
  - Data transformation
  - Validation beyond basic input validation
  - Transaction management

### 3. Repositories Layer (`internal/repositories/`)
- **Purpose**: Data access abstraction
- **Responsibilities**:
  - Database queries (using GORM)
  - CRUD operations
  - Complex queries
  - Data mapping between database and models

### 4. Models Layer (`internal/models/`)
- **Purpose**: Data structures
- **Responsibilities**:
  - Define database schema (GORM tags)
  - Define relationships (has many, belongs to, etc.)
  - Model validation rules

## Key Components

### Authentication Flow
1. User logs in → `auth_handler.go`
2. Validates credentials → `auth_service.go`
3. Checks database → `user_repository.go`
4. Generates JWT → `utils/jwt.go`
5. Returns token to client

### Authorization Flow
1. Request with JWT token
2. `auth.go` middleware validates token
3. `role_check.go` middleware checks permissions
4. Handler processes request if authorized

### Live Scoring Flow
1. Admin adds match event → `match_handler.go`
2. `match_service.go` processes event
3. `match_repository.go` saves to database
4. Redis pub/sub broadcasts update
5. WebSocket sends to connected clients

## Data Flow Example: Creating a Match

```
1. Flutter App → POST /api/v1/matches
   ↓
2. match_handler.go (CreateMatch)
   - Validates request body
   - Extracts user from context (middleware)
   ↓
3. match_service.go (CreateMatch)
   - Validates business rules (teams exist, dates valid)
   - Checks permissions (org admin can only create for their org)
   ↓
4. match_repository.go (Create)
   - Inserts into database using GORM
   ↓
5. Database (PostgreSQL)
   ↓
6. Response flows back up
   ↓
7. Handler returns JSON response
```

## Best Practices

1. **Separation of Concerns**: Each layer has a specific responsibility
2. **Dependency Injection**: Services receive repositories as dependencies
3. **Error Handling**: Use custom error types from `pkg/errors`
4. **Validation**: Validate at handler level (input) and service level (business rules)
5. **Transactions**: Use database transactions for multi-step operations
6. **Logging**: Log important events and errors
7. **Testing**: Unit tests for services, integration tests for handlers

## Next Steps

1. Create the directory structure
2. Set up configuration management
3. Create database models
4. Implement authentication
5. Build CRUD operations for each entity
6. Add business logic
7. Implement live scoring
8. Integrate GoAdmin
9. Add WebSocket support
10. Write tests

