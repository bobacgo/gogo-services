package config

import (
	"fmt"
	"github.com/gogoclouds/gogo-services/common-lib/pkg"
	"github.com/polarismesh/polaris-go"
	"github.com/polarismesh/polaris-go/pkg/model"
	"log"
)

type RemoteConfigFile[T comparable] struct {
	Namespace   string
	FileGroup   string
	FileNameSet pkg.Set[T]
}

type ServerCenter struct{}

func (sc ServerCenter) LoadConfig(dnsConfigFilePath string, remoteConfigFile []string) {
	if dnsConfigFilePath == "" {
		log.Panicf("dns config file path: %s\n", dnsConfigFilePath)
	}
	// TODO
	_, err := polaris.NewConfigAPIByFile("./configs/polaris.yaml")
	if err != nil {
		log.Panicln(err)
	}

}

func main() {
	//db.OpenMySQL()
	configAPI, err := polaris.NewConfigAPIByFile("./configs/polaris.yaml")
	if err != nil {
		log.Panicln(err)
	}

	// 获取远程的配置文件
	namespace := "default"
	fileGroup := "gogo_v1.0.0"
	fileName := "test.yaml"

	configFile, err := configAPI.GetConfigFile(namespace, fileGroup, fileName)
	if err != nil {
		log.Panicln(err)
	}

	// 打印配置文件内容
	log.Printf("config content: %s", configFile.GetContent())
	configFile.AddChangeListener(changeListener)
	select {}
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