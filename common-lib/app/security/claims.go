package security

import (
	"context"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
	UserID   string   `json:"userID"`
	Username string   `json:"username"`
	Nickname string   `json:"nickname"`
	Roles    []string `json:"roles"`
}

func (o *Claims) SetCtx(ctx context.Context) {
	context.WithValue(ctx, ClaimsKey, o)
}
