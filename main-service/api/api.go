package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/framework/g"
	"github.com/gogoclouds/gogo/web/gin/reply"
)

const AdminService = "admin-service"

func Router(e *gin.Engine) {
	e.GET("/ping", func(c *gin.Context) {
		instance, err := g.Service.GetOneInstance(AdminService)
		if err != nil {
			reply.FailMsg(c, err.Error())
			return
		}
		serviceInfo := fmt.Sprintf("%s: ip: %s, port: %d", instance.GetService(), instance.GetHost(), instance.GetPort())
		reply.SuccessMsg(c, serviceInfo)
	})
}
