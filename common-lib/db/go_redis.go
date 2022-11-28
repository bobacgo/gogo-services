package db

import (
	"github.com/go-redis/redis/v8"
	"github.com/gogoclouds/gogo-services/common-lib/dns/config"
	"golang.org/x/net/context"
	"time"
)

var Redis = redisServer{}

type redisServer struct{}

func (redisServer) Open(ctx context.Context, conf *config.Configuration) (*redis.ClusterClient, error) {
	rConf := conf.Redis()
	readTimeout, err := time.ParseDuration(rConf.ReadTimeout)
	if err != nil {
		return nil, err
	}
	writeTimeout, err := time.ParseDuration(rConf.WriteTimeout)
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        rConf.Addr,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	})
	err = rdb.ForEachShard(ctx, func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	})
	return rdb, err
}