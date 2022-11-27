package server

import "github.com/gin-gonic/gin"

type registerHttpFn func(e *gin.Engine)

func RunHttpServer(addr string, register registerHttpFn) {
	e := gin.New()
	e.Use(gin.Logger()) // TODO -> zap.Logger
	e.Use(gin.Recovery())
	register(e)
	if err := e.Run(addr); err != nil {
		panic(err)
	}
}
