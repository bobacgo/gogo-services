package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system"
	"github.com/gogoclouds/gogo-services/common-lib/app"
	"net/http"
)

func Init(app *app.App, e *gin.Engine) {
	e.StaticFS("./web", http.Dir("./web"))
	system.Register(app, e)
}
