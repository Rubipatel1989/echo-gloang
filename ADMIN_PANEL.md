# Admin Panel Documentation

## Overview

The admin panel provides a RESTful API for managing the basketball application. It includes authentication, authorization, and role-based access control.

## Features Implemented

### ✅ Authentication System
- User login with JWT tokens
- User registration
- Token refresh mechanism
- Password hashing with bcrypt
- JWT token validation

### ✅ Permission System
- Role-based access control (RBAC)
- Four user roles:
  - `super_admin`: Full system access
  - `org_admin`: Organization admin
  - `team_member`: Team member
  - `public`: Public user
- Middleware for role checking
- Organization-level access control

### ✅ Admin Routes
- Admin dashboard
- User management (CRUD)
- Organization management (CRUD)
- Protected routes with authentication

## Authentication Flow

### 1. Login
```
POST /api/v1/auth/login
Body: { "email": "...", "password": "..." }
Response: { access_token, refresh_token, user }
```

### 2. Using Token
```
Header: Authorization: Bearer <access_token>
```

### 3. Refresh Token
```
POST /api/v1/auth/refresh
Body: { "refresh_token": "..." }
Response: { access_token, refresh_token }
```

## Permission Middleware

### Available Middleware

1. **AuthMiddleware()** - Validates JWT token
2. **RequireRole(roles...)** - Checks if user has required role
3. **RequireAdmin()** - Requires super_admin role
4. **RequireOrgAdmin()** - Requires org_admin or super_admin
5. **CheckOrganizationAccess()** - Checks organization access

### Usage Example

```go
// Require authentication
protected := api.Group("")
protected.Use(middleware.AuthMiddleware())

// Require admin role
admin := protected.Group("/admin")
admin.Use(middleware.RequireAdmin())
```

## Admin Endpoints

### Dashboard
- `GET /api/v1/admin/dashboard` - Admin dashboard statistics

### User Management
- `GET /api/v1/admin/users` - List all users
- `GET /api/v1/admin/users/:id` - Get user details
- `POST /api/v1/admin/users` - Create new user
- `PUT /api/v1/admin/users/:id` - Update user
- `DELETE /api/v1/admin/users/:id` - Delete user

### Organization Management
- `GET /api/v1/admin/organizations` - List all organizations
- `GET /api/v1/admin/organizations/:id` - Get organization details
- `POST /api/v1/admin/organizations` - Create organization
- `PUT /api/v1/admin/organizations/:id` - Update organization
- `DELETE /api/v1/admin/organizations/:id` - Delete organization

## Default Admin User

On first application start, a default admin user is created:

- **Email**: `admin@basketball.com`
- **Password**: `admin123`
- **Role**: `super_admin`

⚠️ **Change the password immediately after first login!**

## Security Features

1. **Password Hashing**: bcrypt with default cost
2. **JWT Tokens**: 
   - Access token: 15 minutes
   - Refresh token: 7 days
3. **Token Validation**: Validates on every request
4. **User Status Check**: Only active users can login
5. **Role Verification**: Middleware checks roles before access

## Response Format

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

## Error Codes

- `BAD_REQUEST` - Invalid request data
- `UNAUTHORIZED` - Authentication required or invalid
- `FORBIDDEN` - Insufficient permissions
- `NOT_FOUND` - Resource not found
- `INTERNAL_ERROR` - Server error

## Testing Admin Panel

### 1. Login
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@basketball.com","password":"admin123"}'
```

### 2. Get Dashboard
```bash
curl -X GET http://localhost:8080/api/v1/admin/dashboard \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### 3. List Users
```bash
curl -X GET http://localhost:8080/api/v1/admin/users \
  -H "Authorization: Bearer YOUR_TOKEN"
```

## Next Steps

1. ✅ Authentication & Permissions - Complete
2. ⏳ Complete CRUD operations for users
3. ⏳ Complete CRUD operations for organizations
4. ⏳ Add pagination and filtering
5. ⏳ Add search functionality
6. ⏳ Integrate GoAdmin UI (frontend)
7. ⏳ Add audit logging
8. ⏳ Add email verification

## File Structure

```
internal/
├── admin/
│   └── admin.go              # Admin routes and handlers
├── handlers/
│   └── auth_handler.go       # Authentication handlers
├── middleware/
│   ├── auth.go               # JWT authentication
│   ├── permission.go         # Role-based access control
│   └── cors.go               # CORS configuration
├── services/
│   └── auth_service.go       # Authentication business logic
├── repositories/
│   └── user_repository.go    # User data access
├── models/
│   ├── user.go               # User model
│   └── organization.go       # Organization model
└── utils/
    ├── jwt.go                # JWT utilities
    ├── password.go           # Password hashing
    └── response.go           # Response helpers
```

