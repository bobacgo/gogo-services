package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system"
)

func Registers(e *gin.Engine) {
	system.Register(e)
}
