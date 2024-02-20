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
	conf *Config

	appid     string
	endpoints []*url.URL

	sigs            []os.Signal
	registrar       registry.ServiceRegistrar
	registryTimeout time.Duration
	httpServer      func(a *App, e *gin.Engine)
	rpcServer       func(a *App, s *grpc.Server)

	// Before and After hook
	beforeStart, beforeStop, afterStart, afterStop []func(context.Context) error

	DB    *gorm.DB
	Redis redis.UniversalClient
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
		o.conf, err = conf.Load[Config](filename, func(e fsnotify.Event) {
			//logger.S(config.Conf.Logger.Level)
		})
		if err != nil {
			panic(err)
		}
	}
}

func WithLogger() Option {
	return func(o *options) {
		o.conf.Logger.Filename = o.conf.Name
		o.conf.Logger.TimeFormat = o.conf.TimeFormat
		logger.InitZapLogger(o.conf.Logger)
	}
}

func WithDB(tables ...[]string) Option {
	// TODO gorm.AutoMerge
	return func(o *options) {
		newDB, err := db.NewDB(mysql.Open(o.conf.DB.Source), o.conf.DB)
		if err != nil {
			logger.Panic(err.Error())
		}
		o.DB = newDB
	}
}

func WithRedis() Option {
	return func(o *options) {
		newRedis, err := cache.NewRedis(o.conf.Redis)
		if err != nil {
			logger.Panic(err.Error())
		}
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
