package app

import (
	"context"
	"log"
	"log/slog"
	"net/url"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/common-lib/app/cache"
	"github.com/gogoclouds/gogo-services/common-lib/app/conf"
	"github.com/gogoclouds/gogo-services/common-lib/app/db"
	"github.com/gogoclouds/gogo-services/common-lib/app/logger"
	"github.com/gogoclouds/gogo-services/common-lib/app/registry"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Option func(o *Options)

type Options struct {
	// 内部属性不能直接开放
	conf       *conf.BasicConfig
	localCache cache.Cache
	db         *gorm.DB
	redis      redis.UniversalClient

	appid           string
	endpoints       []*url.URL
	sigs            []os.Signal
	registrar       registry.ServiceRegistrar
	registryTimeout time.Duration

	httpServer func(e *gin.Engine, a *Options)

	rpcServer func(s *grpc.Server, a *Options)
	// Before and After hook
	beforeStart, beforeStop, afterStart, afterStop []func(context.Context) error
}

// Conf 获取公共配置(eg app info、logger config、db config 、redis config)
func (o Options) Conf() *conf.BasicConfig {
	return o.conf
}

// LocalCache 获取本地缓存 Interface
func (o Options) LocalCache() cache.Cache {
	return o.localCache
}

// DB 获取数据库默认client
func (o Options) DB() *gorm.DB {
	return o.db
}

// DBByKey 获取数据库client
func (o Options) DBByKey(key string) *gorm.DB {
	// TODO 多个数据源时
	return o.db
}

// Redis 获取redis client
func (o Options) Redis() redis.UniversalClient {
	return o.redis
}

func WithAppId(id string) Option {
	return func(o *Options) {
		o.appid = id
	}
}

func WithEndpoints(endpoints []*url.URL) Option {
	return func(o *Options) {
		o.endpoints = endpoints
	}
}

func WithSignal(sigs []os.Signal) Option {
	return func(o *Options) {
		o.sigs = sigs
	}
}

func WithRegistrar(registrar registry.ServiceRegistrar) Option {
	return func(o *Options) {
		o.registrar = registrar
	}
}

func WithRegistrarTimeout(rt time.Duration) Option {
	return func(o *Options) {
		o.registryTimeout = rt
	}
}

func WithMustConfig[T any](filename string, fn func(cfg *conf.ServiceConfig[T])) Option {
	return func(o *Options) {
		cfg, err := conf.Load[conf.ServiceConfig[T]](filename, func(e fsnotify.Event) {
			//logger.S(config.Conf.Logger.Level)
		})
		if err != nil {
			log.Panic(err)
		}
		fn(cfg)
		o.conf = &cfg.BasicConfig
		conf.Conf = &cfg.BasicConfig
	}
}

func WithLogger() Option {
	return func(o *Options) {
		o.conf.Logger = logger.NewConfig()
		o.conf.Logger.Filename = o.conf.Name
		//o.Conf.Logger.TimeFormat = o.Conf.TimeFormat
		logger.InitZapLogger(o.conf.Logger)
		slog.Info("[logger] init done.")
	}
}

func WithMustLocalCache() Option {
	return func(o *Options) {
		var err error
		o.localCache, err = cache.DefaultCache()
		if err != nil {
			log.Panic(err)
		}
		slog.Info("[local_cache] init done.")
	}
}

func WithMustDB(tables ...[]string) Option {
	// TODO gorm.AutoMerge
	return func(o *Options) {
		var err error
		o.db, err = db.NewDB(mysql.Open(o.conf.DB.Source), o.conf.DB)
		if err != nil {
			log.Panic(err.Error())
		}
		slog.Info("[mysql] init done.")
	}
}

func WithMustRedis() Option {
	return func(o *Options) {
		var err error
		o.redis, err = cache.NewRedis(o.conf.Redis)
		if err != nil {
			log.Panic(err.Error())
		}
		slog.Info("[redis] init done.")
	}
}

func WithGinServer(router func(e *gin.Engine, a *Options)) Option {
	return func(o *Options) {
		o.httpServer = router
	}
}

func WithGrpcServer(svr func(rpcServer *grpc.Server, a *Options)) Option {
	return func(o *Options) {
		o.rpcServer = svr
	}
}

// Before and Afters

// WithBeforeStart run funcs before app starts
func WithBeforeStart(fn func(context.Context) error) Option {
	return func(o *Options) {
		o.beforeStart = append(o.beforeStart, fn)
	}
}

// WithBeforeStop run funcs before app stops
func WithBeforeStop(fn func(context.Context) error) Option {
	return func(o *Options) {
		o.beforeStop = append(o.beforeStop, fn)
	}
}

// WithAfterStart run funcs after app starts
func WithAfterStart(fn func(context.Context) error) Option {
	return func(o *Options) {
		o.afterStart = append(o.afterStart, fn)
	}
}

// WithAfterStop run funcs after app stops
func WithAfterStop(fn func(context.Context) error) Option {
	return func(o *Options) {
		o.afterStop = append(o.afterStop, fn)
	}
}
