package middleware

import (
	slogcontext "github.com/PumpkinSeed/slog-context"
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/common-lib/pkg/uid"
)

const (
	TraceID = "traceID"
	Span    = "span"
)

func Trace(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetString(TraceID)
		if traceID == "" {
			traceID = uid.UUID()
		}
		ctx := slogcontext.WithValue(c, TraceID, traceID)
		ctx = slogcontext.WithValue(ctx, Span, serviceName) // TODO 根据配置获取 service name

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
