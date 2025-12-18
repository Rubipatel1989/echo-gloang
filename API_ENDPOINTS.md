# API Endpoints Reference

Base URL: `http://localhost:8080/api/v1`

## Authentication Endpoints

| Method | Endpoint | Description | Auth Required | Role |
|--------|----------|-------------|---------------|------|
| POST | `/auth/register` | Register new user | No | - |
| POST | `/auth/login` | Login user | No | - |
| POST | `/auth/refresh` | Refresh access token | Yes | Any |
| POST | `/auth/logout` | Logout user | Yes | Any |
| GET | `/auth/me` | Get current user info | Yes | Any |

## Organization Endpoints

| Method | Endpoint | Description | Auth Required | Role |
|--------|----------|-------------|---------------|------|
| GET | `/organizations` | List all organizations | Yes | Super Admin |
| GET | `/organizations/:id` | Get organization details | Yes | Any |
| POST | `/organizations` | Create organization | Yes | Super Admin |
| PUT | `/organizations/:id` | Update organization | Yes | Super Admin, Org Admin |
| DELETE | `/organizations/:id` | Delete organization | Yes | Super Admin |

## Team Endpoints

| Method | Endpoint | Description | Auth Required | Role |
|--------|----------|-------------|---------------|------|
| GET | `/teams` | List teams (with filters) | No | - |
| GET | `/teams/:id` | Get team details | No | - |
| POST | `/teams` | Create team | Yes | Org Admin |
| PUT | `/teams/:id` | Update team | Yes | Org Admin |
| DELETE | `/teams/:id` | Delete team | Yes | Org Admin |
| GET | `/teams/:id/players` | Get team players | No | - |
| GET | `/teams/:id/matches` | Get team matches | No | - |
| GET | `/teams/:id/statistics` | Get team statistics | No | - |

## Player Endpoints

| Method | Endpoint | Description | Auth Required | Role |
|--------|----------|-------------|---------------|------|
| GET | `/players` | List players (with filters) | No | - |
| GET | `/players/:id` | Get player details | No | - |
| POST | `/players` | Create player | Yes | Org Admin |
| PUT | `/players/:id` | Update player | Yes | Org Admin |
| DELETE | `/players/:id` | Delete player | Yes | Org Admin |
| GET | `/players/:id/statistics` | Get player statistics | No | - |

## Match Endpoints

| Method | Endpoint | Description | Auth Required | Role |
|--------|----------|-------------|---------------|------|
| GET | `/matches` | List matches (with filters) | No | - |
| GET | `/matches/:id` | Get match details | No | - |
| POST | `/matches` | Create match | Yes | Org Admin |
| PUT | `/matches/:id` | Update match | Yes | Org Admin |
| DELETE | `/matches/:id` | Delete match | Yes | Org Admin |
| GET | `/matches/:id/live` | Get live match data | No | - |
| POST | `/matches/:id/events` | Add match event (live scoring) | Yes | Org Admin |
| GET | `/matches/:id/statistics` | Get match statistics | No | - |
| GET | `/matches/upcoming` | Get upcoming matches | No | - |
| GET | `/matches/live` | Get live matches | No | - |
| GET | `/matches/completed` | Get completed matches | No | - |

## Tournament Endpoints

| Method | Endpoint | Description | Auth Required | Role |
|--------|----------|-------------|---------------|------|
| GET | `/tournaments` | List tournaments | No | - |
| GET | `/tournaments/:id` | Get tournament details | No | - |
| POST | `/tournaments` | Create tournament | Yes | Super Admin, Org Admin |
| PUT | `/tournaments/:id` | Update tournament | Yes | Super Admin, Org Admin |
| DELETE | `/tournaments/:id` | Delete tournament | Yes | Super Admin |
| GET | `/tournaments/:id/standings` | Get tournament standings | No | - |
| GET | `/tournaments/:id/matches` | Get tournament matches | No | - |
| GET | `/tournaments/:id/teams` | Get tournament teams | No | - |

## Statistics Endpoints

| Method | Endpoint | Description | Auth Required | Role |
|--------|----------|-------------|---------------|------|
| GET | `/statistics/teams` | Team statistics | No | - |
| GET | `/statistics/players` | Player statistics | No | - |
| GET | `/statistics/leaders` | Statistical leaders | No | - |
| GET | `/statistics/leaders/points` | Points leaders | No | - |
| GET | `/statistics/leaders/rebounds` | Rebound leaders | No | - |
| GET | `/statistics/leaders/assists` | Assist leaders | No | - |

## WebSocket Endpoints

| Endpoint | Description | Auth Required |
|----------|-------------|---------------|
| `/ws/matches/:id/live` | Subscribe to live match updates | No |

## Request/Response Examples

### Register User
```http
POST /api/v1/auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123",
  "full_name": "John Doe",
  "role": "org_admin",
  "organization_id": "uuid-here"
}
```

### Create Team
```http
POST /api/v1/teams
Authorization: Bearer <token>
Content-Type: application/json

{
  "organization_id": "uuid",
  "name": "Lakers",
  "coach_name": "Coach Smith",
  "coach_email": "coach@example.com",
  "coach_phone": "+1234567890"
}
```

### Create Match
```http
POST /api/v1/matches
Authorization: Bearer <token>
Content-Type: application/json

{
  "tournament_id": "uuid",
  "team1_id": "uuid",
  "team2_id": "uuid",
  "scheduled_date": "2024-12-25",
  "scheduled_time": "18:00:00",
  "venue": "Main Court"
}
```

### Add Match Event (Live Scoring)
```http
POST /api/v1/matches/:id/events
Authorization: Bearer <token>
Content-Type: application/json

{
  "event_type": "point",
  "team_id": "uuid",
  "player_id": "uuid",
  "points": 2,
  "quarter": 1,
  "time_remaining": "08:30",
  "description": "2-point field goal"
}
```

### Success Response Format
```json
{
  "success": true,
  "data": {
    "id": "uuid",
    "name": "Example",
    ...
  },
  "message": "Operation successful"
}
```

### Error Response Format
```json
{
  "success": false,
  "error": {
    "code": "UNAUTHORIZED",
    "message": "Invalid credentials",
    "details": null
  }
}
```

## Query Parameters

### Pagination
- `page` - Page number (default: 1)
- `limit` - Items per page (default: 10, max: 100)

### Filtering
- `status` - Filter by status (active, inactive, etc.)
- `organization_id` - Filter by organization
- `team_id` - Filter by team
- `tournament_id` - Filter by tournament
- `date_from` - Filter matches from date
- `date_to` - Filter matches to date

### Sorting
- `sort_by` - Field to sort by (created_at, name, etc.)
- `order` - Sort order (asc, desc)

### Example
```
GET /api/v1/matches?status=live&page=1&limit=20&sort_by=created_at&order=desc
```

## Status Codes

- `200 OK` - Success
- `201 Created` - Resource created
- `400 Bad Request` - Invalid input
- `401 Unauthorized` - Authentication required
- `403 Forbidden` - Insufficient permissions
- `404 Not Found` - Resource not found
- `409 Conflict` - Resource conflict
- `500 Internal Server Error` - Server error

