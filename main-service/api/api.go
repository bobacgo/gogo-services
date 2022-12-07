package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/common-lib/g"
	"github.com/gogoclouds/gogo-services/common-lib/web/r"
)

const AdminService = "admin-service"

func Router(e *gin.Engine) {
	e.GET("/ping", func(c *gin.Context) {
		instance, err := g.Service.GetOneInstance(AdminService)
		if err != nil {
			c.JSON(http.StatusOK, r.FailMsg(err.Error()))
			return
		}
		serviceInfo := fmt.Sprintf("%s: ip: %s, port: %d", instance.GetService(), instance.GetHost(), instance.GetPort())
		c.JSON(http.StatusOK, r.SuccessMsg(serviceInfo))
	})
}