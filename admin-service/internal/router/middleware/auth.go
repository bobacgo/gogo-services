package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/admin-service/api/errs"
	"github.com/gogoclouds/gogo-services/common-lib/app/security"
	"github.com/gogoclouds/gogo-services/common-lib/web/r"
	"strings"
)

const (
	tokenPrefix = "Bearer "
	AuthHeader  = "Authorization"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader(AuthHeader)
		if bearerToken == "" {
			c.Abort()
			r.Reply(c, errs.TokenMiss)
			return
		}
		token, found := strings.CutPrefix(bearerToken, tokenPrefix)
		if !found {
			c.Abort()
			r.Reply(c, errs.TokenInvalid.WithDetails(fmt.Errorf("not have prefix [%s]", tokenPrefix)))
			return
		}

		claims, err := security.JwtHelper.Verify(token)
		if err != nil {
			c.Abort()
			if security.JwtHelper.ValidationErrorExpired(err) {
				r.Reply(c, errs.TokenExpired)
				return
			}
			r.Reply(c, errs.TokenInvalid.WithDetails(err))
			return
		}

		tokenID, err := security.JwtHelper.GetTokenID(c, claims.Username)
		if err != nil || tokenID != claims.Id { // 只能单一登录(TODO支持配置可设置)
			c.Abort()
			r.Reply(c, errs.TokenOut.WithDetails(fmt.Errorf("token id not match")))
			return
		}

		// TODO 权限校验
		c.Set(security.ClaimsKey, claims)
		c.Next()
	}
}
