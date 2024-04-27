package middleware

import (
	slogcontext "github.com/PumpkinSeed/slog-context"
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/framework/pkg/uid"
)

// X-Request-ID

const (
	TraceID = "traceID"
	Span    = "span"
)

func Trace(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetString(TraceID)
		if traceID == "" {
			traceID = uid.UUID()
			c.Header(TraceID, traceID)
		}

		ctx := slogcontext.WithValue(c.Request.Context(), TraceID, traceID)
		ctx = slogcontext.WithValue(ctx, Span, serviceName)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
