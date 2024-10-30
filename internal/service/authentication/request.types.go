package authentication

import (
	"github.com/golang-jwt/jwt/v5"
)

type customClaim struct {
	UserID int64
	jwt.RegisteredClaims
}
