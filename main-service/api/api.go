package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/common-lib/web/r"
	"net/http"
)

func Router(h http.Handler) {
	e := h.(*gin.Engine)
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, r.SuccessMsg("pong"))
	})
}
