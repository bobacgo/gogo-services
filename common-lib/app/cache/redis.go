package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	maxRetry = 10 // number of retries
)

// NewRedis Initialize redis connection.
func NewRedis(cfg RedisConf) (redis.UniversalClient, error) {
	if len(cfg.Addrs) == 0 {
		return nil, errors.New("redis address is empty")
	}
	var rdb redis.UniversalClient
	if len(cfg.Addrs) > 1 {
		rdb = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:      cfg.Addrs,
			Username:   cfg.Username,
			Password:   cfg.Password, // no password set
			PoolSize:   cfg.PoolSize, // 50
			MaxRetries: maxRetry,
			// ReadTimeout: cfg.ReadTimeout,
			// WriteTimeout: cfg.WriteTimeout,
		})
	} else {
		rdb = redis.NewClient(&redis.Options{
			Addr:       cfg.Addrs[0],
			Username:   cfg.Username,
			Password:   cfg.Password, // no password set
			DB:         int(cfg.DB),  // use default DB
			PoolSize:   cfg.PoolSize, // connection pool size 100
			MaxRetries: maxRetry,
			// ReadTimeout: cfg.ReadTimeout,
			// WriteTimeout: cfg.WriteTimeout,
		})
	}

	var err error
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err = rdb.Ping(ctx).Err()
	if err != nil {
		return nil, fmt.Errorf("redis ping %w", err)
	}

	return rdb, err
}
