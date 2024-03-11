package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/admin-service/internal/config"
	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system"
	"github.com/gogoclouds/gogo-services/common-lib/app"
	"github.com/gogoclouds/gogo-services/common-lib/app/security"
	"net/http"
)

func Init(app *app.App, e *gin.Engine) {
	e.StaticFS("./web", http.Dir("./web"))
	//authRouter := e.Use(middleware.Auth())
	//_ = authRouter

	security.JwtHelper = security.NewJWT(&config.Conf.Security.Jwt, app.Opts.Redis)

	system.Register(app, e)
}
