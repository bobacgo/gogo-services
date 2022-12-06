package app

import (
	"context"

	"github.com/gogoclouds/gogo-services/common-lib/g"
	"github.com/gogoclouds/gogo-services/common-lib/internal/db"
	"github.com/gogoclouds/gogo-services/common-lib/internal/db/orm"
	"github.com/gogoclouds/gogo-services/common-lib/internal/dns"
	"github.com/gogoclouds/gogo-services/common-lib/internal/dns/config"
	"github.com/gogoclouds/gogo-services/common-lib/internal/logger"
	"github.com/gogoclouds/gogo-services/common-lib/internal/server"
	"github.com/gogoclouds/gogo-services/common-lib/pkg"
	"github.com/polarismesh/polaris-go"
)

type app struct {
	ctx        context.Context
	conf       *config.Configuration
	configMD   *config.FileMetadata
	enableRpc  bool
	enableHttp bool
}

// New().OpenDB().OpenCacheDB().RunXxx()

// New 这个函数调用之后会阻塞
// 1. 从配置中心拉取配置文件
// 2. 启动服务
// 3. 注册服务
// 4. 初始必要的全局参数
func New(ctx context.Context, configFilePath string) *app {
	configMD := config.Bootstrap.Unmarshal(configFilePath)
	server := new(dns.ConfigServer)
	// 拉取配置
	server.LoadConfig(configFilePath, configMD)
	g.Log = logger.New(g.Conf.App().Name, g.Conf.Log())
	return &app{ctx: ctx, conf: g.Conf, configMD: configMD}
}

// OpenDB connect DB
//
// tableModel struct 数据库表
func (s *app) OpenDB(tableModel []any) *app {
	var err error
	if g.DB, err = orm.Server.NewDB(s.ctx, s.conf); err != nil {
		panic(err)
	}
	if err = orm.Server.AutoMigrate(g.DB, tableModel); err != nil {
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

func (s *app) CreateHttpServer(router server.RegisterHttpFn) *app {
	httpConf := s.conf.App().Server.Http
	s.enableHttp = true
	go server.RunHttpServer(httpConf.Addr, router)
	return s
}

func (s *app) CreateRpcServer(router server.RegisterRpcFn) *app {
	rpcConf := s.conf.App().Server.Rpc
	s.enableRpc = true
	go server.RunRpcServer(rpcConf.Addr, router)
	return s
}

func (s *app) Run() {
	var port uint16
	if s.enableHttp {
		_, port = pkg.Addr.Parse((s.conf.App().Server.Http.Addr))
	} else if s.enableRpc { // rpc > http
		_, port = pkg.Addr.Parse((s.conf.App().Server.Rpc.Addr))
	} else { // no server run
		return
	}
	ip, err := pkg.Addr.GetOutBoundIP()
	if err != nil {
		panic(err)
	}
	s.registerServer(ip, port)
}

func (s *app) registerServer(ip string, port uint16) {
	pa, err := polaris.NewProviderAPI()
	if err != nil {
		panic(err)
	}
	defer pa.Destroy()
	server := &dns.DiscoverServer{
		Provider:  pa,
		Namespace: s.configMD.Namespace,
		Service:   s.conf.App().Name,
		Host:      ip,
		Port:      port,
	}

	server.Register()
	// block
	server.RunMainLoop()
}
