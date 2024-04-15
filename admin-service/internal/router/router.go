package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/admin-service/internal/config"
	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system"
	"github.com/gogoclouds/gogo-services/common-lib/app"
	"github.com/gogoclouds/gogo-services/common-lib/app/security"
)

func Init(e *gin.Engine, app *app.Options) {
	e.StaticFS("./web", http.Dir("./web"))
	//authRouter := e.Use(middleware.Auth())
	//_ = authRouter

	security.JwtHelper = security.NewJWT(&config.Conf.Security.Jwt, app.GetRedis())

	system.Register(app, e)
}
