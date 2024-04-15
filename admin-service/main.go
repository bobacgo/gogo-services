package main

import (
	"flag"

	"github.com/gogoclouds/gogo-services/admin-service/internal/config"
	"github.com/gogoclouds/gogo-services/admin-service/internal/router"
	"github.com/gogoclouds/gogo-services/common-lib/app"
	"github.com/gogoclouds/gogo-services/common-lib/app/conf"
	"github.com/gogoclouds/gogo-services/common-lib/app/logger"
)

var filepath = flag.String("config", "./config.yaml", "config file path")

func init() {
	flag.String("name", "admin-service", "service name")
	flag.String("env", "dev", "run config context")
	flag.String("logger.level", "info", "logger level")
	flag.Int("port", 8080, "http port 8080, rpc port 9080")
	conf.BindPFlags()
}

func main() {
	newApp := app.New(
		app.WithConfig(*filepath, func(cfg *conf.ServiceConfig[config.Service]) {
			config.Conf = cfg
		}),
		app.WithLogger(),
		app.WithLocalCache(),
		app.WithDB(),
		app.WithRedis(),
		app.WithGinServer(router.Init),
		// app.WithGrpcServer(domain.RegisterServer),
		// app.WithRegistrar(etcd.New(etcdClient)),
		//app.WithAfterStart()
	)
	if err := newApp.Run(); err != nil {
		logger.Panic(err.Error())
	}
}
