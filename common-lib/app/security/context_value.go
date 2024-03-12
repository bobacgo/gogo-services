package security

import (
	"context"
	"strconv"
)

const (
	ClaimsKey = "claims"
)

func GetClaims(ctx context.Context) *Claims {
	claims, _ := ctx.Value(ClaimsKey).(*Claims)
	return claims
}

func GetUserID(ctx context.Context) string {
	claims := GetClaims(ctx)
	if claims == nil {
		return ""
	}
	return claims.UserID
}

// GetUserIntID 如果存储的UID是数字类型，那么可以通过该方法获取
func GetUserIntID(ctx context.Context) int64 {
	id, _ := strconv.ParseInt(GetUserID(ctx), 10, 0)
	return id
}

func GetUsername(ctx context.Context) string {
	claims := GetClaims(ctx)
	if claims == nil {
		return ""
	}
	return claims.Username
}

func GetNickname(ctx context.Context) string {
	claims := GetClaims(ctx)
	if claims == nil {
		return ""
	}
	return claims.Nickname
}

func GetRoles(ctx context.Context) []string {
	claims := GetClaims(ctx)
	if claims == nil {
		return nil
	}
	return claims.Roles
}

// GetRole 单个角色设计时
func GetRole(ctx context.Context) string {
	roles := GetRoles(ctx)
	if len(roles) > 0 {
		return roles[0]
	}
	return ""
}
