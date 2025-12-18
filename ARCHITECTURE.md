# Basketball App Architecture

## Overview
A comprehensive basketball management system similar to Cricbuzz, with multi-tenant support for colleges/organizations, team management, match tracking, and live scores.

## Technology Stack

### Backend
- **Framework**: Gin (Golang) - High performance, popular framework
- **Database**: MySQL 8.0+ (PostgreSQL also supported via GORM)
- **ORM**: GORM (supports both MySQL and PostgreSQL)
- **Authentication**: JWT (JSON Web Tokens)
- **Caching**: Redis (for live scores and frequently accessed data)
- **File Storage**: Local filesystem or AWS S3 (for team logos, player photos)
- **Admin Panel**: GoAdmin

### Frontend
- **Mobile App**: Flutter
- **Admin Panel**: GoAdmin (Go-based admin interface)

## User Roles & Permissions

### 1. Super Admin
- Full system access
- Manage all colleges/organizations
- Manage all teams and matches
- System configuration
- User management

### 2. College/Organization Admin
- Manage their own organization
- Create and manage teams
- Add/remove players
- Schedule matches
- Update match scores
- View statistics for their teams

### 3. Team Members/Coaches
- View their team's matches
- View team statistics
- View player statistics
- View upcoming fixtures

### 4. Public Users
- View live scores
- View match schedules
- View standings/leaderboards
- View team and player statistics

## Database Schema

### Core Tables

#### Organizations
- id (UUID, Primary Key)
- name
- email
- phone
- address
- logo_url
- admin_user_id (FK to Users)
- status (active/inactive)
- created_at
- updated_at

#### Users
- id (UUID, Primary Key)
- email (unique)
- password (hashed)
- role (super_admin, org_admin, team_member, public)
- organization_id (FK, nullable)
- full_name
- phone
- profile_image_url
- status (active/inactive)
- created_at
- updated_at

#### Teams
- id (UUID, Primary Key)
- organization_id (FK)
- name
- logo_url
- coach_name
- coach_phone
- coach_email
- description
- status (active/inactive)
- created_at
- updated_at

#### Players
- id (UUID, Primary Key)
- team_id (FK)
- jersey_number
- full_name
- position (PG, SG, SF, PF, C)
- height
- weight
- date_of_birth
- photo_url
- status (active/inactive)
- created_at
- updated_at

#### Tournaments/Leagues
- id (UUID, Primary Key)
- name
- description
- start_date
- end_date
- status (upcoming, ongoing, completed)
- created_by (FK to Users)
- created_at
- updated_at

#### Matches
- id (UUID, Primary Key)
- tournament_id (FK, nullable)
- team1_id (FK)
- team2_id (FK)
- scheduled_date
- scheduled_time
- venue
- status (scheduled, live, completed, cancelled)
- team1_score
- team2_score
- winner_team_id (FK, nullable)
- referee_name
- notes
- created_by (FK to Users)
- created_at
- updated_at

#### Match_Events (for live scoring)
- id (UUID, Primary Key)
- match_id (FK)
- event_type (point, foul, timeout, substitution, quarter_end)
- team_id (FK)
- player_id (FK, nullable)
- points (for point events)
- quarter
- time_remaining
- description
- created_at

#### Player_Statistics
- id (UUID, Primary Key)
- player_id (FK)
- match_id (FK)
- points
- rebounds
- assists
- steals
- blocks
- turnovers
- fouls
- minutes_played
- field_goals_made
- field_goals_attempted
- three_pointers_made
- three_pointers_attempted
- free_throws_made
- free_throws_attempted
- created_at
- updated_at

#### Standings
- id (UUID, Primary Key)
- tournament_id (FK)
- team_id (FK)
- wins
- losses
- points_for
- points_against
- win_percentage
- rank
- updated_at

## API Structure

### Base URL
```
/api/v1
```

### Authentication Endpoints
```
POST   /auth/register          - Register new user
POST   /auth/login             - Login
POST   /auth/refresh            - Refresh token
POST   /auth/logout             - Logout
GET    /auth/me                 - Get current user
```

### Organization Endpoints
```
GET    /organizations           - List all organizations (admin only)
GET    /organizations/:id       - Get organization details
POST   /organizations           - Create organization (admin only)
PUT    /organizations/:id       - Update organization
DELETE /organizations/:id       - Delete organization (admin only)
```

### Team Endpoints
```
GET    /teams                   - List teams (with filters)
GET    /teams/:id               - Get team details
POST   /teams                   - Create team (org admin)
PUT    /teams/:id               - Update team
DELETE /teams/:id               - Delete team
GET    /teams/:id/players       - Get team players
GET    /teams/:id/matches       - Get team matches
GET    /teams/:id/statistics    - Get team statistics
```

### Player Endpoints
```
GET    /players                 - List players (with filters)
GET    /players/:id             - Get player details
POST   /players                 - Create player (org admin)
PUT    /players/:id             - Update player
DELETE /players/:id             - Delete player
GET    /players/:id/statistics  - Get player statistics
```

### Match Endpoints
```
GET    /matches                 - List matches (with filters)
GET    /matches/:id             - Get match details
POST   /matches                 - Create match (org admin)
PUT    /matches/:id             - Update match
DELETE /matches/:id             - Delete match
GET    /matches/:id/live        - Get live match data
POST   /matches/:id/events      - Add match event (live scoring)
GET    /matches/:id/statistics  - Get match statistics
```

### Tournament Endpoints
```
GET    /tournaments             - List tournaments
GET    /tournaments/:id         - Get tournament details
POST   /tournaments             - Create tournament (admin)
PUT    /tournaments/:id         - Update tournament
DELETE /tournaments/:id         - Delete tournament
GET    /tournaments/:id/standings - Get tournament standings
GET    /tournaments/:id/matches - Get tournament matches
```

