package app

import (
	"context"
	"github.com/gogoclouds/gogo-services/common-lib/g"
	"github.com/gogoclouds/gogo-services/common-lib/internal/db"
	"github.com/gogoclouds/gogo-services/common-lib/internal/db/orm"
	"github.com/gogoclouds/gogo-services/common-lib/internal/dns"
	"github.com/gogoclouds/gogo-services/common-lib/internal/dns/config"
	"github.com/gogoclouds/gogo-services/common-lib/internal/logger"
	server2 "github.com/gogoclouds/gogo-services/common-lib/internal/server"
)

type app struct {
	ctx  context.Context
	conf *config.Configuration
}

// New().OpenDB().OpenCacheDB.RunXxx()

// New 这个函数调用之后会阻塞
// 1. 从配置中心拉取配置文件
// 2. 启动服务
// 3. 注册服务
// 4. 初始必要的全局参数
func New(ctx context.Context, configFilePath string) *app {
	configMD := config.Bootstrap.Unmarshal(configFilePath)
	// 拉取配置
	dns.Server.LoadConfig(configFilePath, configMD)
	g.Log = logger.New(g.Conf.App().Name, g.Conf.Log())
	return &app{ctx: ctx, conf: g.Conf}
}

// OpenDB connect DB
//
// tableModel struct 数据库表
func (s *app) OpenDB(tableModel []any) *app {
	var err error
	if g.DB, err = orm.Server.NewDB(s.ctx, s.conf); err != nil {
		panic(err)
	}
	if err = orm.Server.AutoMigrate(g.DB, tableModel...); err != nil {
		panic(err)
	}
	return s
}

func (s *app) OpenCacheDB() *app {
	var err error
	if g.CacheDB, err = db.Redis.Open(s.ctx, s.conf); err != nil {
		panic(err)
	}
	return s
}

func (s *app) RunHttp(router server2.RegisterHttpFn) {
	httpConf := s.conf.App().Server.Http
	server2.RunHttpServer(httpConf.Addr, router)
}

func (s *app) RunRPC(router server2.RegisterRpcFn) {
	rpcConf := s.conf.App().Server.Rpc
	server2.RunRpcServer(rpcConf.Addr, router)
}