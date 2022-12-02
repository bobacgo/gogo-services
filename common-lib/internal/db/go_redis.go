package db

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	config2 "github.com/gogoclouds/gogo-services/common-lib/internal/dns/config"
	"github.com/gogoclouds/gogo-services/common-lib/pkg/stream"
	"golang.org/x/net/context"
	"time"
)

var Redis = redisServer{}

type redisServer struct {
	addrs              []string
	rTimeout, wTimeout time.Duration
}

func (s *redisServer) Open(ctx context.Context, conf *config2.Configuration) (cmd redis.Cmdable, err error) {
	rConf := conf.Redis()
	if s.rTimeout, err = time.ParseDuration(rConf.ReadTimeout); err != nil {
		return nil, err
	}
	if s.wTimeout, err = time.ParseDuration(rConf.WriteTimeout); err != nil {
		return nil, err
	}
	s.addrs = stream.New(rConf.Addr).Distinct().List()

	if len(rConf.Addr) > 1 { // Multiple nodes are cluster mode
		return s.clusterMode(ctx)
	}
	return s.standaloneMode(ctx, rConf)
}

func (s *redisServer) standaloneMode(ctx context.Context, conf config2.Redis) (redis.Cmdable, error) {
	if len(s.addrs) == 0 {
		return nil, fmt.Errorf("addr number is zero")
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Addr[0],
		Password: conf.Password, // no password set
		DB:       0,             // use default DB
	})
	err := rdb.Ping(ctx).Err()
	return rdb, err
}

func (s *redisServer) clusterMode(ctx context.Context) (redis.Cmdable, error) {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        s.addrs,
		ReadTimeout:  s.rTimeout,
		WriteTimeout: s.wTimeout,
	})
	err := rdb.ForEachShard(ctx, func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	})
	return rdb, err
}