package db

import (
	"github.com/gogoclouds/gogo-services/common-lib/db/mysql"
	"github.com/gogoclouds/gogo-services/common-lib/db/redis"
	"github.com/gogoclouds/gogo-services/common-lib/g"
	"golang.org/x/net/context"
)

func OpenMySQL() {
	g.DB = mysql.Open()
}

func OpenRedis(ctx context.Context) {
	g.CacheDB = redis.Open(ctx)
}