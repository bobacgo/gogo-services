package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system"
	"github.com/gogoclouds/gogo-services/admin-service/internal/router/middleware"
	"github.com/gogoclouds/gogo-services/common-lib/app"
	"github.com/gogoclouds/gogo-services/common-lib/app/security"
)

func Init(e *gin.Engine, app *app.Options) {
	e.StaticFS("./web", http.Dir("./web"))
	//authRouter := e.Use(middleware.Auth())
	//_ = authRouter
	cfg := app.Conf()

	security.JwtHelper = security.NewJWT(&cfg.Security.Jwt, app.Redis())
	e.Use(middleware.HeaderToContext(), middleware.Trace(cfg.Name))
	system.Register(app, e)
}
