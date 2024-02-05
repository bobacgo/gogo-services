package main

import (
	"context"
	"flag"
	"github.com/gogoclouds/gogo-services/admin-service/internal/router"

	"github.com/gogoclouds/gogo-services/common-lib/app"
)

var config = flag.String("config", "./polaris.yaml", "config file path")

func main() {
	flag.Parse()
	app.New(context.Background(), *config).
		//OpenDB(model.Tables).
		//OpenCacheDB().
		CreateHttpServer(router.Registers).
		//CreateRpcServer(api.Router).
		Run()
}
