package config

import "github.com/gogoclouds/gogo-services/common-lib/app/conf"

var Conf *conf.ServiceConfig[Service]

type Service struct {
	ErrAttemptLimit int `napstructure:"errAttemptLimit"` // 密码错误次数限制
}

// TODO 配置检查
