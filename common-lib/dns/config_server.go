package dns

import (
	"fmt"
	"github.com/gogoclouds/gogo-services/common-lib/pkg/mapset"
	"github.com/polarismesh/polaris-go"
	"github.com/polarismesh/polaris-go/pkg/model"
	"log"
)

type FileMetadata struct {
	Namespace   string
	FileGroup   string
	FileNameSet mapset.Set[string] // 要拉起的文件集合
}

// LoadConfig
//
// eg: dnsConfigFilePath = "./configs/polaris.yaml"
func (sc serverCenter) LoadConfig(dnsConfigFilePath string, remoteConfigFile *FileMetadata) {
	if dnsConfigFilePath == "" {
		log.Panicf("dns config file path: %s\n", dnsConfigFilePath)
	}

	configApi, err := polaris.NewConfigAPIByFile("./configs/polaris.yaml")
	if err != nil {
		log.Panicln(err)
	}

	remoteConfigFile.FileNameSet.Each(func(filename string) {
		// 获取远程的配置文件
		configFile, err := configApi.GetConfigFile(remoteConfigFile.Namespace, remoteConfigFile.FileGroup, filename)
		if err != nil {
			log.Panicln(err)
		}

		// 打印配置文件内容
		fmt.Printf("%s\n", configFile.GetContent())
		configFile.AddChangeListener(changeListener)
	})
}

func changeListener(event model.ConfigFileChangeEvent) {
	log.Printf("config change: %+v\n", event.ConfigFileMetadata)
	log.Printf("change type : %d\n", event.ChangeType)
	fmt.Println("-------------------------------------------- ")
	fmt.Println("------------- old value ------------------- ")
	fmt.Println(event.OldValue)
	fmt.Println("------------- new value ------------------- ")
	fmt.Println(event.NewValue)
	fmt.Println("-------------------------------------------- ")
}
