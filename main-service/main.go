package main

import (
	"flag"

	"github.com/gogoclouds/gogo-services/common-lib/app"
	"github.com/gogoclouds/gogo-services/common-lib/app/conf"
	"github.com/gogoclouds/gogo-services/common-lib/app/logger"
)

var filepath = flag.String("config", "config/config.yaml", "config file path")

func init() {
	flag.String("name", "main-service", "service name")
	flag.String("env", "dev", "run config context")
	flag.String("logger.level", "info", "logger level")
	flag.Int("port", 8081, "http port 8080, rpc port 9080")
	conf.BindPFlags()
}

func main() {
	newApp := app.New(
		app.WithConfig(*filepath),
		app.WithLogger(),
		app.WithDB(),
		app.WithRedis(),
		//app.WithGinServer(nil),
		// app.WithGrpcServer(domain.RegisterServer),
		// app.WithRegistrar(etcd.New(etcdClient)),
	)
	if err := newApp.Run(); err != nil {
		logger.Panic(err.Error())
	}
}