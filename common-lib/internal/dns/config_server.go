package dns

import (
	"fmt"
	"log"

	"github.com/gogoclouds/gogo-services/common-lib/internal/dns/config"
	"github.com/gogoclouds/gogo-services/common-lib/pkg/stream"
	"gopkg.in/yaml.v3"

	"github.com/gogoclouds/gogo-services/common-lib/g"
	"github.com/polarismesh/polaris-go"
	"github.com/polarismesh/polaris-go/pkg/model"
)

// LoadConfig
//
// eg: dnsConfigFilePath = "./configs/polaris.yaml"
func (sc serverCenter) LoadConfig(dnsConfigFilePath string, remoteConfigFile *config.FileMetadata) {
	configApi, err := polaris.NewConfigAPIByFile(dnsConfigFilePath)
	if err != nil {
		panic(err)
	}

	g.Conf = config.New()
	stream.New(remoteConfigFile.Filenames).Distinct(). // 有序去重
								Each(func(idx int, filename string) {
			// 获取远程的配置文件
			configFile, err := configApi.GetConfigFile(remoteConfigFile.Namespace, remoteConfigFile.Group, filename)
			if err != nil {
				panic(err)
			}
			if err = g.Conf.Sync([]byte(configFile.GetContent())); err != nil {
				panic(err)
			}
			configFile.AddChangeListener(sc.changeListener)
		})
	// 打印读取到的配置信息
	sc.printConfInfo()
}

func (sc serverCenter) changeListener(event model.ConfigFileChangeEvent) {
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

func (sc serverCenter) printConfInfo() {
	// 打印读取到的配置信息
	printConf, _ := yaml.Marshal(g.Conf.Config())
	log.Println("======================= config info ========================")
	fmt.Println(string(printConf))
	log.Println("======================= config info end ====================")
}
