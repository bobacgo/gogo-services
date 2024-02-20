package conf_test

import (
	"flag"
	"fmt"
	"testing"

	"github.com/fsnotify/fsnotify"
	"github.com/gogoclouds/gogo-services/common-lib/app/conf"
	"github.com/gogoclouds/gogo-services/common-lib/app/logger"
)

func TestLoad(t *testing.T) {
	//config := conf.Load("./config.yaml")
	//config := conf.Load("")
	flag.String("filename", "project", "help message for flagname")
	conf.BindPFlags()
	config, err := conf.Load[logger.Config]("../../config/config.yaml", func(e fsnotify.Event) {})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", config)
	//select {}
}
