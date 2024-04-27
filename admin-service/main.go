package main

import (
	"flag"
	"log"

	"github.com/gogoclouds/gogo-services/admin-service/internal/config"
	"github.com/gogoclouds/gogo-services/admin-service/internal/router"
	"github.com/gogoclouds/gogo-services/framework/app"
	"github.com/gogoclouds/gogo-services/framework/app/conf"
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
		app.WithMustConfig(*filepath, func(cfg *conf.ServiceConfig[config.Service]) {
			config.Cfg = cfg
		}),
		app.WithLogger(),
		app.WithMustLocalCache(),
		app.WithMustDB(),
		app.WithMustRedis(),
		app.WithGinServer(router.Init),
		// app.WithGrpcServer(domain.RegisterServer),
		// app.WithRegistrar(etcd.New(etcdClient)),
	)
	if err := newApp.Run(); err != nil {
		log.Panic(err.Error())
	}
}
