package middleware

import (
	"log/slog"
	"strings"

	"github.com/gin-gonic/gin"
)

type Header struct {
	Auth    string `header:"Authorization"`
	TraceID string `header:"TraceID"`
}

func HeaderToContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := new(Header)
		if err := c.ShouldBindHeader(header); err != nil {
			slog.Warn("header parse fail", slog.String("err", err.Error()))
			c.Next()
			return
		}

		c.Set(AuthHeader, header.Auth)
		c.Set(TraceID, header.TraceID)
	}
}

func TrimToken(bearerToken string) string {
	if bearerToken == "" {
		return ""
	}
	token, found := strings.CutPrefix(bearerToken, tokenPrefix)
	if !found {
		return ""
	}
	return token
}
