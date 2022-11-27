package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/common-lib/app"
	"github.com/gogoclouds/gogo-services/common-lib/dns"
	"github.com/gogoclouds/gogo-services/common-lib/g"
	"github.com/gogoclouds/gogo-services/common-lib/pkg/mapset"
	"github.com/gogoclouds/gogo-services/common-lib/server"
)

var (
	dnsConfigFilePath = "../configs/polaris.yaml"
	dnsNamespace      = "default"
	fileGroup         = "gogo_v1.0.0"
	configFilenames   = []string{
		"admin-service.yaml", "common.yaml", "mysql.yaml", "redis.yaml", "test.yaml",
	}
)

func main() {
	fileMetadata := &dns.FileMetadata{
		Namespace: dnsNamespace, FileGroup: fileGroup, FileNameSet: mapset.Of(configFilenames...),
	}
	app.Init(dnsConfigFilePath, fileMetadata)
	appConf := g.Conf.App()
	server.RunHttpServer(appConf.Server.Http.Addr, router)
}

func router(e *gin.Engine) {
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"code": 0,
			"msg":  "ok",
		})
	})
}
