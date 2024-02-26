package app

import (
	"context"
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

type Option func(o *options)

type options struct {
	Conf  *Config
	DB    *gorm.DB
	Redis redis.UniversalClient

	appid           string
	endpoints       []*url.URL
	sigs            []os.Signal
	registrar       registry.ServiceRegistrar
	registryTimeout time.Duration

	httpServer func(a *App, e *gin.Engine)

	rpcServer func(a *App, s *grpc.Server)
	// Before and After hook
	beforeStart, beforeStop, afterStart, afterStop []func(context.Context) error
}

func WithAppId(id string) Option {
	return func(o *options) {
		o.appid = id
	}
}

func WithEndpoints(endpoints []*url.URL) Option {
	return func(o *options) {
		o.endpoints = endpoints
	}
}

func WithSignal(sigs []os.Signal) Option {
	return func(o *options) {
		o.sigs = sigs
	}
}

func WithRegistrar(registrar registry.ServiceRegistrar) Option {
	return func(o *options) {
		o.registrar = registrar
	}
}

func WithRegistrarTimeout(rt time.Duration) Option {
	return func(o *options) {
		o.registryTimeout = rt
	}
}

func WithConfig(filename string) Option {
	return func(o *options) {
		var err error
		o.Conf, err = conf.Load[Config](filename, func(e fsnotify.Event) {
			//logger.S(config.Conf.Logger.Level)
		})
		if err != nil {
			panic(err)
		}
	}
}

func WithLogger() Option {
	return func(o *options) {
		o.Conf.Logger.Filename = o.Conf.Name
		o.Conf.Logger.TimeFormat = o.Conf.TimeFormat
		logger.InitZapLogger(o.Conf.Logger)
		logger.Info("logger init done...")
	}
}

func WithDB(tables ...[]string) Option {
	// TODO gorm.AutoMerge
	return func(o *options) {
		newDB, err := db.NewDB(mysql.Open(o.Conf.DB.Source), o.Conf.DB)
		if err != nil {
			logger.Panic(err.Error())
		}
		logger.Info("mysql init done...")
		o.DB = newDB
	}
}

func WithRedis() Option {
	return func(o *options) {
		newRedis, err := cache.NewRedis(o.Conf.Redis)
		if err != nil {
			logger.Panic(err.Error())
		}
		logger.Info("redis init done...")
		o.Redis = newRedis
	}
}

func WithGinServer(router func(a *App, e *gin.Engine)) Option {
	return func(o *options) {
		o.httpServer = router
	}
}

func WithGrpcServer(svr func(a *App, rpcServer *grpc.Server)) Option {
	return func(o *options) {
		o.rpcServer = svr
	}
}

// Before and Afters

// BeforeStart run funcs before app starts
func BeforeStart(fn func(context.Context) error) Option {
	return func(o *options) {
		o.beforeStart = append(o.beforeStart, fn)
	}
}

// BeforeStop run funcs before app stops
func BeforeStop(fn func(context.Context) error) Option {
	return func(o *options) {
		o.beforeStop = append(o.beforeStop, fn)
	}
}

// AfterStart run funcs after app starts
func AfterStart(fn func(context.Context) error) Option {
	return func(o *options) {
		o.afterStart = append(o.afterStart, fn)
	}
}

// AfterStop run funcs after app stops
func AfterStop(fn func(context.Context) error) Option {
	return func(o *options) {
		o.afterStop = append(o.afterStop, fn)
	}
}