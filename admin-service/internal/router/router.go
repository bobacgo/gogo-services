package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system"
	"github.com/gogoclouds/gogo-services/admin-service/internal/router/middleware"
	app "github.com/gogoclouds/gogo-services/framework/app"
	"github.com/gogoclouds/gogo-services/framework/app/security"
)

func Init(e *gin.Engine, app *app.Options) {
	e.StaticFS("./web", http.Dir("./web"))

	cfg := app.Conf()

	security.JwtHelper = security.NewJWT(&cfg.Security.Jwt, app.Redis())
	e.Use(middleware.HeaderToContext(), middleware.Trace(cfg.Name))
	system.Register(app, e)
}