### Statistics Endpoints
```
GET    /statistics/teams        - Team statistics
GET    /statistics/players      - Player statistics
GET    /statistics/leaders      - Statistical leaders
```

## Project Structure

```
basketball-app/
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go              # Configuration management
│   ├── database/
│   │   ├── connection.go          # DB connection
│   │   └── migrations/             # Database migrations
│   ├── models/
│   │   ├── user.go
│   │   ├── organization.go
│   │   ├── team.go
│   │   ├── player.go
│   │   ├── match.go
│   │   ├── tournament.go
│   │   └── statistics.go
│   ├── handlers/
│   │   ├── auth_handler.go
│   │   ├── organization_handler.go
│   │   ├── team_handler.go
│   │   ├── player_handler.go
│   │   ├── match_handler.go
│   │   ├── tournament_handler.go
│   │   └── statistics_handler.go
│   ├── services/
│   │   ├── auth_service.go
│   │   ├── organization_service.go
│   │   ├── team_service.go
│   │   ├── player_service.go
│   │   ├── match_service.go
│   │   ├── tournament_service.go
│   │   └── statistics_service.go
│   ├── repositories/
│   │   ├── user_repository.go
│   │   ├── organization_repository.go
│   │   ├── team_repository.go
│   │   ├── player_repository.go
│   │   ├── match_repository.go
│   │   ├── tournament_repository.go
│   │   └── statistics_repository.go
│   ├── middleware/
│   │   ├── auth.go                # JWT authentication
│   │   ├── cors.go                # CORS handling
│   │   ├── logger.go              # Request logging
│   │   └── role_check.go          # Role-based access control
│   ├── utils/
│   │   ├── jwt.go                 # JWT utilities
│   │   ├── password.go            # Password hashing
│   │   ├── validator.go           # Input validation
│   │   └── response.go            # Standardized API responses
│   └── admin/
│       └── admin.go               # GoAdmin setup
├── pkg/
│   └── errors/
│       └── errors.go              # Custom error types
├── migrations/
│   └── *.sql                      # SQL migration files
├── uploads/                       # File uploads directory
├── .env                           # Environment variables
├── .env.example                   # Example env file
├── go.mod
├── go.sum
├── docker-compose.yml             # For local development
└── README.md
```

## Authentication & Authorization

### JWT Token Structure
```json
{
  "user_id": "uuid",
  "email": "user@example.com",
  "role": "org_admin",
  "organization_id": "uuid",
  "exp": 1234567890
}
```

### Middleware Flow
1. Extract token from Authorization header
2. Validate token
3. Check user status (active/inactive)
4. Attach user context to request
5. Role-based route protection

## Real-time Features

### Live Score Updates
- WebSocket connection for live matches
- Redis pub/sub for score updates
- Match events stored in real-time
- Flutter app subscribes to match events

### WebSocket Endpoints
```
WS /ws/matches/:id/live    - Subscribe to live match updates
```

**Note:** Gin uses `gorilla/websocket` for WebSocket support, which is the industry standard and very reliable.

## Admin Panel Integration (GoAdmin)

### Admin Routes
- `/admin` - GoAdmin dashboard
- Admin panel will use the same API endpoints
- Role-based access in admin panel
- Custom admin views for:
  - Organizations management
  - Teams management
  - Matches management
  - User management
  - Statistics dashboard

## Security Considerations

1. **Password Hashing**: bcrypt with salt rounds
2. **JWT Expiration**: Access token (15min), Refresh token (7 days)
3. **Rate Limiting**: Prevent API abuse
4. **Input Validation**: Validate all inputs
5. **SQL Injection**: Use parameterized queries (GORM handles this)
6. **CORS**: Configure for Flutter app domain
7. **File Upload**: Validate file types and sizes
8. **HTTPS**: Use in production

## Deployment Architecture

```
┌─────────────┐
│   Flutter   │
│     App     │
└──────┬──────┘
       │ HTTPS
       │
┌──────▼──────────────────┐
│   Load Balancer (Nginx) │
└──────┬──────────────────┘
       │
┌──────▼──────────────┐
│   Gin API Server    │
└──────┬──────────────┘
       │
┌──────▼──────┐  ┌──────────┐
│    MySQL    │  │  Redis   │
└─────────────┘  └──────────┘
```

## Development Phases

### Phase 1: Foundation
- Project setup
- Database schema
- Authentication system
- Basic CRUD for organizations, teams, players

### Phase 2: Core Features
- Match management
- Tournament system
- Basic statistics

### Phase 3: Advanced Features
- Live scoring
- Real-time updates (WebSocket)
- Advanced statistics
- Standings calculation

### Phase 4: Admin Panel
- GoAdmin integration
- Admin dashboard
- Reporting features

### Phase 5: Optimization
- Caching strategy
- Performance optimization
- API documentation (Swagger)

## API Response Format

### Success Response
```json
{
  "success": true,
  "data": { ... },
  "message": "Operation successful"
}
```

### Error Response
```json
{
  "success": false,
  "error": {
    "code": "ERROR_CODE",
    "message": "Error message",
    "details": { ... }
  }
}
```

## Environment Variables

```env
# Server
PORT=8080
ENV=development

# Database (MySQL)
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=basketball_db
DB_CHARSET=utf8mb4
DB_PARSE_TIME=true
DB_LOC=Local

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

# JWT
JWT_SECRET=your-secret-key
JWT_EXPIRATION=15m
JWT_REFRESH_EXPIRATION=168h

# File Upload
UPLOAD_DIR=./uploads
MAX_UPLOAD_SIZE=10485760

# Admin
ADMIN_USERNAME=admin
ADMIN_PASSWORD=admin123
```

