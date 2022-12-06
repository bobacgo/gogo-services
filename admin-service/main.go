package main

import (
	"context"
	"flag"

	"github.com/gogoclouds/gogo-services/admin-service/api"
	"github.com/gogoclouds/gogo-services/admin-service/internal/model"
	"github.com/gogoclouds/gogo-services/common-lib/app"
)

var config = flag.String("config", "./polaris.yaml", "config file path")

func main() {
	flag.Parse()
	app.New(context.Background(), *config).
		OpenDB(model.Tables).
		OpenCacheDB().
		CreateRpcServer(api.Router).
		Run()
}
