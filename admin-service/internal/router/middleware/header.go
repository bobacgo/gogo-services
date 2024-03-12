package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func GetAuthHeader(ctx *gin.Context) string {
	bearerToken := ctx.GetHeader(AuthHeader)
	if bearerToken == "" {
		return ""
	}
	token, _ := strings.CutPrefix(bearerToken, tokenPrefix)
	return token
}
