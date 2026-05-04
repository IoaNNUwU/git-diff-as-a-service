package core_auth

import "context"

var (
	SessionCookie = "session_key"

	UserID   = "user_id"
	UserRole = "user_role"
)

var (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

type userRoleAcess struct{}

var userRoleKey userRoleAcess

func RoleFromContext(ctx context.Context) string {
	role, ok := ctx.Value(userRoleKey).(string)
	if !ok {
		panic("unable to get logger from context. perhaps this function was called before Auth middleware")
	}
	return role
}

type userIDAcess struct{}

var userIDKey userIDAcess

func UserIDFromContext(ctx context.Context) int {
	userID, ok := ctx.Value(userIDKey).(int)
	if !ok {
		panic("unable to get logger from context. perhaps this function was called before Auth middleware")
	}
	return userID
}
