package app

import (
	"github.com/gogoclouds/gogo-services/common-lib/dns"
	"github.com/gogoclouds/gogo-services/common-lib/g"
	"github.com/gogoclouds/gogo-services/common-lib/logger"
)

// Init 这个函数调用之后会阻塞
// 1. 从配置中心拉取配置文件
// 2. 启动服务
// 3. 注册服务
// 4. 初始必要的全局参数
func Init(dnsConfigFilePath string, remoteConfigFile *dns.FileMetadata) {
	dns.Server.LoadConfig(dnsConfigFilePath, remoteConfigFile)
	logger.Init(g.Conf.App().Name, g.Conf.Log())
}
