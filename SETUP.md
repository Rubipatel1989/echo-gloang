# Setup Guide

## Prerequisites

- Go 1.24+
- MySQL 8.0+
- Redis (optional, for caching)

## Installation Steps

### 1. Install Dependencies

```bash
# Install required Go packages
go get github.com/gin-gonic/gin
go get github.com/gin-contrib/cors
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
go get github.com/google/uuid
go get github.com/joho/godotenv
go get gorm.io/gorm
go get gorm.io/driver/mysql
go get github.com/gorilla/websocket
```

Or run:
```bash
go mod tidy
```

### 2. Setup Environment Variables

Create a `.env` file in the root directory:

```env
# Server
PORT=8080
ENV=development

# Database
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=basketball_db
DB_CHARSET=utf8mb4

# JWT
JWT_SECRET=your-super-secret-jwt-key-change-in-production
JWT_EXPIRATION=15m
JWT_REFRESH_EXPIRATION=168h

# File Upload
UPLOAD_DIR=./uploads
MAX_UPLOAD_SIZE=10485760

# CORS
CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:8080
```

### 3. Setup Database

#### Option 1: Using Docker Compose

```bash
docker-compose up -d
```

This will start MySQL and Redis containers.

#### Option 2: Manual MySQL Setup

1. Create database:
```sql
CREATE DATABASE basketball_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

2. Update `.env` with your MySQL credentials

### 4. Run the Application

```bash
go run cmd/server/main.go
```

The server will start on `http://localhost:8080`

### 5. Default Admin Credentials

On first run, a default admin user is created:

- **Email**: `admin@basketball.com`
- **Password**: `admin123`

**⚠️ IMPORTANT**: Change the password immediately after first login!

## API Endpoints

### Authentication

- `POST /api/v1/auth/login` - Login
- `POST /api/v1/auth/register` - Register
- `POST /api/v1/auth/refresh` - Refresh token
- `GET /api/v1/auth/me` - Get current user (protected)

### Admin (Protected - Admin Only)

- `GET /api/v1/admin/dashboard` - Admin dashboard
- `GET /api/v1/admin/users` - List users
- `GET /api/v1/admin/users/:id` - Get user
- `POST /api/v1/admin/users` - Create user
- `PUT /api/v1/admin/users/:id` - Update user
- `DELETE /api/v1/admin/users/:id` - Delete user
- `GET /api/v1/admin/organizations` - List organizations
- `GET /api/v1/admin/organizations/:id` - Get organization
- `POST /api/v1/admin/organizations` - Create organization
- `PUT /api/v1/admin/organizations/:id` - Update organization
- `DELETE /api/v1/admin/organizations/:id` - Delete organization

## Testing the API

### 1. Login as Admin

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@basketball.com",
    "password": "admin123"
  }'
```

Response:
```json
{
  "success": true,
  "data": {
    "user": {
      "id": "...",
      "email": "admin@basketball.com",
      "role": "super_admin",
      ...
    },
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_in": 900
  },
  "message": "Login successful"
}
```

### 2. Get Current User

```bash
curl -X GET http://localhost:8080/api/v1/auth/me \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

### 3. Access Admin Dashboard

```bash
curl -X GET http://localhost:8080/api/v1/admin/dashboard \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

## User Roles

- **super_admin**: Full system access
- **org_admin**: Organization admin (can manage their organization)
- **team_member**: Team member (can view team data)
- **public**: Public user (read-only access)

## Troubleshooting

### Database Connection Error

- Check MySQL is running
- Verify credentials in `.env`
- Ensure database exists

### Port Already in Use

- Change `PORT` in `.env`
- Or kill the process using the port

### JWT Token Invalid

- Check `JWT_SECRET` in `.env`
- Ensure token hasn't expired (default: 15 minutes)
- Use refresh token to get new access token

## Next Steps

1. ✅ Authentication system
2. ✅ Permission system
3. ✅ Admin panel routes
4. ⏳ Complete admin CRUD operations
5. ⏳ Add more admin features
6. ⏳ Implement GoAdmin UI integration

