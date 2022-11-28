package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/common-lib/web/r"
)

func Router(e *gin.Engine) {
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(200, r.SuccessMsg("pong"))
	})
}