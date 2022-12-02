package main

import (
	"context"
	"flag"
	"github.com/gogoclouds/gogo-services/common-lib/app"
	"github.com/gogoclouds/gogo-services/main-service/api"
	"github.com/gogoclouds/gogo-services/main-service/internal/model"
)

var config = flag.String("config", "./configs/polaris.yaml", "config file path")

func main() {
	flag.Parse()
	ctx := context.Background()
	app.New(ctx, *config).
		OpenDB(model.Tables).
		OpenCacheDB().
		RunHttp(api.Router)
}