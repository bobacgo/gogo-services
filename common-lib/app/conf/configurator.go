package conf

import (
	"flag"

	"github.com/fsnotify/fsnotify"
	"github.com/gogoclouds/gogo-services/common-lib/app/logger"
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
		return cf, err
	}
	if err := vpr.Unmarshal(cf); err != nil {
		return cf, err
	}
	vpr.OnConfigChange(func(e fsnotify.Event) {
		if err := vpr.Unmarshal(cf); err != nil {
			logger.Error(err.Error())
		}
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
