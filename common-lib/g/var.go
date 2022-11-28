package g

import (
	redis "github.com/go-redis/redis/v8"
	"github.com/gogoclouds/gogo-services/common-lib/dns/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

// Conf All 配置
var Conf *config.Configuration

// Log 打印日志
var Log *zap.SugaredLogger

// CacheLocal 一级缓存 变动小、容量少。容量固定，有淘汰策略。
var CacheLocal sync.Map

// CacheDB 二级缓存 容量大，有网络IO延迟
var CacheDB *redis.ClusterClient

// DB gorm 关系型数据库 -- 持久化
var DB *gorm.DB