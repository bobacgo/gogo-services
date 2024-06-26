package dns

import (
	"fmt"
	"log"

	"github.com/gogoclouds/gogo-services/framework/g"
	"github.com/gogoclouds/gogo-services/framework/internal/dns/config"
	"github.com/gogoclouds/gogo-services/framework/pkg/stream"
	"gopkg.in/yaml.v3"

	"github.com/polarismesh/polaris-go"
	"github.com/polarismesh/polaris-go/api"
	"github.com/polarismesh/polaris-go/pkg/model"
)

// ConfigServer 配置中心
type configServer struct {
	ctx api.SDKContext
}

// Load load remote config to g.Conf
func (c configServer) Load(remoteConfigFile *config.FileMetadata) {
	configApi := polaris.NewConfigAPIByContext(c.ctx)
	g.Conf = config.New()
	stream.New(remoteConfigFile.Filenames).
		Distinct(). // 有序去重
		Each(func(idx int, filename string) {
			// 获取远程的配置文件
			configFile, err := configApi.GetConfigFile(remoteConfigFile.Namespace, remoteConfigFile.Group, filename)
			if err != nil {
				panic(err)
			}
			if err = g.Conf.Sync([]byte(configFile.GetContent())); err != nil {
				panic(err)
			}
			configFile.AddChangeListener(c.changeListener)
		})
	// 打印读取到的配置信息
	go c.printConfInfo()
}

func (c configServer) changeListener(event model.ConfigFileChangeEvent) {
	log.Printf("config change: %+v\n", event.ConfigFileMetadata)
	log.Printf("change type : %d\n", event.ChangeType)
	fmt.Println("-------------------------------------------- ")
	fmt.Println("------------- old value ------------------- ")
	fmt.Println(event.OldValue)
	fmt.Println("------------- new value ------------------- ")
	fmt.Println(event.NewValue)
	fmt.Println("-------------------------------------------- ")
	if err := g.Conf.Sync([]byte(event.NewValue)); err != nil {
		g.Log.Errorf("sync config: %v", err)
	}
}

func (c configServer) printConfInfo() {
	// 打印读取到的配置信息
	printConf, _ := yaml.Marshal(g.Conf.Config())
	log.Println("======================= load config info ========================")
	fmt.Println(string(printConf))
}
