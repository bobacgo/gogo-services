package conf

import (
	"github.com/gogoclouds/gogo-services/common-lib/app/cache"
	"github.com/gogoclouds/gogo-services/common-lib/app/db"
	"github.com/gogoclouds/gogo-services/common-lib/app/logger"
	"github.com/gogoclouds/gogo-services/common-lib/app/security/config"
)

var Conf *BasicConfig

type EnvType string

const (
	EnvDev  EnvType = "dev"
	EnvTest EnvType = "test"
	EnvProd EnvType = "prod"
)

type ServiceConfig[T any] struct {
	BasicConfig `mapstructure:",squash"`
	Service     T `mapstructure:"service"`
}

type BasicConfig struct {
	Name       string  `mapstructure:"name" validate:"required"`    // 服务名
	Version    string  `mapstructure:"version" validate:"required"` // 版本号
	Env        EnvType `mapstructure:"env" validate:"required"`
	TimeFormat string  `mapstructure:"timeFormat"`
	Server     struct {
		Http Transport `mapstructure:"http"`
		Rpc  Transport `mapstructure:"rpc"`
	}
	Security config.Config   `mapstructure:"security"`
	Logger   logger.Config   `mapstructure:"logger"`
	Registry Transport       `mapstructure:"registry"`
	DB       db.Config       `mapstructure:"db"`
	Redis    cache.RedisConf `mapstructure:"redis"`
}

// Transport 传输协议
type Transport struct {
	Addr    string `mapstructure:"addr"`    // 0.0.0.0:8000
	Timeout string `mapstructure:"timeout"` // 1s
}
