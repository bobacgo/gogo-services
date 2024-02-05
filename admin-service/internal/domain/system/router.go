package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system/handler"
)

func Register(r gin.IRouter) {
	adminServer := handler.NewAdminServer()
	r.GET("/register", adminServer.Register)
	r.GET("/login", adminServer.Login)
	r.POST("/logout", adminServer.Logout)
	r.POST("/refreshToken", adminServer.RefreshToken)
}
