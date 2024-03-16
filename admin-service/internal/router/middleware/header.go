package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAuthHeader(ctx *gin.Context) string {
	bearerToken := ctx.GetHeader(AuthHeader)
	if bearerToken == "" {
		return ""
	}
	token, found := strings.CutPrefix(bearerToken, tokenPrefix)
	if !found {
		return ""
	}
	return token
}
