package conf

import (
	"flag"
	"github.com/fsnotify/fsnotify"
	"github.com/gogoclouds/gogo-services/common-lib/app/logger"
	"github.com/gogoclouds/gogo-services/common-lib/app/validator"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// LoadDefault ./config.yaml
func LoadDefault[T any](onChange func(e fsnotify.Event)) (*T, error) {
	return Load[T](".", onChange)
}

func Load[T any](filepath string, onChange func(e fsnotify.Event)) (*T, error) {
	vpr := viper.GetViper()
	if filepath != "" {
		vpr.SetConfigFile(filepath)
	} else {
		vpr.AddConfigPath(".") // "./config.yaml"
	}
	cf := new(T)
	if err := vpr.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := vpr.Unmarshal(cf); err != nil {
		return nil, err
	}
	if err := validator.Struct(cf); err != nil {
		return nil, err
	}

	vpr.OnConfigChange(func(e fsnotify.Event) {
		newCfg := new(T)
		if err := vpr.Unmarshal(newCfg); err != nil {
			logger.Error(err.Error())
			return
		}
		if err := validator.Struct(newCfg); err != nil {
			logger.Error(err.Error())
			return
		}
		cf = newCfg
		onChange(e)
	})
	vpr.WatchConfig()
	return cf, nil
}

func BindPFlags() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)
}
