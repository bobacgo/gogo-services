package app

import (
	"context"
	"strconv"
	"strings"

	"github.com/gogoclouds/gogo-services/common-lib/g"
	"github.com/gogoclouds/gogo-services/common-lib/internal/db"
	"github.com/gogoclouds/gogo-services/common-lib/internal/db/orm"
	"github.com/gogoclouds/gogo-services/common-lib/internal/dns"
	"github.com/gogoclouds/gogo-services/common-lib/internal/dns/config"
	"github.com/gogoclouds/gogo-services/common-lib/internal/logger"
	"github.com/gogoclouds/gogo-services/common-lib/internal/server"
	"github.com/polarismesh/polaris-go"
)

type app struct {
	ctx      context.Context
	conf     *config.Configuration
	configMD *config.FileMetadata
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

func (s *app) CreateHttpServer(router server.HttpHandlerFn) *app {
	httpConf := s.conf.App().Server.Http
	go server.RunHttpServer(httpConf.Addr, router)
	return s
}

func (s *app) CreateRpcServer(router server.HttpHandlerFn) *app {
	rpcConf := s.conf.App().Server.Rpc
	go server.RunRpcServer(rpcConf.Addr, router)
	return s
}

func (s *app) Run() {
	pa, err := polaris.NewProviderAPI()
	if err != nil {
		panic(err)
	}
	defer pa.Destroy()

	// TODO
	_, port := addrSplitHostPort(s.conf.App().Server.Rpc.Addr)
	server := &dns.DiscoverServer{
		Provider:  pa,
		Namespace: s.configMD.Namespace,
		Service:   s.conf.App().Name,
		Host:      "127.0.0.1", // TODO
		Port:      port,
	}

	server.Register()
	// block
	server.RunMainLoop()
}

// TODO
func addrSplitHostPort(addr string) (string, uint16) {
	if addr != "" {
		return "", 0
	}
	ipAndPort := strings.Split(addr, ":")
	if len(ipAndPort) < 1 {
		return "", 0
	}
	port, err := strconv.Atoi(ipAndPort[1])
	if err != nil {
		return "", 0
	}
	return ipAndPort[0], uint16(port)
}
