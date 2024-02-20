package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system"
	"github.com/gogoclouds/gogo-services/common-lib/app"
)

func Init(app *app.App, e *gin.Engine) {
	system.Register(app.Opts.DB, e)
}
