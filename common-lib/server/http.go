package server

import "github.com/gin-gonic/gin"

type RegisterHttpFn func(e *gin.Engine)

func RunHttpServer(addr string, register RegisterHttpFn) {
	e := gin.New()
	e.Use(gin.Logger()) // TODO -> zap.Logger
	e.Use(gin.Recovery())
	register(e)
	if err := e.Run(addr); err != nil {
		panic(err)
	}
}