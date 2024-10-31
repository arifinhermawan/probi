package auth

import "github.com/golang-jwt/jwt/v5"

type contextKey string

const (
	ContextKeyUserID contextKey = "user_id"
)

type customClaim struct {
	UserID int64
	jwt.RegisteredClaims
}
