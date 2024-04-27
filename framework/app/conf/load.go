package conf

import (
	"flag"
	"log/slog"

	"github.com/fsnotify/fsnotify"
	"github.com/gogoclouds/gogo-services/framework/app/validator"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// LoadService 加载配置文件
// 优先级: (相同key)
//
//	1.主配置文件优先级最高
//	2.configs 数组索引越小优先级越高
func LoadService[T any](filepath string, onChange func(e fsnotify.Event)) (*ServiceConfig[T], error) {

	// vpr := viper.GetViper()
	// if filepath != "" {
	// 	vpr.SetConfigFile(filepath)
	// } else {
	// 	vpr.AddConfigPath(".") // "./config.yaml"
	// }
	cfg := new(ServiceConfig[T])
	if err := Load(filepath, cfg, onChange); err != nil {
		return nil, err
	}
	// 加载其他配置文件
	// configs 数组索引越小优先级越高
	for i := len(cfg.Configs) - 1; i >= 0; i-- {
		if err := Load(cfg.Configs[i], cfg, onChange); err != nil {
			return nil, err
		}
	}
	if len(cfg.Configs) > 0 {
		if err := Load(filepath, cfg, onChange); err != nil { // 主配置文件优先级最高, 覆盖其他配置文件
			return nil, err
		}
	}
	return cfg, nil
}

// LoadDefault ./config.yaml
func LoadDefault[T any](onChange func(e fsnotify.Event)) (*T, error) {
	cfg := new(T)
	err := Load(".", cfg, onChange)
	return cfg, err
}

func Load[T any](filepath string, cfg *T, onChange func(e fsnotify.Event)) error {
	vpr := viper.New()
	vpr.SetConfigFile(filepath)
	vpr.ReadInConfig()
	if err := vpr.ReadInConfig(); err != nil {
		return err
	}
	if err := vpr.Unmarshal(cfg); err != nil {
		return err
	}
	if err := validator.Struct(cfg); err != nil {
		return err
	}
	vpr.WatchConfig()
	vpr.OnConfigChange(func(e fsnotify.Event) {
		newCfg := new(T)
		if err := vpr.Unmarshal(newCfg); err != nil {
			slog.Error(err.Error())
			return
		}
		// if err := validator.Struct(newCfg); err != nil {
		// 	slog.Error(err.Error())
		// 	return
		// }
		// TODO merge 判断优先级

		slog.Info("Config file changed: " + e.String())
		onChange(e)
	})
	return nil
}

func BindPFlags() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)
}
