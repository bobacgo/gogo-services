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
	ctx := context.Background()
	app.New(ctx, *config).
		OpenDB(model.Tables).
		OpenCacheDB().
		CreateRpcServer(api.Router).
		Run()
}
