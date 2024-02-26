package app

import (
	"github.com/gogoclouds/gogo-services/common-lib/app/cache"
	"github.com/gogoclouds/gogo-services/common-lib/app/db"
	"github.com/gogoclouds/gogo-services/common-lib/app/logger"
	"github.com/gogoclouds/gogo-services/common-lib/app/security"
)

type EnvType string

const (
	EnvDev  EnvType = "dev"
	EnvTest EnvType = "test"
	EnvProd EnvType = "prod"
)

type Config struct {
	Name       string  `yaml:"name"`    // 服务名
	Version    string  `yaml:"version"` // 版本号
	Env        EnvType `yaml:"env"`
	TimeFormat string  `yaml:"timeFormat"`
	Server     struct {
		Http Transport `yaml:"http"`
		Rpc  Transport `yaml:"rpc"`
	}
	Logger   logger.Config   `yaml:"logger"`
	Registry Transport       `yaml:"registry"`
	Jwt      security.Config `yaml:"jwt"`
	DB       db.Config       `yaml:"db"`
	Redis    cache.RedisConf `yaml:"redis"`
}

// Transport 传输协议
type Transport struct {
	Addr    string `yaml:"addr"`    // 0.0.0.0:8000
	Timeout string `yaml:"timeout"` // 1s
}