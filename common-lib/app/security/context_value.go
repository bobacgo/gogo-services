package security

import "context"

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

// GetRole 获取slice[0]角色
func GetRole(ctx context.Context) string {
	claims := GetClaims(ctx)
	if claims == nil {
		return ""
	}
	if len(claims.Roles) > 0 {
		return claims.Roles[0]
	}
	return ""
}
