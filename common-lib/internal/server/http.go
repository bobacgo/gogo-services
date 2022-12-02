package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/common-lib/g"
	"github.com/gogoclouds/gogo-services/common-lib/web/r"
	"net/http"
)

type RegisterHttpFn func(e *gin.Engine)

func RunHttpServer(addr string, register RegisterHttpFn) {
	e := gin.New()
	e.Use(gin.Logger()) // TODO -> zap.Logger
	e.Use(gin.Recovery())
	healthApi(e) // provide health API
	register(e)
	if err := e.Run(addr); err != nil {
		panic(err)
	}
}

func healthApi(e *gin.Engine) {
	e.GET("/health", func(c *gin.Context) {
		appConf := g.Conf.App()
		msg := fmt.Sprintf("%s %s, is active", appConf.Name, appConf.Version)
		c.JSON(http.StatusOK, r.SuccessMsg(msg))
	})
}