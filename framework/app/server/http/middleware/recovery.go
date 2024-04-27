package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/framework/web/r"
)

func Recovery() func(c *gin.Context) {
	return gin.CustomRecovery(func(c *gin.Context, err any) {
		r.Reply(c, err)
	})
}
