package main

import (
	"flag"

	"github.com/gogoclouds/gogo-services/main-service/internal/config"

	"github.com/gogoclouds/gogo-services/framework/app"
	"github.com/gogoclouds/gogo-services/framework/app/conf"
	"github.com/gogoclouds/gogo-services/framework/app/logger"
)

var filepath = flag.String("config", "./config.yaml", "config file path")

func init() {
	flag.String("name", "main-service", "service name")
	flag.String("env", "dev", "run config context")
	flag.String("logger.level", "info", "logger level")
	flag.Int("port", 8081, "http port 8080, rpc port 9080")
	conf.BindPFlags()
}

func main() {
	newApp := app.New(
		app.WithMustConfig(*filepath, func(cfg *conf.ServiceConfig[config.Service]) {
			config.Conf = cfg
		}),
		app.WithLogger(),
		app.WithMustLocalCache(),
		app.WithMustDB(),
		app.WithMustRedis(),
		//app.WithGinServer(nil),
		// app.WithGrpcServer(domain.RegisterServer),
		// app.WithRegistrar(etcd.New(etcdClient)),
	)
	if err := newApp.Run(); err != nil {
		logger.Panic(err.Error())
	}
}
